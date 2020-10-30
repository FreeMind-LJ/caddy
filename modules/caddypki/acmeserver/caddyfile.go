// Copyright 2015 Matthew Holt and The Caddy Authors
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

package acmeserver

import (
	"github.com/freemind-lj/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/freemind-lj/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("acme_server", parseACMEServer)
}

// parseACMEServer sets up an ACME server handler from Caddyfile tokens.
//
//     acme_server [<matcher>]
//
func parseACMEServer(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var as Handler
	for h.Next() {
		if h.NextArg() {
			return nil, h.ArgErr()
		}
	}
	return as, nil
}
