// Code generated by protoc-gen-go.
// source: google/logging/v2/logging.proto
// DO NOT EDIT!

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/googleapis/proto-client-go/api"
import google_api1 "github.com/googleapis/proto-client-go/api"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The parameters to DeleteLog.
type DeleteLogRequest struct {
	// Required. The resource name of the log to delete.  Example:
	// `"projects/my-project/logs/syslog"`.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
}

func (m *DeleteLogRequest) Reset()                    { *m = DeleteLogRequest{} }
func (m *DeleteLogRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogRequest) ProtoMessage()               {}
func (*DeleteLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// Optional. A default log resource name for those log entries in `entries`
	// that do not specify their own `logName`.  Example:
	// `"projects/my-project/logs/syslog"`.  See
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Optional. A default monitored resource for those log entries in `entries`
	// that do not specify their own `resource`.
	Resource *google_api1.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// Optional. User-defined `key:value` items that are added to
	// the `labels` field of each log entry in `entries`, except when a log
	// entry specifies its own `key:value` item with the same key.
	// Example: `{ "size": "large", "color":"red" }`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Required. The log entries to write. The log entries must have values for
	// all required fields.
	Entries []*LogEntry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
}

func (m *WriteLogEntriesRequest) Reset()                    { *m = WriteLogEntriesRequest{} }
func (m *WriteLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesRequest) ProtoMessage()               {}
func (*WriteLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *WriteLogEntriesRequest) GetResource() *google_api1.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Result returned from WriteLogEntries.
type WriteLogEntriesResponse struct {
}

func (m *WriteLogEntriesResponse) Reset()                    { *m = WriteLogEntriesResponse{} }
func (m *WriteLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesResponse) ProtoMessage()               {}
func (*WriteLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

// The parameters to `ListLogEntries`.
type ListLogEntriesRequest struct {
	// Required. One or more project IDs or project numbers from which to retrieve
	// log entries.  Examples of a project ID: `"my-project-1A"`, `"1234567890"`.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds" json:"project_ids,omitempty"`
	// Optional. An [advanced logs filter](/logging/docs/view/advanced_filters).
	// The filter is compared against all log entries in the projects specified by
	// `projectIds`.  Only entries that match the filter are retrieved.  An empty
	// filter matches all log entries.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// Optional. How the results should be sorted.  Presently, the only permitted
	// values are `"timestamp"` (default) and `"timestamp desc"`.  The first
	// option returns entries in order of increasing values of
	// `LogEntry.timestamp` (oldest first), and the second option returns entries
	// in order of decreasing timestamps (newest first).  Entries with equal
	// timestamps are returned in order of `LogEntry.insertId`.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
	// Optional. The maximum number of results to return from this request.  Fewer
	// results might be returned. You must check for the `nextPageToken` result to
	// determine if additional results are available, which you can retrieve by
	// passing the `nextPageToken` value in the `pageToken` parameter to the next
	// request.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If the `pageToken` request parameter is supplied, then the next
	// page of results in the set are retrieved.  The `pageToken` parameter must
	// be set with the value of the `nextPageToken` result parameter from the
	// previous request.  The values of `projectIds`, `filter`, and `orderBy` must
	// be the same as in the previous request.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLogEntriesRequest) Reset()                    { *m = ListLogEntriesRequest{} }
func (m *ListLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesRequest) ProtoMessage()               {}
func (*ListLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

// Result returned from `ListLogEntries`.
type ListLogEntriesResponse struct {
	// A list of log entries.
	Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// given a value in the response.  To get the next batch of results, call
	// this method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogEntriesResponse) Reset()                    { *m = ListLogEntriesResponse{} }
func (m *ListLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesResponse) ProtoMessage()               {}
func (*ListLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ListLogEntriesResponse) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// The parameters to ListMonitoredResourceDescriptors
type ListMonitoredResourceDescriptorsRequest struct {
	// Optional. The maximum number of results to return from this request.  Fewer
	// results might be returned. You must check for the `nextPageToken` result to
	// determine if additional results are available, which you can retrieve by
	// passing the `nextPageToken` value in the `pageToken` parameter to the next
	// request.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If the `pageToken` request parameter is supplied, then the next
	// page of results in the set are retrieved.  The `pageToken` parameter must
	// be set with the value of the `nextPageToken` result parameter from the
	// previous request.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
}
func (m *ListMonitoredResourceDescriptorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{5}
}

// Result returned from ListMonitoredResourceDescriptors.
type ListMonitoredResourceDescriptorsResponse struct {
	// A list of resource descriptors.
	ResourceDescriptors []*google_api1.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=resource_descriptors,json=resourceDescriptors" json:"resource_descriptors,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// returned in the response.  To get the next batch of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
}
func (m *ListMonitoredResourceDescriptorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{6}
}

func (m *ListMonitoredResourceDescriptorsResponse) GetResourceDescriptors() []*google_api1.MonitoredResourceDescriptor {
	if m != nil {
		return m.ResourceDescriptors
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteLogRequest)(nil), "google.logging.v2.DeleteLogRequest")
	proto.RegisterType((*WriteLogEntriesRequest)(nil), "google.logging.v2.WriteLogEntriesRequest")
	proto.RegisterType((*WriteLogEntriesResponse)(nil), "google.logging.v2.WriteLogEntriesResponse")
	proto.RegisterType((*ListLogEntriesRequest)(nil), "google.logging.v2.ListLogEntriesRequest")
	proto.RegisterType((*ListLogEntriesResponse)(nil), "google.logging.v2.ListLogEntriesResponse")
	proto.RegisterType((*ListMonitoredResourceDescriptorsRequest)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsRequest")
	proto.RegisterType((*ListMonitoredResourceDescriptorsResponse)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for LoggingServiceV2 service

type LoggingServiceV2Client interface {
	// Deletes a log and all its log entries.
	// The log will reappear if it receives new entries.
	//
	DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
	// Writes log entries to Cloud Logging.
	// All log entries in Cloud Logging are written by this method.
	//
	WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from Cloud
	// Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	//
	ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error)
	// Lists monitored resource descriptors that are used by Cloud Logging.
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
}

type loggingServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewLoggingServiceV2Client(cc *grpc.ClientConn) LoggingServiceV2Client {
	return &loggingServiceV2Client{cc}
}

func (c *loggingServiceV2Client) DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/DeleteLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error) {
	out := new(WriteLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/WriteLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error) {
	out := new(ListLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoggingServiceV2 service

type LoggingServiceV2Server interface {
	// Deletes a log and all its log entries.
	// The log will reappear if it receives new entries.
	//
	DeleteLog(context.Context, *DeleteLogRequest) (*google_protobuf4.Empty, error)
	// Writes log entries to Cloud Logging.
	// All log entries in Cloud Logging are written by this method.
	//
	WriteLogEntries(context.Context, *WriteLogEntriesRequest) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from Cloud
	// Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	//
	ListLogEntries(context.Context, *ListLogEntriesRequest) (*ListLogEntriesResponse, error)
	// Lists monitored resource descriptors that are used by Cloud Logging.
	ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error)
}

func RegisterLoggingServiceV2Server(s *grpc.Server, srv LoggingServiceV2Server) {
	s.RegisterService(&_LoggingServiceV2_serviceDesc, srv)
}

func _LoggingServiceV2_DeleteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/DeleteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, req.(*DeleteLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_WriteLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/WriteLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, req.(*WriteLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, req.(*ListLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitoredResourceDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, req.(*ListMonitoredResourceDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.LoggingServiceV2",
	HandlerType: (*LoggingServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteLog",
			Handler:    _LoggingServiceV2_DeleteLog_Handler,
		},
		{
			MethodName: "WriteLogEntries",
			Handler:    _LoggingServiceV2_WriteLogEntries_Handler,
		},
		{
			MethodName: "ListLogEntries",
			Handler:    _LoggingServiceV2_ListLogEntries_Handler,
		},
		{
			MethodName: "ListMonitoredResourceDescriptors",
			Handler:    _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor3 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0x12, 0x5f,
	0x14, 0xcf, 0x85, 0x7e, 0x71, 0xfa, 0xff, 0xdb, 0x7a, 0x6d, 0x91, 0x0e, 0x36, 0xa5, 0xd3, 0x68,
	0x29, 0x09, 0x83, 0x62, 0x9a, 0x58, 0x8c, 0x9b, 0xa6, 0x5d, 0x98, 0x50, 0xd3, 0x4c, 0x8d, 0x26,
	0x8d, 0x09, 0x19, 0xe0, 0x74, 0xbc, 0x76, 0x98, 0x8b, 0x33, 0x17, 0x2a, 0x1a, 0x37, 0x6e, 0x5c,
	0xb8, 0xf4, 0x21, 0xdc, 0xe9, 0x7b, 0xb8, 0xf5, 0x15, 0x7c, 0x00, 0x97, 0x2e, 0xbd, 0x33, 0x73,
	0x87, 0x52, 0xc0, 0x96, 0xb8, 0x82, 0x73, 0xce, 0xef, 0x7c, 0xff, 0xce, 0x1d, 0x58, 0xb3, 0x39,
	0xb7, 0x1d, 0x2c, 0x39, 0xdc, 0xb6, 0x99, 0x6b, 0x97, 0xba, 0xe5, 0xf8, 0xaf, 0xd1, 0xf6, 0xb8,
	0xe0, 0xf4, 0x7a, 0x04, 0x30, 0x62, 0x6d, 0xb7, 0xac, 0xdd, 0x52, 0x3e, 0x56, 0x9b, 0x95, 0x2c,
	0xd7, 0xe5, 0xc2, 0x12, 0x8c, 0xbb, 0x7e, 0xe4, 0xa0, 0x6d, 0x0c, 0x58, 0x5b, 0xdc, 0x65, 0x82,
	0x7b, 0xd8, 0xac, 0x79, 0xe8, 0xf3, 0x8e, 0xd7, 0x40, 0x05, 0x5a, 0x1f, 0x9b, 0xb6, 0x86, 0xae,
	0xf0, 0x7a, 0x0a, 0x92, 0x55, 0x90, 0x50, 0xaa, 0x77, 0x4e, 0x4a, 0xd8, 0x6a, 0x0b, 0x65, 0xd4,
	0x8b, 0xb0, 0xb8, 0x87, 0x0e, 0x0a, 0xac, 0x72, 0xdb, 0xc4, 0xd7, 0x1d, 0xf4, 0x05, 0x5d, 0x81,
	0xb9, 0x20, 0x86, 0x6b, 0xb5, 0x30, 0x43, 0x72, 0x24, 0x9f, 0x32, 0x67, 0xa5, 0xfc, 0x44, 0x8a,
	0xfa, 0xb7, 0x04, 0xa4, 0x9f, 0x7b, 0x2c, 0x84, 0xef, 0xcb, 0x1c, 0x0c, 0xfd, 0xab, 0xbd, 0xe8,
	0x0e, 0xcc, 0xc5, 0x65, 0x67, 0x12, 0xd2, 0x34, 0x5f, 0x5e, 0x35, 0xd4, 0x34, 0x64, 0x73, 0xc6,
	0x41, 0xdc, 0x9c, 0xa9, 0x40, 0x66, 0x1f, 0x4e, 0x0f, 0x60, 0xc6, 0xb1, 0xea, 0xe8, 0xf8, 0x99,
	0x64, 0x2e, 0x29, 0x1d, 0xb7, 0x8d, 0x91, 0x31, 0x1a, 0xe3, 0x0b, 0x32, 0xaa, 0xa1, 0x5f, 0xa0,
	0xec, 0x99, 0x2a, 0x08, 0xdd, 0x86, 0x59, 0x8c, 0x50, 0x99, 0xa9, 0x30, 0x5e, 0x76, 0x4c, 0x3c,
	0x15, 0xaa, 0x67, 0xc6, 0x58, 0x6d, 0x07, 0xe6, 0x07, 0xa2, 0xd1, 0x45, 0x48, 0x9e, 0x62, 0x4f,
	0x75, 0x19, 0xfc, 0xa5, 0x4b, 0x30, 0xdd, 0xb5, 0x9c, 0x4e, 0xd4, 0x5e, 0xca, 0x8c, 0x84, 0x4a,
	0xe2, 0x01, 0xd1, 0x57, 0xe0, 0xe6, 0x48, 0x7d, 0x7e, 0x5b, 0x6e, 0x19, 0xf5, 0x2f, 0x04, 0x96,
	0xab, 0xcc, 0x17, 0xa3, 0xb3, 0x5c, 0x83, 0x79, 0xb9, 0x9e, 0x57, 0xd8, 0x10, 0x35, 0xd6, 0xf4,
	0x65, 0xa2, 0xa4, 0x0c, 0x0a, 0x4a, 0xf5, 0xb8, 0xe9, 0xd3, 0x34, 0xcc, 0x9c, 0x30, 0x47, 0xa0,
	0xa7, 0x12, 0x2a, 0x29, 0x58, 0x02, 0xf7, 0x9a, 0xe8, 0xd5, 0xea, 0x3d, 0x39, 0xb0, 0x70, 0x09,
	0xa1, 0xbc, 0xdb, 0xa3, 0x59, 0x48, 0xb5, 0x2d, 0x1b, 0x6b, 0x3e, 0x7b, 0x8b, 0xb2, 0x79, 0x92,
	0x9f, 0x36, 0xe7, 0x02, 0xc5, 0x91, 0x94, 0xe9, 0x2a, 0x40, 0x68, 0x14, 0xfc, 0x14, 0xdd, 0xcc,
	0x74, 0xe8, 0x19, 0xc2, 0x9f, 0x06, 0x0a, 0xfd, 0x0c, 0xd2, 0xc3, 0x85, 0x46, 0x3d, 0x0c, 0x0e,
	0x94, 0x4c, 0x3e, 0x50, 0x7a, 0x07, 0x16, 0x5c, 0x7c, 0x23, 0x6a, 0x03, 0x49, 0xa3, 0x46, 0xfe,
	0x0f, 0xd4, 0x87, 0xfd, 0xc4, 0x08, 0x9b, 0x41, 0xe2, 0x11, 0x86, 0xec, 0xa1, 0xdf, 0xf0, 0x58,
	0x5b, 0xea, 0xfa, 0x33, 0xbb, 0xd0, 0x1f, 0xb9, 0xb4, 0xbf, 0xc4, 0x70, 0x7f, 0x5f, 0x09, 0xe4,
	0xaf, 0xce, 0xa3, 0x5a, 0x3e, 0x86, 0xa5, 0x98, 0x9e, 0xb5, 0xe6, 0xb9, 0x5d, 0xf5, 0xbf, 0x79,
	0x29, 0xb3, 0xcf, 0xe3, 0x99, 0x37, 0xbc, 0xd1, 0x1c, 0x93, 0xce, 0xa5, 0xfc, 0x6b, 0x0a, 0x16,
	0xab, 0xd1, 0x80, 0x8f, 0xd0, 0xeb, 0xb2, 0x06, 0x3e, 0x2b, 0xd3, 0x33, 0x48, 0xf5, 0x6f, 0x99,
	0x6e, 0x8c, 0xd9, 0xc3, 0xf0, 0xa5, 0x6b, 0xe9, 0x18, 0x14, 0xbf, 0x0d, 0xc6, 0x7e, 0xf0, 0x36,
	0xe8, 0xc5, 0x0f, 0x3f, 0x7e, 0x7e, 0x4e, 0x6c, 0x16, 0x6e, 0xcb, 0xf7, 0xa4, 0x8e, 0xc2, 0xba,
	0x57, 0x7a, 0x17, 0xdf, 0xf6, 0x23, 0xc5, 0x42, 0xbf, 0x54, 0x08, 0x5e, 0x1a, 0xf9, 0xf3, 0x9e,
	0x7e, 0x22, 0xb0, 0x30, 0x44, 0x72, 0xba, 0x35, 0xf1, 0xa1, 0x6a, 0x85, 0x49, 0xa0, 0xea, 0x66,
	0xd6, 0xc3, 0xca, 0xb2, 0x7a, 0xba, 0x5f, 0x99, 0xa2, 0x54, 0xe5, 0x2c, 0xf0, 0xa8, 0x90, 0x02,
	0xfd, 0x48, 0xe0, 0xda, 0x45, 0xb6, 0xd2, 0xfc, 0x38, 0x52, 0x8e, 0xbb, 0x3c, 0x6d, 0x6b, 0x02,
	0xa4, 0x2a, 0x25, 0x17, 0x96, 0xa2, 0xe9, 0xcb, 0x23, 0xa5, 0x38, 0xd2, 0x21, 0xa8, 0xe4, 0x3b,
	0x81, 0xdc, 0x55, 0xb4, 0xa2, 0x95, 0xbf, 0x64, 0x9c, 0x80, 0xf3, 0xda, 0xc3, 0x7f, 0xf2, 0x55,
	0xf5, 0xab, 0x25, 0xd3, 0xf3, 0x25, 0xb7, 0x2e, 0x71, 0xdb, 0x7d, 0x01, 0xcb, 0x0d, 0xde, 0x1a,
	0x4d, 0xb8, 0xfb, 0x9f, 0x22, 0xe2, 0x61, 0xc0, 0xa1, 0x43, 0x72, 0x7c, 0xd7, 0x66, 0xe2, 0x65,
	0xa7, 0x6e, 0x48, 0x74, 0x29, 0x42, 0xcb, 0x53, 0xf0, 0xa3, 0xcf, 0x4f, 0xb1, 0xe1, 0x30, 0x39,
	0xa5, 0xa2, 0xcd, 0x07, 0xbe, 0x58, 0xbf, 0x09, 0xa9, 0xcf, 0x84, 0xe6, 0xfb, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0x70, 0x19, 0x9a, 0x47, 0x07, 0x00, 0x00,
}
