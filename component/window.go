package component

import (
    "github.com/faiface/pixel/pixelgl"
)

const WindowSignature = 1 << WINDOW

type Window struct {
    *pixelgl.Window
    // delete window in order to remove dependency from external lib (faiface)
    Scale float64
}

func (j *Window) Signature() uint64 {
    return WindowSignature
}
