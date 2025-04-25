/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-25 11:21:55
 * @LastEditTime: 2025-04-25 11:57:06
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /goqueue/queue_test.go
 */

package goqueue

import "testing"

const TEST_SCALE int = 100000000

func TestQueue(t *testing.T) {
	data := make([]int, TEST_SCALE)
	q := make(Queue[int], 0)
	for i := 0; i < TEST_SCALE; i++ {
		data[i] = i
		q.Push(i)
		front := q.Front()
		back := q.Back()
		if front != data[0] {
			t.Errorf("Front should be %d, but got %d", data[0], front)
		}
		if back != data[i] {
			t.Errorf("Front should be %d, but got %d", data[i], back)
		}
		if q.Empty() {
			t.Errorf("Queue is not empty, but got it is")
		}
	}
	for i := 0; i < TEST_SCALE; i++ {
		if q.Empty() {
			t.Errorf("Queue is not empty, but got it is")
		}
		front := q.Front()
		back := q.Back()
		if back != data[TEST_SCALE-1] {
			t.Errorf("Back should be %d, but got %d", data[TEST_SCALE-1], back)
		}
		if front != data[i] {
			t.Errorf("Front should be %d, but got %d", data[i], front)
		}
		poped := q.Pop()
		if poped != data[i] {
			t.Errorf("Poped should be %d, but got %d", data[i], poped)
		}
	}
	if !q.Empty() {
		t.Errorf("Queue is empty, but got it not")
	}
}
