package config

type Config struct{}

var conf Config

func Apply(c Config) {
	conf = c
}

func Get() Config {
	return conf
}
