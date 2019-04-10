package camera

import (
    "github.com/faiface/pixel/pixelgl"
    "github.com/liborio7/npong/core"
)

type Manager struct {
    ComponentManager *componentManager
    EntityManager    *entityManager
}

func NewManager() *Manager {
    componentManager := newComponentManager()
    entityManager := newEntityManager()

    return &Manager{
        ComponentManager: componentManager,
        EntityManager:    entityManager,
    }
}

func (m *Manager) NewEntity(win *pixelgl.Window, jointEntity *core.Entity) *core.Entity {
    return m.EntityManager.newEntity(m, win, jointEntity)
}
