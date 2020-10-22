package m_regex

import (
	"fmt"
	"regexp"
)

const text = `
my email is ss@qq.com
email2 is a@aa.com
email3 is b@bb.com.cn
`

func regexDemo() {

	// .+@.+\\..+  .+代表任何字符
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//re.FindString()

	//allString := re.FindAllString(text, -1) // 匹配所有n=-1 | 匹配n个

	submatch := re.FindAllStringSubmatch(text, -1)
	for _, s := range submatch {
		fmt.Println(s)
	}
	/*
		[ss@qq.com ss qq .com]
		[a@aa.com a aa .com]
		[b@bb.com.cn b bb .com.cn]
	*/
}
