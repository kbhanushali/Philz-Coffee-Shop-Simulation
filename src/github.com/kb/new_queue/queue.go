package queue
 
type Cust struct {
	ip string
    name string
    port string
}
 

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Cust, size),
		size:  size,
	}
}
 
// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Cust
	size  int
	head  int
	tail  int
	count int
}
 
// Push adds a node to the queue.
func (q *Queue) Push(ip, name, port string) {
	n:= new(Cust)
	n.ip=ip
	n.name=name
	n.port=port
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Cust, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}
 
// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() (string,string,string) {
	if q.count == 0 {
		return "", "", ""
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node.ip, node.name, node.port
}

//Is empty check for queue
func (q *Queue) IsEmpty() (bool) {
	if q.count == 0 {
		return true
	} else {
		return false
	}
}

	//check for queue Size
func (q *Queue) Size() (int) {
		return q.count
}
