//*
// API to manage members of an organization.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: priv/members/v1alpha1/members.proto

package membersv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "go.jetpack.io/pkg/api/gen/priv/members/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MembersServiceName is the fully-qualified name of the MembersService service.
	MembersServiceName = "priv.members.v1alpha1.MembersService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MembersServiceGetMemberProcedure is the fully-qualified name of the MembersService's GetMember
	// RPC.
	MembersServiceGetMemberProcedure = "/priv.members.v1alpha1.MembersService/GetMember"
	// MembersServiceCreateMemberProcedure is the fully-qualified name of the MembersService's
	// CreateMember RPC.
	MembersServiceCreateMemberProcedure = "/priv.members.v1alpha1.MembersService/CreateMember"
	// MembersServiceDeleteMemberProcedure is the fully-qualified name of the MembersService's
	// DeleteMember RPC.
	MembersServiceDeleteMemberProcedure = "/priv.members.v1alpha1.MembersService/DeleteMember"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	membersServiceServiceDescriptor            = v1alpha1.File_priv_members_v1alpha1_members_proto.Services().ByName("MembersService")
	membersServiceGetMemberMethodDescriptor    = membersServiceServiceDescriptor.Methods().ByName("GetMember")
	membersServiceCreateMemberMethodDescriptor = membersServiceServiceDescriptor.Methods().ByName("CreateMember")
	membersServiceDeleteMemberMethodDescriptor = membersServiceServiceDescriptor.Methods().ByName("DeleteMember")
)

// MembersServiceClient is a client for the priv.members.v1alpha1.MembersService service.
type MembersServiceClient interface {
	// Get logged-in member
	//
	// Retrieves the details of an existing member identified by its unique
	// member id.
	GetMember(context.Context, *connect.Request[v1alpha1.GetMemberRequest]) (*connect.Response[v1alpha1.GetMemberResponse], error)
	CreateMember(context.Context, *connect.Request[v1alpha1.CreateMemberRequest]) (*connect.Response[v1alpha1.CreateMemberResponse], error)
	DeleteMember(context.Context, *connect.Request[v1alpha1.DeleteMemberRequest]) (*connect.Response[v1alpha1.DeleteMemberResponse], error)
}

// NewMembersServiceClient constructs a client for the priv.members.v1alpha1.MembersService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMembersServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MembersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &membersServiceClient{
		getMember: connect.NewClient[v1alpha1.GetMemberRequest, v1alpha1.GetMemberResponse](
			httpClient,
			baseURL+MembersServiceGetMemberProcedure,
			connect.WithSchema(membersServiceGetMemberMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createMember: connect.NewClient[v1alpha1.CreateMemberRequest, v1alpha1.CreateMemberResponse](
			httpClient,
			baseURL+MembersServiceCreateMemberProcedure,
			connect.WithSchema(membersServiceCreateMemberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteMember: connect.NewClient[v1alpha1.DeleteMemberRequest, v1alpha1.DeleteMemberResponse](
			httpClient,
			baseURL+MembersServiceDeleteMemberProcedure,
			connect.WithSchema(membersServiceDeleteMemberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// membersServiceClient implements MembersServiceClient.
type membersServiceClient struct {
	getMember    *connect.Client[v1alpha1.GetMemberRequest, v1alpha1.GetMemberResponse]
	createMember *connect.Client[v1alpha1.CreateMemberRequest, v1alpha1.CreateMemberResponse]
	deleteMember *connect.Client[v1alpha1.DeleteMemberRequest, v1alpha1.DeleteMemberResponse]
}

// GetMember calls priv.members.v1alpha1.MembersService.GetMember.
func (c *membersServiceClient) GetMember(ctx context.Context, req *connect.Request[v1alpha1.GetMemberRequest]) (*connect.Response[v1alpha1.GetMemberResponse], error) {
	return c.getMember.CallUnary(ctx, req)
}

// CreateMember calls priv.members.v1alpha1.MembersService.CreateMember.
func (c *membersServiceClient) CreateMember(ctx context.Context, req *connect.Request[v1alpha1.CreateMemberRequest]) (*connect.Response[v1alpha1.CreateMemberResponse], error) {
	return c.createMember.CallUnary(ctx, req)
}

// DeleteMember calls priv.members.v1alpha1.MembersService.DeleteMember.
func (c *membersServiceClient) DeleteMember(ctx context.Context, req *connect.Request[v1alpha1.DeleteMemberRequest]) (*connect.Response[v1alpha1.DeleteMemberResponse], error) {
	return c.deleteMember.CallUnary(ctx, req)
}

// MembersServiceHandler is an implementation of the priv.members.v1alpha1.MembersService service.
type MembersServiceHandler interface {
	// Get logged-in member
	//
	// Retrieves the details of an existing member identified by its unique
	// member id.
	GetMember(context.Context, *connect.Request[v1alpha1.GetMemberRequest]) (*connect.Response[v1alpha1.GetMemberResponse], error)
	CreateMember(context.Context, *connect.Request[v1alpha1.CreateMemberRequest]) (*connect.Response[v1alpha1.CreateMemberResponse], error)
	DeleteMember(context.Context, *connect.Request[v1alpha1.DeleteMemberRequest]) (*connect.Response[v1alpha1.DeleteMemberResponse], error)
}

// NewMembersServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMembersServiceHandler(svc MembersServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	membersServiceGetMemberHandler := connect.NewUnaryHandler(
		MembersServiceGetMemberProcedure,
		svc.GetMember,
		connect.WithSchema(membersServiceGetMemberMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	membersServiceCreateMemberHandler := connect.NewUnaryHandler(
		MembersServiceCreateMemberProcedure,
		svc.CreateMember,
		connect.WithSchema(membersServiceCreateMemberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	membersServiceDeleteMemberHandler := connect.NewUnaryHandler(
		MembersServiceDeleteMemberProcedure,
		svc.DeleteMember,
		connect.WithSchema(membersServiceDeleteMemberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/priv.members.v1alpha1.MembersService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MembersServiceGetMemberProcedure:
			membersServiceGetMemberHandler.ServeHTTP(w, r)
		case MembersServiceCreateMemberProcedure:
			membersServiceCreateMemberHandler.ServeHTTP(w, r)
		case MembersServiceDeleteMemberProcedure:
			membersServiceDeleteMemberHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMembersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMembersServiceHandler struct{}

func (UnimplementedMembersServiceHandler) GetMember(context.Context, *connect.Request[v1alpha1.GetMemberRequest]) (*connect.Response[v1alpha1.GetMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.members.v1alpha1.MembersService.GetMember is not implemented"))
}

func (UnimplementedMembersServiceHandler) CreateMember(context.Context, *connect.Request[v1alpha1.CreateMemberRequest]) (*connect.Response[v1alpha1.CreateMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.members.v1alpha1.MembersService.CreateMember is not implemented"))
}

func (UnimplementedMembersServiceHandler) DeleteMember(context.Context, *connect.Request[v1alpha1.DeleteMemberRequest]) (*connect.Response[v1alpha1.DeleteMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.members.v1alpha1.MembersService.DeleteMember is not implemented"))
}