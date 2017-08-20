// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: empty/empty.proto

/*
	Package empty is a generated protocol buffer package.

	It is generated from these files:
		empty/empty.proto

	It has these top-level messages:
		Empty
*/
package empty

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
}

// MarshalToWriter marshals Empty to the provided writer.
func (m *Empty) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	return
}

// Marshal marshals Empty to a slice of bytes.
func (m *Empty) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Empty from the provided reader.
func (m *Empty) UnmarshalFromReader(reader jspb.Reader) *Empty {
	for reader.Next() {
		if m == nil {
			m = &Empty{}
		}

		switch reader.GetFieldNumber() {
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Empty from a slice of bytes.
func (m *Empty) Unmarshal(rawBytes []byte) (*Empty, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
