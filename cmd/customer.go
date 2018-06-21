// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"

	"gitlab.com/will.wang1/hotrod-customer/customer"
)

// customerCmd represents the customer command
var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Starts Customer service",
	Long:  `Starts Customer service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := customer.NewServer(
			net.JoinHostPort(customerOptions.serverInterface, strconv.Itoa(customerOptions.serverPort)),
		)
		if err != nil {
			return logError(server.Run())
		}
		return logError(server.Run())
	},
}

var (
	customerOptions struct {
		serverInterface string
		serverPort      int
	}
)

func init() {
	RootCmd.AddCommand(customerCmd)

	customerCmd.Flags().StringVarP(&customerOptions.serverInterface, "bind", "", "0.0.0.0", "interface to which the Customer server will bind")
	customerCmd.Flags().IntVarP(&customerOptions.serverPort, "port", "p", 8081, "port on which the Customer server will listen")
}
