package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logrus.New()

// Config is the central configuration object
var Config Configuration

// Execute -- application entrypoint
func Execute() error {

	var rootCmd = &cobra.Command{
		Use:              Config.Binaryname,
		Short:            Config.Shortdescription,
		Long:             Config.Longdescription,
		PersistentPreRun: persistentPreRun,
		Run: func(cmd *cobra.Command, args []string) {
			log.Debug("command running...")
		},
	}

	level, err := logrus.ParseLevel(viper.GetString("loglevel"))
	if err != nil {
		log.Fatalf("Error parsing %s as valid logrus.LogLevel", viper.GetString("loglevel"))
	}

	log.SetLevel(level)

	return rootCmd.Execute()

}

func init() {
	initConfig()
}

func initConfig() {

	// Look for a config file in the working directory
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())

		marshalErr := viper.Unmarshal(&Config)
		if marshalErr != nil {
			log.Fatal("Unable to decode config into struct")
		}

		log.Debug("Read config file successfully")

	} else {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Error reading config file")
	}
}

// Perform any validation on the presence / type of require Config attributes in this function
func persistentPreRun(cmd *cobra.Command, args []string) {

}
