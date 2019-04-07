package main

type SymfonyJsonProxy struct {
	Tld     string            `json:"tld"`
	Port    int               `json:"port"`
	Domains map[string]string `json:"domains"`
}

type PomdokYamlConfig struct {
	Pomdok struct {
		Tld      string
		Projects []struct {
			Domain string `yaml:"domain"`
			Path   string `yaml:"path"`
		}
	}
}
