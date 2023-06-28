package gun

import "WB2/pattern/06_factory_method/interfaces"

type ak47 struct {
	gun
}

func NewAk47() interfaces.Gun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
