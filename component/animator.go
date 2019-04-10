package component

import (
    "github.com/faiface/pixel"
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
)

const AnimatorSignature = 1 << ANIMATOR

type Animator struct {
    Pics       map[string]pixel.Picture
    Frames     map[string]*r2.Rect
    SpritesNum map[string]float64
    Animate    func(*core.Entity) string
}

func (Animator) Signature() uint64 {
    return AnimatorSignature
}