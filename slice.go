package ext

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	slops "github.com/mus-format/mus-go/options/slice"
	"github.com/mus-format/mus-go/varint"
)

// NewSliceSer returns a new slice serializer with the given element serializer.
func NewSliceSer[T any](elemProtobuf mus.Serializer[T]) sliceSer[T] {
	return sliceSer[T]{elemProtobuf}
}

// NewValidSliceSer returns a new valid slice serializer.
func NewValidSliceSer[T any](elemProtobuf mus.Serializer[T],
	ops ...slops.SetOption[T],
) validSliceSer[T] {
	o := slops.Options[T]{}
	slops.Apply(ops, &o)

	var (
		lenVl  com.Validator[int]
		elemVl com.Validator[T]
	)
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	if o.ElemVl != nil {
		elemVl = o.ElemVl
	}
	return validSliceSer[T]{
		sliceSer: NewSliceSer(elemProtobuf),
		lenVl:    lenVl,
		elemVl:   elemVl,
	}
}

// sliceSer implements the mus.Serializer interface for slices.
type sliceSer[T any] struct {
	elemProtobuf mus.Serializer[T]
}

func (s sliceSer[T]) Marshal(sl []T, bs []byte) (n int) {
	length := len(sl)
	if length > 0 {
		n += varint.PositiveInt.Marshal(s.size(sl), bs[n:])
		for i := range sl {
			n += s.elemProtobuf.Marshal(sl[i], bs[n:])
		}
	}
	return
}

func (s sliceSer[T]) Unmarshal(bs []byte) (sl []T, n int, err error) {
	var (
		n1 int
		e  T
	)
	sl = []T{}
	size, n, err := varint.PositiveInt.Unmarshal(bs)
	if err != nil {
		return
	}
	if len(bs) < size {
		err = com.ErrOverflow
		return
	}
	for n < size {
		e, n1, err = s.elemProtobuf.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		sl = append(sl, e)
	}
	return
}

func (s sliceSer[T]) Size(sl []T) (size int) {
	size = s.size(sl)
	return size + varint.PositiveInt.Size(size)
}

func (s sliceSer[T]) Skip(bs []byte) (n int, err error) {
	l, n, err := varint.PositiveInt.Unmarshal(bs)
	if err != nil {
		return
	}
	n += l
	return
}

func (s sliceSer[T]) size(sl []T) (size int) {
	for i := range len(sl) {
		size += s.elemProtobuf.Size(sl[i])
	}
	return
}

type validSliceSer[T any] struct {
	sliceSer[T]
	lenVl  com.Validator[int]
	elemVl com.Validator[T]
}

func (s validSliceSer[T]) Unmarshal(bs []byte) (sl []T, n int, err error) {
	var (
		n1 int
		e  T
	)
	sl = []T{}
	size, n, err := varint.PositiveInt.Unmarshal(bs)
	if err != nil {
		return
	}
	if len(bs) < size {
		err = com.ErrOverflow
		return
	}
	if s.lenVl != nil {
		if err = s.lenVl.Validate(size); err != nil {
			return
		}
	}
	for n < size {
		e, n1, err = s.elemProtobuf.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if s.elemVl != nil {
			if err = s.elemVl.Validate(e); err != nil {
				return
			}
		}
		sl = append(sl, e)
	}
	return
}
