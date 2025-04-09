package ext

// MarshallerProtobuf interface wraps the MarshalProtobuf and SizeProtobuf
// methods. It is intended for use with DTS.
type MarshallerTypedProtobuf interface {
	MarshalTypedProtobuf(bs []byte) (n int)
	SizeTypedProtobuf() (size int)
}
