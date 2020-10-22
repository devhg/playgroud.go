package algorithm

//14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	var slen = len(strs)

	if slen == 0 {
		return ""
	}

	minLen := len(strs[0])

	for i := 1; i < slen; i++ {
		strLen := len(strs[i])
		if strLen < minLen {
			if strLen == 0 {
				return ""
			}
			minLen = strLen
		}
	}

	//获取相同部分的长度
	i := 0
I:
	for ; i < minLen; i++ {
		for j := 1; j < slen; j++ {
			if strs[j][i] != strs[j-1][i] {
				break I
			}
		}
	}

	return strs[0][0:i]
}
