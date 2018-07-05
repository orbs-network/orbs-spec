// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Processor

type Processor interface {
	ProcessCall(*ProcessCallInput) (*ProcessCallOutput, error)
	DeployNativeService(*DeployNativeServiceInput) (*DeployNativeServiceOutput, error)
	RegisterContractSdkCallHandler(handlers.ContractSdkCallHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallInput

// reader

type ProcessCallInput struct {
	message membuffers.Message
}

var m_ProcessCallInput_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessageArray,membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_ProcessCallInput_Unions = [][]membuffers.FieldType{}

func ProcessCallInputReader(buf []byte) *ProcessCallInput {
	x := &ProcessCallInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ProcessCallInput_Scheme, m_ProcessCallInput_Unions)
	return x
}

func (x *ProcessCallInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ProcessCallInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ProcessCallInput) ContextId() uint32 {
	return x.message.GetUint32(0)
}

func (x *ProcessCallInput) RawContextId() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ProcessCallInput) MutateContextId(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *ProcessCallInput) Contract() *protocol.ContractAddress {
	b, s := x.message.GetMessage(1)
	return protocol.ContractAddressReader(b[:s])
}

func (x *ProcessCallInput) RawContract() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ProcessCallInput) Method() *protocol.MethodAddress {
	b, s := x.message.GetMessage(2)
	return protocol.MethodAddressReader(b[:s])
}

func (x *ProcessCallInput) RawMethod() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *ProcessCallInput) InputArgumentIterator() *ProcessCallInputInputArgumentIterator {
	return &ProcessCallInputInputArgumentIterator{iterator: x.message.GetMessageArrayIterator(3)}
}

type ProcessCallInputInputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *ProcessCallInputInputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ProcessCallInputInputArgumentIterator) NextInputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *ProcessCallInput) RawInputArgumentArray() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *ProcessCallInput) AccessScope() AccessScope {
	return AccessScope(x.message.GetUint16(4))
}

func (x *ProcessCallInput) RawAccessScope() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *ProcessCallInput) MutateAccessScope(v AccessScope) error {
	return x.message.SetUint16(4, uint16(v))
}

func (x *ProcessCallInput) PermissionScope() PermissionScope {
	return PermissionScope(x.message.GetUint16(5))
}

func (x *ProcessCallInput) RawPermissionScope() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *ProcessCallInput) MutatePermissionScope(v PermissionScope) error {
	return x.message.SetUint16(5, uint16(v))
}

func (x *ProcessCallInput) CallingService() *protocol.ContractAddress {
	b, s := x.message.GetMessage(6)
	return protocol.ContractAddressReader(b[:s])
}

func (x *ProcessCallInput) RawCallingService() []byte {
	return x.message.RawBufferForField(6, 0)
}

func (x *ProcessCallInput) TransactionSender() *protocol.Sender {
	b, s := x.message.GetMessage(7)
	return protocol.SenderReader(b[:s])
}

func (x *ProcessCallInput) RawTransactionSender() []byte {
	return x.message.RawBufferForField(7, 0)
}

// builder

type ProcessCallInputBuilder struct {
	builder membuffers.Builder
	ContextId uint32
	Contract *protocol.ContractAddressBuilder
	Method *protocol.MethodAddressBuilder
	InputArgument []*protocol.MethodArgumentBuilder
	AccessScope AccessScope
	PermissionScope PermissionScope
	CallingService *protocol.ContractAddressBuilder
	TransactionSender *protocol.SenderBuilder
}

func (w *ProcessCallInputBuilder) arrayOfInputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.InputArgument))
	for i, v := range w.InputArgument {
		res[i] = v
	}
	return res
}

func (w *ProcessCallInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ContextId)
	err = w.builder.WriteMessage(buf, w.Contract)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.Method)
	if err != nil {
		return
	}
	err = w.builder.WriteMessageArray(buf, w.arrayOfInputArgument())
	if err != nil {
		return
	}
	w.builder.WriteUint16(buf, uint16(w.AccessScope))
	w.builder.WriteUint16(buf, uint16(w.PermissionScope))
	err = w.builder.WriteMessage(buf, w.CallingService)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.TransactionSender)
	if err != nil {
		return
	}
	return nil
}

func (w *ProcessCallInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ProcessCallInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ProcessCallInputBuilder) Build() *ProcessCallInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ProcessCallInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallOutput

// reader

type ProcessCallOutput struct {
	message membuffers.Message
}

var m_ProcessCallOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeUint16,}
var m_ProcessCallOutput_Unions = [][]membuffers.FieldType{}

func ProcessCallOutputReader(buf []byte) *ProcessCallOutput {
	x := &ProcessCallOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ProcessCallOutput_Scheme, m_ProcessCallOutput_Unions)
	return x
}

func (x *ProcessCallOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ProcessCallOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ProcessCallOutput) OutputArgumentIterator() *ProcessCallOutputOutputArgumentIterator {
	return &ProcessCallOutputOutputArgumentIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type ProcessCallOutputOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *ProcessCallOutputOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ProcessCallOutputOutputArgumentIterator) NextOutputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *ProcessCallOutput) RawOutputArgumentArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ProcessCallOutput) CallStatus() protocol.CallMethodStatus {
	return protocol.CallMethodStatus(x.message.GetUint16(1))
}

func (x *ProcessCallOutput) RawCallStatus() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ProcessCallOutput) MutateCallStatus(v protocol.CallMethodStatus) error {
	return x.message.SetUint16(1, uint16(v))
}

// builder

type ProcessCallOutputBuilder struct {
	builder membuffers.Builder
	OutputArgument []*protocol.MethodArgumentBuilder
	CallStatus protocol.CallMethodStatus
}

func (w *ProcessCallOutputBuilder) arrayOfOutputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
		res[i] = v
	}
	return res
}

func (w *ProcessCallOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessageArray(buf, w.arrayOfOutputArgument())
	if err != nil {
		return
	}
	w.builder.WriteUint16(buf, uint16(w.CallStatus))
	return nil
}

func (w *ProcessCallOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ProcessCallOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ProcessCallOutputBuilder) Build() *ProcessCallOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ProcessCallOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message DeployNativeServiceInput

// reader

type DeployNativeServiceInput struct {
	message membuffers.Message
}

var m_DeployNativeServiceInput_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_DeployNativeServiceInput_Unions = [][]membuffers.FieldType{}

func DeployNativeServiceInputReader(buf []byte) *DeployNativeServiceInput {
	x := &DeployNativeServiceInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_DeployNativeServiceInput_Scheme, m_DeployNativeServiceInput_Unions)
	return x
}

func (x *DeployNativeServiceInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *DeployNativeServiceInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *DeployNativeServiceInput) ContextId() uint32 {
	return x.message.GetUint32(0)
}

func (x *DeployNativeServiceInput) RawContextId() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *DeployNativeServiceInput) MutateContextId(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *DeployNativeServiceInput) Contract() *protocol.ContractAddress {
	b, s := x.message.GetMessage(1)
	return protocol.ContractAddressReader(b[:s])
}

func (x *DeployNativeServiceInput) RawContract() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *DeployNativeServiceInput) AccessScope() AccessScope {
	return AccessScope(x.message.GetUint16(2))
}

func (x *DeployNativeServiceInput) RawAccessScope() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *DeployNativeServiceInput) MutateAccessScope(v AccessScope) error {
	return x.message.SetUint16(2, uint16(v))
}

func (x *DeployNativeServiceInput) PermissionScope() PermissionScope {
	return PermissionScope(x.message.GetUint16(3))
}

func (x *DeployNativeServiceInput) RawPermissionScope() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *DeployNativeServiceInput) MutatePermissionScope(v PermissionScope) error {
	return x.message.SetUint16(3, uint16(v))
}

func (x *DeployNativeServiceInput) CallingService() *protocol.ContractAddress {
	b, s := x.message.GetMessage(4)
	return protocol.ContractAddressReader(b[:s])
}

func (x *DeployNativeServiceInput) RawCallingService() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *DeployNativeServiceInput) TransactionSender() *protocol.Sender {
	b, s := x.message.GetMessage(5)
	return protocol.SenderReader(b[:s])
}

func (x *DeployNativeServiceInput) RawTransactionSender() []byte {
	return x.message.RawBufferForField(5, 0)
}

// builder

type DeployNativeServiceInputBuilder struct {
	builder membuffers.Builder
	ContextId uint32
	Contract *protocol.ContractAddressBuilder
	AccessScope AccessScope
	PermissionScope PermissionScope
	CallingService *protocol.ContractAddressBuilder
	TransactionSender *protocol.SenderBuilder
}

func (w *DeployNativeServiceInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ContextId)
	err = w.builder.WriteMessage(buf, w.Contract)
	if err != nil {
		return
	}
	w.builder.WriteUint16(buf, uint16(w.AccessScope))
	w.builder.WriteUint16(buf, uint16(w.PermissionScope))
	err = w.builder.WriteMessage(buf, w.CallingService)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.TransactionSender)
	if err != nil {
		return
	}
	return nil
}

func (w *DeployNativeServiceInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *DeployNativeServiceInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *DeployNativeServiceInputBuilder) Build() *DeployNativeServiceInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return DeployNativeServiceInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message DeployNativeServiceOutput

// reader

type DeployNativeServiceOutput struct {
	message membuffers.Message
}

var m_DeployNativeServiceOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_DeployNativeServiceOutput_Unions = [][]membuffers.FieldType{}

func DeployNativeServiceOutputReader(buf []byte) *DeployNativeServiceOutput {
	x := &DeployNativeServiceOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_DeployNativeServiceOutput_Scheme, m_DeployNativeServiceOutput_Unions)
	return x
}

func (x *DeployNativeServiceOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *DeployNativeServiceOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *DeployNativeServiceOutput) Status() DeployServiceStatus {
	return DeployServiceStatus(x.message.GetUint16(0))
}

func (x *DeployNativeServiceOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *DeployNativeServiceOutput) MutateStatus(v DeployServiceStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type DeployNativeServiceOutputBuilder struct {
	builder membuffers.Builder
	Status DeployServiceStatus
}

func (w *DeployNativeServiceOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.Status))
	return nil
}

func (w *DeployNativeServiceOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *DeployNativeServiceOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *DeployNativeServiceOutputBuilder) Build() *DeployNativeServiceOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return DeployNativeServiceOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type AccessScope uint16

const (
	ACCESS_SCOPE_RESERVED AccessScope = 0
	ACCESS_SCOPE_READ_ONLY AccessScope = 1
	ACCESS_SCOPE_READ_WRITE AccessScope = 2
)

type PermissionScope uint16

const (
	PERMISSION_SCOPE_RESERVED PermissionScope = 0
	PERMISSION_SCOPE_SYSTEM PermissionScope = 1
	PERMISSION_SCOPE_SERVICE PermissionScope = 2
)

type DeployServiceStatus uint16

const (
	DEPLOY_SERVICE_STATUS_RESERVED DeployServiceStatus = 0
	DEPLOY_SERVICE_STATUS_SUCCESS DeployServiceStatus = 1
	DEPLOY_SERVICE_STATUS_FAILED DeployServiceStatus = 2
)

