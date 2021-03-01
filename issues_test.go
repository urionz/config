package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urionz/config"
	"github.com/urionz/config/yaml"
	"github.com/urionz/config/yamlv3"
	"github.com/urionz/goutil/dump"
)

// https://github.com/gookit/config/issues/37
func TestIssues37(t *testing.T) {
	is := assert.New(t)

	c := config.New("test")
	c.AddDriver(yaml.Driver)

	err := c.LoadStrings(config.JSON, `
{
    "lang": {
        "allowed": {
            "en": "ddd"
        }
    }
}
`)
	is.NoError(err)
	dump.Println(c.Data())

	is.Panics(func() {
		_ = c.LoadStrings(config.Yaml, `
lang:
  allowed:
    en: "666"
`)
	})
}

// https://github.com/gookit/config/issues/37
func TestIssues37_yaml_v3(t *testing.T) {
	is := assert.New(t)

	c := config.New("test")
	c.AddDriver(yamlv3.Driver)

	err := c.LoadStrings(config.JSON, `
{
    "lang": {
        "allowed": {
            "en": "ddd"
        }
    }
}
`)
	is.NoError(err)
	dump.Println(c.Data())

	err = c.LoadStrings(config.Yaml, `
lang:
  newKey: hhh
  allowed:
    en: "666"
`)
	is.NoError(err)
	dump.Println(c.Data())
}
