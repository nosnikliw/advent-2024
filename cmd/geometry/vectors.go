package geometry

type Line struct {
	P1 Vector
	P2 Vector
}

type VectorBox struct {
	V1 Vector
	V2 Vector
}

func (s VectorBox) Contains(v Vector) bool {
	return ((s.V1.X-v.X)*(s.V2.X-v.X) <= 0) && ((s.V1.Y-v.Y)*(s.V2.Y-v.Y) <= 0)
}

type Vector struct {
	X int
	Y int
}

func (dv1 Vector) IsParallelTo(dv2 Vector) bool {
	return float64(dv1.X)/float64(dv1.Y) == float64(dv2.X)/float64(dv2.Y)
}

func (from Vector) To(to Vector) Vector {
	return to.Subtract(from)
}

func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Vector) Subtract(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Vector) Equals(b Vector) bool {
	return a.X == b.X && a.Y == b.Y
}

func Box[T any](area [][]T) VectorBox {
	return VectorBox{
		V1: Vector{X: 0, Y: 0},
		V2: Vector{X: len(area) - 1, Y: len(area[0]) - 1},
	}
}
