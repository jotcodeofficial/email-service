package sendgrid

import (
	"fmt"

	"github.com/spf13/viper"
)

// ConfigApplication - Here is where this microservice get the ENV valuess
type ConfigApplication struct {
	AppName     string `mapstructure:"APP_NAME"`
	AppEnv      string `mapstructure:"APP_ENV"`
	Port        string `mapstructure:"APP_PORT"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	SendGridAPI string `mapstructure:"SENDGRID_API_KEY"`
}

var config ConfigApplication

func startViperConfiguration() {
	// set defaults
	setDefaults()
	defineConfiguration()
	bindEnvReader()
}

// Don't put sensitive vaules here
func setDefaults() {
	viper.SetDefault("APP_NAME", "email-service")
	viper.SetDefault("APP_ENV", "localhost")
	viper.SetDefault("APP_PORT", "8081")
	viper.SetDefault("LOG_LEVEl", "ERROR")
}

func defineConfiguration() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("env")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")     // path to look for the config file. can have multiple lines here to search
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("Fatal error config file: %s", err))
		} else {
			// Config file was found but another error was produced
			fmt.Println("Config file was found but another error was produced")
			fmt.Println(err)
		}
	}

}

func bindEnvReader() {

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	fmt.Println("no error: ")
	fmt.Println("%+v\n", config)
}
