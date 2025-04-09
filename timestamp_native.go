package ext

import (
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var TimestampNativeProtobuf = timestampNativeProtobuf{}

var (
	secondsFieldTag = protowire.EncodeTag(1, protowire.VarintType)
	nanosFieldTag   = protowire.EncodeTag(2, protowire.VarintType)
)

// timestampNativeProtobuf implements the mus.Serializer interface for
// timestamppb.Timestamp.
type timestampNativeProtobuf struct{}

func (s timestampNativeProtobuf) Marshal(tm *timestamppb.Timestamp, bs []byte) (n int) {
	size := s.size(tm)
	if size > 0 {
		n += varint.PositiveInt.Marshal(size, bs[n:])
		if tm.Seconds != 0 {
			n += varint.Uint64.Marshal(secondsFieldTag, bs[n:])
			n += varint.PositiveInt64.Marshal(tm.Seconds, bs[n:])
		}
		if tm.Nanos != 0 {
			n += varint.Uint64.Marshal(nanosFieldTag, bs[n:])
			n += varint.PositiveInt32.Marshal(tm.Nanos, bs[n:])
		}
	}
	return
}

func (timestampNativeProtobuf) Unmarshal(bs []byte) (tm *timestamppb.Timestamp, n int,
	err error) {
	size, n, err := varint.PositiveInt.Unmarshal(bs)
	if err != nil {
		return
	}
	if len(bs) < size {
		err = mus.ErrTooSmallByteSlice
		return
	}
	var (
		n1  int
		tag uint64
	)
	tm = &timestamppb.Timestamp{}
	for n < size {
		tag, n1, err = varint.Uint64.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case secondsFieldTag:
			tm.Seconds, n1, err = varint.PositiveInt64.Unmarshal(bs[n:])
		case nanosFieldTag:
			tm.Nanos, n1, err = varint.PositiveInt32.Unmarshal(bs[n:])
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func (s timestampNativeProtobuf) Size(tm *timestamppb.Timestamp) (size int) {
	size = s.size(tm)
	return size + varint.PositiveInt.Size(size)
}

func (timestampNativeProtobuf) Skip(bs []byte) (n int, err error) {
	size, n, err := varint.PositiveInt.Unmarshal(bs)
	if err != nil {
		return
	}
	if len(bs) < size {
		err = mus.ErrTooSmallByteSlice
		return
	}
	var (
		n1  int
		tag uint64
	)
	for n < size {
		tag, n1, err = varint.Uint64.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case secondsFieldTag:
			n1, err = varint.PositiveInt64.Skip(bs[n:])
		case nanosFieldTag:
			n1, err = varint.PositiveInt32.Skip(bs[n:])
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func (s timestampNativeProtobuf) size(tm *timestamppb.Timestamp) (size int) {
	if tm.Seconds != 0 {
		size += varint.Uint64.Size(secondsFieldTag)
		size += varint.PositiveInt64.Size(tm.Seconds)
	}
	if tm.Nanos != 0 {
		size += varint.Uint64.Size(nanosFieldTag)
		size += varint.PositiveInt32.Size(tm.Nanos)
	}
	return
}
