# ext-protobuf-go
Provides a [mus-go](https://github.com/mus-format/mus-go) serializer extension
for the Protobuf format.

Includes the `MarshallerProtobuf` interface, which represents a type that can 
marshal itself into the Protobuf format, along with the generic 
`MarshalProtobuf` function. Also includes the `MarshallerTypedProtobuf` 
interface and the `MarshalTypedProtobuf` function, intended for use with 
[DTS](https://github.com/mus-format/mus-dts-go).

Contains serializers for string, slice and timestamp types.