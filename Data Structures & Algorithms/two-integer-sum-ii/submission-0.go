func twoSum(numbers []int, target int) []int {
	//target 7
	// [0,1,2,3,4,5]
left:=0
right:=len(numbers)-1
for left<right{
	//if the sum of the left +right > target decrease the right pointer
	//else increase the left pointer
	if numbers[left]+numbers[right] > target{
		right--
	}else if numbers[left]+numbers[right]<target{
		left++
	}else{
		break
	}
	

}
return []int{left+1,right+1}
}
