package config


type Config struct {
	ServerAddr string 
	AssetsFS   string
}

func NewConfig() Config {
	return Config{
		ServerAddr: ":8080",
		AssetsFS:  "../web/public/assets",
	}
}