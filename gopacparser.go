package gopacparser

import (
	"errors"
	"fmt"
	"github.com/jackwakefield/gopac"
	neturl "net/url"
	"strings"
)

func getHostname(value string) string {
	result, err := neturl.Parse(value)
	hostname := ""
	if err == nil {
		hostname = result.Hostname()
	}
	return hostname
}

// Приводит в удобный вид значения, которые возвращает gopac
// value - значение для парсинга: "DIRECT", "PROXY example.local:8080", "SOCKS example.local:8080"
func proxyUrl(value string) (string, error) {
	if strings.ToUpper(value) == "DIRECT" {
		return "DIRECT", nil
	}
	parts := strings.Split(value, " ")

	if len(parts) == 2 {
		keyword, proxy := strings.ToUpper(parts[0]), parts[1]
		if keyword == "PROXY" {
			return "http://" + proxy, nil
		}
		if keyword == "SOCKS" {
			return "socks5://" + proxy, nil
		}
	}
	return "", errors.New(fmt.Sprintf("unrecognized proxy config value '%s'", value))
}

// Парсит строку возвращенную gopac
// value это строка вида:
// - "PROXY example.local:8080; DIRECT"
// - "DIRECT"
func parsePacValue(value string) []string {
	var result []string
	for _, element := range strings.Split(value, ";") {
		element = strings.TrimSpace(element)
		if len(element) == 0 {
			continue
		} else {
			url, err := proxyUrl(element)
			if err != nil {
				continue
			}
			result = append(result, url)
		}
	}
	return result
}

func FindProxy(pacPath string, url string) (map[string]*neturl.URL, error) {
	defer func() {
		if err := recover(); err != nil {
			err = errors.New("unexpected error when retrieving a proxy")
		}
	}()

	pacParser := new(gopac.Parser)
	if err := pacParser.ParseUrl(pacPath); err != nil {
		err = errors.New("error getting access to the PAC file")
		return map[string]*neturl.URL{}, err
	}

	pacData, err := pacParser.FindProxy(url, getHostname(url))
	if err != nil {
		err = errors.New("failed to find a proxy for: " + url)
		return map[string]*neturl.URL{}, err
	}

	proxies := parsePacValue(pacData)
	proxy := ""
	if len(proxies) != 0 {
		proxy = proxies[0]
	}
	if proxy == "" {
		return map[string]*neturl.URL{}, errors.New("no proxy configured or available for: " + url)
	}

	if proxy == "DIRECT" {
		return map[string]*neturl.URL{}, nil
	} else {
		parsedUrl, err := neturl.Parse(proxy)
		if err != nil {
			return map[string]*neturl.URL{}, err
		}
		return map[string]*neturl.URL{
			"http":  parsedUrl,
			"https": parsedUrl,
		}, nil
	}
}
