# TransIP module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with TransIP. In order to use this module, you must build caddy with [xcaddy](https://github.com/caddyserver/xcaddy)

```
xcaddy build --with github.com/nierot/caddy-dns-transip
```

## Caddy module name

```
dns.providers.transip
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "transip",
				"account_name": "TRANSIP_ACCOUNT_NAME",
				"private_key_path": "PATH_TO_PRIVATE_KEY"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns transip ...
}
```

```
# one site
tls {
	dns transip ...
}
```
