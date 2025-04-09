package ext

// MarshalProtobuf creates and returns a byte slice filled with the serialized
// data.
func MarshalProtobuf(v MarshallerProtobuf) (bs []byte) {
	bs = make([]byte, v.SizeProtobuf())
	v.MarshalProtobuf(bs)
	return
}
