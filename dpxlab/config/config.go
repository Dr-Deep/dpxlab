package config

import _ "gopkg.in/yaml.v3"

type Configuration struct {
	WWW struct {
		Socket string
	}

	Controller struct{}

	Engine struct{}

	Storage struct{}

	Database struct{}

	Logger struct {
		File  string
		Level string
	}
}
