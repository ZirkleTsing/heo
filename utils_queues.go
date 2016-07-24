package acogo

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	var queue = &Queue{
		Nodes: make([]interface{}, size),
		Size:  size,
	}

	return queue
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	Nodes []interface{}
	Size  int
	Head  int
	Tail  int
	Count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n interface{}) {
	if q.Head == q.Tail && q.Count > 0 {
		nodes := make([]interface{}, len(q.Nodes) + q.Size)
		copy(nodes, q.Nodes[q.Head:])
		copy(nodes[len(q.Nodes) - q.Head:], q.Nodes[:q.Head])
		q.Head = 0
		q.Tail = len(q.Nodes)
		q.Nodes = nodes
	}
	q.Nodes[q.Tail] = n
	q.Tail = (q.Tail + 1) % len(q.Nodes)
	q.Count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() interface{} {
	if q.Count == 0 {
		return nil
	}
	node := q.Nodes[q.Head]
	q.Head = (q.Head + 1) % len(q.Nodes)
	q.Count--
	return node
}

func (q *Queue) Peek() interface{} {
	return q.Nodes[q.Head]
}
