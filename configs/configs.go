package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var vp *viper.Viper

func LoadConfig(cn string, ct string, cp string) (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName(cn)
	vp.SetConfigType(ct)
	vp.AddConfigPath(cp)
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// ReadInput reads input from user
func LoadConfigFromInput() (Config, error) {
	var cn, ct, cp string
	var err error

	fmt.Println("path to config folder:  ")
	_, err = fmt.Scanln(&cp)
	if err != nil {
		fmt.Printf("No path was provided...\n")
		log.Fatal(err)
	} else {
		// Check if directory exists
		_, err := os.Stat(cp)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("config file name: ")
	_, err = fmt.Scanln(&cn)
	if err != nil {
		fmt.Printf("No name was provided...\n")
		log.Fatal(err)
	}

	fmt.Println("config file type: ")
	_, err = fmt.Scanln(&ct)
	if err != nil {
		fmt.Printf("No type was provided...\n")
		log.Fatal(err)
	}

	conf, err := LoadConfig(cn, ct, cp)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
