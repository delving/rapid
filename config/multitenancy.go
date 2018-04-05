// Copyright © 2017 Delving B.V. <info@delving.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// TenantConfig holds the configuration for each enabled organisation
// This data is loaded from support TOML files
type Tenant struct {
	OrgID string `json:"orgID"` // organisation ID
	LOD   `json:"lod"`
}

// LOD holds all the configuration for the Linked Open Data (LOD) functionality
type LOD struct {
	// make the lod endpoint available
	enabled bool `json:"enabled"`

	// the 303 redirect entry point.  This is where the content negotiation happens
	Resource string `json:"resource"`

	// the endpoint that renders the data as formatted HTML
	HTML string `json:"html"`

	// the endpoint that renders the RDF data in the requested RDF format. Currently, JSON-LD and N-triples are supported
	RDF string `json:"rdf"`

	// the regular expression to convert the subject uri to the uri for the external Page view
	HTMLRedirectRegex string `json:"redirectregex"`
}

// NewTenant creates a new instance of the Tenant
func NewTenant(cfg *viper.Viper) (*Tenant, error) {
	// setDefaults
	setTenantDefaults(cfg)

	// create tenant
	var t Tenant
	err := cfg.Unmarshal(&t)
	if err != nil {
		return nil, err
	}
	if t.OrgID == "" {
		return nil, fmt.Errorf("unable to find Tenant Configuration.\n")
	}
	return &t, nil
}

// setTenantDefaults sets the sub-tenants defaults
func setTenantDefaults(cfg *viper.Viper) {

	// lod
	cfg.SetDefault("LOD.enabled", true)
	cfg.SetDefault("LOD.html", "page")
	cfg.SetDefault("LOD.rdf", "data")
	cfg.SetDefault("LOD.resource", "resource")
	cfg.SetDefault("LOD.redirectregex", "")
}
