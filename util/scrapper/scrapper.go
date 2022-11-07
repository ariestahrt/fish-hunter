package scrapper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type Scrapper interface {
	GetPhishUrl(source string) ([]string, error)
}

type UrlScrapper struct {
}

func NewUrlScrapper() Scrapper {
	return &UrlScrapper{}
}

func (u *UrlScrapper) GetPhishUrl(source string) ([]string, error) {
	switch source {
	case "phishtank":
		return get_phishtank_feed()
	case "openphish":
		return get_openphish_feed()
	default:
		return nil, fmt.Errorf("source %s not found", source)
	}
}

func get_phishtank_feed() ([]string, error) {
	return nil, errors.New("not implemented")
}

func get_openphish_feed() ([]string, error) {
	// Http requests to openphish.com
	req, err := http.NewRequest("GET", "https://openphish.com/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,id;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp_str := string(b)

	// Parse the response
	result := make([]string, 0)
	var re = regexp.MustCompile(`(?m)class="url_entry">(.*?)</td>`)


	for _, match := range re.FindAllStringSubmatch(resp_str, -1) {
		result = append(result, match[1])
	}

	return result, nil
}