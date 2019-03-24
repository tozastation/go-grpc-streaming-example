// Code generated by protoc-gen-go. DO NOT EDIT.
// source: music.proto

package music

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_music_e9a3f117305b58aa, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type StreamingMusic struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingMusic) Reset()         { *m = StreamingMusic{} }
func (m *StreamingMusic) String() string { return proto.CompactTextString(m) }
func (*StreamingMusic) ProtoMessage()    {}
func (*StreamingMusic) Descriptor() ([]byte, []int) {
	return fileDescriptor_music_e9a3f117305b58aa, []int{1}
}
func (m *StreamingMusic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingMusic.Unmarshal(m, b)
}
func (m *StreamingMusic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingMusic.Marshal(b, m, deterministic)
}
func (dst *StreamingMusic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingMusic.Merge(dst, src)
}
func (m *StreamingMusic) XXX_Size() int {
	return xxx_messageInfo_StreamingMusic.Size(m)
}
func (m *StreamingMusic) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingMusic.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingMusic proto.InternalMessageInfo

func (m *StreamingMusic) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "music.Empty")
	proto.RegisterType((*StreamingMusic)(nil), "music.StreamingMusic")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MusicClient is the client API for Music service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MusicClient interface {
	Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Music_StartClient, error)
}

type musicClient struct {
	cc *grpc.ClientConn
}

func NewMusicClient(cc *grpc.ClientConn) MusicClient {
	return &musicClient{cc}
}

func (c *musicClient) Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Music_StartClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Music_serviceDesc.Streams[0], "/music.Music/Start", opts...)
	if err != nil {
		return nil, err
	}
	x := &musicStartClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Music_StartClient interface {
	Recv() (*StreamingMusic, error)
	grpc.ClientStream
}

type musicStartClient struct {
	grpc.ClientStream
}

func (x *musicStartClient) Recv() (*StreamingMusic, error) {
	m := new(StreamingMusic)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MusicServer is the server API for Music service.
type MusicServer interface {
	Start(*Empty, Music_StartServer) error
}

func RegisterMusicServer(s *grpc.Server, srv MusicServer) {
	s.RegisterService(&_Music_serviceDesc, srv)
}

func _Music_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MusicServer).Start(m, &musicStartServer{stream})
}

type Music_StartServer interface {
	Send(*StreamingMusic) error
	grpc.ServerStream
}

type musicStartServer struct {
	grpc.ServerStream
}

func (x *musicStartServer) Send(m *StreamingMusic) error {
	return x.ServerStream.SendMsg(m)
}

var _Music_serviceDesc = grpc.ServiceDesc{
	ServiceName: "music.Music",
	HandlerType: (*MusicServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Start",
			Handler:       _Music_Start_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "music.proto",
}

func init() { proto.RegisterFile("music.proto", fileDescriptor_music_e9a3f117305b58aa) }

var fileDescriptor_music_e9a3f117305b58aa = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2d, 0x2d, 0xce,
	0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xd8, 0xb9, 0x58, 0x5d,
	0x73, 0x0b, 0x4a, 0x2a, 0x95, 0x54, 0xb8, 0xf8, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x33, 0xf3,
	0xd2, 0x7d, 0x41, 0x52, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x91, 0x25, 0x17, 0x2b, 0x44, 0xd2, 0x80, 0x8b, 0x35, 0xb8, 0x24,
	0xb1, 0xa8, 0x44, 0x88, 0x47, 0x0f, 0x62, 0x2a, 0xd8, 0x14, 0x29, 0x51, 0x28, 0x0f, 0xd5, 0x28,
	0x25, 0x06, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0x84, 0xcf, 0x8d, 0x86, 0x00, 0x00, 0x00,
}
