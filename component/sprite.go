package component

import (
    "github.com/faiface/pixel"
    "github.com/golang/geo/r2"
    "image/color"
    "time"
)

const SpriteSignature = 1 << SPRITE

type Sprite struct {
    RGBA       *color.RGBA
    Pic        pixel.Picture
    PicId      string
    Frame      *r2.Rect
    FrameIndex float64
    LastUpdate time.Time
}

func (t *Sprite) Signature() uint64 {
    return SpriteSignature
}
