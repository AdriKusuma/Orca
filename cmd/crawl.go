package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"net/url"
	"strings"
)

var crawl string

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Input target URL for crawling",
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
	fmt.Println("Target: ", crawl)
	fmt.Println("========================================")
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
func normalizeURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}

	u.Fragment = ""

	q := u.Query()
	for key := range q {
		q.Set(key, "")
	}
	u.RawQuery = q.Encode()

	return u.String()
}

func RunCrawler(target string) {
	if target == "" {
		fmt.Println("Empty target URL")
	}

	parsed, err := url.Parse(target)
	if err != nil {
		fmt.Println("Invalid URL")
	}

	host := parsed.Hostname()

	c := colly.NewCollector(
		colly.AllowedDomains(host),
		colly.MaxDepth(3),
		colly.Async(true),
	)

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	c.OnResponse(func(r *colly.Response) {
		code := r.StatusCode
		url := r.Request.URL.String()

		switch {
		case code >= 200 && code < 300:
			fmt.Println(green("[", code, "]"), url)
		case code >= 300 && code < 400:
			fmt.Println(yellow("[", code, "]"), url)
		default:
			fmt.Println(red("[", code, "]"), url)
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link == "" {
			return
		}

		abs := e.Request.AbsoluteURL(link)
		if abs == "" {
			return
		}

		u, err := url.Parse(abs)
		if err != nil {
			return
		}
		if !strings.Contains(u.Hostname(), host) {
			return
		}

		clean := normalizeURL(abs)
		if clean == "" {
			return
		}

		e.Request.Visit(clean)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(red("[ERR]"), r.Request.URL, err)
	})

	fmt.Println("Start crawling:", target)
	c.Visit(target)
	c.Wait()
}
