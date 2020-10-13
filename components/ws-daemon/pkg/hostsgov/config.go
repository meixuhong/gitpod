// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package hostsgov

// Config configures the hosts governer
type Config struct {
	NodeHostsFile string            `json:"nodeHostsFile"`
	FromNodeIPs   map[string]string `json:"fromPodNodeIP"`
	ServiceProxy  struct {
		Enabled     bool `json:"enabled,omitempty"`
		PortMapping []struct {
			Selector  string `json:"selector"`
			Alias     string `json:"alias"`
			ProxyPort int    `json:"proxyPort"`
		} `json:"mapping"`
	} `json:"serviceProxy,omitempty"`
}
