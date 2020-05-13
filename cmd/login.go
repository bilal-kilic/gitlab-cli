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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab-cli/utils"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Used to login gitlab-cli",
	Long:  `Used to login gitlab-cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		url  := utils.GetPromptString("url", false)
		userName := utils.GetPromptString("username", false)
		password := utils.GetPromptString("password", true)

		//initialize client with url first time logging in
		initGitlabClient()
		tokenResponse := gitlabClient.GetToken(userName, password)
		viper.Set("url", url)
		viper.Set("token", tokenResponse)
		_ = viper.SafeWriteConfig()
		println("Login Successful!")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
