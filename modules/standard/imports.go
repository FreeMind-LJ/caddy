package standard

import (
	// standard Caddy modules
	_ "github.com/freemind-lj/caddy/v2/caddyconfig/caddyfile"
	_ "github.com/freemind-lj/caddy/v2/modules/caddyhttp/standard"
	_ "github.com/freemind-lj/caddy/v2/modules/caddypki"
	_ "github.com/freemind-lj/caddy/v2/modules/caddypki/acmeserver"
	_ "github.com/freemind-lj/caddy/v2/modules/caddytls"
	_ "github.com/freemind-lj/caddy/v2/modules/caddytls/distributedstek"
	_ "github.com/freemind-lj/caddy/v2/modules/caddytls/standardstek"
	_ "github.com/freemind-lj/caddy/v2/modules/filestorage"
	_ "github.com/freemind-lj/caddy/v2/modules/logging"
	_ "github.com/freemind-lj/caddy/v2/modules/metrics"
)
