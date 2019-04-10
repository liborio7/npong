package terrain

import (
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/component"
    "image/color"
)

type componentManager struct {
}

func newComponentManager() *componentManager {
    return &componentManager{}
}

func (m *componentManager) newGround() *component.Ground {
    return &component.Ground{
        Type: component.TerrainGround,
    }
}

func (m *componentManager) newTransform(point *r2.Point) *component.Transform {
    return &component.Transform{Point: point}
}

func (m *componentManager) newCollider(rect *r2.Rect) *component.Collider {
    return &component.Collider{
        Rect: rect,
        Type: component.RigidbodyCollider,
    }
}

func (m *componentManager) newSprite(rgba *color.RGBA, rect *r2.Rect) *component.Sprite {
    return &component.Sprite{
        RGBA:  rgba,
        Frame: rect,
    }
}
