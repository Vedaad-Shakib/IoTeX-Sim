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
	// EmptyAddress is the empty string
	EmptyAddress = ""
	// ExecutionDataGas represents the execution data gas per uint
	ExecutionDataGas = uint64(100)
	// ExecutionBaseIntrinsicGas represents the base intrinsic gas for execution
	ExecutionBaseIntrinsicGas = uint64(10000)
)

// Execution defines the struct of account-based contract execution
type Execution struct {
	AbstractAction
	amount *big.Int
	data   []byte
}

// NewExecution returns a Execution instance
func NewExecution(executorAddress string, contractAddress string, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (*Execution, error) {
	if executorAddress == "" {
		return nil, errors.Wrap(ErrAddress, "address of the executor is empty")
	}

	return &Execution{
		AbstractAction: AbstractAction{
			version:  version.ProtocolVersion,
			nonce:    nonce,
			srcAddr:  executorAddress,
			dstAddr:  contractAddress,
			gasLimit: gasLimit,
			gasPrice: gasPrice,
		},
		amount: amount,
		data:   data,
	}, nil
}

// Executor returns an executor address
func (ex *Execution) Executor() string {
	return ex.SrcAddr()
}

// ExecutorPublicKey returns the executor's public key
func (ex *Execution) ExecutorPublicKey() keypair.PublicKey {
	return ex.SrcPubkey()
}

// SetExecutorPublicKey sets the executor's public key
func (ex *Execution) SetExecutorPublicKey(executorPubkey keypair.PublicKey) {
	ex.SetSrcPubkey(executorPubkey)
}

// Contract returns a contract address
func (ex *Execution) Contract() string {
	return ex.DstAddr()
}

// Amount returns the amount
func (ex *Execution) Amount() *big.Int {
	return ex.amount
}

// Data returns the data bytes
func (ex *Execution) Data() []byte {
	return ex.data
}

// TotalSize returns the total size of this Execution
func (ex *Execution) TotalSize() uint32 {
	size := ex.BasicActionSize()
	if ex.amount != nil && len(ex.amount.Bytes()) > 0 {
		size += uint32(len(ex.amount.Bytes()))
	}

	return size + uint32(len(ex.data))
}

// ByteStream returns a raw byte stream of this Transfer
func (ex *Execution) ByteStream() []byte {
	stream := ex.BasicActionByteStream()
	if ex.amount != nil && len(ex.amount.Bytes()) > 0 {
		stream = append(stream, ex.amount.Bytes()...)
	}
	// Signature = Sign(hash(ByteStream())), so not included
	return append(stream, ex.data...)
}

// Proto converts Execution to protobuf's ActionPb
func (ex *Execution) Proto() *iproto.ActionPb {
	act := &iproto.ActionPb{
		Action: &iproto.ActionPb_Execution{
			Execution: &iproto.ExecutionPb{
				Contract: ex.dstAddr,
				Data:     ex.data,
			},
		},
		Version:      ex.version,
		Sender:       ex.srcAddr,
		SenderPubKey: ex.srcPubkey[:],
		Nonce:        ex.nonce,
		GasLimit:     ex.gasLimit,
		Signature:    ex.signature,
	}
	if ex.amount != nil && len(ex.amount.Bytes()) > 0 {
		act.GetExecution().Amount = ex.amount.Bytes()
	}
	if ex.gasPrice != nil && len(ex.gasPrice.Bytes()) > 0 {
		act.GasPrice = ex.gasPrice.Bytes()
	}
	return act
}

// ToJSON converts Execution to ExecutionJSON
func (ex *Execution) ToJSON() (*explorer.Execution, error) {
	execution := &explorer.Execution{
		Version:        int64(ex.version),
		Nonce:          int64(ex.nonce),
		Executor:       ex.srcAddr,
		Contract:       ex.dstAddr,
		ExecutorPubKey: keypair.EncodePublicKey(ex.srcPubkey),
		GasLimit:       int64(ex.gasLimit),
		Data:           hex.EncodeToString(ex.data),
		Signature:      hex.EncodeToString(ex.signature),
	}
	if ex.amount != nil && len(ex.amount.String()) > 0 {
		execution.Amount = ex.amount.String()
	}
	if ex.gasPrice != nil && len(ex.gasPrice.String()) > 0 {
		execution.GasPrice = ex.gasPrice.String()
	}
	return execution, nil
}

// Serialize returns a serialized byte stream for the Execution
func (ex *Execution) Serialize() ([]byte, error) {
	return proto.Marshal(ex.Proto())
}

// LoadProto converts a protobuf's ActionPb to Execution
func (ex *Execution) LoadProto(pbAct *iproto.ActionPb) error {
	if pbAct == nil {
		return errors.New("empty action proto to load")
	}
	if ex == nil {
		return errors.New("nil action to load proto")
	}
	*ex = Execution{}
	srcPub, err := keypair.BytesToPublicKey(pbAct.SenderPubKey)
	if err != nil {
		return err
	}
	pbExecution := pbAct.GetExecution()
	if pbExecution == nil {
		return errors.New("empty Execution action proto to load")
	}

	ab := &Builder{}
	act := ab.SetVersion(pbAct.Version).
		SetNonce(pbAct.Nonce).
		SetSourceAddress(pbAct.Sender).
		SetSourcePublicKey(srcPub).
		SetGasLimit(pbAct.GasLimit).
		SetGasPriceByBytes(pbAct.GasPrice).
		SetDestinationAddress(pbExecution.Contract).
		Build()
	act.SetSignature(pbAct.Signature)
	ex.AbstractAction = act

	ex.data = pbExecution.GetData()
	ex.amount = big.NewInt(0)
	if len(pbExecution.Amount) > 0 {
		ex.amount.SetBytes(pbExecution.GetAmount())
	}
	return nil
}

// NewExecutionFromJSON creates a new Execution from ExecutionJSON
func NewExecutionFromJSON(jsonExecution *explorer.Execution) (*Execution, error) {
	ex := &Execution{}
	ex.version = uint32(jsonExecution.Version)
	ex.nonce = uint64(jsonExecution.Nonce)
	ex.srcAddr = jsonExecution.Executor
	ex.dstAddr = jsonExecution.Contract
	ex.gasLimit = uint64(jsonExecution.GasLimit)
	executorPubKey, err := keypair.StringToPubKeyBytes(jsonExecution.ExecutorPubKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode executor public key")
	}
	copy(ex.srcPubkey[:], executorPubKey)
	amount, ok := big.NewInt(0).SetString(jsonExecution.Amount, 10)
	if !ok {
		return nil, errors.New("failed to set amount of execution")
	}
	ex.amount = amount
	data, err := hex.DecodeString(jsonExecution.Data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode execution data")
	}
	ex.data = data
	gasPrice, ok := big.NewInt(0).SetString(jsonExecution.GasPrice, 10)
	if !ok {
		return nil, errors.New("failed to set gas price of execution")
	}
	ex.gasPrice = gasPrice
	signature, err := hex.DecodeString(jsonExecution.Signature)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode execution signature")
	}
	ex.signature = signature

	return ex, nil
}

// Deserialize parse the byte stream into Execution
func (ex *Execution) Deserialize(buf []byte) error {
	pbAction := &iproto.ActionPb{}
	if err := proto.Unmarshal(buf, pbAction); err != nil {
		return err
	}
	return ex.LoadProto(pbAction)
}

// Hash returns the hash of the Execution
func (ex *Execution) Hash() hash.Hash32B {
	return blake2b.Sum256(ex.ByteStream())
}

// IntrinsicGas returns the intrinsic gas of an execution
func (ex *Execution) IntrinsicGas() (uint64, error) {
	dataSize := uint64(len(ex.Data()))
	if (math.MaxUint64-ExecutionBaseIntrinsicGas)/ExecutionDataGas < dataSize {
		return 0, ErrOutOfGas
	}

	return dataSize*ExecutionDataGas + ExecutionBaseIntrinsicGas, nil
}

// Cost returns the cost of an execution
func (ex *Execution) Cost() (*big.Int, error) {
	maxExecFee := big.NewInt(0).Mul(ex.GasPrice(), big.NewInt(0).SetUint64(ex.GasLimit()))
	return big.NewInt(0).Add(ex.Amount(), maxExecFee), nil
}
