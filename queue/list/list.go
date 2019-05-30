package list

//Node node
type Node struct {
	value interface{}
	next  *Node
}

//NewNode create one node
func NewNode(value interface{}) *Node {
	var node Node
	node.value = value
	node.next = nil
	return &node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q Queue) GetSize() int {
	return q.size
}

func NewQueue() *Queue {
	var queue Queue
	queue.head = nil
	queue.tail = nil
	queue.size = 0
	return &queue
}
func (q *Queue) Push(value interface{}) {
	if q.head == nil || q.tail == nil {
		q.head = NewNode(value)
		q.tail = q.head
		return
	}
	temp := q.tail
	temp.next = NewNode(value)
	q.tail = temp.next
	q.size++
}
func (q *Queue) Pop() *Node {
	if q.head == nil || q.tail == nil {
		return nil
	}
	temp := q.head
	q.head = temp.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return temp
}
func (q Queue)IsEmpty()bool{
	if q.head == nil || q.tail == nil {
		return true
	}
	return false
}
