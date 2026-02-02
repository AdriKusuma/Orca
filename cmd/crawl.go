package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var crawl string

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Input target URL for crawling",
	Run: func(cmd *cobra.Command, args []string) {
		RunCrawler(crawl)
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

func RunCrawler(target string) {
	if target == "" {
		fmt.Println("Empty URL")
		return
	}

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200{
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Println(green(r.StatusCode), r.Request.URL)
		}else if r.StatusCode == 302 {
			yellow := color.New(color.FgYellow).SprintFunc()
			fmt.Println(yellow(r.StatusCode), r.Request.URL)
		}else{
			red:= color.New(color.FgRed).SprintFunc()
			fmt.Println(red(r.StatusCode), r.Request.URL)
		}
		
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) { 
		e.Request.Visit(e.Attr("href")) 
	})

	c.Visit(target)
}



