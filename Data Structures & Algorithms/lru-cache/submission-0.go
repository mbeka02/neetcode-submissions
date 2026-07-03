type Node struct {
next *Node //points to the next node
prev *Node //points to the previous node
key int
value int
}

type LRUCache struct {
capacity int //max number of elements the cache can hold 
cacheMap map[int]*Node
head *Node
tail *Node

    
}

func Constructor(capacity int) LRUCache {
return LRUCache{
	capacity: capacity,
	cacheMap: make(map[int]*Node),
	head: nil,
	tail:nil,
}    
}
func (this *LRUCache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		this.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		this.tail = node.prev
	}

	node.prev = nil
	node.next = nil
}
func (this *LRUCache) addToFront(node *Node){
	node.prev=nil
	node.next=this.head

	if this.head !=nil{
		this.head.prev=node
		}
   this.head=node
   //if the list was empty also set the tail
   if this.tail==nil {
	this.tail=node
   }
}
func (this *LRUCache)moveToFront(node *Node){
	this.removeNode(node)
	this.addToFront(node)
}
func (this *LRUCache) Get(key int) int {

	//check the cache map 
	node,ok:=this.cacheMap[key]

  //return -1 of the key doesn't exist
  if !ok{
	return -1
  }
  //move the node to the front
this.moveToFront(node)
//return the value
  return node.value    
}

func (this *LRUCache) Put(key int, value int) {
    //check if key exists , if it does update the value else add it.
	if node, ok := this.cacheMap[key]; ok {
    node.value = value
    this.moveToFront(node)
    return
}
	newNode:=Node{key:key,value:value}
	this.addToFront(&newNode)
	this.cacheMap[key]=&newNode
	if len(this.cacheMap)>this.capacity {
		//evict the least recently used value
		evictedKey:=this.tail.key
		this.removeNode(this.tail)
		delete(this.cacheMap,evictedKey)

	}

}
