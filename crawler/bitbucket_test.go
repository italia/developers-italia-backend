package crawler

import (
	"io/ioutil"
	"testing"

	log "github.com/sirupsen/logrus"
)

// IsBitbucket returns "true" if the url can use Bitbucket API.
func TestIsBitbucket(t *testing.T) {
	// Disablle log output for this function
	log.SetOutput(ioutil.Discard)

	links := []struct {
		in  string
		out bool
	}{
		// {"https://bitbucket.org/Soft", true},
		// {"https://gitlab.com/Soft", false},
		// {"https://github.com/Soft", false},
		// {"", false},
		// {"invalidUrl", false},
		// {"example.example", false},
		// {":unparsable", false},
	}

	for _, l := range links {
		if IsBitbucket(l.in) != l.out {
			t.Logf("Expected %s == %t.", l.in, l.out)
			t.Fail()
		}
	}

}

// GenerateBitbucketAPIURL returns the api url of given Bitbucket  organization link.
// IN: https://bitbucket.org/Soft
// OUT:https://api.bitbucket.org/2.0/repositories/Soft?pagelen=100
func TestGenerateBitbucketAPIURL(t *testing.T) {
	// Disablle log output for this function
	log.SetOutput(ioutil.Discard)

	links := []struct {
		in  string
		out string
	}{
		{"https://bitbucket.org/Soft", "https://api.bitbucket.org/2.0/repositories/Soft"},
		{":unparsable", ":unparsable"},
	}

	for _, l := range links {
		genURL := GenerateBitbucketAPIURL()
		if out, err := genURL(l.in); out != l.out {
			t.Logf("Expected %s == %s: %v ", out, l.out, err)
			t.Fail()
		}
	}

}
