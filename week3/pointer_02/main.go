package main

import "fmt"

type Point struct {
	x int
	y int
}

func scalePoint(p *Point, factor float64) {
	p.x = int(float64(p.x) * factor)
	p.y = int(float64(p.x) * factor)
}

func main() {
	// 创建 Point 实例
	p := Point{x: 3, y: 4}
	fmt.Printf("原始坐标：x=%d, y=%d\n", p.x, p.y)

	// 调用 scalePoint 进行缩放
	scalePoint(&p, 2.5)

	// 打印修改后的结果
	fmt.Printf("缩放后坐标：x=%d, y=%d\n", p.x, p.y)
}
