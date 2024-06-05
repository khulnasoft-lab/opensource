//*
// API to manage token service

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: priv/tokenservice/v1alpha1/tokenservice.proto

package tokenservicev1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "go.jetpack.io/pkg/api/gen/priv/tokenservice/v1alpha1"
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
	// TokenServiceName is the fully-qualified name of the TokenService service.
	TokenServiceName = "priv.tokenservice.v1alpha1.TokenService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TokenServiceGetAccessTokenProcedure is the fully-qualified name of the TokenService's
	// GetAccessToken RPC.
	TokenServiceGetAccessTokenProcedure = "/priv.tokenservice.v1alpha1.TokenService/GetAccessToken"
	// TokenServiceCreatePATProcedure is the fully-qualified name of the TokenService's CreatePAT RPC.
	TokenServiceCreatePATProcedure = "/priv.tokenservice.v1alpha1.TokenService/CreatePAT"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	tokenServiceServiceDescriptor              = v1alpha1.File_priv_tokenservice_v1alpha1_tokenservice_proto.Services().ByName("TokenService")
	tokenServiceGetAccessTokenMethodDescriptor = tokenServiceServiceDescriptor.Methods().ByName("GetAccessToken")
	tokenServiceCreatePATMethodDescriptor      = tokenServiceServiceDescriptor.Methods().ByName("CreatePAT")
)

// TokenServiceClient is a client for the priv.tokenservice.v1alpha1.TokenService service.
type TokenServiceClient interface {
	GetAccessToken(context.Context, *connect.Request[v1alpha1.GetAccessTokenRequest]) (*connect.Response[v1alpha1.GetAccessTokenResponse], error)
	CreatePAT(context.Context, *connect.Request[v1alpha1.CreatePATRequest]) (*connect.Response[v1alpha1.CreatePATResponse], error)
}

// NewTokenServiceClient constructs a client for the priv.tokenservice.v1alpha1.TokenService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTokenServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TokenServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tokenServiceClient{
		getAccessToken: connect.NewClient[v1alpha1.GetAccessTokenRequest, v1alpha1.GetAccessTokenResponse](
			httpClient,
			baseURL+TokenServiceGetAccessTokenProcedure,
			connect.WithSchema(tokenServiceGetAccessTokenMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createPAT: connect.NewClient[v1alpha1.CreatePATRequest, v1alpha1.CreatePATResponse](
			httpClient,
			baseURL+TokenServiceCreatePATProcedure,
			connect.WithSchema(tokenServiceCreatePATMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tokenServiceClient implements TokenServiceClient.
type tokenServiceClient struct {
	getAccessToken *connect.Client[v1alpha1.GetAccessTokenRequest, v1alpha1.GetAccessTokenResponse]
	createPAT      *connect.Client[v1alpha1.CreatePATRequest, v1alpha1.CreatePATResponse]
}

// GetAccessToken calls priv.tokenservice.v1alpha1.TokenService.GetAccessToken.
func (c *tokenServiceClient) GetAccessToken(ctx context.Context, req *connect.Request[v1alpha1.GetAccessTokenRequest]) (*connect.Response[v1alpha1.GetAccessTokenResponse], error) {
	return c.getAccessToken.CallUnary(ctx, req)
}

// CreatePAT calls priv.tokenservice.v1alpha1.TokenService.CreatePAT.
func (c *tokenServiceClient) CreatePAT(ctx context.Context, req *connect.Request[v1alpha1.CreatePATRequest]) (*connect.Response[v1alpha1.CreatePATResponse], error) {
	return c.createPAT.CallUnary(ctx, req)
}

// TokenServiceHandler is an implementation of the priv.tokenservice.v1alpha1.TokenService service.
type TokenServiceHandler interface {
	GetAccessToken(context.Context, *connect.Request[v1alpha1.GetAccessTokenRequest]) (*connect.Response[v1alpha1.GetAccessTokenResponse], error)
	CreatePAT(context.Context, *connect.Request[v1alpha1.CreatePATRequest]) (*connect.Response[v1alpha1.CreatePATResponse], error)
}

// NewTokenServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTokenServiceHandler(svc TokenServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tokenServiceGetAccessTokenHandler := connect.NewUnaryHandler(
		TokenServiceGetAccessTokenProcedure,
		svc.GetAccessToken,
		connect.WithSchema(tokenServiceGetAccessTokenMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceCreatePATHandler := connect.NewUnaryHandler(
		TokenServiceCreatePATProcedure,
		svc.CreatePAT,
		connect.WithSchema(tokenServiceCreatePATMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/priv.tokenservice.v1alpha1.TokenService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TokenServiceGetAccessTokenProcedure:
			tokenServiceGetAccessTokenHandler.ServeHTTP(w, r)
		case TokenServiceCreatePATProcedure:
			tokenServiceCreatePATHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTokenServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTokenServiceHandler struct{}

func (UnimplementedTokenServiceHandler) GetAccessToken(context.Context, *connect.Request[v1alpha1.GetAccessTokenRequest]) (*connect.Response[v1alpha1.GetAccessTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.tokenservice.v1alpha1.TokenService.GetAccessToken is not implemented"))
}

func (UnimplementedTokenServiceHandler) CreatePAT(context.Context, *connect.Request[v1alpha1.CreatePATRequest]) (*connect.Response[v1alpha1.CreatePATResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.tokenservice.v1alpha1.TokenService.CreatePAT is not implemented"))
}
