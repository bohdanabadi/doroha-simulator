package datastructures

import (
	"simulator/internal/dto"
)

type PointNodeWrapper struct {
	Node                     *dto.PointNode
	TotalCostFromStart       float64
	MinRemainingCostToTarget float64
	CameFrom                 *PointNodeWrapper
	CostSum                  float64
	Index                    int
}

func (p *PointNodeWrapper) SetValues(node *dto.PointNode, totalCostFromStart, minRemainingCostToTarget, costSum float64,
	index int, cameFrom *PointNodeWrapper) {
	p.Node = node
	p.TotalCostFromStart = totalCostFromStart
	p.MinRemainingCostToTarget = minRemainingCostToTarget
	p.CostSum = costSum
	p.Index = index
	p.CameFrom = cameFrom
}

func (p *PointNodeWrapper) SetTotalCostAndMinRemaining(totalCostFromStart, minRemainingCostToTarget float64) {
	p.TotalCostFromStart = totalCostFromStart
	p.MinRemainingCostToTarget = minRemainingCostToTarget
	p.CostSum = p.TotalCostFromStart + p.MinRemainingCostToTarget
}

type PointNodeWrapperHeap []*PointNodeWrapper

func (h PointNodeWrapperHeap) Len() int {
	return len(h)
}

func (h PointNodeWrapperHeap) Less(i, j int) bool {
	// We want the minimum CostSum at the top
	return h[i].CostSum < h[j].CostSum
}

func (h PointNodeWrapperHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *PointNodeWrapperHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*PointNodeWrapper)
	item.Index = n
	*h = append(*h, item)
}

func (h *PointNodeWrapperHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*h = old[0 : n-1]
	return item
}
