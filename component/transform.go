package component

import (
    "github.com/golang/geo/r2"
)

const TransformSignature = 1 << TRANSFORM

type Transform struct {
    *r2.Point
    OnUpdate func(*r2.Point)
}

func (t *Transform) Signature() uint64 {
    return TransformSignature
}
