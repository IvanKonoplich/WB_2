package gun

import "WB2/patterns/factory/interfaces"

type maverick struct {
	gun
}

func NewMaverick() interfaces.Gun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
