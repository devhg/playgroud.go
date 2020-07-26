package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	l1 *ListNode
	l2 *ListNode
)

func init() {
}
func Delete() int {
	var nums = []int{3, 3, 3, 3}
	val := 3
	left := 0
	n := len(nums)
	for right := 0; right < n; right++ {
		if nums[right] == val {
			continue
		} else {
			nums[left] = nums[right]
			left++
		}
	}
	for i := 0; i < left; i++ {
		println(nums[i])
	}
	return left + 1
}

func Convert() string {
	var strs = []string{"flower", "flow", "flight"}
	var slen = len(strs)
	minLen := len(strs[0])

	if slen == 0 {
		return ""
	}

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
