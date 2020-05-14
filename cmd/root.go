/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab-cli/client"
	"gitlab-cli/configuration"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitlab-cli",
	Short: "A cli tool to manage gitlab operations",
	Long: `gitlab-cli is a cli tool to help you manage gitlab operation without leaving the comfort of your terminal 
such as creating merge requests, cloning projects and many more.`,
}

var gitlabClient *client.GitlabApiClient

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initGitlabClient)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitlab-cli.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initGitlabClient() {
	gitlabClient = client.NewGitlabClient(&configuration.GitlabConfiguration{
		Url:   viper.GetString("url"),
		Token: viper.GetString("token"),
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gitlab-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gitlab-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("json")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		println(err)
	}
}
