package config

func SetUser(username string, cfg *Config) error {

	cfg.CurrentUserName = username

	return write(*cfg)

}
