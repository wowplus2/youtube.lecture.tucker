package dataStruct

import "fmt"

/*
	Heap Tree 에서...
	n번째 node의
		왼쪽 자식 node 값  = 2 * n + 1
		오른쪽 자식 node 값 = 2 * n + 2
		부모 node 값 = n - 1 / 2
*/
type Heap struct {
	list []int
}

func (h *Heap) PushHeapElement(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 { // 위로 올라가면서 비교
		pidx := (idx - 1) / 2 // parent node index
		if pidx < 0 {         // 나의 위치가 root 이다.
			break
		}

		if h.list[idx] < h.list[pidx] { // 나의 값이 부모값보다 작으면 swap
			h.list[idx], h.list[pidx] = h.list[pidx], h.list[idx]
			idx = pidx
		} else {
			break
		}
	}
}

func (h *Heap) PopHeapElement() int {
	if len(h.list) == 0 { // pop할 대상이 없다.
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return top
	}

	h.list[0] = last
	idx := 0
	for idx < len(h.list) { // 하향식 검사.
		swapidx := -1
		lidx := idx*2 + 1 // left child node index

		if lidx >= len(h.list) { // 왼쪽 자식 node가 없다.
			break
		}

		if h.list[lidx] < h.list[idx] { // 왼쪽 자식 node가 현재 값보다 작다 -> swap target
			swapidx = lidx
		}

		ridx := idx*2 + 2       // right child node index
		if ridx < len(h.list) { // 오른쪽 자식 node가 있다.
			if h.list[ridx] < h.list[idx] { // 왼쪽 node와 비교한다.
				if swapidx < 0 || h.list[swapidx] > h.list[ridx] {
					swapidx = ridx
				}
			}
		}

		if swapidx < 0 { // swap 대상이 없다.
			break
		}

		h.list[idx], h.list[swapidx] = h.list[swapidx], h.list[idx]
		idx = swapidx
	}

	return top
}

func (h *Heap) PrintHeapElements() {
	fmt.Println(h.list)
}

func (h *Heap) GetHeapCount() int {
	return len(h.list)
}
