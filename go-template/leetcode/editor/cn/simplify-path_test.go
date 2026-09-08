package leetcode_solutions

import (
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func simplifyPath(path string) string {
	res := ""

	words := make([]string, 0)
	validWords := make([]string, 0)

	words = strings.Split(path, "/")

	//fmt.Println(words)

	for i := 0; i < len(words); i++ {
		if words[i] == "." || words[i] == "" {
			continue
		} else if words[i] == ".." {
			if len(validWords) > 0 {
				validWords = validWords[:len(validWords)-1]
			}
			continue
		}
		validWords = append(validWords, words[i])
	}

	for i := range validWords {
		res = res + "/" + validWords[i]
	}
	if res == "" {
		res = "/"
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSimplifyPath(t *testing.T) {
	// your test code here
	simplifyPath("/../")
}
