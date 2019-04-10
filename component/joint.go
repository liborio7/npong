package component

import (
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
)

const JointSignature = 1 << JOINT

type Joint struct {
    Entity *core.Entity
    Anchor *r2.Point
}

func (j *Joint) Signature() uint64 {
    return JointSignature
}
