package tank

import "errors"

// Create creates tank instance, based on given name
func Create(name string) (ITank, error) {
	switch name {
	case "M1":
		return newM1(), nil
	case "M2":
		return newM2(), nil
	default:
		return &tank{}, errors.New("There is no tank with given name")
	}
}
