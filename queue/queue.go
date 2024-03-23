package queue

type Queue struct {
	bound int
	data  []int
}

func New(bound int) *Queue {
	return &Queue{
		bound: bound,
		data:  make([]int, bound),
	}
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() {
	q.data = q.data[1:]
}

func (q *Queue) GetData() []int {
	return q.data
}

func (q *Queue) IsQueueFull() bool {
	return len(q.data) == q.bound
}
