package main

type PriorityQueue[T any] struct {
	elements []T
	priorities []float64
}

func (Q *PriorityQueue[T]) Peek() (T, float64) {
	return Q.elements[0], Q.priorities[0]
}

func (Q *PriorityQueue[T]) Queue(element T, priority float64) {

	index := Q.FindSlot(priority)

	if index == 0 {
		Q.elements = append([]T {element}, Q.elements...)
		Q.priorities = append([]float64 {priority}, Q.priorities...)
	} else if index == len(Q.elements) {
		Q.elements = append(Q.elements, element)
		Q.priorities = append(Q.priorities, priority)
	} else {
		Q.elements = append(Q.elements[:index+1], Q.elements[index:]...)
		Q.elements[index] = element
		Q.priorities = append(Q.priorities[:index+1], Q.priorities[index:]...)
		Q.priorities[index] = priority
	}
}

func (Q *PriorityQueue[T]) Dequeue() (T, float64) {
	element := Q.elements[0]
	priority := Q.priorities[0]

	Q.elements = Q.elements[1:]
	Q.priorities = Q.priorities[1:]

	return element, priority
}

func (Q *PriorityQueue[T]) Any() bool {
	return len(Q.elements) > 0
}

func (Q *PriorityQueue[T]) Size() int {
	return len(Q.elements)
}

func (Q *PriorityQueue[T]) ToStack() Stack[T] {
	var stack Stack[T]

	stack.elements = Q.elements

	return stack
}

func (Q *PriorityQueue[T]) FindSlot(priority float64) int {

	size := len(Q.elements)

	index := 0

	for i := range size {
		if priority >= Q.priorities[i] {
			index = i
			break
		} else {
			index++
		}
	}

	return index
}