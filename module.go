// Package connbyte provides a Caddy DNS provider module for ConnByte (https://connbyte.com/).
package connbyte

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsconnbyte "github.com/libdns/connbyte"
)

// Provider lets Caddy read and manipulate DNS records hosted by ConnByte.
type Provider struct{ *libdnsconnbyte.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.connbyte",
		New: func() caddy.Module { return &Provider{new(libdnsconnbyte.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
// It resolves placeholders such as {env.CONNBYTE_TOKEN}.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.APIToken = repl.ReplaceAll(p.Provider.APIToken, "")
	p.Provider.BaseURL = repl.ReplaceAll(p.Provider.BaseURL, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	connbyte [<api_token>] {
//	    api_token <api_token>
//	    base_url <url>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIToken = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "base_url":
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.BaseURL = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
