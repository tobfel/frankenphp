package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here.
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/dunglas/caddy-cbrotli"
	_ "github.com/dunglas/frankenphp/caddy"
	_ "github.com/dunglas/mercure/caddy"
	_ "github.com/dunglas/vulcain/caddy"
	_ "github.com/abiosoft/caddy-exec"
	_ "github.com/greenpau/caddy-security"
	_ "github.com/hslatman/caddy-crowdsec-bouncer"
	_ "github.com/caddyserver/transform-encoder"
	_ "github.com/mholt/caddy-events-exec"
	_ "github.com/hslatman/caddy-crowdsec-bouncer/http@main"
    _ "github.com/hslatman/caddy-crowdsec-bouncer/appsec@main"
    _ "github.com/hslatman/caddy-crowdsec-bouncer/layer4@main"
	_ "github.com/hslatman/caddy-crowdsec-bouncer/layer4@main"
	_ "github.com/caddyserver/replace-response"
	_ "github.com/mholt/caddy-webdav"
	_ "github.com/mholt/caddy-dynamicdns"
	_ "github.com/corazawaf/coraza-caddy/v2"
    _ "github.com/porech/caddy-maxmind-geolocation"
	_ "github.com/baldinof/caddy-supervisor"
	_ "github.com/caddy-dns/desec"
	_ "github.com/sagikazarmark/caddy-fs-s3"	
)

func main() {
	caddycmd.Main()
}
