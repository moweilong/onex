// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package app

import (
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	genericapiserver "k8s.io/apiserver/pkg/server"
	logsapi "k8s.io/component-base/logs/api/v1"

	"github.com/onexstack/onex/cmd/onex-pump/app/options"
	"github.com/onexstack/onex/internal/pkg/feature"
	"github.com/onexstack/onex/internal/pump"
	"github.com/onexstack/onex/pkg/app"
)

func init() {
	utilruntime.Must(logsapi.AddFeatureGates(feature.DefaultMutableFeatureGate))
}

const commandDesc = `Pump is a pluggable analytics purger to move Analytics generated by your onex nodes to any back-end.

Find more onex-pump information at:
    https://github.com/onexstack/onex/blob/master/docs/guide/en-US/cmd/onex-pump.md`

// NewApp creates an App object with default parameters.
func NewApp() *app.App {
	opts := options.NewOptions()
	application := app.NewApp("onex-pump", "Launch a onex pump server",
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
		app.WithHealthCheckFunc(func() error {
			go opts.HealthOptions.ServeHealthCheck()
			return nil
		}),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func() error {
		cfg, err := opts.Config()
		if err != nil {
			return err
		}

		return Run(cfg, genericapiserver.SetupSignalHandler())
	}
}

// Run runs the specified APIServer. This should never exit.
func Run(c *pump.Config, stopCh <-chan struct{}) error {
	server, err := c.Complete().New()
	if err != nil {
		return err
	}

	return server.PrepareRun().Run(stopCh)
}
