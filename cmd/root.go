package cmd

import (
    "log"
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

    config.SetDefault("app-name", DEFAULT_APP_NAME)

	if cfgFile != "" {
		config.SetConfigFile(cfgFile)
        config.ReadInConfig()
    } else {
		home, err := homedir.Dir()
        if err == nil {
            config.AddConfigPath(home)
        }

        repoDirectory, err = utils.FindRepoDir()
        if err == nil {
            config.AddConfigPath(repoDirectory)
        }

        config.ReadInConfig()
        log.Printf("read configuration from %s", config.ConfigFileUsed())
    }
}
