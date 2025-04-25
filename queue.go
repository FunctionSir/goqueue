/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-25 11:00:01
 * @LastEditTime: 2025-04-25 12:03:28
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /goqueue/queue.go
 */

package goqueue

type Queue[T any] []T

// Get size of the queue.
func (q Queue[T]) Size() int {
	return len(q)
}

// Check the queue is empty or not.
func (q Queue[T]) Empty() bool {
	return q.Size() == 0
}

// Push new element into queue.
func (q *Queue[T]) Push(x T) {
	(*q) = append((*q), x)
}

// Pop the front of queue.
func (q *Queue[T]) Pop() T {
	poped := (*q)[0]
	(*q) = (*q)[1:]
	return poped
}

// Get the first element of queue.
func (q Queue[T]) Front() T {
	return q[0]
}

// Get the last element of queue.
func (q Queue[T]) Back() T {
	return q[q.Size()-1]
}
