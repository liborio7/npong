package component

const GroundSignature = 1 << GROUND

const (
    TerrainGround = iota
)

type Ground struct {
    Type       int
    Friction   float64
    Bounciness float64
}

func (Ground) Signature() uint64 {
    return GroundSignature
}
