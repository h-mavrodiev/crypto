package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	vp   *viper.Viper
	Conf Config
)

func LoadConfig(cn string, ct string, cp string) error {
	vp = viper.New()

	vp.SetConfigName(cn)
	vp.SetConfigType(ct)
	vp.AddConfigPath(cp)
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}

	err = vp.Unmarshal(&Conf)
	if err != nil {
		return err
	}

	return nil
}

// ReadInput reads input from user
func LoadConfigFromInput() error {
	var cn, ct, cp string
	var err error

	fmt.Println("path to config folder:  ")
	_, err = fmt.Scanln(&cp)
	if err != nil {
		fmt.Printf("No path was provided...\n")
		return err
	} else {
		// Check if directory exists
		_, err := os.Stat(cp)
		if err != nil {
			return err
		}
	}

	fmt.Println("config file name: ")
	_, err = fmt.Scanln(&cn)
	if err != nil {
		fmt.Printf("No name was provided...\n")
		return err
	}

	fmt.Println("config file type: ")
	_, err = fmt.Scanln(&ct)
	if err != nil {
		fmt.Printf("No type was provided...\n")
		return err
	}

	err = LoadConfig(cn, ct, cp)
	if err != nil {
		fmt.Printf("No name was provided...\n")
		return err
	}

	return nil
}
