package tank

type m2 struct {
	tank
}

func newM2() ITank {
	return &m2{
		tank{
			name: "M2",
			power: 11,
			speed: 43,
		},
	}
}
