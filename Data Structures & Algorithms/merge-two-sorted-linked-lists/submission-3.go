/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
var (
	arr1 []int
	arr2 []int
)
//convert both lists into arrays
for current:=list1;current!=nil;current=current.Next{
arr1=append(arr1,current.Val)
}
for current:=list2;current!=nil;current=current.Next{
arr2=append(arr2,current.Val)
}
//append
mergedArray:=append(arr1,arr2...)
//sort
sort.Ints(mergedArray)
// edge case : empty lists passed as input
if len(mergedArray)==0{
return nil
}
//convert back to a linked list.
head:=&ListNode{Val:mergedArray[0]}
current:=head
for i:=1;i<len(mergedArray);i++{
current.Next=&ListNode{Val:mergedArray[i]}
current=current.Next // move forward
}



return head

}
