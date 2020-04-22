package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// Set the application configuration file name
	viper.SetConfigName("appConfig")

	viper.SetConfigType("yml")

	// Add the folder path for the application configuration file
	viper.AddConfigPath(".")

	// Viper package to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Extract configuration values using Viper (ignore config.go)
	fmt.Println("Extract configuration values using Viper (ignore config.go)")
	fmt.Println("Database is\t", viper.GetString("database.hostname"))
	fmt.Println("DB User is\t", viper.GetString("database.dbuser"))
	fmt.Println("DB Password is\t", viper.GetString("database.dbpassword"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	fmt.Println("Max Workers is\t", viper.GetInt("maxworkers"))
	fmt.Println("Max Queues is\t", viper.GetInt("maxqueues"))

	// Set default values for one of the parameters
	viper.SetDefault("database.hostname", "db_host_name")

	var configValues Configurations

	err := viper.Unmarshal(&configValues)
	if err != nil {
		fmt.Printf("Unmarshalling issue with configValues - Viper: , %v", err)
	}

	// Extract configuration values using Viper and Configurations struct (config.go)
	fmt.Println("\nExtract configuration values using Viper and Configurations struct (config.go):")
	fmt.Println("Database is\t", configValues.Database.HostName)
	fmt.Println("DB User is\t", configValues.Database.DBUser)
	fmt.Println("DB Password is\t", configValues.Database.DBPassword)
	fmt.Println("Port is\t\t", configValues.Server.Port)
	fmt.Println("Max Workers is\t", configValues.MaxWorkers)
	fmt.Println("Max Queues is\t", configValues.MaxQueues)

}
