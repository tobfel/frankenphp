package main

// OPC-spezifische Caddy-Module. Bewusst eigene Datei (statt in main.go), damit
// main.go identisch zu php/frankenphp bleibt und Upstream-Merges nicht kollidieren.
// cbrotli/mercure/vulcain kommen bereits ueber die Upstream-Dateien cbrotli.go etc.
import (
	_ "github.com/abiosoft/caddy-exec"
	_ "github.com/baldinof/caddy-supervisor"
	_ "github.com/caddy-dns/desec"
	//_ "github.com/caddyserver/replace-response"
	_ "github.com/caddyserver/transform-encoder"
	//_ "github.com/corazawaf/coraza-caddy/v2"
	_ "github.com/anujc4/caddy-geoblock"
	_ "github.com/fabriziosalmi/caddy-mib"
	_ "github.com/fabriziosalmi/caddy-waf"
	_ "github.com/hslatman/caddy-crowdsec-bouncer"
	_ "github.com/hslatman/caddy-crowdsec-bouncer/appsec"
	_ "github.com/hslatman/caddy-crowdsec-bouncer/http"
	_ "github.com/hslatman/caddy-crowdsec-bouncer/layer4"
	_ "github.com/jpillora/ipfilter-caddy"
	_ "github.com/mholt/caddy-dynamicdns"
	_ "github.com/mholt/caddy-events-exec"
	_ "github.com/mholt/caddy-ratelimit"
	_ "github.com/mholt/caddy-webdav"
	_ "github.com/sagikazarmark/caddy-fs-s3"
	_ "github.com/steffenbusch/caddy-postauth-2fa"
)
