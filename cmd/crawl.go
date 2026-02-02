package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var crawl string

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Input target URL for crawling",
	Run: func(cmd *cobra.Command, args []string) {

		if crawl == "" {
			fmt.Println("Empty URL, Use -h or --help for help")
			return
		}
		fmt.Println("Target:", crawl)
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)

	crawlCmd.Flags().StringVarP(
		&crawl,
		"url",
		"u",
		"",
		"Target URL",
	)

	crawlCmd.MarkFlagRequired("url")
}
