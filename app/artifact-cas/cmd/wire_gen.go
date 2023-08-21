// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chainloop-dev/chainloop/app/artifact-cas/internal/conf"
	"github.com/chainloop-dev/chainloop/app/artifact-cas/internal/server"
	"github.com/chainloop-dev/chainloop/app/artifact-cas/internal/service"
	"github.com/chainloop-dev/chainloop/internal/blobmanager/oci"
	"github.com/chainloop-dev/chainloop/internal/credentials"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, auth *conf.Auth, reader credentials.Reader, logger log.Logger) (*kratos.App, func(), error) {
	backendProvider := oci.NewBackendProvider(reader)
	v := serviceOpts(logger)
	byteStreamService := service.NewByteStreamService(backendProvider, v...)
	resourceService := service.NewResourceService(backendProvider, v...)
	grpcServer, err := server.NewGRPCServer(confServer, auth, byteStreamService, resourceService, logger)
	if err != nil {
		return nil, nil, err
	}
	downloadService := service.NewDownloadService(backendProvider, v...)
	httpServer, err := server.NewHTTPServer(confServer, auth, downloadService, logger)
	if err != nil {
		return nil, nil, err
	}
	httpMetricsServer, err := server.NewHTTPMetricsServer(confServer)
	if err != nil {
		return nil, nil, err
	}
	app := newApp(logger, grpcServer, httpServer, httpMetricsServer)
	return app, func() {
	}, nil
}

// wire.go:

func serviceOpts(l log.Logger) []service.NewOpt {
	return []service.NewOpt{service.WithLogger(l)}
}
