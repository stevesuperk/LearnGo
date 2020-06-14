package main

import "fmt"

func twoSum1(nums []int, target int) []int {
	//思路：循环数组，取出第一个值，
	use := []int{}
	var ret []int
	for i:=0;i<len(nums)-1;i++{
		//如果已使用跳过
		if use[i] != 0 {
			continue
		}
		//大于目标跳过
		if nums[i]>target {
			use[i] = 1
			continue
		}

		for j:=i+1;j<len(nums);j++{
			if nums[i]+nums[j] == target {
				use[i] = 1
				use[j] = 1
				fmt.Println(use)
				ret = append(ret,nums[i])
				ret = append(ret,nums[j])
			}
		}
	}
	return ret
}
func twoSum2(nums []int, target int) []int {
	//思路：循环数组，取出第一个值，
	for i:=0;i<len(nums)-1;i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]+nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}
func twoSum(nums []int, target int) []int {
	v := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dif := target - nums[i]
		c, ok := v[dif]
		if ok != false {
			return []int{c, i}
		}
		v[nums[i]] = i
	}
	return []int{-1,-1}
}


func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("打印值%#v",twoSum(nums,target))
	nums1 := []int{-1,-2,-3,-4,-5}
	target1 := -8
	fmt.Printf("打印值%#v",twoSum(nums1,target1))
}