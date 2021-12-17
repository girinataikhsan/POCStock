package configuration

// Configuration stores global configuration loaded from json file
type ConfigurationApp struct {
	App
	Database
	Redis
}

type App struct {
	ListenPort    string `split_words:"true" default:":9500"`
	AppName       string `split_words:"true" default:"Stock"`
	Version       string `split_words:"true" default:"0.0.1"`
	RootURL       string `split_words:"true" default:"/v1/stock"`
	Timeout       int64  `split_words:"true" default:"60000"`

}

type Redis struct {
	Host                         string   `split_words:"true" default:"127.0.0.1:6379"`
	DB                     		 int   `split_words:"true" default:"0"`
	Password                     string   `split_words:"true" default:""`
}

type Database struct {
	Host                         string   `split_words:"true" default:"10.0.148.216"`
	Port                         int      `split_words:"true" default:"5432"`
	Username                     string   `split_words:"true" default:"postgres"`
	Password                     string   `split_words:"true" default:"Smartfren2021"`
	Name                         string   `split_words:"true" default:"product_catalog"`
	SSLMode                      string   `split_words:"true" default:"disable"`
	LogMode                      bool     `split_words:"true" default:"true"`
}

var ConfigApp App
