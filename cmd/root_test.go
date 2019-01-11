// Copyright © 2019 openSUSE opensuse-project@opensuse.org
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

package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

const (
	configContent = `---
apiVersion: kubic.suse.com/v1alpha2
kind: KubicInitConfiguration
features:
  PSP: true
runtime:
  engine: crio
paths:
  kubeadm: /usr/bin/kubeadm
auth:
  oidc:
   issuer: https://some.name.com:32000
   clientID: kubernetes
   ca: /etc/kubernetes/pki/ca.crt
   username: email
   groups: groups
certificates:
  directory: /etc/kubernetes/pki
  caCrt:
  caCrtHash:
etcd:
  local:
    serverCertSANs: []
    peerCertSANs: []
manager:
  image: "kubic-init:latest"
clusterFormation:
  seeder: some-node.com
  token: 94dcda.c271f4ff502789ca
  autoApprove: false
network:
  bind:
    address: 0.0.0.0
    interface: eth0
  podSubnet: "172.16.0.0/13"
  serviceSubnet: "172.24.0.0/16"
  proxy:
    http: my-proxy.com:8080
    https: my-proxy.com:8080
    noProxy: localdomain.com
    systemwide: false
  dns:
    domain: someDomain.local
    externalFqdn: some.name.com
  cni:
    driver: flannel
    image: registry.opensuse.org/devel/caasp/kubic-container/container/kubic/flannel:0.9.1
bootstrap:
  registries:
    - prefix: https://registry.suse.com
      mirrors:
        - url: https://airgapped.suse.com
        - url: https://airgapped2.suse.com
          certificate: "-----BEGIN CERTIFICATE-----
MIIGJzCCBA+gAwIBAgIBATANBgkqhkiG9w0BAQUFADCBsjELMAkGA1UEBhMCRlIx
DzANBgNVBAgMBkFsc2FjZTETMBEGA1UEBwwKU3RyYXNib3VyZzEYMBYGA1UECgwP
hnx8SB3sVJZHeer8f/UQQwqbAO+Kdy70NmbSaqaVtp8jOxLiidWkwSyRTsuU6D8i
DiH5uEqBXExjrj0FslxcVKdVj5glVcSmkLwZKbEU1OKwleT/iXFhvooWhQ==
-----END CERTIFICATE-----"
          fingerprint: "E8:73:0C:C5:84:B1:EB:17:2D:71:54:4D:89:13:EE:47:36:43:8D:BF:5D:3C:0F:5B:FC:75:7E:72:28:A9:7F:73"
          hashalgorithm: "SHA256"
    - prefix: https://registry.io
      mirrors:
        - url: https://airgapped.suse.com
        - url: https://airgapped2.suse.com
          certificate: "-----BEGIN CERTIFICATE-----
MIIGJzCCBA+gAwIBAgIBATANBgkqhkiG9w0BAQUFADCBsjELMAkGA1UEBhMCRlIx
DzANBgNVBAgMBkFsc2FjZTETMBEGA1UEBwwKU3RyYXNib3VyZzEYMBYGA1UECgwP
hnx8SB3sVJZHeer8f/UQQwqbAO+Kdy70NmbSaqaVtp8jOxLiidWkwSyRTsuU6D8i
DiH5uEqBXExjrj0FslxcVKdVj5glVcSmkLwZKbEU1OKwleT/iXFhvooWhQ==
-----END CERTIFICATE-----"
          fingerprint: "E8:73:0C:C5:84:B1:EB:17:2D:71:54:4D:89:13:EE:47:36:43:8D:BF:5D:3C:0F:5B:FC:75:7E:72:28:A9:7F:73"
          hashalgorithm: "SHA256"
`
	configContentWithError = `---
bootstrap:
  registries:
    - 
    prefix: https://registry.suse.com
    mirrors:
        - url: https://airgapped.suse.com
        - url: https://airgapped2.suse.com
          certificate: "-----BEGIN CERTIFICATE-----
MIIGJzCCBA+gAwIBAgIBATANBgkqhkiG9w0BAQUFADCBsjELMAkGA1UEBhMCRlIx
DzANBgNVBAgMBkFsc2FjZTETMBEGA1UEBwwKU3RyYXNib3VyZzEYMBYGA1UECgwP
hnx8SB3sVJZHeer8f/UQQwqbAO+Kdy70NmbSaqaVtp8jOxLiidWkwSyRTsuU6D8i
DiH5uEqBXExjrj0FslxcVKdVj5glVcSmkLwZKbEU1OKwleT/iXFhvooWhQ==
-----END CERTIFICATE-----"
          fingerprint: "E8:73:0C:C5:84:B1:EB:17:2D:71:54:4D:89:13:EE:47:36:43:8D:BF:5D:3C:0F:5B:FC:75:7E:72:28:A9:7F:73"
          hashalgorithm: "SHA256"
    - prefix: https://registry.io
      mirrors:
        - url: https://airgapped.suse.com
        - url: https://airgapped2.suse.com
          certificate: "-----BEGIN CERTIFICATE-----
MIIGJzCCBA+gAwIBAgIBATANBgkqhkiG9w0BAQUFADCBsjELMAkGA1UEBhMCRlIx
DzANBgNVBAgMBkFsc2FjZTETMBEGA1UEBwwKU3RyYXNib3VyZzEYMBYGA1UECgwP
hnx8SB3sVJZHeer8f/UQQwqbAO+Kdy70NmbSaqaVtp8jOxLiidWkwSyRTsuU6D8i
DiH5uEqBXExjrj0FslxcVKdVj5glVcSmkLwZKbEU1OKwleT/iXFhvooWhQ==
-----END CERTIFICATE-----"
          fingerprint: "E8:73:0C:C5:84:B1:EB:17:2D:71:54:4D:89:13:EE:47:36:43:8D:BF:5D:3C:0F:5B:FC:75:7E:72:28:A9:7F:73"
          hashalgorithm: "SHA256"
`
	filename           = "kubic-init.mirrors.yaml"
	filenameWithErrors = "kubic-init.errors.yaml"
)

func Test_runE(t *testing.T) {
	c := &cobra.Command{}
	err := ioutil.WriteFile(filename, []byte(configContent), os.FileMode(0644))
	if err != nil {
		t.Fatalf("faliled to write config file: %s", err)
	}
	err = ioutil.WriteFile(filenameWithErrors, []byte(configContentWithError), os.FileMode(0644))
	if err != nil {
		t.Fatalf("faliled to write config file: %s", err)
	}
	tmpDir, err := ioutil.TempDir("", "caasp-init-certs")
	if err != nil {
		t.Fatalf("creating tmp dir: %s", err)
	}
	defer os.RemoveAll(tmpDir)
	defer os.RemoveAll(filename)
	defer os.RemoveAll(filenameWithErrors)
	type args struct {
		cmd        *cobra.Command
		args       []string
		folder     string
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{c, []string{}, tmpDir, filename}, true},
		{"2", args{c, []string{}, tmpDir, filenameWithErrors}, true},
		{"3", args{c, []string{}, "/sys/temp", filenameWithErrors}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registryConfigFolder = tt.args.folder
			cfgFile = tt.args.configFile
			if err := runE(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("runE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
