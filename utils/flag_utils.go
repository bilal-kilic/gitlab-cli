package utils

import (
	"github.com/spf13/pflag"
	"gitlab-cli/cmd/models"
)

func IsFlagMode(flags *pflag.FlagSet, localFlags []models.Flag) bool {
	hasAnyLocalFlag := false
	for _, localFlag := range localFlags {
		flagName, _ := flags.GetString(localFlag.Name)
		if flagName != "" {
			hasAnyLocalFlag = true
			break
		}
	}

	return hasAnyLocalFlag
}

func ValidateAllRequiredFlags(flags *pflag.FlagSet, localFlags []models.Flag) {
	for _, localFlag := range localFlags {
		flagName, _ := flags.GetString(localFlag.Name)
		if flagName == "" && localFlag.Required {
			println("Parameter %s is missing. (--%s -%s)", localFlag.Name, localFlag.Name, localFlag.Shorthand)
			panic("When in parameter mode you must enter all required parameters.")
		}
	}
}
