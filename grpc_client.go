package authzed

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	api "github.com/authzed/authzed-go/arrakisapi/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GrpcClientOptions represents the options that can be used to configure the client.
type GrpcClientOptions struct {
	// Endpoint is the grpc DNS name at which to find the service. Defaults to the production endpoint.
	Endpoint string

	// Certificate pool from which to load certificates. Defaults to the system cert pool.
	Certificates *x509.CertPool

	// Token is the token which is used to authorized all requests.
	Token string
}

// NewClientOptions returns a ClientOptions with defaults set.
func NewGrpcClientOptions(token string) (GrpcClientOptions, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return GrpcClientOptions{}, err
	}

	return GrpcClientOptions{
		Endpoint:     "grpc.authzed.com:443",
		Certificates: certPool,
		Token:        token,
	}, nil
}

// NewClient creates a new Authzed client authorized with the given token.
func NewGrpcClient(options GrpcClientOptions) (api.ACLServiceClient, error) {
	conn, err := getConnection(options)
	if err != nil {
		return nil, err
	}

	return api.NewACLServiceClient(conn), nil
}

// NewManagementClient creates a new Authzed management client with the given token.
func NewGrpcManagementClient(options GrpcClientOptions) (api.NamespaceServiceClient, error) {
	conn, err := getConnection(options)
	if err != nil {
		return nil, err
	}

	return api.NewNamespaceServiceClient(conn), nil
}

type presharedKeyCredentials struct {
	metadata map[string]string
}

// NewPresharedKeyCredentials creates a new credentials type which applies a
// fixed authorization metadataum of the form `authorization: Bearer presharedKey`
func NewGrpcPresharedKeyCredentials(presharedKey string) credentials.PerRPCCredentials {
	return presharedKeyCredentials{
		metadata: map[string]string{
			"authorization": "Bearer " + presharedKey,
		},
	}
}

func (pskc presharedKeyCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return pskc.metadata, nil
}

func (pskc presharedKeyCredentials) RequireTransportSecurity() bool {
	return false
}

func getConnection(options GrpcClientOptions) (*grpc.ClientConn, error) {
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            options.Certificates,
	})

	return grpc.Dial(
		options.Endpoint,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewGrpcPresharedKeyCredentials(options.Token)),
	)
}