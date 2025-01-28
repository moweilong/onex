// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package app

import (
	"context"
	"fmt"

	"github.com/onexstack/onexstack/pkg/app"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/onexstack/onex/cmd/onex-gateway/app/options"
	"github.com/onexstack/onex/internal/gateway"
	"github.com/onexstack/onex/internal/pkg/contextx"
	"github.com/onexstack/onex/internal/pkg/known"
)

const commandDesc = `The gateway server is the back-end portal server of onex. All 
requests from the front-end will arrive at the gateway, requests will be uniformly processed 
and distributed by the gateway.`

// NewApp creates and returns a new App object with default parameters.
func NewApp() *app.App {
	opts := options.NewServerOptions()
	application := app.NewApp(
		gateway.Name,
		"Launch a onex gateway server",
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
		app.WithLoggerContextExtractor(map[string]func(context.Context) string{
			known.XTraceID: contextx.TraceID,
			known.XUserID:  contextx.UserID,
		}),
	)

	return application
}

// run contains the main logic for initializing and running the server.
func run(opts *options.ServerOptions) app.RunFunc {
	return func() error {
		// Load the configuration options
		cfg, err := opts.Config()
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		ctx := genericapiserver.SetupSignalContext()

		// Build the server using the configuration
		server, err := cfg.NewServer(ctx)
		if err != nil {
			return fmt.Errorf("failed to create server: %w", err)
		}

		// Run the server with signal context for graceful shutdown
		return server.Run(ctx)
	}
}
