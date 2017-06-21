package replace

import (
	"os"
	"bufio"
	"io"
	"strings"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/One2r/ahocorasick"
)

//是否开启替换词功能
var Enable bool = true

var m *ahocorasick.Matcher
var replaceMap map[string]string
var replace []string

func init() {
	UpdateReplaceWords()
}

//读取替换词词典
func ReadDict() map[string]string {
	replaceWords := make(map[string]string)
	f, err := os.Open(filepath.Join(beego.AppPath, "conf", beego.AppConfig.String("replace_words_dict")))
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			line = strings.TrimSpace(line)
			strArr := strings.Split(line, "|")
			replaceWords[strArr[0]] = strArr[1]
		}
	}
	return replaceWords
}

//替换文本中的词
func Replace(content string) string {
	if m != nil {
		hits := m.Match([]byte(content))
		if len(hits) > 0 {
			for _, v := range hits {
				content = strings.Replace(content, replace[v], replaceMap[replace[v]], 1)
			}
		}
	}
	return content
}

func UpdateReplaceWords() bool {
	replaceMap = ReadDict()
	if len(replaceMap) > 0 {
		for k,_ := range replaceMap {
			replace =  append(replace,k)
		}
		m = ahocorasick.NewStringMatcher(replace)
		m.GreedyMatch = true
		return true
	}else{
		return false
	}
}