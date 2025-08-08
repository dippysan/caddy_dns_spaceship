package spaceship

import (
	"fmt"
	"strings"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/dippysan/libdns_spaceship"
)

func TestGoodConfig(t *testing.T) {
	fmt.Println("Testing Good Caddyfile Notation...")

	api_key := "api_key"
	api_secret := "api_secret"

	config := fmt.Sprintf(`spaceship {
		api_key %s
    api_secret %s
	}`, api_key, api_secret)

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&spaceship.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err != nil {
		t.Errorf("UnmarshalCaddyfile failed with %v", err)
		return
	}

	{
		expected := api_key
		actual := p.Provider.APIKey
		if expected != actual {
			t.Errorf("Expected APIKey to be '%s' but got '%s'", expected, actual)
		}
	}

	{
		expected := api_secret
		actual := p.Provider.APISecret
		if expected != actual {
			t.Errorf("Expected APISecret to be '%s' but got '%s'", expected, actual)
		}
	}
}

func TestBadConfig(t *testing.T) {
	fmt.Println("Testing Bad Caddyfile Notation...")

	tests := []struct{ config, expected string }{
		// invalid syntax
		{"spaceship api_key", "syntax error: unexpected token 'api_key', expecting '{}'"},
		// missing required credentials
		{`spaceship {
      api_key api_key
    }`, "missing one or more required api credentials: api_key or api_secret"},
		// unrecognized subdirective
		{`spaceship {
      api_key api_key
      api_secret api_secret
      another_credential another_credential
    }`, "unrecognized subdirective 'another_credential'"},
	}

	for _, test := range tests {
		dispenser := caddyfile.NewTestDispenser(test.config)
		p := Provider{&spaceship.Provider{}}

		err := p.UnmarshalCaddyfile(dispenser)
		if err == nil || !strings.Contains(strings.ToLower(err.Error()), test.expected) {
			t.Errorf("Expected error message to be '%s' but got '%s'", test.expected, err.Error())
		}
	}

}
