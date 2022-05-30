package math

import "sort"

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	var res float64 = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 0; k < n; k++ {
				current := area(points[i], points[j], points[k])
				if current >= res {
					res = current
				}
			}
		}
	}
	return res
}

func area(points ...[]int) float64 {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	leftPoint := points[0]
	rightPoint := points[2]

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	buttonPoint := points[0]
	topPoint := points[2]

	allArea := float64(abs((topPoint[1] - buttonPoint[1]) * (rightPoint[0] - leftPoint[0])))
	area1 := float64(abs((topPoint[1] - rightPoint[1]) * (rightPoint[0] - topPoint[0])))
	area2 := float64(abs((topPoint[1] - leftPoint[1]) * (leftPoint[0] - topPoint[0])))
	area3 := float64(abs((leftPoint[0] - rightPoint[0]) * (leftPoint[1] - rightPoint[1])))

	res := allArea - (area1+area2+area3)/2
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
