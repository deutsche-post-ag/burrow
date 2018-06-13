package txs

import (
	"fmt"

	acm "github.com/hyperledger/burrow/account"
	"github.com/hyperledger/burrow/account/state"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/txs/payload"
)

type Encoder interface {
	EncodeTx(envelope *Envelope) ([]byte, error)
}

type Decoder interface {
	DecodeTx(txBytes []byte) (*Envelope, error)
}

// An envelope contains both the signable Tx and the signatures for each input (in signatories)
type Envelope struct {
	Signatories []Signatory
	Tx          *Tx
}

// Enclose a Payload in an Envelope so it is ready to be signed by first wrapping the Payload
// as a Tx (including ChainID) and writing it to the Tx field of the Envelope
func Enclose(chainID string, payload payload.Payload) *Envelope {
	body := NewTx(payload)
	body.ChainID = chainID
	return body.Enclose()
}

func (txEnv *Envelope) String() string {
	return fmt.Sprintf("TxEnvelope{Signatures: %v, Tx: %s}", len(txEnv.Signatories), txEnv.Tx)
}

// Signatory contains signature and one or both of Address and PublicKey to identify the signer
type Signatory struct {
	Address   *crypto.Address
	PublicKey *crypto.PublicKey
	crypto.Signature
}

// Attempts to 'realise' the PublicKey and Address of a Signatory possibly referring to state
// in the case where the Signatory contains an Address by no PublicKey. Checks consistency in other
// cases, possibly generating the Address from the PublicKey
func (s *Signatory) RealisePublicKey(getter state.AccountGetter) error {
	const errPrefix = "could not realise public key for signatory"
	if s.PublicKey == nil {
		if s.Address == nil {
			return fmt.Errorf("%s: address not provided", errPrefix)
		}
		acc, err := getter.GetAccount(*s.Address)
		if err != nil {
			return fmt.Errorf("%s: could not get account %v: %v", errPrefix, *s.Address, err)
		}
		publicKey := acc.PublicKey()
		s.PublicKey = &publicKey
	}
	if !s.PublicKey.IsValid() {
		return fmt.Errorf("%s: public key %v is invalid", errPrefix, *s.PublicKey)
	}
	address := s.PublicKey.Address()
	if s.Address == nil {
		s.Address = &address
	} else if address != *s.Address {
		return fmt.Errorf("address %v provided with signatory does not match address generated from "+
			"public key %v", *s.Address, address)
	}
	return nil
}

// Verifies the validity of the Signatories' Signatures in the Envelope. The Signatories must
// appear in the same order as the inputs as returned by Tx.GetInputs().
func (txEnv *Envelope) Verify(getter state.AccountGetter) error {
	errPrefix := fmt.Sprintf("could not verify transaction %X", txEnv.Tx.Hash())
	inputs := txEnv.Tx.GetInputs()
	if len(inputs) != len(txEnv.Signatories) {
		return fmt.Errorf("%s: number of inputs (= %v) should equal number of signatories (= %v)",
			errPrefix, len(inputs), len(txEnv.Signatories))
	}
	signBytes, err := txEnv.Tx.SignBytes()
	if err != nil {
		return fmt.Errorf("%s: could not generate SignBytes: %v", errPrefix, err)
	}
	// Expect order to match (we could build lookup but we want Verify to be quicker than Sign which does order sigs)
	for i, s := range txEnv.Signatories {
		if inputs[i].Address != *s.Address {
			return fmt.Errorf("signatory %v has address %v but input %v has address %v",
				i, *s.Address, i, inputs[i].Address)
		}
		if !s.PublicKey.Verify(signBytes, s.Signature) {
			return fmt.Errorf("invalid signature in signatory %v ", *s.Address)
		}
	}
	return nil
}

// Sign the Tx Envelope by adding Signatories containing the signatures for each TxInput.
// signing accounts for each input must be provided (in any order).
func (txEnv *Envelope) Sign(signingAccounts ...acm.AddressableSigner) error {
	// Clear any existing
	txEnv.Signatories = nil
	signBytes, err := txEnv.Tx.SignBytes()
	if err != nil {
		return err
	}
	signingAccountMap := make(map[crypto.Address]acm.AddressableSigner)
	for _, sa := range signingAccounts {
		signingAccountMap[sa.Address()] = sa
	}
	// Sign in order of inputs
	for i, in := range txEnv.Tx.GetInputs() {
		sa, ok := signingAccountMap[in.Address]
		if !ok {
			return fmt.Errorf("account to sign %v (position %v) not passed to Sign, passed: %v", in, i, signingAccounts)
		}
		sig, err := sa.Sign(signBytes)
		if err != nil {
			return err
		}
		address := sa.Address()
		publicKey := sa.PublicKey()
		txEnv.Signatories = append(txEnv.Signatories, Signatory{
			Address:   &address,
			PublicKey: &publicKey,
			Signature: sig,
		})
	}
	return nil
}
