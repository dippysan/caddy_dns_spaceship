Spaceship module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Spaceship](https://spaceship.com/).

## Caddy module name

```
dns.providers.spaceship
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "spaceship",
        "api_key": "YOUR_API_KEY",
        "api_secret": "YOUR_API_SECRET"
      }
    }
  }
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns spaceship {
		api_key {env.API_KEY}
		api_secret {env.API_SECRET}
	}
}
```

```
# one site
tls {
	dns spaceship {
		api_key {env.API_KEY}
		api_secret {env.API_SECRET}
	}
}
```
You can replace `{env.API_KEY}`, `{env.API_SECRET}` with the actual auth credentials in the `""` if you prefer to put it directly in your config instead of an environment variable.

## Recommended additional configuration

To improve DNS propagation handling, it is recommended to add the following settings to your configuration:

https://caddyserver.com/docs/caddyfile/directives/tls#acme

```
# Maximum time to wait for DNS record propagation before timing out.
propagation_timeout 10m

# Initial delay before checking for DNS propagation.
propagation_delay 5m

# DNS resolvers to use for propagation checks (Google IPv6 DNS and Cloudflare DNS).
resolvers 2001:4860:4860::8888 1.1.1.1
```

## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/spaceship) for important information about credentials.
