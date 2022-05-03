package coding

import (
	"sort"
	"strings"
)

// 输入：logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
//
// 输出：["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
//
// 解释：
// 字母日志的内容都不同，所以顺序为 "art can", "art zero", "own kit dig" 。
//
// 数字日志保留原来的相对顺序 "dig1 8 1 5 1", "dig2 3 6"
func ReorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		// 字母 vs 数字
		if isLetter(logs[i]) && !isLetter(logs[j]) {
			return true
		}
		// 数字 vs 字母
		if !isLetter(logs[i]) && isLetter(logs[j]) {
			return false
		}
		// 数字 vs 数字
		if !isLetter(logs[i]) && !isLetter(logs[j]) {
			return false
		}
		// 如果都是字母，在内容不同时，忽略标识符后，按内容字母顺序排序；在内容相同时，按标识符排序。
		c1, c2 := getContent(logs[i]), getContent(logs[j])
		if c1 == c2 {
			return getIdentity(logs[i]) < getIdentity(logs[j])
		}
		return c1 < c2
	})
	return logs
}

func isLetter(log string) bool {
	chars := strings.Split(log, " ")
	return chars[1][0] >= 'a' && chars[1][0] <= 'z'
}

func getContent(log string) string {
	return strings.Join(strings.Split(log, " ")[1:], " ")
}

func getIdentity(log string) string {
	return strings.Split(log, " ")[0]
}
