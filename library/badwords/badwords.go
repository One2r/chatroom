package badwords

import (
	"os"
	"bufio"
	"io"
	"strings"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/One2r/ahocorasick"
)

var m *ahocorasick.Matcher

func init() {
	badword := ReadDict()
	m = ahocorasick.NewStringMatcher(badword)
}

//字符串content是否含有敏感关键词
func HasBadWord(content string) bool {
	hits := m.Match([]byte(content))
	if len(hits) > 0 {
		return true
	} else {
		return false
	}
}

//读取敏感关键词词典
func ReadDict() []string {
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
	return badword
}

//更新敏感关键词
func UpdateBadword() bool {
	badword := ReadDict()
	if len(badword) > 0 {
		m = ahocorasick.NewStringMatcher(badword)
	}
	return true
}