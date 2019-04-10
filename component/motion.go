package component

const MotionSignature = 1 << MOTION

type Motion struct {
    // UseGravity bool ?
    SpeedX        *MotionSpeed
    SpeedY        *MotionSpeed
    AccelerationX *MotionAcceleration
    AccelerationY *MotionAcceleration
}

type MotionSpeed struct {
    Max, Min, Value float64
}

type MotionAcceleration struct {
    Max, Min, Value float64
}

func (t *Motion) Signature() uint64 {
    return MotionSignature
}
