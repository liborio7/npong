package terrain

import (
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
    "image/color"
)

type Manager struct {
    ComponentManager *componentManager
    EntityManager    *entityManager
}

func NewManager() *Manager {
    componentManager := newComponentManager()
    entityManager := newEntityManager()

    return &Manager{ComponentManager: componentManager, EntityManager: entityManager}
}

func (m *Manager) NewEntity(rgba *color.RGBA, point *r2.Point, rect *r2.Rect) *core.Entity {
    return m.EntityManager.NewEntity(m, rgba, point, rect)
}
