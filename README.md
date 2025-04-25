<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-25 10:29:33
 * @LastEditTime: 2025-04-25 11:59:54
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /goqueue/README.md
-->

# goqueue

Queue in Golang.

Well tested.

## Usage

``` go
import "github.com/FunctionSir/goqueue"
// ... //
q := make(goqueue.Queue[int], 0) // In most cases, must be 0.
if q.Empty() {
    // ... //
}
if q.Size() > 10 {
    // ... //
}
q.Push(1)
poped := q.Pop()
front := q.Front()
back := q.Back()
// ... //
```
