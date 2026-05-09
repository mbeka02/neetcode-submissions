func trap(height []int) int {
/*
left,right:=0,len(height)-1
width:=1
trappedWater:=0
to get the amount of water at point p , the greater element on the left and right are crucial
so get min(height[left],height[right])
area=width * min(height[left],height[right])-height[i] 
*/
n:=len(height)
  if n == 0 {
        return 0
    }
//precompute the leftMax and rightMax for each index
leftMax:=make([]int,n)
rightMax:=make([]int,n)
//fill left max
leftMax[0]=height[0]
for i:=1;i<n;i++{
leftMax[i]=max(leftMax[i-1],height[i])
}
//fill right max
rightMax[n-1]=height[n-1]
for i:=n-2;i>=0;i--{
  rightMax[i]=max(rightMax[i+1],height[i])
}
trappedWater:=0
for i:=0;i<n;i++{
trappedWater+= min(leftMax[i],rightMax[i])-height[i]
}
return trappedWater
}


func min(a,b int)int{
	if a < b{
	return a
	} 
	return b
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}