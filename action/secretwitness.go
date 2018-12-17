// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package action

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"

	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/enc"
	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/hash"
	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/keypair"
	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/version"
	"github.com/Vedaad-Shakib/IoTeX-Sim/proto"
)

// SecretWitness defines the struct of DKG secret witness
type SecretWitness struct {
	AbstractAction
	witness [][]byte
}

// NewSecretWitness returns a SecretWitness instance
func NewSecretWitness(
	nonce uint64,
	sender string,
	witness [][]byte,
) (*SecretWitness, error) {
	if len(sender) == 0 {
		return nil, errors.Wrap(ErrAddress, "address of sender is empty")
	}
	return &SecretWitness{
		AbstractAction: AbstractAction{
			version: version.ProtocolVersion,
			nonce:   nonce,
			srcAddr: sender,
		},
		witness: witness,
	}, nil
}

// Witness returns the witness
func (sw *SecretWitness) Witness() [][]byte { return sw.witness }

// ByteStream returns a raw byte stream of this SecretWitness
func (sw *SecretWitness) ByteStream() []byte {
	stream := make([]byte, 4)
	enc.MachineEndian.PutUint32(stream, sw.version)
	temp := make([]byte, 8)
	enc.MachineEndian.PutUint64(temp, sw.nonce)
	stream = append(stream, temp...)
	stream = append(stream, sw.srcAddr...)
	for _, w := range sw.witness {
		stream = append(stream, w...)
	}
	return stream
}

// Proto converts SecretWitness to protobuf's ActionPb
func (sw *SecretWitness) Proto() *iproto.ActionPb {
	// used by account-based model
	act := &iproto.ActionPb{
		Action: &iproto.ActionPb_SecretWitness{
			SecretWitness: &iproto.SecretWitnessPb{
				Witness: sw.witness,
			},
		},
		Version:      sw.version,
		Sender:       sw.srcAddr,
		SenderPubKey: sw.srcPubkey[:],
		Nonce:        sw.nonce,
	}
	return act
}

// Serialize returns a serialized byte stream for the SecretWitness
func (sw *SecretWitness) Serialize() ([]byte, error) {
	return proto.Marshal(sw.Proto())
}

// LoadProto converts a protobuf's ActionPb to SecretWitness
func (sw *SecretWitness) LoadProto(pbAct *iproto.ActionPb) error {
	if pbAct == nil {
		return errors.New("empty action proto to load")
	}
	srcPub, err := keypair.BytesToPublicKey(pbAct.SenderPubKey)
	if err != nil {
		return err
	}
	if sw == nil {
		return errors.New("nil action to load proto")
	}
	*sw = SecretWitness{}
	pbSecretWitness := pbAct.GetSecretWitness()
	if pbSecretWitness == nil {
		return errors.New("empty CreateDeposit action proto to load")
	}

	ab := &Builder{}
	act := ab.SetVersion(pbAct.Version).
		SetNonce(pbAct.Nonce).
		SetSourceAddress(pbAct.Sender).
		SetSourcePublicKey(srcPub).
		Build()
	act.SetSignature(pbAct.Signature)
	sw.AbstractAction = act
	sw.witness = pbSecretWitness.Witness
	return nil
}

// Deserialize parses the byte stream into SecretWitness
func (sw *SecretWitness) Deserialize(buf []byte) error {
	pbAct := &iproto.ActionPb{}
	if err := proto.Unmarshal(buf, pbAct); err != nil {
		return err
	}
	return sw.LoadProto(pbAct)
}

// Hash returns the hash of the SecretWitness
func (sw *SecretWitness) Hash() hash.Hash32B {
	return blake2b.Sum256(sw.ByteStream())
}

// IntrinsicGas returns the intrinsic gas of a secret witness
func (sw *SecretWitness) IntrinsicGas() (uint64, error) { return 0, nil }
