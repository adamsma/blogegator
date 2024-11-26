package config

func (cfg *Config) SetUser(username string) error {

	cfg.CurrentUserName = username

	return write(*cfg)

}
