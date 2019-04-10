package component

import (
    "time"
)

const PlayerSignature = 1 << PLAYER

type Player struct {
    Channel map[string]chan time.Time
}

func (i *Player) Signature() uint64 {
    return PlayerSignature
}
