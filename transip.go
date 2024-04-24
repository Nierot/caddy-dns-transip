package transip

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnstransip "github.com/libdns/transip"
)

const TRANSIP_API_URL = "https://api.transip.nl/v6"

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnstransip.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.transip",
		New: func() caddy.Module {
			return &Provider{new(libdnstransip.Provider)} 
		},
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.AccountName    = caddy.NewReplacer().ReplaceAll(p.Provider.AccountName, "")
	p.Provider.PrivateKeyPath = caddy.NewReplacer().ReplaceAll(p.Provider.PrivateKeyPath, "")

	return nil
}

// TODO: This is just an example. Update accordingly.
// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// transip {
//     account_name <account_name>
//     private_key_path <path_to_private_key>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "account_name":
				if p.Provider.AccountName != "" {
					return d.Err("Account Name already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.AccountName = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "private_key_path":
				if p.Provider.PrivateKeyPath != "" {
					return d.Err("Private Key Path already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.PrivateKeyPath = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.AccountName == "" {
		return d.Err("missing API token")
	}
	if p.Provider.PrivateKeyPath == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
