package yagrt

import (
	"fmt"
	"sort"
)

type BVHNode struct {
	Left      *BVHNode
	Right     *BVHNode
	Shape     Shape
	splitAxis Axis
}

type shapesOnX []Shape

func (shapes shapesOnX) Len() int {
	return len(shapes)
}

func (shapes shapesOnX) Less(i, j int) bool {
	ci := shapes[i].BoundingBox().Center()
	cj := shapes[j].BoundingBox().Center()
	return ci.X < cj.X
}

func (shapes shapesOnX) Swap(i, j int) {
	shapes[i], shapes[j] = shapes[j], shapes[i]
}

type shapesOnY []Shape

func (shapes shapesOnY) Len() int {
	return len(shapes)
}

func (shapes shapesOnY) Less(i, j int) bool {
	ci := shapes[i].BoundingBox().Center()
	cj := shapes[j].BoundingBox().Center()
	return ci.Y < cj.Y
}

func (shapes shapesOnY) Swap(i, j int) {
	shapes[i], shapes[j] = shapes[j], shapes[i]
}

type shapesOnZ []Shape

func (shapes shapesOnZ) Len() int {
	return len(shapes)
}

func (shapes shapesOnZ) Less(i, j int) bool {
	ci := shapes[i].BoundingBox().Center()
	cj := shapes[j].BoundingBox().Center()
	return ci.Z < cj.Z
}

func (shapes shapesOnZ) Swap(i, j int) {
	shapes[i], shapes[j] = shapes[j], shapes[i]
}

func NewBVH(shapes []Shape) *BVHNode {
	if len(shapes) == 0 {
		return nil
	}
	if len(shapes) == 1 {
		return &BVHNode{nil, nil, shapes[0], 0}
	}
	bb := shapes[0].BoundingBox()
	for _, s := range shapes {
		bb.Extend(s.BoundingBox())
	}
	axis := bb.SplitAxis()
	switch axis {
	case AxisX:
		sort.Sort(shapesOnX(shapes))
	case AxisY:
		sort.Sort(shapesOnY(shapes))
	case AxisZ:
		sort.Sort(shapesOnZ(shapes))
	default:
		sort.Sort(shapesOnX(shapes))
	}
	return &BVHNode{
		Left:      NewBVH(shapes[:int(len(shapes)/2)]),
		Right:     NewBVH(shapes[int(len(shapes)/2):]),
		Shape:     &bb,
		splitAxis: axis,
	}
}

func (n *BVHNode) Intersect(r Ray, hit *Hit) bool {
	if n == nil {
		hit.Shape = nil
		return false
	}
	if isHit := n.Shape.Intersect(r, hit); !isHit || (n.Left == nil && n.Right == nil) {
		return true
	}

	var otherHit Hit
	rayDir := 0.0
	switch n.splitAxis {
	case AxisX:
		rayDir = r.Dir.X
	case AxisY:
		rayDir = r.Dir.Y
	case AxisZ:
		rayDir = r.Dir.Z
	}
	var firstHit, secondHit bool
	if rayDir >= 0 {
		firstHit = n.Left.Intersect(r, hit)
		secondHit = n.Right.Intersect(r, &otherHit)
	} else {
		firstHit = n.Right.Intersect(r, hit)
		secondHit = n.Left.Intersect(r, &otherHit)
	}
	if secondHit && (!firstHit || hit.T > otherHit.T) {
		*hit = otherHit
	}
	return true
}

func (n *BVHNode) BoundingBox() Box {
	return n.Shape.BoundingBox()
}

func (n *BVHNode) DebugPrint(indentation int) {
	for i := 0; i < indentation; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("Node: %T\n", n.Shape)
	if n.Left != nil {
		n.Left.DebugPrint(indentation + 1)
	}
	if n.Right != nil {
		n.Right.DebugPrint(indentation + 1)
	}
}
