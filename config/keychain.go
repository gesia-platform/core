package config

type Keychain struct {
	Host string `toml:"host"`

	Password string `toml:"password"`
}
