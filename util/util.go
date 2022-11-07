package util

import "github.com/spf13/viper"

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile("C:\\code\\Alta\\fish-hunter\\.env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return viper.GetString(key)
}