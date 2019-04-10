package terrain

import (
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
    "image/color"
)

type entityManager struct {
}

func newEntityManager() *entityManager {
    return &entityManager{}
}

func (em *entityManager) NewEntity(m *Manager, rgba *color.RGBA, point *r2.Point, rect *r2.Rect) *core.Entity {
    ground := m.ComponentManager.newGround()
    transform := m.ComponentManager.newTransform(point)
    collision := m.ComponentManager.newCollider(rect)
    sprite := m.ComponentManager.newSprite(rgba, rect)

    entity := core.NewEntity()
    entity.AddComponents(ground, transform, collision, sprite, )

    return entity
}
