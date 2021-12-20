package cmd
/*
Copyright 2022 Absa Group Limited

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/

import (
	"fmt"
	"os"

	"github.com/kuritka/golic-1/impl/update"

	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var injectOptions update.Options

var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Injects licenses",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(injectOptions.LicIgnore); os.IsNotExist(err) {
			logger.Error().Msgf("invalid license path '%s'", injectOptions.LicIgnore)
			_ = cmd.Help()
			os.Exit(1)
		}
		if masterconfig == "" {
			logger.Error().Msgf("invalid master config. ")
			_ = cmd.Help()
			os.Exit(1)
		}
		injectOptions.MasterConfig = masterconfig
		injectOptions.Type = update.LicenseInject
		i := update.New(ctx, injectOptions)
		exitCode = Command(i).MustRun()
		if exitCode == 0 {
			fmt.Printf(" %s %s\n", emoji.Rocket, aurora.BrightWhite("done"))
		} else {
			fmt.Printf(" %s %s\n", emoji.FaceScreamingInFear, aurora.BrightWhite("found files with missing a license, exit"))
		}
	},
}

func init() {
	injectCmd.Flags().BoolVarP(&injectOptions.ModifiedExitStatus, "modified-exit", "x", false,
		"If enabled, exits with status 1 when any file is modified. The settings is used by CI")
	injectCmd.Flags().StringVarP(&injectOptions.LicIgnore, "licignore", "l", ".licignore", ".licignore path")
	injectCmd.Flags().StringVarP(&injectOptions.Template, "template", "t", "apache2", "license key")
	injectCmd.Flags().StringVarP(&injectOptions.Copyright, "copyright", "c", "2022 MyCompany",
		"company initials entered into license")
	injectCmd.Flags().BoolVarP(&injectOptions.Dry, "dry", "d", false, "dry run")
	injectCmd.Flags().StringVarP(&injectOptions.ConfigPath, "config-path", "p", ".golic.yaml", "path to the local configuration overriding config-url")
	rootCmd.AddCommand(injectCmd)
}
