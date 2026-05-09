func maxArea(heights []int) int {
// I need to find the 2 largest values in the array.
//constraints
/*
2 <= height.length <= 1000
0 <= height[i] <= 1000
*/
/*
pseudo-code:
left:=0
right:=len(heights)-1
*/
left,right:=0,len(heights)-1
res:=0
for left<right{
area:=min(heights[left],heights[right])* (right-left)
if area > res {
 res=area
}
if heights[left]<=heights[right]{
	left++
}else{
	right--
}
}
return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
