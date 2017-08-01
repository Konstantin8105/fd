package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/output/displacement"
	"github.com/Konstantin8105/GoFea/output/forceLocal"
	"github.com/Konstantin8105/GoFea/output/reaction"
)

type forceCase2d struct {
	indexCase int

	// input data
	//TODO: gravityForces []gravityForce2d
	nodeForces []nodeForce2d

	// statis property
	static staticTypes

	// dynamic property
	dynamicType  dynamicTypes
	dynamicValue []float64

	// output data
	globalDisplacements []displacement.PointDim2
	localDisplacement   []displacement.BeamDim2
	localForces         []forceLocal.BeamDim2
	reactions           []reaction.Dim2
}

type staticTypes bool

const (
	linear staticTypes = false
	nolinear
)

type dynamicTypes int

const (
	none dynamicTypes = iota
	naturalFrequency
	bucklingFactors
)

// GetGlobalDisplacement - return global displacement
func (f *forceCase2d) GetGlobalDisplacement(pointIndex point.Index) (d displacement.Dim2, err error) {
	for _, g := range f.globalDisplacements {
		if g.Index == pointIndex {
			return g.Dim2, nil
		}
	}
	return d, fmt.Errorf("Cannot found point")
}

// GetLocalForce - return local force of beam
func (f *forceCase2d) GetLocalForce(beamIndex element.Index) (begin, end forceLocal.Dim2, err error) {
	for _, l := range f.localForces {
		if l.Index == beamIndex {
			return l.Begin, l.End, nil
		}
	}
	return begin, end, fmt.Errorf("Cannot found beam")
}

// GetReaction - return reaction of support
func (f *forceCase2d) GetReaction(pointIndex point.Index) (r reaction.Dim2, err error) {
	for _, g := range f.reactions {
		if g.Index == pointIndex {
			return g, nil
		}
	}
	return r, fmt.Errorf("Cannot found point")
}

// GetNaturalFrequency - return natural frequency
func (f *forceCase2d) GetNaturalFrequency() (hz []float64, err error) {
	if f.dynamicType != naturalFrequency {
		return hz, fmt.Errorf("Natural frequency is not calculate for that case")
	}
	return f.dynamicValue, nil
}

func (f *forceCase2d) check() (err error) {
	err = isUniqueIndexes(nodeForceByPoint(f.nodeForces))
	if err != nil {
		return fmt.Errorf("Errors in case %v in node forces:\n%v", f.indexCase, err)
	}
	return nil
}

/*
func zeroCopy(f forceCase2d) (result forceCase2d) {
	result.indexCase = f.indexCase
	result.gravityForces = make([]gravityForce2d, len(f.gravityForces))
	result.nodeForce2d = make([]nodeForce2d, len(f.nodeForces))
	return
}

func delta(a, b forceCase2d) (d float64) {
	for i := range a.gravityForces {
		d = math.Max(d, math.Abs(a.gravityForces[i]-b.gravityForces[i])/math.Max(math.Abs(a.gravityForces[i]), math.Abs(b.gravityForces[i])))
	}
	for i := range a.nodeForces {
		d = math.Max(d, math.Abs(a.nodeForces[i]-b.nodeForces[i])/math.Max(math.Abs(a.nodeForces[i]), math.Abs(b.nodeForces[i])))
	}
	return
}
*/
