package practics

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.removeNode(node)
		c.addNode(node)
		return node.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		c.removeNode(node)
		c.addNode(node)
		node.Val = value
		return
	}

	node := &Node{
		Key: key,
		Val: value,
	}

	if c.capacity == 0 {
		tail := c.tail.Prev
		c.removeNode(tail)
		delete(c.cache, tail.Key)
	} else {
		c.capacity--
	}

	c.cache[key] = node
	c.addNode(node)
}

func (c *LRUCache) addNode(node *Node) {
	node.Prev = c.head
	node.Next = c.head.Next
	node.Next.Prev = node
	c.head.Next = node
}

func (c *LRUCache) removeNode(node *Node) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}
