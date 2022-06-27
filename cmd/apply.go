/*
 * Copyright contributors to the Hyperledgendary Fabric Sail project
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 * 	  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/hyperledgendary/fabric-sail/sail"
	"log"

	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a Fabric sail",
	Long: `Construct a Fabric network by applying a sail configuration file to a 
Kubernetes cluster.  The sail will launch fabric-operator in a target namespace,
creating peers, orderers, CAs, channels, and chaincode as specified in the sail 
configuration document.  Resources on the cluster will be created if they do 
not already exist, and additional sail applications will apply a delta update to 
minimize the scope of changes required to the cluster.

After applying a sail, a local directory will be created with certificates, 
connection profiles, and metadata suitable for running Fabric CLI binaries against 
the generated network.

Examples: 
  # create the fabric-samples test-network, mychannel, and asset transfer chaincode:
  sail apply -f samples/asset-transfer-basic.yaml

  # Apply a sail configuration via stdin:
  cat sail.yaml | sail apply -f -
`,
	Run: func(cmd *cobra.Command, args []string) {

		filename, _ := cmd.Flags().GetString("filename")

		log.Printf("Applying sail from configuration file %s\n", filename)

		sail, err := sail.Loft(filename)
		if err != nil {
			log.Fatalf("Could not load sail from configuration %s: %s\n", filename, err)
		}

		if err = sail.Sail(); err != nil {
			log.Fatalf("error: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	applyCmd.Flags().StringP("filename", "f", "", "that contains the configuration to apply")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
