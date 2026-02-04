package crawler

import (
	"fmt"
	"net/url"
	"orca/internal/agent"
	"orca/internal/output"
	"strings"
	"time"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
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


func Run(target string, rate int, out *output.Writer, parallelism int, uaFile string) error {
	parsed, err := url.Parse(target)
	if err != nil {
		return err
	}
	host := parsed.Hostname()

	userAgents, err := agent.Load(uaFile)
	if err != nil {
		return err
	}

	c := colly.NewCollector(
		colly.AllowedDomains(host),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: parallelism,
		Delay:       time.Second / time.Duration(rate),
	})

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", agent.Random(userAgents))
	})

	c.OnResponse(func(r *colly.Response) {
		code := r.StatusCode
		link := r.Request.URL.String()
		var line string
		switch {
		case code >= 200 && code < 300:
			line = fmt.Sprintf("%s %s", green(fmt.Sprintf("[%d]", code)), link)
		case code >= 300 && code < 400:
			line = fmt.Sprintf("%s %s", yellow(fmt.Sprintf("[%d]", code)), link)
		default:
			line = fmt.Sprintf("%s %s", red(fmt.Sprintf("[%d]", code)), link)
		}
		out.Write(line)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		abs := e.Request.AbsoluteURL(href)
		u, err := url.Parse(abs)
		if err != nil || !strings.Contains(u.Hostname(), host) {
			return
		}
		clean := normalizeURL(abs)
		e.Request.Visit(clean)
	})

	c.Visit(target)
	c.Wait()
	return nil
}