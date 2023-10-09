package main

import (
	"net"
	pb "simplebank/pb"
	service "simplebank/service/auth"
	"testing"

	_ "github.com/lib/pq"
)

func Test_runRESTServer(t *testing.T) {
	type args struct {
		authServer     pb.AuthServiceServer
		dumpServer     pb.DumpServiceServer
		documentServer pb.DocumentServiceServer
		kplusServer    pb.KPlusServiceServer
		jwtManager     *service.JWTManager
		enableTLS      bool
		listener       net.Listener
		grpcEndpoint   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runRESTServer(tt.args.authServer, tt.args.dumpServer, tt.args.documentServer, tt.args.kplusServer, tt.args.jwtManager, tt.args.enableTLS, tt.args.listener, tt.args.grpcEndpoint); (err != nil) != tt.wantErr {
				t.Errorf("runRESTServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
