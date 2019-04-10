package camera

import (
    "github.com/faiface/pixel/pixelgl"
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
)

type componentManager struct {
}

func newComponentManager() *componentManager {
    return &componentManager{}
}

func (m *componentManager) newJoint(entity *core.Entity) *component.Joint {
    return &component.Joint{
        Entity: entity,
        Anchor: &r2.Point{},
    }
}

func (m *componentManager) newTransform() *component.Transform {
    return &component.Transform{Point: &r2.Point{}}
}

func (m *componentManager) newWindow(win *pixelgl.Window) *component.Window {
    // defined by state
    return &component.Window{
        Window: win,
        Scale:  1,
    }
}
