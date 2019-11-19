// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"bytes"
	"context"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
)

/////////////////////////////////////////////////////////////////////////////
// service Indexer

type Indexer interface {
	GetEvents(ctx context.Context, input *GetEventsInput) (*GetEventsOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message GetEventsInput

// reader

type GetEventsInput struct {
	// ClientRequest client.IndexerRequest

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetEventsInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientRequest:%s,}", x.StringClientRequest())
}

var _GetEventsInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
var _GetEventsInput_Unions = [][]membuffers.FieldType{}

func GetEventsInputReader(buf []byte) *GetEventsInput {
	x := &GetEventsInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetEventsInput_Scheme, _GetEventsInput_Unions)
	return x
}

func (x *GetEventsInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetEventsInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetEventsInput) Equal(y *GetEventsInput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetEventsInput) ClientRequest() *client.IndexerRequest {
	b, s := x._message.GetMessage(0)
	return client.IndexerRequestReader(b[:s])
}

func (x *GetEventsInput) RawClientRequest() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetEventsInput) RawClientRequestWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetEventsInput) StringClientRequest() string {
	return x.ClientRequest().String()
}

// builder

type GetEventsInputBuilder struct {
	ClientRequest *client.IndexerRequestBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetEventsInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.ClientRequest)
	if err != nil {
		return
	}
	return nil
}

func (w *GetEventsInputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetEventsInput.ClientRequest", w.ClientRequest)
	if err != nil {
		return
	}
	return nil
}

func (w *GetEventsInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetEventsInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetEventsInputBuilder) Build() *GetEventsInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetEventsInputReader(buf)
}

func GetEventsInputBuilderFromRaw(raw []byte) *GetEventsInputBuilder {
	return &GetEventsInputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message GetEventsOutput

// reader

type GetEventsOutput struct {
	// ClientResponse client.IndexerResponse

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetEventsOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

var _GetEventsOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
var _GetEventsOutput_Unions = [][]membuffers.FieldType{}

func GetEventsOutputReader(buf []byte) *GetEventsOutput {
	x := &GetEventsOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetEventsOutput_Scheme, _GetEventsOutput_Unions)
	return x
}

func (x *GetEventsOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetEventsOutput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetEventsOutput) Equal(y *GetEventsOutput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetEventsOutput) ClientResponse() *client.IndexerResponse {
	b, s := x._message.GetMessage(0)
	return client.IndexerResponseReader(b[:s])
}

func (x *GetEventsOutput) RawClientResponse() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetEventsOutput) RawClientResponseWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetEventsOutput) StringClientResponse() string {
	return x.ClientResponse().String()
}

// builder

type GetEventsOutputBuilder struct {
	ClientResponse *client.IndexerResponseBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetEventsOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.ClientResponse)
	if err != nil {
		return
	}
	return nil
}

func (w *GetEventsOutputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetEventsOutput.ClientResponse", w.ClientResponse)
	if err != nil {
		return
	}
	return nil
}

func (w *GetEventsOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetEventsOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetEventsOutputBuilder) Build() *GetEventsOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetEventsOutputReader(buf)
}

func GetEventsOutputBuilderFromRaw(raw []byte) *GetEventsOutputBuilder {
	return &GetEventsOutputBuilder{_overrideWithRawBuffer: raw}
}
