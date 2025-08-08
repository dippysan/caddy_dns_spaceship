package spaceship

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/dippysan/libdns_spaceship"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *spaceship.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.spaceship",
		New: func() caddy.Module { return &Provider{new(spaceship.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIKey = caddy.NewReplacer().ReplaceAll(p.Provider.APIKey, "")
	p.Provider.APISecret = caddy.NewReplacer().ReplaceAll(p.Provider.APISecret, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	spaceship {
//	    api_key <api_key>
//	    api_secret <api_secret>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.SyntaxErr("{}")
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if p.Provider.APIKey != "" {
					return d.Err("APIKey already set")
				}
				if d.NextArg() {
					p.Provider.APIKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_secret":
				if p.Provider.APISecret != "" {
					return d.Err("APISecret already set")
				}
				if d.NextArg() {
					p.Provider.APISecret = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIKey == "" || p.Provider.APISecret == "" {
		return d.Err("missing one or more required API credentials: api_key or api_secret")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
