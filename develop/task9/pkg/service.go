package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"task9/pkg/config"
	"time"

	"github.com/gocolly/colly"
)

func SaveFile(p string, b []byte) error {
	dir, file := path.Split(p)

	if file == "" {
		file = "index.html"
	}

	if err := os.MkdirAll(dir, os.FileMode(0644)); err != nil {
		return err
	}

	p = path.Join(dir, file)

	if err := os.WriteFile(p, b, 0644); err != nil {
		return err
	}

	return nil
}

func Wget(cfg *config.Config, logger *log.Logger) {
	c := colly.NewCollector(
		colly.AllowedDomains(cfg.Url.Host),
		colly.Async(true),
	)

	err := c.Limit(
		&colly.LimitRule{
			Delay: time.Duration(int(time.Second) / cfg.Rps),
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	if cfg.Rec {
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			trimedHref := strings.TrimSpace(e.Attr("href"))
			sep := strings.SplitN(trimedHref, "#", 2)
			if len(sep) == 0 || len(sep[0]) == 0 {
				return
			}
			link := e.Request.AbsoluteURL(sep[0])
			if err := c.Visit(link); err != nil {
				skip := errors.Is(err, colly.ErrAlreadyVisited) ||
					errors.Is(err, colly.ErrForbiddenDomain)

				if !skip {
					logger.Println(err, link)
				}
			}
		})
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		u := path.Join(r.Request.URL.Host, r.Request.URL.Path)

		if err := SaveFile(u, r.Body); err != nil {
			logger.Println(err)
		}
	})

	if err := c.Visit(cfg.Url.String()); err != nil {
		logger.Println(err)
	}

	c.Wait()
}
