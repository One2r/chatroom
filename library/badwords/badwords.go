package badwords

import (
	"github.com/cloudflare/ahocorasick"
)

var m *ahocorasick.Matcher

func init() {
	m = ahocorasick.NewStringMatcher([]string{"Mozilla", "Mac", "Macintosh", "Safari", "Sausage"})
}

func HasBadWord(content string) bool {
	hits := m.Match([]byte(content))
	if len(hits) > 0 {
		return true
	} else {
		return false
	}
}
