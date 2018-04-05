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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	// Config is the general configuration object
	Config RapidConfig

	// CfgFile is the path to the config file
	CfgFile string
)

func init() {
	// make sure the config is initialised first
	InitConfig()
}

// RapidConfig holds all the configuration blocks.
// These are bound from cli, Environment variables or configuration files by
// Viper.
type RapidConfig struct {
	WildcardTenant string   `json:"wildcardOrg"`
	EnabledTenants []string `json:"enableOrgIDs"`
	HTTP           `json:"http"`
	ElasticSearch  `json:"elasticsearch"`
	DevMode        bool `json:"devmode"`
	Tenants        map[string]*Tenant
}

// ElasticSearch holds all the configuration values
// It is bound by Viper.
type ElasticSearch struct {
	Urls        []string `json:"urls"`
	IndexName   string   `json:"index"`
	Proxy       bool     `json:"proxy"`
	EnableV1    bool     `json:"enableV1"` // Enable v1 style index
	EnableTrace bool     `json:"enableTrace"`
}

// HTTP holds all the configuration for the http server subcommand
type HTTP struct {
	Port int `json:"port" mapstructure:"port"`
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "rapid" (without extension).
		viper.AddConfigPath("/etc/default/")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("rapid")
	}

	viper.SetEnvPrefix("RAPID")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	setDefaults()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal(
			fmt.Sprintf("unable to decode into struct, %v", err),
		)
	}

	//for _, tenantID := range Config.EnabledTenants {
	//tenantConfig := viper.Sub(tenantID)
	//_, err := NewTenant(tenantConfig)
	//if err != nil {
	//log.Fatal(
	//fmt.Sprintf("Unable to decode subTenant\n", tenantID),
	//)
	//}
	//Config.Tenants[tenantID] = tenant
	//}

}

// set the RapidConfig default values
func setDefaults() {

	// setting defaults
	viper.SetDefault("HTTP.port", 3001)
	viper.SetDefault("WildcardTenant", "all")
	viper.SetDefault("EnabledTenants", []string{"all", "default"})
	viper.SetDefault("DevMode", false)

	// elastic
	viper.SetDefault("ElasticSearch.urls", []string{"http://localhost:9200"})
	viper.SetDefault("ElasticSearch.enabled", true)
	viper.SetDefault("ElasticSearch.IndexName", "hub3-rapid")
	viper.SetDefault("ElasticSearch.Proxy", false)
	viper.SetDefault("ElasticSearch.IndexV1", false)
	viper.SetDefault("ElasticSearch.EnableTrace", false)
}

// BuildVersionInfo holds all the version information
type BuildVersionInfo struct {
	Version    string `json:"version"`
	Commit     string `json:"commit"`
	BuildAgent string `json:"buildAgent"`
	BuildDate  string `json:"buildDate"`
}

// NewBuildVersionInfo creates a BuildVersionInfo struct
func NewBuildVersionInfo(version, commit, buildagent, builddate string) *BuildVersionInfo {
	if version == "" {
		version = "devBuild"
	}
	return &BuildVersionInfo{
		Version:    version,
		Commit:     commit,
		BuildAgent: buildagent,
		BuildDate:  builddate,
	}
}

// JSON returns a json version of the BuildVersionInfo
func (b BuildVersionInfo) JSON(pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(b, "", "\t")
	}
	return json.Marshal(b)
}
