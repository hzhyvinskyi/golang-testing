package tank

type m1 struct {
	tank
}

func newM1() ITank {
	return &m1{
		tank{
			name: "M1",
			power: 10,
			speed: 45,
		},
	}
}
