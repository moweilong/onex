// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra/doc"
	"k8s.io/kubernetes/cmd/genutils"

	apiservapp "github.com/onexstack/onex/cmd/onex-apiserver/app"
	ctrlmgrapp "github.com/onexstack/onex/cmd/onex-controller-manager/app"
	// fakeserverapp "github.com/onexstack/onex/cmd/onex-fakeserver/app"
	gwapp "github.com/onexstack/onex/cmd/onex-gateway/app"
	watchapp "github.com/onexstack/onex/cmd/onex-nightwatch/app"
	pumpapp "github.com/onexstack/onex/cmd/onex-pump/app"
	toyblcapp "github.com/onexstack/onex/cmd/onex-toyblc/app"
	usercenterapp "github.com/onexstack/onex/cmd/onex-usercenter/app"
	onexctlcmd "github.com/onexstack/onex/internal/onexctl/cmd"
)

func main() {
	// use os.Args instead of "flags" because "flags" will mess up the man pages!
	path, module := "", ""
	if len(os.Args) == 3 {
		path = os.Args[1]
		module = os.Args[2]
	} else {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory] [module] \n", os.Args[0])
		os.Exit(1)
	}

	outDir, err := genutils.OutDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}

	switch module {
	/*
		case "onex-fakeserver":
			// generate docs for onexfakeserver-
			fakeserver := fakeserverapp.NewApp().Command()
			_ = doc.GenMarkdownTree(fakeserver, outDir)
	*/
	case "onex-usercenter":
		// generate docs for onex-usercenter
		usercenter := usercenterapp.NewApp().Command()
		_ = doc.GenMarkdownTree(usercenter, outDir)
	case "onex-apiserver":
		// generate docs for onex-apiserver
		apiserver := apiservapp.NewAPIServerCommand()
		_ = doc.GenMarkdownTree(apiserver, outDir)
	case "onex-gateway":
		// generate docs for onex-gateway
		gwserver := gwapp.NewApp().Command()
		_ = doc.GenMarkdownTree(gwserver, outDir)
	case "onex-nightwatch":
		// generate docs for onex-nightwatch
		nw := watchapp.NewApp("onex-nightwatch").Command()
		_ = doc.GenMarkdownTree(nw, outDir)
	case "onex-pump":
		// generate docs for onex-pump
		pump := pumpapp.NewApp().Command()
		_ = doc.GenMarkdownTree(pump, outDir)
	case "onex-toyblc":
		// generate docs for onex-toyblc
		toyblc := toyblcapp.NewApp().Command()
		_ = doc.GenMarkdownTree(toyblc, outDir)
	case "onex-controller-manager":
		// generate docs for onex-controller-manager
		ctrlmgr := ctrlmgrapp.NewControllerManagerCommand()
		_ = doc.GenMarkdownTree(ctrlmgr, outDir)
	case "onexctl":
		// generate docs for onexctl
		onexctl := onexctlcmd.NewDefaultOneXCtlCommand()
		_ = doc.GenMarkdownTree(onexctl, outDir)
	default:
		fmt.Fprintf(os.Stderr, "Module %s is not supported", module)
		os.Exit(1)
	}
}
