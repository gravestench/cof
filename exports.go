package cof

import (
	"github.com/gravestench/cof/pkg"
)

// these aliases are here to allow importing from repo root

type (
	COF                 = pkg.COF
	CofLayer            = pkg.CofLayer
	CompositeType       = pkg.CompositeType
	AnimationData       = pkg.AnimationData
	AnimationDataRecord = pkg.AnimationDataRecord
	FrameEvent          = pkg.FrameEvent
	WeaponClass         = pkg.WeaponClass
	DrawEffect          = pkg.DrawEffect
)

func Load(data []byte) (*AnimationData, error) {
	return pkg.Load(data)
}

func New() *COF {
	return pkg.New()
}

func Dir64ToCof(direction, numDirections int) int {
	return pkg.Dir64ToCof(direction, numDirections)
}

func Marshal(cof *COF) []byte {
	return pkg.Marshal(cof)
}

func Unmarshal(data []byte) (*COF, error) {
	return pkg.Unmarshal(data)
}

func WeaponClassFromString(weaponClass string) WeaponClass {
	return pkg.WeaponClassFromString(weaponClass)
}
