package main

import (
	"flag"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"example.com/mod/order/internal/config"
	"example.com/mod/order/internal/handlers"
	"example.com/mod/order/internal/service"
	order_v1 "example.com/mod/proto/gen/go"
	"google.golang.org/grpc"
)

func main() {
	configPath := flag.String("config", "/home/kirill/GolandProjects/orderSystem/order/internal/config/config.yaml", "path to config file") // config

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		slog.Info("Error load config", "error", err)
		os.Exit(1)
	}

	// GRPC_HOST := cfg.GRPCServerConfig.Host // configs lines
	GRPC_PORT := cfg.GRPCServerConfig.Port
	DATABASE_URL := cfg.DatabaseConfig.Url

	repo, err := service.NewPostgresOrderRepository(DATABASE_URL) // connect db
	if err != nil {
		slog.Info("Error connect database", "error", err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	order_v1.RegisterOrderServiceServer(server, &handlers.OrderServer{Repo: repo})

	lis, err := net.Listen("tcp", GRPC_PORT) // listening server
	if err != nil {
		slog.Info("Error listening port", "error", err)
		os.Exit(1)
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		slog.Info("Starting gRPC server", "port", GRPC_PORT)
		if err := server.Serve(lis); err != nil {
			slog.Error("Failed to serve", "error", err)
		}
	}()

	<-quit
	slog.Info("Proccesing stopped gRPC server...")

	// todo context with timeout

	server.GracefulStop()
	lis.Close()

	slog.Info("Stopped gRPC server")
}
