package geo

import "k8s.io/klog/v2"

type Geo struct {
	x int
	y int
}

func (g *Geo) New(x int, y int) {
	klog.V(4).Info("Creating Geo")
	g.x = x
	g.y = y
}

func (g *Geo) Multiply() int {
	klog.V(2).Info("Doing multiplication")
	return g.x * g.y
}
