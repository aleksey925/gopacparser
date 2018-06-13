package gopacparser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLocalPacFile(t *testing.T) {
	proxy, err := FindProxy("test_data/antizapret.pac", "http://filmix.me")
	assert.NoError(t, err)
	assert.Equal(t, proxy["http"].String(), "http://proxy.antizapret.prostovpn.org:3128")
	assert.Equal(t, proxy["https"].String(), "http://proxy.antizapret.prostovpn.org:3128")
}

func TestRemotePacFile(t *testing.T) {
	_, err := FindProxy("https://antizapret.prostovpn.org/proxy.pac", "http://filmix.me")
	assert.NoError(t, err)
}

func TestNotValidArgs(t *testing.T) {
	_, err := FindProxy("antizapret.prostovpn.org/proxy.pac", "http://filmix.me")
	assert.Error(t, err)

	_, err = FindProxy("test_data/antizapret.pa", "http://filmix.me")
	assert.Error(t, err)
}
