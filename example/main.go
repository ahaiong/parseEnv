package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
  "github.com/ahaiong/parseEnv"
)

func main() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("the env variable from the os is", os.Getenv("env"))
	fmt.Println("the key first.second.third value from the file is", viper.GetString("first.second.third"))

  parseEnv.ReplaceForViper("__", ".")

	fmt.Println("the key first.second.third has the value", viper.GetString("first.second.third"))
}
