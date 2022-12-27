package queue

type (
	Queue[QT any] struct {
		start, end *node[QT]
		length     int
	}
	node[NT any] struct {
		value NT
		next  *node[NT]
	}
)

// Create a new queue
func New[T any]() *Queue[T] {
	return &Queue[T]{nil, nil, 0}
}

// Take the next item off the front of the queue
func (this *Queue[QT]) Dequeue() (res QT, exists bool) {
	if this.length == 0 {
		exists = false
		return
	}
	exists = true
	res = this.start.value
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return
}

// Put an item on the end of a queue
func (this *Queue[QT]) Enqueue(value QT) {
	n := &node[QT]{value, nil}
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}

// Return the number of items in the queue
func (this *Queue[QT]) Len() int {
	return this.length
}

// Return the first item in the queue without removing it
func (this *Queue[QT]) Peek() (res QT, exists bool) {
	if this.length == 0 {
		exists = false
		return
	}
	return this.start.value, true
}
