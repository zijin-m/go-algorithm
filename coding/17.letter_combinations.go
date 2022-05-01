package coding

var digitMap = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func LetterCombinations(digits string) []string {
	res := []string{}
	letterCombinations(digits, 0, "", &res)
	return res
}

func letterCombinations(digits string, index int, s string, result *[]string) {
	if len(digits) == len(s) {
		*result = append(*result, s)
		return
	}
	letters := digitMap[digits[index]-'0']
	for i := 0; i < len(letters); i++ {
		s += string(letters[i])
		letterCombinations(digits, index+1, s, result)
		s = s[:len(s)-1]
	}
}
