package player

import (
    "github.com/faiface/pixel/pixelgl"
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
)

type Manager struct {
    CommandManager   *commandManager
    ComponentManager *componentManager
    StateManager     *stateManager
    EntityManager    *entityManager
}

func NewManager() *Manager {
    commandManager := newCommandManager()
    componentManager := newComponentManager()
    stateManager := newStateManager(commandManager)
    entityManager := newEntityManager()

    return &Manager{
        CommandManager:   commandManager,
        ComponentManager: componentManager,
        StateManager:     stateManager,
        EntityManager:    entityManager,
    }
}

func (m *Manager) NewEntity(win *pixelgl.Window, point *r2.Point) *core.Entity {
    return m.EntityManager.newEntity(m, win, point)
}
