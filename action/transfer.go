// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package action

import (
	"encoding/hex"
	"math"
	"math/big"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"

	"github.com/Vedaad-Shakib/IoTeX-Sim/explorer/idl/explorer"
	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/hash"
	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/keypair"
	"github.com/Vedaad-Shakib/IoTeX-Sim/pkg/version"
	"github.com/Vedaad-Shakib/IoTeX-Sim/proto"
)

const (
	// TransferPayloadGas represents the transfer payload gas per uint
	TransferPayloadGas = uint64(100)
	// TransferBaseIntrinsicGas represents the base intrinsic gas for transfer
	TransferBaseIntrinsicGas = uint64(10000)
)

// Transfer defines the struct of account-based transfer
type Transfer struct {
	AbstractAction
	amount  *big.Int
	payload []byte
	// Coinbase transfer is not expected to be received from the network but can only be generated by block producer
	isCoinbase bool
}

// NewTransfer returns a Transfer instance
func NewTransfer(
	nonce uint64,
	amount *big.Int,
	sender string,
	recipient string,
	payload []byte,
	gasLimit uint64,
	gasPrice *big.Int,
) (*Transfer, error) {
	if len(sender) == 0 || len(recipient) == 0 {
		return nil, errors.Wrap(ErrAddress, "address of sender or recipient is empty")
	}

	return &Transfer{
		AbstractAction: AbstractAction{
			version:  version.ProtocolVersion,
			nonce:    nonce,
			srcAddr:  sender,
			dstAddr:  recipient,
			gasLimit: gasLimit,
			gasPrice: gasPrice,
		},
		amount:     amount,
		payload:    payload,
		isCoinbase: false,
		// SenderPublicKey and Signature will be populated in Sign()
	}, nil
}

// NewCoinBaseTransfer returns a coinbase Transfer
func NewCoinBaseTransfer(nonce uint64, amount *big.Int, recipient string) *Transfer {
	return &Transfer{
		AbstractAction: AbstractAction{
			nonce:   nonce,
			version: version.ProtocolVersion,
			dstAddr: recipient,
		},
		amount: amount,
		// payload is empty for now
		payload:    []byte{},
		isCoinbase: true,
		// SenderPublicKey and Signature will be populated in Sign()
	}
}

// Amount returns the amount
func (tsf *Transfer) Amount() *big.Int { return tsf.amount }

// Payload returns the payload bytes
func (tsf *Transfer) Payload() []byte { return tsf.payload }

// IsCoinbase returns a boolean value to indicate if a transfer is a coinbase one
func (tsf *Transfer) IsCoinbase() bool { return tsf.isCoinbase }

// Sender returns the sender address. It's the wrapper of Action.SrcAddr
func (tsf *Transfer) Sender() string { return tsf.SrcAddr() }

// SenderPublicKey returns the sender public key. It's the wrapper of Action.SrcPubkey
func (tsf *Transfer) SenderPublicKey() keypair.PublicKey { return tsf.SrcPubkey() }

// SetSenderPublicKey sets the sender public key. It's the wrapper of Action.SetSrcPubkey
func (tsf *Transfer) SetSenderPublicKey(pubkey keypair.PublicKey) { tsf.SetSrcPubkey(pubkey) }

// Recipient returns the recipient address. It's the wrapper of Action.DstAddr
func (tsf *Transfer) Recipient() string { return tsf.DstAddr() }

// IsContract returns true for contract action
func (tsf *Transfer) IsContract() bool {
	return tsf.dstAddr == EmptyAddress
}

// TotalSize returns the total size of this Transfer
func (tsf *Transfer) TotalSize() uint32 {
	size := tsf.BasicActionSize()
	size += uint32(1) // Size of boolean isCoinbase
	if tsf.amount != nil && len(tsf.amount.Bytes()) > 0 {
		size += uint32(len(tsf.amount.Bytes()))
	}

	return size + uint32(len(tsf.payload))
}

// ByteStream returns a raw byte stream of this Transfer
func (tsf *Transfer) ByteStream() []byte {
	stream := tsf.BasicActionByteStream()
	if tsf.amount != nil && len(tsf.amount.Bytes()) > 0 {
		stream = append(stream, tsf.amount.Bytes()...)
	}
	stream = append(stream, tsf.payload...)
	// Signature = Sign(hash(ByteStream())), so not included
	if tsf.isCoinbase {
		stream = append(stream, 1)
	} else {
		stream = append(stream, 0)
	}
	return stream
}

// Proto converts Transfer to protobuf's ActionPb
func (tsf *Transfer) Proto() *iproto.ActionPb {
	// used by account-based model
	act := &iproto.ActionPb{
		Action: &iproto.ActionPb_Transfer{
			Transfer: &iproto.TransferPb{
				Recipient:  tsf.dstAddr,
				Payload:    tsf.payload,
				IsCoinbase: tsf.isCoinbase,
			},
		},
		Version:      tsf.version,
		Sender:       tsf.srcAddr,
		SenderPubKey: tsf.srcPubkey[:],
		Nonce:        tsf.nonce,
		GasLimit:     tsf.gasLimit,
		Signature:    tsf.signature,
	}

	if tsf.amount != nil && len(tsf.amount.Bytes()) > 0 {
		act.GetTransfer().Amount = tsf.amount.Bytes()
	}
	if tsf.gasPrice != nil && len(tsf.gasPrice.Bytes()) > 0 {
		act.GasPrice = tsf.gasPrice.Bytes()
	}
	return act
}

// ToJSON converts Transfer to TransferJSON
func (tsf *Transfer) ToJSON() *explorer.Transfer {
	// used by account-based model
	t := &explorer.Transfer{
		Version:      int64(tsf.version),
		Nonce:        int64(tsf.nonce),
		Sender:       tsf.srcAddr,
		Recipient:    tsf.dstAddr,
		Payload:      hex.EncodeToString(tsf.payload),
		SenderPubKey: keypair.EncodePublicKey(tsf.srcPubkey),
		GasLimit:     int64(tsf.gasLimit),
		Signature:    hex.EncodeToString(tsf.signature),
		IsCoinbase:   tsf.isCoinbase,
	}

	if tsf.amount != nil && len(tsf.amount.String()) > 0 {
		t.Amount = tsf.amount.String()
	}
	if tsf.gasPrice != nil && len(tsf.gasPrice.String()) > 0 {
		t.GasPrice = tsf.gasPrice.String()
	}
	return t
}

// Serialize returns a serialized byte stream for the Transfer
func (tsf *Transfer) Serialize() ([]byte, error) {
	return proto.Marshal(tsf.Proto())
}

// LoadProto converts a protobuf's ActionPb to Transfer
func (tsf *Transfer) LoadProto(pbAct *iproto.ActionPb) error {
	if pbAct == nil {
		return errors.New("empty action proto to load")
	}
	if tsf == nil {
		return errors.New("nil action to load proto")
	}
	*tsf = Transfer{}
	srcPub, err := keypair.BytesToPublicKey(pbAct.SenderPubKey)
	if err != nil {
		return err
	}
	pbTsf := pbAct.GetTransfer()
	if pbTsf == nil {
		return errors.New("empty Transfer action proto to load")
	}

	ab := &Builder{}
	act := ab.SetVersion(pbAct.Version).
		SetNonce(pbAct.Nonce).
		SetSourceAddress(pbAct.Sender).
		SetSourcePublicKey(srcPub).
		SetGasLimit(pbAct.GasLimit).
		SetGasPriceByBytes(pbAct.GasPrice).
		SetDestinationAddress(pbTsf.Recipient).
		Build()
	act.SetSignature(pbAct.Signature)
	tsf.AbstractAction = act

	tsf.amount = big.NewInt(0)
	if len(pbTsf.Amount) > 0 {
		tsf.amount.SetBytes(pbTsf.Amount)
	}
	tsf.payload = pbTsf.Payload
	tsf.isCoinbase = pbTsf.IsCoinbase
	return nil
}

// NewTransferFromJSON creates a new Transfer from TransferJSON
func NewTransferFromJSON(jsonTsf *explorer.Transfer) (*Transfer, error) {
	tsf := &Transfer{}
	tsf.version = uint32(jsonTsf.Version)
	// used by account-based model
	tsf.nonce = uint64(jsonTsf.Nonce)
	tsf.srcAddr = jsonTsf.Sender
	tsf.dstAddr = jsonTsf.Recipient
	tsf.gasLimit = uint64(jsonTsf.GasLimit)
	amount, ok := big.NewInt(0).SetString(jsonTsf.Amount, 10)
	if !ok {
		return nil, errors.New("failed to set amount of transfer")
	}
	tsf.amount = amount
	gasPrice, ok := big.NewInt(0).SetString(jsonTsf.GasPrice, 10)
	if !ok {
		return nil, errors.New("failed to set gas price of transfer")
	}
	tsf.gasPrice = gasPrice
	payload, err := hex.DecodeString(jsonTsf.Payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode transfer payload")
	}
	tsf.payload = payload
	senderPubKey, err := keypair.StringToPubKeyBytes(jsonTsf.SenderPubKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode transfer sender public key")
	}
	copy(tsf.srcPubkey[:], senderPubKey)
	signature, err := hex.DecodeString(jsonTsf.Signature)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode transfer signature")
	}
	tsf.signature = signature
	tsf.isCoinbase = jsonTsf.IsCoinbase

	return tsf, nil
}

// Deserialize parse the byte stream into Transfer
func (tsf *Transfer) Deserialize(buf []byte) error {
	pbAct := &iproto.ActionPb{}
	if err := proto.Unmarshal(buf, pbAct); err != nil {
		return err
	}
	return tsf.LoadProto(pbAct)
}

// Hash returns the hash of the Transfer
func (tsf *Transfer) Hash() hash.Hash32B {
	return blake2b.Sum256(tsf.ByteStream())
}

// IntrinsicGas returns the intrinsic gas of a transfer
func (tsf *Transfer) IntrinsicGas() (uint64, error) {
	payloadSize := uint64(len(tsf.Payload()))
	if (math.MaxUint64-TransferBaseIntrinsicGas)/TransferPayloadGas < payloadSize {
		return 0, ErrOutOfGas
	}

	return payloadSize*TransferPayloadGas + TransferBaseIntrinsicGas, nil
}

// Cost returns the total cost of a transfer
func (tsf *Transfer) Cost() (*big.Int, error) {
	intrinsicGas, err := tsf.IntrinsicGas()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get intrinsic gas for the transfer")
	}
	transferFee := big.NewInt(0).Mul(tsf.GasPrice(), big.NewInt(0).SetUint64(intrinsicGas))
	return big.NewInt(0).Add(tsf.Amount(), transferFee), nil
}
