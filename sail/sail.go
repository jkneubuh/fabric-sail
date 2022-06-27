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

package sail

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Sail struct {
	Network Network `json:"network,omitempty"`
}

type Network struct {
	Name          string         `json:"name,omitempty"`
	Domain        string         `json:"domain,omitempty"`
	Namespace     string         `json:"namespace,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
	Channels      []Channel      `json:"channels,omitempty"`
	Chaincode     []Chaincode    `json:"chaincode,omitempty"`
}

type Organization struct {
	Name     string    `json:"name,omitempty"`
	Orderers []Orderer `json:"orderers,omitempty"`
	Peers    []Peer    `json:"peers,omitempty"`
}

type Orderer struct {
	Name  string `json:"name,omitempty"`
	Count int32  `json:"count,omitempty"`
}

type Peer struct {
	Name   string `json:"name,omitempty"`
	Anchor bool   `json:"anchor,omitempty"`
}

type Channel struct {
	Name          string   `json:"name,omitempty"`
	Organizations []string `json:"organizations,omitempty"`
}

type Chaincode struct {
	Name     string           `json:"name,omitempty"`
	Version  string           `json:"version,omitempty"`
	Package  string           `json:"package,omitempty"`
	Channels []ChannelBinding `json:"channels,omitempty"`
}

type ChannelBinding struct {
	Name              string `json:"name,omitempty"`
	Policy string `json:"policy,omitempty"`
}

func Loft(filename string) (*Sail, error) {

	log.Printf("Lofting sail %v\n", filename)

	rawbytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading %v: %v", filename, err)
	}

	sail := &Sail{}
	err = yaml.Unmarshal(rawbytes, &sail)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return sail, nil
}

func (sail *Sail) Sail() error {
	log.Printf("ahoy!\n")

	d, err := yaml.Marshal(&sail)
	if err != nil {
		return err
	}

	log.Printf("Sail config:\n%v\n\n", string(d))

	return nil
}
