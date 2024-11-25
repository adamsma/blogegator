package config

func SetUser(username string, cfg *Config) error {

	cfg.Current_User_Name = username
	write(*cfg)
	return nil

}
