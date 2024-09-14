package grpc

import (
	"fmt"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/materials"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/supplier"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/warehouse"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type Server struct {
	server          *grpc.Server
	warehouseServer warehouse.WarehouseServiceServer
	supplierServer  supplier.SupplierServiceServer
	materialsServer materials.MaterialServiceServer
}

func New(warehouseServer warehouse.WarehouseServiceServer, supplierServer supplier.SupplierServiceServer, materialsServer materials.MaterialServiceServer) *Server {
	opt := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 100),
		grpc.MaxSendMsgSize(1024 * 1024 * 100),
		grpc.MaxConcurrentStreams(1000),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	}

	return &Server{
		server:          grpc.NewServer(opt...),
		warehouseServer: warehouseServer,
		supplierServer:  supplierServer,
		materialsServer: materialsServer,
	}
}

func (s *Server) Run(port int64) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	warehouse.RegisterWarehouseServiceServer(s.server, s.warehouseServer)
	supplier.RegisterSupplierServiceServer(s.server, s.supplierServer)
	materials.RegisterMaterialServiceServer(s.server, s.materialsServer)

	if err = s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() func() {
	return s.server.GracefulStop
}
