/*
Copyright © 2020 NAME HERE <bkilic61@gmail.com>

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
	"errors"
	"github.com/spf13/cobra"
	"gitlab-cli/cmd/models"
	"gitlab-cli/utils"
)

var localFlags = []models.Flag{
	{Name: "sourceBranch", Shorthand: "s", Description: "Branch to merge.", Value: "", Required: true},
	{Name: "targetBranch", Shorthand: "t", Description: "Branch to merge source branch", Value: "", Required: true},
	{Name: "title", Shorthand: "i", Description: "Title of merge request.", Value: "", Required: true},
	{Name: "description", Shorthand: "d", Description: "Description of merge request.", Value: "", Required: false},
}

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Use to merge given branches with options",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error{
		var createMergeRequest models.CreateMergeRequest
		projectName, err := utils.GetProjectName()

		if err != nil {
			return err
		}

		projects, err := gitlabClient.GetProjects()

		if err != nil {
			return err
		}

		var selectedProject models.ProjectResponse
		for _, project := range projects {
			if project.Path == projectName {
				selectedProject = project
			}
		}

		if selectedProject.Name == "" {
			return errors.New("Project not found. Either it does not exists or you don't have required permissions.")
		}

		println(selectedProject.ID)
		if utils.IsFlagMode(cmd.Flags(), localFlags) {
			utils.ValidateAllRequiredFlags(cmd.Flags(), localFlags)
			sourceBranch, _ := cmd.Flags().GetString("sourceBranch")
			targetBranch, _ := cmd.Flags().GetString("targetBranch")
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")

			createMergeRequest = models.CreateMergeRequest{
				SourceBranch: sourceBranch,
				TargetBranch: targetBranch,
				Title:        title,
				Description:  description,
			}
		} else {
			branches, err := gitlabClient.GetProjectBranches(selectedProject.ID)
			if err != nil {
				 return err
			}

			branchNames := make([]string, len(projects))
			for i, branch := range branches {
				branchNames[i] = branch.Name
			}
			sourceBranch := utils.PromptSearchStrings(branchNames, "Select source branch")
			targetBranch := utils.PromptSearchStrings(branchNames, "Select target branch")
			title := utils.GetPromptString("Title", false)
			description := utils.GetPromptString("Description (optional)", false)

			createMergeRequest = models.CreateMergeRequest{
				SourceBranch: sourceBranch,
				TargetBranch: targetBranch,
				Title:        title,
				Description:  description,
			}
		}

		response, err := gitlabClient.CreateMergeRequest(selectedProject.ID, createMergeRequest)

		if err != nil {
			return err
		}

		println("Merge request create successfully!")
		println(response.WebURL)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	for _, flag := range localFlags {
		mergeCmd.Flags().StringP(flag.Name, flag.Shorthand, flag.Value, flag.Description)
	}
}
