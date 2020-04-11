package cmd

import (
	"github.com/spf13/cobra"
	"mytool/02-craw/t66y/controller"
)

var cmdDownload = &cobra.Command{
	Use:   "download",
	Short: "下载t66y图片",
	Long:  `下载t66y图片`,
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controller.Run(t66yURL, downLoadPath, paral)
	},
}

func init() {
	rootCmd.AddCommand(cmdDownload)
	cmdDownload.PersistentFlags().StringVarP(&t66yURL, "url", "u", t66yURL, "需要下载的地址")
	cmdDownload.PersistentFlags().IntVarP(&paral, "paral", "n", 5, " 下载时的并发线程(默认为5)")
}
