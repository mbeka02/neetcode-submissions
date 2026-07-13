/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
   // left,right ptr , left:=1 , right:=len(arr)-1
   //convert the linked list to a slice.
   //itr through the slice for idx 1...n (while left<right)
   // for each itr:
     /*
	 -push right value to result array
	 -push left value to result array
	 - left ++ 
	 - right --
	 */
	// convert result array back to a linkedlist.
 if head == nil {
	 return 
 }
 values:=make([]*ListNode,0)
 for current:=head;current!=nil;current=current.Next{
 values=append(values,current)
}
left,right:=1,len(values)-1
result:=make([]*ListNode,0)
result=append(result,values[0])
for left<=right{
	result=append(result,values[right])
	right --
	if left <= right {
	result=append(result,values[left])
	left ++
	}
}

for i:=0;i<len(result)-1;i++{
 result[i].Next=result[i+1]
}
result[len(result)-1].Next = nil
}