package tank

type ITank interface {
	getName() string
	getPower() uint64
	getSpeed() uint64
	setName(name string)
	setPower(power uint64)
	setSpeed(speed uint64)
}

type tank struct {
	name  string
	power uint64
	speed uint64
}

func (t *tank) getName() string {
	return t.name
}

func (t *tank) getPower() uint64 {
	return t.power
}

func (t *tank) getSpeed() uint64 {
	return t.speed
}

func (t *tank) setName(name string) {
	t.name = name
}

func (t *tank) setPower(power uint64) {
	t.power = power
}

func (t *tank) setSpeed(speed uint64) {
	t.speed = speed
}
