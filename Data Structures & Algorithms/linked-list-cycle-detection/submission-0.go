/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    // algorithm:
	// m:= make(map[*ListNode]int)
	//iterate through the linkedlist;
   // for current:=head;current!=nil;current=current.Next
	// m[current]++
	//
   m:=make(map[*ListNode]int)
  for current:=head;current!=nil;current=current.Next{
   if _,ok:=m[current];ok{
	return true
   }    
   m[current]++
   }
return false
}