package gun

import "WB2/pattern/06_factory_method/interfaces"

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
