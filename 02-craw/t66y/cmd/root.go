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
	"mytool/02-craw/util"
)

var (
	t66yURL      = "http://t66y.com/htm_data/2004/2/3872569.html"
	downLoadPath = util.GetDownLoadPath()
	paral        =  5
)

var rootCmd = &cobra.Command{
	Use:   "t66y",
	Short: "A tool to download t66y image",
	Long:  `A tool to download t66y image`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(downLoadPath)
	},
}

// Execute executes the root command.
func Execute() error {

	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	rootCmd.AddCommand(cmdDownload)
}
