package wordfilter

import (
	"strings"

	"github.com/qiuchengw/wordfilter/trie"
)

// func apiHelper(w http.ResponseWriter) {
// 	help := make(map[string]string)
// 	help["/v1/query?q={text} [GET,POST] "] = "查找敏感词"

// 	help["/v1/black_words [GET]"] = "查看敏感词"
// 	help["/v1/black_words [POST]"] = "添加敏感词"
// 	help["/v1/black_words [DELETE]"] = "删除敏感词"

// 	help["/v1/white_prefix_words [GET]"] = "查看白名单（前缀）词组"
// 	help["/v1/white_prefix_words [POST]"] = "添加白名单（前缀）词组"

// 	help["/v1/white_suffix_words [GET]"] = "查看白名单（后缀）词组"
// 	help["/v1/white_suffix_words [POST]"] = "添加白名单（后缀）词组"

// 	serveJSON(w, help)
// }

// ReplaceBlackWords 替换关键字a
func ReplaceBlackWords(text string) (ok bool, keywords []string, ret string) {
	ok, keywords, ret = trie.BlackTrie().Query(text)
	return
}

// AddBlackWords 添加敏感词
func AddBlackWords(words []string) int {
	i := 0
	for _, s := range words {
		trie.BlackTrie().Add(strings.Trim(s, " "))
		i++
	}
	return i
}

// DeleteBlackWords 删除敏感词
func DeleteBlackWords(words []string) int {
	i := 0
	for _, s := range words {
		trie.BlackTrie().Del(strings.Trim(s, " "))
		i++
	}
	return i
}

// func showBlackWords(w http.ResponseWriter, r *http.Request) {
// 	words := trie.BlackTrie().ReadAll()
// }

// AddWhitePrefixWords 添加白名单
func AddWhitePrefixWords(words []string) int {
	i := 0
	for _, s := range words {
		trie.WhitePrefixTrie().Add(strings.Trim(s, " "))
		i++
	}
	return i
}

func init() {
	trie.InitAllTrie()
}
