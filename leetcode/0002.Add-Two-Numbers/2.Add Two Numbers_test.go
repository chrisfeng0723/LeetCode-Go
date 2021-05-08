/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/5/8 14:53
 */

package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type question2 struct {
	para2
	ans2
}
type para2 struct {
	one []int
	another []int
}

type ans2 struct {
	one []int
}

func Test_Problem2(t *testing.T){
	qs := []question2{
		{
			para2{[]int{},[]int{}},
			ans2{[]int{}},
		},
		{
			para2{[]int{1},[]int{1}},
			ans2{[]int{2}},
		},
		{
			para2{[]int{1, 8, 3}, []int{7, 1}},
			ans2{[]int{8, 9, 3}},
		},
	}

	for _,q:=range qs{
	 _,p :=q.ans2,q.para2
	 fmt.Printf("[input]:%v  [output]:%v\n",p,structures.List2Ints(addTwoNumbers(structures.Ints2List(p.one),structures.Ints2List(p.another))))

	}
	fmt.Printf("\n\n\n")
}
