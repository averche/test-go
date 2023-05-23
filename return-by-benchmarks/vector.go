package main

import "fmt"

type Vector3D struct {
	x, y, z float64
}

func (p *Vector3D) Add(other *Vector3D) {
	p.x += other.x
	p.y += other.y
	p.z += other.z
}

func (p *Vector3D) AddByValue(other Vector3D) {
	p.x += other.x
	p.y += other.y
	p.z += other.z
}

func AddReturnCopy(v1, v2 Vector3D) Vector3D {
	v1.x += v2.x
	v1.y += v2.y
	v1.z += v2.z

	return v1
}

func (v *Vector3D) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.x, v.y, v.z)
}

func main() {
	v1 := Vector3D{1.0, 2.0, 3.0}
	v2 := Vector3D{4.0, 5.0, 6.0}

	v1.Add(&v2)

	fmt.Println(v1)
}
