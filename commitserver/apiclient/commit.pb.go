// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commitserver/commit/commit.proto

package apiclient

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ManifestsRequest struct {
	// Repo contains repository information including, at minimum, the URL of the repository. Generally it will contain
	// repo credentials.
	Repo *v1alpha1.Repository `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	// SyncBranch is the branch Argo CD syncs from, i.e. the hydrated branch.
	SyncBranch string `protobuf:"bytes,2,opt,name=sync_branch,json=syncBranch,proto3" json:"sync_branch,omitempty"`
	// TargetBranch is the branch Argo CD is committing to, i.e. the branch that will be updated.
	TargetBranch string `protobuf:"bytes,3,opt,name=target_branch,json=targetBranch,proto3" json:"target_branch,omitempty"`
	// DrySha is the commit SHA from the dry branch, i.e. pre-rendered manifest branch.
	DrySha string `protobuf:"bytes,4,opt,name=dry_sha,json=drySha,proto3" json:"dry_sha,omitempty"`
	// CommitMessage is the commit message to use when committing changes.
	CommitMessage string `protobuf:"bytes,5,opt,name=commit_message,json=commitMessage,proto3" json:"commit_message,omitempty"`
	// Paths contains the paths to write hydrated manifests to, along with the manifests and commands to execute.
	Paths                []*PathDetails `protobuf:"bytes,6,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ManifestsRequest) Reset()         { *m = ManifestsRequest{} }
func (m *ManifestsRequest) String() string { return proto.CompactTextString(m) }
func (*ManifestsRequest) ProtoMessage()    {}
func (*ManifestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{0}
}
func (m *ManifestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManifestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestsRequest.Merge(m, src)
}
func (m *ManifestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ManifestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestsRequest proto.InternalMessageInfo

func (m *ManifestsRequest) GetRepo() *v1alpha1.Repository {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *ManifestsRequest) GetSyncBranch() string {
	if m != nil {
		return m.SyncBranch
	}
	return ""
}

func (m *ManifestsRequest) GetTargetBranch() string {
	if m != nil {
		return m.TargetBranch
	}
	return ""
}

func (m *ManifestsRequest) GetDrySha() string {
	if m != nil {
		return m.DrySha
	}
	return ""
}

func (m *ManifestsRequest) GetCommitMessage() string {
	if m != nil {
		return m.CommitMessage
	}
	return ""
}

func (m *ManifestsRequest) GetPaths() []*PathDetails {
	if m != nil {
		return m.Paths
	}
	return nil
}

type PathDetails struct {
	// Path is the path to write the hydrated manifests to.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Manifests contains the manifests to write to the path.
	Manifests []*ManifestDetails `protobuf:"bytes,2,rep,name=manifests,proto3" json:"manifests,omitempty"`
	// Commands contains the commands executed when hydrating the manifests.
	Commands []string `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty"`
	// ReadmeDetails contains the readme info to write to the path.
	ReadmeDetails        *ReadmeDetails `protobuf:"bytes,4,opt,name=readme_details,json=readmeDetails,proto3" json:"readme_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PathDetails) Reset()         { *m = PathDetails{} }
func (m *PathDetails) String() string { return proto.CompactTextString(m) }
func (*PathDetails) ProtoMessage()    {}
func (*PathDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{1}
}
func (m *PathDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PathDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PathDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PathDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathDetails.Merge(m, src)
}
func (m *PathDetails) XXX_Size() int {
	return m.Size()
}
func (m *PathDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_PathDetails.DiscardUnknown(m)
}

var xxx_messageInfo_PathDetails proto.InternalMessageInfo

func (m *PathDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathDetails) GetManifests() []*ManifestDetails {
	if m != nil {
		return m.Manifests
	}
	return nil
}

func (m *PathDetails) GetCommands() []string {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *PathDetails) GetReadmeDetails() *ReadmeDetails {
	if m != nil {
		return m.ReadmeDetails
	}
	return nil
}

// ManifestDetails contains the manifest hydrated manifests.
type ManifestDetails struct {
	Manifest             string   `protobuf:"bytes,1,opt,name=manifest,proto3" json:"manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestDetails) Reset()         { *m = ManifestDetails{} }
func (m *ManifestDetails) String() string { return proto.CompactTextString(m) }
func (*ManifestDetails) ProtoMessage()    {}
func (*ManifestDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{2}
}
func (m *ManifestDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManifestDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDetails.Merge(m, src)
}
func (m *ManifestDetails) XXX_Size() int {
	return m.Size()
}
func (m *ManifestDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDetails proto.InternalMessageInfo

func (m *ManifestDetails) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

// ReadmeDetails contains the readme info to write to the path.
type ReadmeDetails struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadmeDetails) Reset()         { *m = ReadmeDetails{} }
func (m *ReadmeDetails) String() string { return proto.CompactTextString(m) }
func (*ReadmeDetails) ProtoMessage()    {}
func (*ReadmeDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{3}
}
func (m *ReadmeDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadmeDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReadmeDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReadmeDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadmeDetails.Merge(m, src)
}
func (m *ReadmeDetails) XXX_Size() int {
	return m.Size()
}
func (m *ReadmeDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadmeDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ReadmeDetails proto.InternalMessageInfo

// ManifestsResponse is the response to the ManifestsRequest.
type ManifestsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestsResponse) Reset()         { *m = ManifestsResponse{} }
func (m *ManifestsResponse) String() string { return proto.CompactTextString(m) }
func (*ManifestsResponse) ProtoMessage()    {}
func (*ManifestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{4}
}
func (m *ManifestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManifestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestsResponse.Merge(m, src)
}
func (m *ManifestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ManifestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ManifestsRequest)(nil), "ManifestsRequest")
	proto.RegisterType((*PathDetails)(nil), "PathDetails")
	proto.RegisterType((*ManifestDetails)(nil), "ManifestDetails")
	proto.RegisterType((*ReadmeDetails)(nil), "ReadmeDetails")
	proto.RegisterType((*ManifestsResponse)(nil), "ManifestsResponse")
}

func init() { proto.RegisterFile("commitserver/commit/commit.proto", fileDescriptor_cf3a3abbc35e3069) }

var fileDescriptor_cf3a3abbc35e3069 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x34, 0x90, 0x49, 0x93, 0xb6, 0xcb, 0x01, 0x2b, 0x87, 0x60, 0x19, 0x21, 0xe5,
	0xd2, 0xb5, 0x6a, 0xd4, 0x3b, 0x6a, 0x39, 0x70, 0xa0, 0x12, 0x72, 0x6f, 0x08, 0xc9, 0x9a, 0xd8,
	0x83, 0xbd, 0x34, 0xf6, 0x2e, 0xbb, 0x5b, 0x4b, 0xf9, 0x15, 0xce, 0x7c, 0x0c, 0x47, 0x3e, 0x01,
	0xe5, 0x4b, 0x90, 0xd7, 0x71, 0x9b, 0x96, 0x03, 0x27, 0xcf, 0xbc, 0x19, 0xbf, 0x37, 0xfb, 0xf4,
	0x20, 0xc8, 0x64, 0x55, 0x09, 0x6b, 0x48, 0x37, 0xa4, 0xa3, 0xae, 0xd9, 0x7d, 0xb8, 0xd2, 0xd2,
	0xca, 0xf9, 0xc7, 0x42, 0xd8, 0xf2, 0x6e, 0xc5, 0x33, 0x59, 0x45, 0xa8, 0x0b, 0xa9, 0xb4, 0xfc,
	0xe6, 0x8a, 0xb3, 0x2c, 0x8f, 0x9a, 0x38, 0x52, 0xb7, 0x45, 0x84, 0x4a, 0x98, 0x08, 0x95, 0x5a,
	0x8b, 0x0c, 0xad, 0x90, 0x75, 0xd4, 0x9c, 0xe3, 0x5a, 0x95, 0x78, 0x1e, 0x15, 0x54, 0x93, 0x46,
	0x4b, 0x79, 0xc7, 0x16, 0xfe, 0x38, 0x80, 0x93, 0x6b, 0xac, 0xc5, 0x57, 0x32, 0xd6, 0x24, 0xf4,
	0xfd, 0x8e, 0x8c, 0x65, 0x5f, 0x60, 0xa8, 0x49, 0x49, 0xdf, 0x0b, 0xbc, 0xe5, 0x24, 0xfe, 0xc0,
	0x1f, 0x14, 0x79, 0xaf, 0xe8, 0x8a, 0x34, 0xcb, 0x79, 0x13, 0x73, 0x75, 0x5b, 0xf0, 0x56, 0x91,
	0xef, 0x29, 0xf2, 0x5e, 0x91, 0x27, 0xa4, 0xa4, 0x11, 0x56, 0xea, 0x4d, 0xe2, 0x58, 0xd9, 0x2b,
	0x98, 0x98, 0x4d, 0x9d, 0xa5, 0x2b, 0x8d, 0x75, 0x56, 0xfa, 0x07, 0x81, 0xb7, 0x1c, 0x27, 0xd0,
	0x42, 0x97, 0x0e, 0x61, 0xaf, 0x61, 0x6a, 0x51, 0x17, 0x64, 0xfb, 0x95, 0x81, 0x5b, 0x39, 0xea,
	0xc0, 0xdd, 0xd2, 0x4b, 0x78, 0x96, 0xeb, 0x4d, 0x6a, 0x4a, 0xf4, 0x87, 0x6e, 0x3c, 0xca, 0xf5,
	0xe6, 0xa6, 0x44, 0xf6, 0x06, 0x66, 0x9d, 0x5f, 0x69, 0x45, 0xc6, 0x60, 0x41, 0xfe, 0xa1, 0x9b,
	0x4f, 0x3b, 0xf4, 0xba, 0x03, 0x59, 0x08, 0x87, 0x0a, 0x6d, 0x69, 0xfc, 0x51, 0x30, 0x58, 0x4e,
	0xe2, 0x23, 0xfe, 0x09, 0x6d, 0xf9, 0x9e, 0x2c, 0x8a, 0xb5, 0x49, 0xba, 0x51, 0xf8, 0xd3, 0x83,
	0xc9, 0x1e, 0xcc, 0x18, 0x0c, 0xdb, 0x81, 0xf3, 0x65, 0x9c, 0xb8, 0x9a, 0x71, 0x18, 0x57, 0xbd,
	0x7f, 0xfe, 0x81, 0xe3, 0x3a, 0xe1, 0xbd, 0xa3, 0x3d, 0xdf, 0xc3, 0x0a, 0x9b, 0xc3, 0xf3, 0xf6,
	0x10, 0xac, 0x73, 0xe3, 0x0f, 0x82, 0xc1, 0x72, 0x9c, 0xdc, 0xf7, 0xec, 0x02, 0x66, 0x9a, 0x30,
	0xaf, 0x28, 0xcd, 0xbb, 0x1f, 0xdd, 0xd3, 0x26, 0xf1, 0x8c, 0x27, 0x0e, 0xee, 0xe9, 0xa6, 0x7a,
	0xbf, 0x0d, 0xcf, 0xe0, 0xf8, 0x89, 0x60, 0xab, 0xd2, 0x4b, 0xee, 0xae, 0xbd, 0xef, 0xc3, 0x63,
	0x98, 0x3e, 0xa2, 0x0b, 0x5f, 0xc0, 0xe9, 0x5e, 0x04, 0x8c, 0x92, 0xb5, 0xa1, 0xf8, 0x1d, 0x4c,
	0xaf, 0x9c, 0x61, 0x37, 0xa4, 0x1b, 0x91, 0x11, 0x8b, 0x60, 0xd4, 0x01, 0xec, 0x94, 0x3f, 0x4d,
	0xcc, 0x9c, 0xf1, 0x7f, 0x18, 0x2e, 0xaf, 0x7e, 0x6d, 0x17, 0xde, 0xef, 0xed, 0xc2, 0xfb, 0xb3,
	0x5d, 0x78, 0x9f, 0x2f, 0xfe, 0x13, 0xdb, 0x47, 0xb9, 0x47, 0x25, 0xb2, 0xb5, 0xa0, 0xda, 0xae,
	0x46, 0x2e, 0xa6, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x88, 0x24, 0x04, 0x4b, 0x18, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommitServiceClient is the client API for CommitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommitServiceClient interface {
	// Commit commits hydrated manifests to a repository.
	Commit(ctx context.Context, in *ManifestsRequest, opts ...grpc.CallOption) (*ManifestsResponse, error)
}

type commitServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommitServiceClient(cc *grpc.ClientConn) CommitServiceClient {
	return &commitServiceClient{cc}
}

func (c *commitServiceClient) Commit(ctx context.Context, in *ManifestsRequest, opts ...grpc.CallOption) (*ManifestsResponse, error) {
	out := new(ManifestsResponse)
	err := c.cc.Invoke(ctx, "/CommitService/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommitServiceServer is the server API for CommitService service.
type CommitServiceServer interface {
	// Commit commits hydrated manifests to a repository.
	Commit(context.Context, *ManifestsRequest) (*ManifestsResponse, error)
}

// UnimplementedCommitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommitServiceServer struct {
}

func (*UnimplementedCommitServiceServer) Commit(ctx context.Context, req *ManifestsRequest) (*ManifestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}

func RegisterCommitServiceServer(s *grpc.Server, srv CommitServiceServer) {
	s.RegisterService(&_CommitService_serviceDesc, srv)
}

func _CommitService_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitServiceServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommitService/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitServiceServer).Commit(ctx, req.(*ManifestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommitService",
	HandlerType: (*CommitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Commit",
			Handler:    _CommitService_Commit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commitserver/commit/commit.proto",
}

func (m *ManifestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManifestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Paths) > 0 {
		for iNdEx := len(m.Paths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Paths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.CommitMessage) > 0 {
		i -= len(m.CommitMessage)
		copy(dAtA[i:], m.CommitMessage)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.CommitMessage)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DrySha) > 0 {
		i -= len(m.DrySha)
		copy(dAtA[i:], m.DrySha)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.DrySha)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TargetBranch) > 0 {
		i -= len(m.TargetBranch)
		copy(dAtA[i:], m.TargetBranch)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.TargetBranch)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SyncBranch) > 0 {
		i -= len(m.SyncBranch)
		copy(dAtA[i:], m.SyncBranch)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.SyncBranch)))
		i--
		dAtA[i] = 0x12
	}
	if m.Repo != nil {
		{
			size, err := m.Repo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PathDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PathDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PathDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ReadmeDetails != nil {
		{
			size, err := m.ReadmeDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Commands) > 0 {
		for iNdEx := len(m.Commands) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Commands[iNdEx])
			copy(dAtA[i:], m.Commands[iNdEx])
			i = encodeVarintCommit(dAtA, i, uint64(len(m.Commands[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Manifests) > 0 {
		for iNdEx := len(m.Manifests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Manifests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ManifestDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManifestDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Manifest) > 0 {
		i -= len(m.Manifest)
		copy(dAtA[i:], m.Manifest)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.Manifest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReadmeDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadmeDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReadmeDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *ManifestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManifestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommit(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ManifestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Repo != nil {
		l = m.Repo.Size()
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.SyncBranch)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.TargetBranch)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.DrySha)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.CommitMessage)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	if len(m.Paths) > 0 {
		for _, e := range m.Paths {
			l = e.Size()
			n += 1 + l + sovCommit(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PathDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	if len(m.Manifests) > 0 {
		for _, e := range m.Manifests {
			l = e.Size()
			n += 1 + l + sovCommit(uint64(l))
		}
	}
	if len(m.Commands) > 0 {
		for _, s := range m.Commands {
			l = len(s)
			n += 1 + l + sovCommit(uint64(l))
		}
	}
	if m.ReadmeDetails != nil {
		l = m.ReadmeDetails.Size()
		n += 1 + l + sovCommit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManifestDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Manifest)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReadmeDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManifestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommit(x uint64) (n int) {
	return sovCommit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ManifestsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repo == nil {
				m.Repo = &v1alpha1.Repository{}
			}
			if err := m.Repo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncBranch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SyncBranch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBranch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetBranch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrySha", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DrySha = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, &PathDetails{})
			if err := m.Paths[len(m.Paths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PathDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PathDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manifests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Manifests = append(m.Manifests, &ManifestDetails{})
			if err := m.Manifests[len(m.Manifests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commands", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commands = append(m.Commands, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadmeDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReadmeDetails == nil {
				m.ReadmeDetails = &ReadmeDetails{}
			}
			if err := m.ReadmeDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manifest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Manifest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadmeDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadmeDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadmeDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommit = fmt.Errorf("proto: unexpected end of group")
)