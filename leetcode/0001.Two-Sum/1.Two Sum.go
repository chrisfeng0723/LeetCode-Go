/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/27 17:54
 */

package leetcode

func twoSum(nums []int,target int)[]int{
	m :=make(map[int]int,len(nums))
	for key,val :=range nums{
		if idx,ok :=m[target -val];ok{
			return []int{idx,key}
		}
		m[val] = key
	}
	return nil
}

