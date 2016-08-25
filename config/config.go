/*
Shows embedding anonymous fields in structs, composition.

Structs can be composed, "merging" their fields and methods. This is useful
when dealing with structs defined in different packages (libraries). My basic
use case here is having a single config file, some of the values going into
"library" config, and some going into "application" config. It would be easy
enough to parse the same json twice, once for lib config and once for app
config. But going the other way - dumping the config as a single flat json -
would be harder; this is clean.

The Dump() function and method are equivalent. Which to pick?
*/
package config

import "encoding/json"

type LibConfig struct {
	URL         string `json:"url"`
	InsecureTLS bool   `json:"insecure_tls"`
	Debug       bool   `json:"debug"`
}

type AppConfig struct {
	LibConfig
	Source string `json:"source"`
	Debug  bool   `json:"debug"` // duplicate field in LibConfig
}

func Load(b []byte) (*AppConfig, error) {
	c := new(AppConfig)
	err := json.Unmarshal(b, c)
	return c, err
}

func Dump(c *AppConfig) ([]byte, error) {
	return json.Marshal(c)
}

func (c *AppConfig) Dump() ([]byte, error) {
	return json.Marshal(c)
}
