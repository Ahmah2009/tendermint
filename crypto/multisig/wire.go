package multisig

import (
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// TODO: Figure out API for others to either add their own pubkey types, or
// to make verify / marshal accept a cdc.
const (
	ThresholdPubkeyAminoRoute = "tendermint/PubkeyMultisigThreshold"
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ThresholdMultiSignaturePubKey{},
		ThresholdPubkeyAminoRoute, nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
		ed25519.PubKeyAminoRoute, nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{},
		secp256k1.PubKeyAminoRoute, nil)
}
