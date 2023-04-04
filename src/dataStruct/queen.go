package main

type Queue struct {
	elements []interface{}
}

func (q *Queue) Push(value interface{}) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Pop() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
