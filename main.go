// main文件
// author: baoqiang
// time: 2019/2/10 下午8:51
package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/githubao/xiao-gopro/ch02/tempconv"
	"github.com/githubao/xiao-gopro/ch05"
)

func main() {
	//RunCode()
	//RunMath()
	//ch03.RunBasename2()
	//ch03.RunComma()
	//ch04.Dedup()
	//ch04.CountChar()
	//ch05.RunVisit()
	//ch05.RunOutline()
	//ch05.RunVisit2()
	//ch05.RunWaitfor()
	//ch05.RunOutline2()
	//ch05.RunCrawl()
	//ch05.RunTitle()
	//ch05.RunTitle2()
	//ch05.RunFetch()
	ch05.RunTitle3()
}

func RunMath() {
	var max uint64 = math.MaxUint64

	//// int32
	//var max2 int32 = int32(max + 1)
	//
	//fmt.Println(max)
	//fmt.Println(max2)

	limit, _ := strconv.ParseInt(fmt.Sprintf("%d", max), 10, 32)

	if limit <= 0 {
		limit = 5000
	}

	// 2147483647
	//-2147483648
	//-2147483644
	var max2 = int32(limit + 1)

	fmt.Println(limit)
	fmt.Println(max2)
}

func RunCode() {

	if true {
		tempconv.TempconvGo()
	}
}
