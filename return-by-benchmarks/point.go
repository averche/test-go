package main

import "fmt"

type Point struct {
	x, y, z          float64
	str1, str2, str3 string
	i1, i2, i3       int
	b1, b2, b3       bool
}

func (p *Point) Add(other *Point) {
	p.x += other.x
	p.y += other.y
	p.z += other.z
}

func (p *Point) AddByValue(other Point) {
	p.x += other.x
	p.y += other.y
	p.z += other.z
}

func AddReturnCopy(v1, v2 Point) Point {
	v1.x += v2.x
	v1.y += v2.y
	v1.z += v2.z

	return v1
}

func (v *Point) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.x, v.y, v.z)
}

func main() {
	v1 := Point{
		x: 1.0,
		y: 2.0,
		z: 3.0,
	}
	v2 := Point{
		x: 4.0,
		y: 5.0,
		z: 6.0,
	}

	v1.Add(&v2)

	fmt.Println(v1)
}
