/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
To reverse a singly linked list in place, 
you must flip the direction of each node's next pointer 
so that it points to its predecessor instead of its 
successor
*/
func reverseList(head *ListNode) *ListNode {

current:=head
var prev *ListNode
for current !=nil {
nxt:=current.Next //save the forward path so that you don't lose the list
current.Next=prev //flips the direction backwards
prev=current //the current node is now the predecessor for the next step
current=nxt //move forward
}
 return prev

}
