package main

import "fmt"

type Edge struct {
	t, w int
}

type MinHeap struct {
	heap []Edge
	len  int
}

func (h *MinHeap) swap(i, j int) {
	tmp := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = tmp
}

func (h *MinHeap) bubbleUp(i int) {
	parent := i / 2
	if i > 1 && h.heap[i].w < h.heap[parent].w {
		h.swap(i, parent)
		h.bubbleUp(parent)
	}
}

func (h *MinHeap) bubbleDown(i int) {
	left := 2 * i
	right := 2*i + 1
	if left <= h.len {
		child := left
		if right <= h.len && h.heap[right].w < h.heap[left].w {
			child = right
		}
		if h.heap[child].w < h.heap[i].w {
			h.swap(i, child)
			h.bubbleDown(child)
		}
	}
}

func (h *MinHeap) push(e Edge) {
	h.len += 1
	h.heap[h.len] = e
	h.bubbleUp(h.len)
}

func (h *MinHeap) pop() Edge {
	e := h.heap[1]
	h.heap[1] = h.heap[h.len]
	h.len -= 1
	h.bubbleDown(1)
	return e
}

func newHeap(n int) MinHeap {
	h := MinHeap{}
	h.heap = make([]Edge, n+1)
	h.len = 0
	return h
}

// func newMatrix(n, m int) [][]int {
// 	matrix := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			matrix[i] = append(matrix[i], 0)
// 		}
// 	}
// 	return matrix
// }

func aFeastForCats() {
	var m, c int
	fmt.Scan(&m)
	fmt.Scan(&c)
	n := c * (c - 1) / 2
	// adj_matrix := newMatrix(c, c)

	// for i := 0; i < combinations; i++ {
	// 	var x, y, d int
	// 	fmt.Scan(&x)
	// 	fmt.Scan(&y)
	// 	fmt.Scan(&d)
	// 	adj_matrix[x][y] = d
	// 	adj_matrix[y][x] = d
	// }

	adj_list := make([][]Edge, c)
	for i := 0; i < n; i++ {
		var x, y, d int
		fmt.Scan(&x)
		fmt.Scan(&y)
		fmt.Scan(&d)
		adj_list[x] = append(adj_list[x], Edge{t: y, w: d})
		adj_list[y] = append(adj_list[y], Edge{t: x, w: d})
	}
	// fmt.Println(adj_list)

	q := newHeap(n)
	q.push(Edge{t: 0, w: 0})

	visited := make([]bool, c)
	for i := 0; i < c; i++ {
		visited[i] = false
	}

	min_distance := 0

	for i := 0; i < c-1; {
		fmt.Println(q)
		e := q.pop()
		fmt.Println(e)
		if visited[e.t] == false {
			visited[e.t] = true
			fmt.Println(visited)
			min_distance += e.w
			for j := 0; j < len(adj_list[e.t]); j++ {
				q.push(adj_list[e.t][j])
			}
			i++
		}
	}

	if min_distance+c > m {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		aFeastForCats()
	}

	// h := newHeap(10)
	// h.push(Edge{t: 1, w: 42})
	// h.push(Edge{t: 1, w: 56})
	// h.push(Edge{t: 1, w: 45})
	// h.push(Edge{t: 1, w: 11})
	// fmt.Println(h)
	// fmt.Println(h.pop())
	// fmt.Println(h)
}
