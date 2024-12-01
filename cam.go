package cam

import (
	"github.com/Xsf57i2G/geom"
)

type Camera struct {
	Position, Direction, Up, Right geom.Vec
}

func (c *Camera) look(target geom.Vec) {
	c.Direction = target.Sub(c.Position).Norm()
	c.Right = c.Direction.Cross(c.Up).Norm()
	c.Up = c.Right.Cross(c.Direction)
}

func (c *Camera) Project(p geom.Vec) geom.Vec {
	var view = p.Sub(c.Position)
	var depth = view.Dot(c.Direction)
	var right = view.Dot(c.Right)
	var up = view.Dot(c.Up)
	return geom.Vec{
		X: right / depth,
		Y: up / depth,
		Z: depth,
	}
}
