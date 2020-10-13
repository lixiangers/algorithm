package string

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAcAutoMachine_Query(t *testing.T) {
	content := "我是红领巾，祖国未来的花朵王八蛋abcd好吧红领巾"
	ac := NewAcAutoMachine()
	ac.AddPattern("红领巾")
	ac.AddPattern("红领")
	ac.AddPattern("红")
	ac.AddPattern("祖国")
	ac.AddPattern("花")
	ac.AddPattern("花朵")
	ac.AddPattern("王八")
	ac.AddPattern("王八蛋")
	ac.AddPattern("abcd")
	ac.AddPattern("a")
	ac.AddPattern("ab")
	ac.AddPattern("bcd")
	ac.Build()
	results := ac.Query(content)
	fmt.Println("内容: " + content)
	for _, result := range results {
		fmt.Println(result)
	}
}

func BenchmarkAcAutoMatch2(b *testing.B) {
	b.StopTimer()
	// 加载敏感词
	sensitiveWords := loadSensitiveWords("../data/sensitive_word_other_2.txt")
	// 建树
	//buildStartTime:=time.Now()
	ac := NewAcAutoMachine()
	for _, w := range sensitiveWords {
		ac.AddPattern(w)
	}
	ac.Build()
	//fmt.Println("BuildTree:",time.Now().Sub(buildStartTime))

	// 设置fail指针
	//setNodeFailPointTime :=time.Now()
	ac.Build()
	//fmt.Println("setNodeFailPointTime:",time.Now().Sub(setNodeFailPointTime))
	part := "你是抓不是的中国很好水电费加二米几十分是丹佛奇偶偶没说明佛惊个附加奥奇建安费劲阿斯顿发力看见哦按点说"
	message := ""
	for i := 0; i < 2; i++ {
		message += part
	}
	//println("message len:",len(message))
	b.StartTimer()
	// 查找
	//searchStartTime:=time.Now()
	for i := 0; i < b.N; i++ {
		//GetPrimes(1000)
		//AcAutoMatch("腐败中国，你说是抓不是这样的。天擎很好，是你们佛按玉穴是否阿萨德爱上i")
		//AcAutoMatch(message)
		ac.Query(message)
		//fmt.Println()
	}
	//fmt.Println("search Cost:",time.Now().Sub(searchStartTime))
}

func TestAcAutoMachine_Query_2(t *testing.T) {
	orgidNum, _ := strconv.ParseInt("", 10, 64)
	fmt.Println("orgidNum", orgidNum)
	//fmt.Println(string([]rune("腐败中国")[1:3]))
	//加载敏感词
	sensitiveWords := loadSensitiveWords("../data/sensitive_word_other_2.txt")
	// 建树
	//buildStartTime:=time.Now()
	ac := NewAcAutoMachine()
	for _, w := range sensitiveWords {
		ac.AddPattern(w)
	}
	ac.Build()
	results := ac.Query("你是个王八蛋嘛,fuck you")
	for _, result := range results {
		fmt.Println(result)
	}
}

func loadSensitiveWords(path string) []string {
	result := make([]string, 0)
	filePaths := strings.Split(path, ",")
	if len(filePaths) < 1 {
		panic("no file")
	}
	for _, filePaths := range filePaths {
		fi, err := os.Open(filePaths)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		defer fi.Close()

		br := bufio.NewReader(fi)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			//fmt.Println(string(a))
			line := strings.ReplaceAll(string(a), " ", "")
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			result = append(result, line)
		}
	}

	//fmt.Println("words:",len(result))
	return result
}
