package cmd

import (
	"fmt"
	"orca/internal/crawler"
	"orca/internal/output"
	"time"
	"github.com/spf13/cobra"
)

var (
	target string
	outputFile string
)

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl target website",
	Run: func(cmd *cobra.Command, args []string) {	
    fmt.Println(`
		⢀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠺⢿⣿⣿⣿⣿⣿⣿⣷⣦⣠⣤⣤⣤⣄⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣄⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⢀⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⣿⣿⣷⣄⠀⠀
	⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣀⠀⠀⠀⣀⣿⣿⣿⣆⠀
	⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄
	⠀⠀⠀⠀⣾⣿⣿⡿⠋⠁⣀⣠⣬⣽⣿⣿⣿⣿⣿⣿⠿⠿⠿⠿⠿⠿⠿⠿⠟⠁
	⠀⠀⠀⢀⣿⣿⡏⢀⣴⣿⠿⠛⠉⠉⠀⢸⣿⣿⠿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⢸⣿⣿⢠⣾⡟⠁⠀⠀⠀⠀⠀⠈⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⢸⣿⣿⣾⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⣸⣿⣿⣿⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⢠⣾⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⢰⣿⡿⠛⠉⠀⠀⠀⠈⠙⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠈⠁  ___      ____       __       ____ 
	/$$$$$$  /$$$$$$$   /$$$$$$   /$$$$$$ 
	/$$__  $$| $$__  $$ /$$__  $$ /$$__  $$
	| $$  \ $$| $$  \ $$| $$  \__/| $$  \ $$
	| $$  | $$| $$$$$$$/| $$      | $$$$$$$$
	| $$  | $$| $$__  $$| $$      | $$__  $$
	| $$  | $$| $$  \ $$| $$    $$| $$  | $$
	|  $$$$$$/| $$  | $$|  $$$$$$/| $$  | $$
	\______/ |__/  |__/ \______/ |__/  |__/                                       
											
	OFFENSIVE SECURITY TOOL BY Adri Kusuma`)
 	fmt.Println("========================================")
	fmt.Println("Target    : ", target)
	fmt.Println("Rate limit: ", time.Second/time.Duration(rate))
	fmt.Println("========================================")
	if target == "" {
		fmt.Println("Target URL required")
		return
	}

	if rate <= 0 {
		rate = 5
	}

	writer, err := output.New(outputFile)
	if err != nil {
		fmt.Println("Output error:", err)
		return
	}
	defer writer.Close()

	crawler.Run(target, rate, writer)
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)

	crawlCmd.Flags().StringVarP(
		&target,
		"url",
		"u",
		"",
		"target URL",
	)

	crawlCmd.MarkFlagRequired("url")
}



