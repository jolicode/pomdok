package main

// SymfonyJSONProxy is a representation of Symfony CLI's configuration file
type SymfonyJSONProxy struct {
	Tld     string            `json:"tld"`
	Port    int               `json:"port"`
	Domains map[string]string `json:"domains"`
	Ports   map[string]int    `json:"ports"`
}

// PomdokYamlConfig is a representation of Pomdok's configuration file
type PomdokYamlConfig struct {
	Pomdok struct {
		Tld      string
		Projects []struct {
			Domain string `yaml:"domain"`
			Path   string `yaml:"path"`
			Port   int    `yaml:"port"`
		}
	}
}
