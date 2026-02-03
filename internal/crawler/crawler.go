package crawler

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"orca/internal/output"
)

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

func Run(target string, rate int,out *output.Writer){
	parse,_:= url.Parse(target)
	host:= parse.Hostname()

	c:=colly.NewCollector(
		colly.AllowedDomains(host),
		colly.Async(true),)

	c.Limit(&colly.LimitRule{
			DomainGlob: "*",
			Parallelism: 1,
			Delay: time.Second/time.Duration(rate),
		},
	)
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	c.OnResponse(func(r *colly.Response) {
		code := r.StatusCode
		url := r.Request.URL.String()

		switch {
		case code >= 200 && code < 300:
			line := fmt.Sprintf("%s %s", green(fmt.Sprintf("[%d]", code)), url)
			out.Write(line)
		case code >= 300 && code < 400:
			line := fmt.Sprintf("%s %s", yellow(fmt.Sprintf("[%d]", code)), url)
			out.Write(line)
		default:
			line := fmt.Sprintf("%s %s", red(fmt.Sprintf("[%d]", code)), url)
			out.Write(line)
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