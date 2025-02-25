package httptest

import "github.com/crossplane/upjet/pkg/config"

// Configure customizes individual resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("thousandeyes_http_server", func(r *config.Resource) {
		r.ShortGroup = "httptest"
	})
}
