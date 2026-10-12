package connbyte

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsconnbyte "github.com/libdns/connbyte"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		token   string
		baseURL string
		wantErr bool
	}{
		{name: "inline token", input: `connbyte secret`, token: "secret"},
		{name: "block token", input: "connbyte {\n api_token secret\n base_url http://127.0.0.1:8080\n}", token: "secret", baseURL: "http://127.0.0.1:8080"},
		{name: "env placeholder", input: `connbyte {env.CONNBYTE_TOKEN}`, token: "{env.CONNBYTE_TOKEN}"},
		{name: "missing token", input: `connbyte`, wantErr: true},
		{name: "token twice", input: "connbyte a {\n api_token b\n}", wantErr: true},
		{name: "too many args", input: `connbyte a b`, wantErr: true},
		{name: "unknown subdirective", input: "connbyte {\n foo bar\n}", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Provider{new(libdnsconnbyte.Provider)}
			err := p.UnmarshalCaddyfile(caddyfile.NewTestDispenser(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && (p.APIToken != tt.token || p.BaseURL != tt.baseURL) {
				t.Fatalf("got token %q base %q", p.APIToken, p.BaseURL)
			}
		})
	}
}

func TestModuleID(t *testing.T) {
	if id := (Provider{}).CaddyModule().ID; id != "dns.providers.connbyte" {
		t.Fatalf("module ID = %s", id)
	}
}
