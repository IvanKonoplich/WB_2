package gun

import "WB2/patterns/factory/interfaces"

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