ConnByte module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [ConnByte](https://connbyte.com/), for example to get wildcard certificates with the ACME DNS challenge.

It uses [libdns/connbyte](https://github.com/libdns/connbyte) and the [ConnByte DNS API](https://api.connbyte.com/).

## Caddy module name

```
dns.providers.connbyte
```

## Build

```
xcaddy build --with github.com/caddy-dns/connbyte
```

## Config examples

Create an API token in your [ConnByte](https://connbyte.com/) account.

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "connbyte",
				"api_token": "{env.CONNBYTE_TOKEN}"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns connbyte {env.CONNBYTE_TOKEN}
}
```

```
# one site
*.example.com {
	tls {
		dns connbyte {env.CONNBYTE_TOKEN}
	}
}
```

Block syntax:

```
dns connbyte {
	api_token {env.CONNBYTE_TOKEN}
}
```
