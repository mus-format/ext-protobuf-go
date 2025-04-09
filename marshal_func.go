package ext

// MarshalProtobuf creates and returns a byte slice filled with the serialized
// data.
func MarshalProtobuf(v MarshallerProtobuf) (bs []byte) {
	bs = make([]byte, v.SizeProtobuf())
	v.MarshalProtobuf(bs)
	return
}

// MarshalTypedProtobuf creates and returns a byte slice filled with the
// serialized data.
func MarshalTypedProtobuf(v MarshallerTypedProtobuf) (bs []byte) {
	bs = make([]byte, v.SizeTypedProtobuf())
	v.MarshalTypedProtobuf(bs)
	return
}
