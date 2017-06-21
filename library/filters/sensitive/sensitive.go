package sensitive

import (
	"os"
	"bufio"
	"io"
	"strings"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/One2r/ahocorasick"
)

//是否开启敏感词检测功能
var Enable bool = true

var m *ahocorasick.Matcher

func init() {
	UpdateSensitiveWords()
}

//字符串content是否含有敏感关键词
func HasSensitiveWords(content string) bool {
	if m != nil {
		hits := m.Match([]byte(content))
		if len(hits) > 0 {
			return true
		}
	}
	return false
}

//读取敏感关键词词典
func ReadDict() []string {
	var sensitiveWords []string
	f, err := os.Open(filepath.Join(beego.AppPath, "conf", beego.AppConfig.String("sensitive_words_dict")))
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			line = strings.TrimSpace(line)  
			sensitiveWords = append(sensitiveWords,line)
		}
	}
	return sensitiveWords
}

//更新敏感关键词
func UpdateSensitiveWords() bool {
	sensitiveWords := ReadDict()
	if len(sensitiveWords) > 0 {
		m = ahocorasick.NewStringMatcher(sensitiveWords)
		return true
	}else{
		return false
	}
}