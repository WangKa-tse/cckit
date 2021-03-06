// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list.proto

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type List struct {
	Items []*google_protobuf.Any `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *List) Reset()                    { *m = List{} }
func (m *List) String() string            { return proto.CompactTextString(m) }
func (*List) ProtoMessage()               {}
func (*List) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *List) GetItems() []*google_protobuf.Any {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*List)(nil), "cckit.state.schema.List")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0x2c, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0x4e, 0xce, 0xce, 0x2c, 0xd1, 0x2b, 0x2e,
	0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x94, 0x92, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0xab, 0x48, 0x2a, 0x4d, 0xd3, 0x4f, 0xcc, 0xab, 0x84, 0x28, 0x57, 0x32,
	0xe2, 0x62, 0xf1, 0xc9, 0x2c, 0x2e, 0x11, 0xd2, 0xe2, 0x62, 0xcd, 0x2c, 0x49, 0xcd, 0x2d, 0x96,
	0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd1, 0x83, 0x68, 0xd1, 0x83, 0x69, 0xd1, 0x73, 0xcc,
	0xab, 0x0c, 0x82, 0x28, 0x71, 0xe2, 0x88, 0x62, 0x83, 0x18, 0x9c, 0xc4, 0x06, 0x96, 0x36, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x3a, 0xd2, 0x60, 0x81, 0x00, 0x00, 0x00,
}
