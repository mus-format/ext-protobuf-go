package mock

import (
	"github.com/ymz-ncnk/mok"
)

type MarshalTypedProtobufFn[T any] func(bs []byte) (n T, err error)
type SizeTypedProtobufFn func() (size int)

func NewMarshallerTypedProtobuf[T any]() MarshallerTypedProtobuf[T] {
	return MarshallerTypedProtobuf[T]{mok.New("MarshallerTypedProtobuf")}
}

type MarshallerTypedProtobuf[T any] struct {
	*mok.Mock
}

func (m MarshallerTypedProtobuf[T]) RegisterMarshalTypedProtobuf(
	fn MarshalTypedProtobufFn[T]) MarshallerTypedProtobuf[T] {
	m.Register("MarshalTypedProtobuf", fn)
	return m
}

func (m MarshallerTypedProtobuf[T]) RegisterSizeTypedProtobuf(
	fn SizeTypedProtobufFn) MarshallerTypedProtobuf[T] {
	m.Register("SizeTypedProtobuf", fn)
	return m
}

func (m MarshallerTypedProtobuf[T]) MarshalTypedProtobuf(bs []byte) (n int,
	err error) {
	result, err := m.Call("MarshalTypedProtobuf", bs)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m MarshallerTypedProtobuf[T]) SizeTypedProtobuf() (size int) {
	result, err := m.Call("SizeTypedProtobuf")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
