# ext-protobuf-go

Provides a [mus-go](https://github.com/mus-format/mus-go) serializer
extension for the Protobuf format.

This package includes:

- `MarshallerProtobuf` - an interface for types that can marshal themselves
  into the Protobuf format.
- `MarshalProtobuf` - a generic function to marshal values implementing
  `MarshallerProtobuf`.
- `MarshallerTypedProtobuf` — an interface for types that support typed Protobuf
  serialization (designed for use with [DTS](https://github.com/mus-format/dts-go)).
- `MarshallerTypedProtobuf` - a generic function to marshal values implementing
  `MarshallerTypedProtobuf`.
- Serializers for string, slice and timestamp types.
