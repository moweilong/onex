// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package info print the host information.
package info

import (
	"fmt"
	"reflect"
	"strconv"

	hoststat "github.com/likexian/host-stat-go"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"

	cmdutil "github.com/onexstack/onex/internal/onexctl/cmd/util"
	"github.com/onexstack/onex/internal/onexctl/util/templates"
	iputil "github.com/onexstack/onexstack/pkg/util/ip"
)

// Info defines the host information struct.
type Info struct {
	HostName  string
	IPAddress string
	OSRelease string
	CPUCore   uint64
	MemTotal  string
	MemFree   string
}

// InfoOptions is an options struct to support 'info' sub command.
type InfoOptions struct {
	genericiooptions.IOStreams
}

var infoExample = templates.Examples(`
		# Print the host information
		onexctl info`)

// NewInfoOptions returns an initialized InfoOptions instance.
func NewInfoOptions(ioStreams genericiooptions.IOStreams) *InfoOptions {
	return &InfoOptions{
		IOStreams: ioStreams,
	}
}

// NewCmdInfo returns new initialized instance of 'info' sub command.
func NewCmdInfo(f cmdutil.Factory, ioStreams genericiooptions.IOStreams) *cobra.Command {
	o := NewInfoOptions(ioStreams)

	cmd := &cobra.Command{
		Use:                   "info",
		DisableFlagsInUseLine: true,
		Aliases:               []string{},
		Short:                 "Print the host information",
		Long:                  "Print the host information.",
		Example:               infoExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Run(args))
		},
		SuggestFor: []string{},
	}

	return cmd
}

// Run executes an info sub command using the specified options.
func (o *InfoOptions) Run(args []string) error {
	var info Info

	hostInfo, err := hoststat.GetHostInfo()
	if err != nil {
		return fmt.Errorf("get host info failed!error:%w", err)
	}

	info.HostName = hostInfo.HostName
	info.OSRelease = hostInfo.Release + " " + hostInfo.OSBit

	memStat, err := hoststat.GetMemStat()
	if err != nil {
		return fmt.Errorf("get mem stat failed!error:%w", err)
	}

	info.MemTotal = strconv.FormatUint(memStat.MemTotal, 10) + "M"
	info.MemFree = strconv.FormatUint(memStat.MemFree, 10) + "M"
	info.IPAddress = iputil.GetLocalIP()

	cpuStat, err := hoststat.GetCPUInfo()
	if err != nil {
		return fmt.Errorf("get cpu stat failed!error:%w", err)
	}

	info.CPUCore = cpuStat.CoreCount

	s := reflect.ValueOf(&info).Elem()
	typeOfInfo := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		v := fmt.Sprintf("%v", f.Interface())
		if v != "" {
			fmt.Fprintf(o.Out, "%12s %v\n", typeOfInfo.Field(i).Name+":", f.Interface())
		}
	}

	return nil
}
