package m_common

import "fmt"

func TestSwitch(score int) {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score %d\n", score))
	case score < 60:
		g = "不及格"
	case score > 60:
		g = "及格"
	}
	fmt.Println(g)
}

func TestSwitch1(score string) {
	g := ""
	switch score {
	case "1":
		fallthrough
	case "3":
		fallthrough
	case "5":
		fallthrough
	case "7":
		fallthrough
	case "8":
		fallthrough
	case "10":
		fallthrough
	case "12":
		fmt.Println("31天")
	case "4":
		fallthrough
	case "6":
		fallthrough
	case "9":
		fallthrough
	case "11":
		fmt.Println("30天")
	case "2":
		fmt.Println("28、29")
	}
	fmt.Println(g)
}
