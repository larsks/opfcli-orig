package cmd

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/operate-first/opfcli/utils"
)

var config = viper.New()

var cfgFile string
var appName string
var repoDirectory string

var rootCmd = &cobra.Command{
	Use:   "opfcli",
	Short: "A command line tool for Operate First GitOps",
	Long: `A command line tool for Operate First GitOps.

Use opfcli to interact with an Operate First style Kubernetes
configuration repository.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	loglevel, err := strconv.Atoi(os.Getenv("OPF_LOGLEVEL"))
	if err != nil {
		loglevel = 1
	}

	if loglevel >= 2 {
		log.SetLevel(log.DebugLevel)
	} else if loglevel >= 1 {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(
		&cfgFile, "config", "f", "", "configuration file")
	rootCmd.PersistentFlags().StringVarP(
		&appName, "app-name", "a", "", "application name")

	config.BindPFlag("app-name", rootCmd.PersistentFlags().Lookup("app-name"))
}

func initConfig() {
	config.SetEnvPrefix("opf")
	config.AutomaticEnv()
	config.SetConfigName(".opfcli")

	replacer := strings.NewReplacer("-", "_")
	config.SetEnvKeyReplacer(replacer)

	config.SetDefault("app-name", defaultAppName)

	if cfgFile != "" {
		config.SetConfigFile(cfgFile)
		config.ReadInConfig()
	} else {
		home, err := homedir.Dir()
		if err == nil {
			config.AddConfigPath(home)
		}

		repoDirectory, err = utils.FindRepoDir()
		if err != nil {
			repoDirectory, err = filepath.Abs(".")
			if err != nil {
				log.Fatalf("failed to determine repository directory: %v", err)
			}
		}
		log.Debugf("using %s as repository directory", repoDirectory)

		config.AddConfigPath(repoDirectory)
		config.ReadInConfig()
		if config.ConfigFileUsed() != "" {
			log.Printf("read configuration from %s", config.ConfigFileUsed())
		}
	}
}
