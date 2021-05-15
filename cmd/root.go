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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "opfcli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(
        &cfgFile, "config", "f", "", "configuration file")
    rootCmd.PersistentFlags().StringVarP(
        &appName, "app-name", "a", "", "application name")

    config.BindPFlag("app-name", rootCmd.PersistentFlags().Lookup("app-name"))
}

// initConfig reads in config file and ENV variables if set.
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

        repodir, err := utils.FindRepoDir()
        if err == nil {
            config.AddConfigPath(repodir)
        }

        config.ReadInConfig()
        log.Printf("read configuration from %s", config.ConfigFileUsed())
    }
}
