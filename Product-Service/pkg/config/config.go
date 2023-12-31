package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {

	// database
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

var envs = []string{

	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", //database
	// etc...
}

var config Config // create instence of config

// func to get env variable and store it on struct Config and retuen it with error as nil or error
func LoadConfig() (Config, error) {

	// vipen setup
	viper.AddConfigPath("./")   // add config path
	viper.SetConfigFile(".env") //setup file name to viper
	viper.ReadInConfig()        // read .env file

	// range through the envNames and take each envName and bind that env variable to viper

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err // error handling
		}
	}

	// then unmarshel the viper into config variable

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	// atlast validate the config file using validator pakage
	// create instance and direct validte
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	//successfully loaded the env values into struct config
	return config, nil
}

func GetCofig() Config {
	return config

}

