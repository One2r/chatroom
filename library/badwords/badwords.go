package badwords

import (
	"os"
	"bufio"
	"io"
	"strings"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/cloudflare/ahocorasick"
)

var m *ahocorasick.Matcher

func init() {
	var badword []string
	f, err := os.Open(filepath.Join(beego.AppPath, "conf", beego.AppConfig.String("badword_dict")))
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			line = strings.TrimSpace(line)  
			badword = append(badword,line)
		}
	}
	m = ahocorasick.NewStringMatcher(badword)
}

func HasBadWord(content string) bool {
	hits := m.Match([]byte(content))
	if len(hits) > 0 {
		return true
	} else {
		return false
	}
}
