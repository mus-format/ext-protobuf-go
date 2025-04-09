package ext

// MarshallerProtobuf interface wraps MarhsalProtobuf and SizeProtobuf methods.
type MarshallerProtobuf interface {
	MarshalProtobuf(bs []byte) (n int)
	SizeProtobuf() (size int)
}
