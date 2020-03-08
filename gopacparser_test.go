package gopacparser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"reflect"
	"testing"
)


func TestFindProxy_SiteNeedsToBeHttpProxy(t *testing.T) {
	expected := map[string]*url.URL{
		"http": &url.URL{
			Scheme: "http",
			Host:   "proxy.antizapret.prostovpn.org:3128",
		},
		"https": &url.URL{
			Scheme: "http",
			Host:   "proxy.antizapret.prostovpn.org:3128",
		},
	}

	proxy, err := FindProxy("test_data/antizapret.pac", "http://filmix.me")

	isEquals := reflect.DeepEqual(proxy, expected)

	assert.NoError(t, err)
	if !isEquals {
		assert.Fail(t, fmt.Sprintf("Expected:\n%s \nActual:\n%s", fmt.Sprint(expected), fmt.Sprint(proxy)))
	}
}

func TestFindProxy_SiteDoesntHttpProxy(t *testing.T) {
	expected := map[string]*url.URL{}

	proxy, err := FindProxy("test_data/antizapret.pac", "https://google.com")

	isEquals := reflect.DeepEqual(proxy, expected)

	assert.NoError(t, err)
	if !isEquals {
		assert.Fail(t, fmt.Sprintf("Expected:\n%s \nActual:\n%s", fmt.Sprint(expected), fmt.Sprint(proxy)))
	}
}

func TestFindProxy_(t *testing.T) {
	expected := map[string]*url.URL{}

	proxy, err := FindProxy("test_data/not-valid-file.pac", "https://google.com")

	isEquals := reflect.DeepEqual(proxy, expected)

	assert.Error(t, err)
	if !isEquals {
		assert.Fail(t, fmt.Sprintf("Expected:\n%s \nActual:\n%s", fmt.Sprint(expected), fmt.Sprint(proxy)))
	}
}

func TestFindProxy_SendNotValidPathToPacFile(t *testing.T) {
	_, err := FindProxy("test_data/antizapret.pa", "http://filmix.me")
	assert.Error(t, err)
}

func TestGetHostname(t *testing.T) {
	testArgs := []string{"https://tour.golang.org/moretypes/16", "https://filmix.co/drama/135023-21-most-19.html", "some-not-valid-url"}
	expectedValues := []string{"tour.golang.org", "filmix.co", ""}

	for i := 0; i < len(testArgs); i++ {
		arg := testArgs[i]
		expected := expectedValues[i]

		baseUrl := getHostname(arg)

		assert.Equal(t, expected, baseUrl, "")
	}
}

func TestProxyUrl_RunWithValidData(t *testing.T) {
	testArgs := []string{"DIRECT", "PROXY proxy-nossl.antizapret.prostovpn.org:29976", "SOCKS 164.22.93.93:45344"}
	expectedValues := []string{"DIRECT", "http://proxy-nossl.antizapret.prostovpn.org:29976", "socks5://164.22.93.93:45344"}

	for i := 0; i < len(testArgs); i++ {
		arg := testArgs[i]
		expected := expectedValues[i]

		urlToProxy, err := proxyUrl(arg)

		assert.Equal(t, expected, urlToProxy, "")
		assert.Nil(t, err, "")
	}
}

func TestProxyUrl_RunWithNotValidData(t *testing.T) {
	urlToProxy, err := proxyUrl("some-not-valid-url")

	assert.Equal(t, "", urlToProxy, "")
	assert.Error(t, err, "")
}

func TestParsePacValue(t *testing.T) {
	testArgs := []string{
		"DIRECT",
		"PROXY proxy-nossl.antizapret.prostovpn.org:29976; DIRECT",
		"SOCKS 164.22.93.93:45344; SOCKS 164.22.39.89:45344",
		"SOCKS 164.22.93.93:45344; SOCKS 164.22.39.89:45344; ",
		"some-not-valid-url; SOCKS 164.22.39.89:45344"}
	expectedValues := [][]string{
		{"DIRECT"},
		{"http://proxy-nossl.antizapret.prostovpn.org:29976", "DIRECT"},
		{"socks5://164.22.93.93:45344", "socks5://164.22.39.89:45344"},
		{"socks5://164.22.93.93:45344", "socks5://164.22.39.89:45344"},
		{"socks5://164.22.39.89:45344"},}

	for i := 0; i < len(testArgs); i += 1 {
		arg := testArgs[i]
		expected := expectedValues[i]

		proxies := parsePacValue(arg)

		assert.ElementsMatch(t, expected, proxies)
	}
}
