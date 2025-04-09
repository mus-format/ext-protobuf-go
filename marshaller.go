package ext

// MarshallerProtobuf interface wraps MarhsalProtobuf and SizeProtobuf methods.
type MarshallerProtobuf interface {
	MarshalProtobuf(bs []byte) (n int)
	SizeProtobuf() (size int)
}

// MarshallerProtobuf interface wraps the MarshalProtobuf and SizeProtobuf
// methods. It is intended for use with DTS.
type MarshallerTypedProtobuf interface {
	MarshalTypedProtobuf(bs []byte) (n int)
	SizeTypedProtobuf() (size int)
}
