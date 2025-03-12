package main

import ("fmt")

type Node struct {
	key int
	value int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	cache map[int]*Node
	head *Node
	tail *Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		head: head,
		tail: tail,
	}
}

func (l *LRUCache) addNode(node *Node) {
	nextNode:= l.head.next
	l.head.next = node
	node.prev = l.head
	node.next = nextNode
	nextNode.prev = node
}

func (l *LRUCache) removeNode(node *Node){
	prevNode := node.prev
	nextNode := node.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (l *LRUCache) Get(key int) int {
	if node, exists := l.cache[key]; exists {
		l.removeNode(node)
		l.addNode(node)
		return node.value
	}
	return -1
}
func (l *LRUCache) Put(key int, value int) {
	if node, exists := l.cache[key]; exists {
		l.removeNode(node)
		delete(l.cache, key)
	}
	if len(l.cache) >= l.capacity {
		lruNode := l.tail.prev
		l.removeNode(lruNode)
		delete(l.cache, lruNode.key)
	}
	newNode := &Node{key: key, value: value}
	l.cache[key] = newNode
	l.addNode(newNode)
}

func main (){
	var capacity int;
	fmt.Print("Enter cache capacity: ");
	fmt.Scan(&capacity);

	cache:= NewLRUCache(capacity);

	for {
		fmt.Println("\n1. Get Value")
		fmt.Println("2. Put key-value pair")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter key to get: ")
			var key int
			fmt.Scan(&key)
			value := cache.Get(key)
			if value != -1 {
				fmt.Printf("Value for key %d: %d\n", key, value)
			} else {
				fmt.Printf("Key %d not found\n", key)
			}
		case 2:
			fmt.Print("Enter key and value to put: ")
			var key, value int
			fmt.Scan(&key, &value)
			cache.Put(key, value)
			fmt.Printf("Inserted key=%d, value=%d\n", key, value)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}	
}