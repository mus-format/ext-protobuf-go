package ext

import (
	"github.com/mus-format/mus-go"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
)

var (
	LenSer       = lenSer{}
	String       = ord.NewStringSer(strops.WithLenSer(LenSer))
	StringUnsafe = unsafe.NewStringSer(strops.WithLenSer(LenSer))
)

// NewValidStringProtobuf returns a new valid string serializer.
func NewValidStringProtobuf(ops ...strops.SetOption) mus.Serializer[string] {
	ops = append(ops, strops.WithLenSer(LenSer))
	return ord.NewValidStringSer(ops...)
}

// NewValidStringUnsafeProtobuf returns a new valid string serializer.
func NewValidStringUnsafeProtobuf(ops ...strops.SetOption) mus.Serializer[string] {
	ops = append(ops, strops.WithLenSer(LenSer))
	return unsafe.NewValidStringSer(ops...)
}

// lenSer implements the mus.Serializer interface for length.
type lenSer struct{}

func (lenSer) Marshal(v int, bs []byte) (n int) {
	return varint.PositiveInt32.Marshal(int32(v), bs)
}

func (lenSer) Unmarshal(bs []byte) (v int, n int, err error) {
	v32, n, err := varint.PositiveInt32.Unmarshal(bs)
	v = int(v32)
	return
}

func (lenSer) Size(v int) (size int) {
	return varint.PositiveInt32.Size(int32(v))
}

func (lenSer) Skip(bs []byte) (n int, err error) {
	return varint.PositiveInt32.Skip(bs)
}
