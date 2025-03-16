package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type Point struct {
	x, y int
}

type Item struct {
	point    Point
	priority int
	index    int
}

// utils untuk interface heap di golang
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

type KotakKatik struct {
	papan       [][]int
	solusiJalan []Point
}

func NewKotakKatik(papan [][]int) *KotakKatik {
	return &KotakKatik{
		papan: papan,
	}
}

func (k *KotakKatik) temukanJalan() []Point { // menggunakan algoritna djikstra
	rows, cols := len(k.papan), len(k.papan[0])
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // arah kanan, bawah, kiri, atas

	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = 1<<31 - 1
		}
	}
	prev := make(map[Point]Point)

	dist[0][0] = 0
	pq := &PriorityQueue{}
	heap.Push(pq, &Item{Point{0, 0}, 0, 0})

	for pq.Len() > 0 {
		// implementasi djikstra
		current := heap.Pop(pq).(*Item)
		curPoint := current.point

		if curPoint == (Point{rows - 1, cols - 1}) {
			break
		}
		for _, d := range directions {
			nx, ny := curPoint.x+d.x, curPoint.y+d.y
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols { // koordinat point harus berada di dalam papan
				newCost := dist[curPoint.x][curPoint.y] + k.papan[nx][ny]
				if newCost < dist[nx][ny] {
					dist[nx][ny] = newCost
					prev[Point{nx, ny}] = curPoint
					heap.Push(pq, &Item{Point{nx, ny}, newCost, 0})
				}
			}
		}
	}

	path := []Point{}
	at := Point{rows - 1, cols - 1}

	for at != (Point{0, 0}) {
		path = append([]Point{{at.y, at.x}}, path...)

		prevAt, exists := prev[at]
		if !exists {
			break
		}
		at = prevAt
	}

	if at == (Point{0, 0}) {
		path = append([]Point{{0, 0}}, path...)
		k.solusiJalan = path
	}

	return path
}

func (k *KotakKatik) totalEnergi() (int, error) {
	if len(k.solusiJalan) == 0 {
		return 0, errors.New("jalan belum ditemukan, jalankan temukanJalan() terlebih dahulu")
	}
	total := 0
	for _, p := range k.solusiJalan {
		total += k.papan[p.y][p.x]
	}
	return total, nil
}

func main() {
	// Tambahan testcase untuk mengecek kinerja algoritma
	testCases := []struct {
		papan [][]int
	}{
		{
			papan: [][]int{
				{0, 1, 1, 7, 6, 4},
				{4, 6, 2, 8, 6, 1},
				{2, 1, 1, 1, 8, 4},
				{8, 7, 4, 9, 1, 1},
				{8, 8, 6, 7, 9, 2},
				{8, 8, 5, 2, 6, 0},
			},
		},
		{
			papan: [][]int{
				{0, 2, 3, 4, 5, 6},
				{1, 2, 3, 4, 5, 6},
				{1, 1, 1, 1, 1, 1},
				{9, 8, 7, 6, 5, 4},
				{9, 8, 7, 6, 5, 4},
				{9, 8, 7, 6, 5, 0},
			},
		},
		{
			papan: [][]int{
				{0, 9, 9, 9, 9, 9},
				{1, 1, 1, 1, 9, 9},
				{9, 9, 9, 1, 9, 9},
				{9, 2, 1, 1, 9, 9},
				{9, 1, 9, 9, 9, 9},
				{9, 1, 1, 1, 1, 0},
			},
		},
	}

	for index, tc := range testCases {
		fmt.Println("====================================")
		fmt.Println("Test Case:", "Papan", index+1)

		permainan := NewKotakKatik(tc.papan)
		jalan := permainan.temukanJalan()
		fmt.Println("Jalan ditemukan:", jalan)

		total, err := permainan.totalEnergi()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Total energi:", total)
		}
	}
}
