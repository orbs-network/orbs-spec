package protocol

import (
	"fmt"
	"math/big"
)

func ArgumentBuilderFromNative(arg interface{}) (res *ArgumentBuilder, err error) {
	switch arg.(type) {
	case uint32:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_UINT_32_VALUE, Uint32Value: arg.(uint32)}
	case uint64:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_UINT_64_VALUE, Uint64Value: arg.(uint64)}
	case string:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_STRING_VALUE, StringValue: arg.(string)}
	case []byte:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_BYTES_VALUE, BytesValue: arg.([]byte)}
	case bool:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_BOOL_VALUE, BoolValue: arg.(bool)}
	case *big.Int:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_UINT_256_VALUE, Uint256Value: arg.(*big.Int)}
	case [20]byte:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_BYTES_20_VALUE, Bytes20Value: arg.([20]byte)}
	case [32]byte:
		res = &ArgumentBuilder{Type: ARGUMENT_TYPE_BYTES_32_VALUE, Bytes32Value: arg.([32]byte)}
	default:
		err = fmt.Errorf("argument has unsupported type (%T), supported: (uint32) (uint64) (string) ([]byte) (bool) (uint256) ([20]byte) ([32]byte)", arg)
	}
	return
}

func ArgumentBuildersFromNatives(args []interface{}) ([]*ArgumentBuilder, error) {
	res := make([]*ArgumentBuilder, 0, len(args))
	for index, arg := range args {
		resArg, err := ArgumentBuilderFromNative(arg)
		if err != nil {
			return nil, fmt.Errorf("given argument %d: %s", index, err.Error())
		}
		res = append(res, resArg)
	}
	return res, nil
}

func ArgumentsFromNatives(args []interface{}) ([]*Argument, error) {
	res := make([]*Argument, 0, len(args))
	for index, arg := range args {
		resArg, err := ArgumentBuilderFromNative(arg)
		if err != nil {
			return nil, fmt.Errorf("given argument %d: %s", index, err.Error())
		}
		res = append(res, resArg.Build())
	}
	return res, nil
}

func ArgumentArrayFromNatives(args []interface{}) (*ArgumentArray, error) {
	builders, err := ArgumentBuildersFromNatives(args)
	if err != nil {
		return nil, err
	}
	return (&ArgumentArrayBuilder{Arguments: builders}).Build(), nil
}

// the func input is array of go natives (for contract) and encoded to ArgumentArray packed bytes *without* header (length)
func PackedInputArgumentsFromNatives(args []interface{}) ([]byte, error) {
	argArray, err := ArgumentArrayFromNatives(args)
	if err != nil {
		return nil, err
	}
	return argArray.RawArgumentsArray(), nil
}

// the func input is ArgumentArray output (of contract/event) as packed bytes *with* header (length) that is decoded to array of go natives
func PackedOutputArgumentsToNatives(buf []byte) (res []interface{}, err error) {
	return ArgumentArrayReader(buf).ToNatives()
}

func (x *ArgumentArray) ToNatives() (res []interface{}, err error) {
	res = []interface{}{}
	index := 0
	for i := x.ArgumentsIterator(); i.HasNext(); {
		argument := i.NextArguments()
		switch argument.Type() {
		case ARGUMENT_TYPE_UINT_32_VALUE:
			res = append(res, argument.Uint32Value())
		case ARGUMENT_TYPE_UINT_64_VALUE:
			res = append(res, argument.Uint64Value())
		case ARGUMENT_TYPE_STRING_VALUE:
			res = append(res, argument.StringValue())
		case ARGUMENT_TYPE_BYTES_VALUE:
			res = append(res, argument.BytesValue())
		case ARGUMENT_TYPE_BOOL_VALUE:
			res = append(res, argument.BoolValue())
		case ARGUMENT_TYPE_UINT_256_VALUE:
			res = append(res, argument.Uint256Value())
		case ARGUMENT_TYPE_BYTES_20_VALUE:
			res = append(res, argument.Bytes20Value())
		case ARGUMENT_TYPE_BYTES_32_VALUE:
			res = append(res, argument.Bytes32Value())
		default:
			err = fmt.Errorf("argument %d has unsupported type: %s", index, argument.StringType())
			return
		}
		index++
	}
	return
}
