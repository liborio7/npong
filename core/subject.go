package core

import (
    "github.com/satori/go.uuid"
    "sync"
)

type Subject struct {
    observersMutex sync.RWMutex
    Observers      map[uint64]map[uuid.UUID]Observer
}

func (ms *Subject) AddObserver(o Observer) {
    ms.observersMutex.Lock()
    if os := ms.Observers[o.Signature()]; os == nil {
        ms.Observers[o.Signature()] = make(map[uuid.UUID]Observer, 0)
    }
    ms.Observers[o.Signature()][o.Id()] = o
    ms.observersMutex.Unlock()
}

func (ms *Subject) RemoveObserver(o Observer) {
    ms.observersMutex.Lock()
    delete(ms.Observers[o.Signature()], o.Id())
    ms.observersMutex.Unlock()
}

func (ms *Subject) Notify(e *Entity, signature uint64) {
    ms.observersMutex.RLock()
    for mapSignature, mapO := range ms.Observers {
        if signature&mapSignature == signature {
            for _, o := range mapO {
                o.OnEvent(e)
            }
        }
    }
    ms.observersMutex.RUnlock()
}
