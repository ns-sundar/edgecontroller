// Copyright 2019 Intel Corporation. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cnca

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var tac int

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register controller to AF services registry",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var s LocationService
		s.TAC = tac
		s.DNAI, _ = cmd.Flags().GetString("dnai")
		s.DNN, _ = cmd.Flags().GetString("dnn")
		s.PriDNS, _ = cmd.Flags().GetString("priDns")
		s.SecDNS, _ = cmd.Flags().GetString("secDns")
		s.UPFIP, _ = cmd.Flags().GetString("upfIp")
		s.SNSSAI, _ = cmd.Flags().GetString("snssai")

		srv, err := json.Marshal(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		// register AF service
		afID, err := OAM5gRegisterAFService(srv)
		if err != nil {
			klog.Info(err)
			return
		}

		fmt.Printf("Service `%s` registered successfully\n", afID)
	},
}

func init() {

	const help = `Register controller to NGC AF services registry

Usage:
  cnca register --dnai=<DNAI> --dnn=<DNN> --priDns=<pri-DNS> --secDns=<sec-DNS> --upfIp=<UPF-IP> --snssai=<SNSSAI>

Flags:
  -h, --help       help
      --dnai       Identifies DNAI
      --dnn        Identifies data network name
      --tac        Identifies Tracking Area Code (TAC)
      --priDns     Identifies primary DNS
      --secDns     Identifies secondary DNS
      --upfIp      Identifies UPF IP address
      --snssai     Identifies SNSSAI"
`
	// add `register` command
	cncaCmd.AddCommand(registerCmd)
	registerCmd.Flags().String("dnai", "", "Identifies DNAI")
	registerCmd.MarkFlagRequired("dnai")
	registerCmd.Flags().String("dnn", "", "Identifies data network name")
	registerCmd.MarkFlagRequired("dnn")
	registerCmd.Flags().IntVar(&tac, "tac", 0, "Identifies Tracking Area Code (TAC)")
	registerCmd.MarkFlagRequired("tac")
	registerCmd.Flags().String("priDns", "", "Identifies primary DNS")
	registerCmd.MarkFlagRequired("priDns")
	registerCmd.Flags().String("secDns", "", "Identifies secondary DNS")
	registerCmd.MarkFlagRequired("secDns")
	registerCmd.Flags().String("upfIp", "", "Identifies UPF IP address")
	registerCmd.MarkFlagRequired("upfIp")
	registerCmd.Flags().String("snssai", "", "Identifies SNSSAI")
	registerCmd.MarkFlagRequired("snssai")
	registerCmd.SetHelpTemplate(help)
}
