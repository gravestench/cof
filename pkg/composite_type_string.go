// Code generated by "stringer -linecomment -type CompositeType -output composite_type_string.go"; DO NOT EDIT.
package pkg

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CompositeTypeHead-0]
	_ = x[CompositeTypeTorso-1]
	_ = x[CompositeTypeLegs-2]
	_ = x[CompositeTypeRightArm-3]
	_ = x[CompositeTypeLeftArm-4]
	_ = x[CompositeTypeRightHand-5]
	_ = x[CompositeTypeLeftHand-6]
	_ = x[CompositeTypeShield-7]
	_ = x[CompositeTypeSpecial1-8]
	_ = x[CompositeTypeSpecial2-9]
	_ = x[CompositeTypeSpecial3-10]
	_ = x[CompositeTypeSpecial4-11]
	_ = x[CompositeTypeSpecial5-12]
	_ = x[CompositeTypeSpecial6-13]
	_ = x[CompositeTypeSpecial7-14]
	_ = x[CompositeTypeSpecial8-15]
	_ = x[CompositeTypeMax-16]
}

const _CompositeType_name = "HDTRLGRALARHLHSHS1S2S3S4S5S6S7S8CompositeTypeMax"

var _CompositeType_index = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 48}

func (i CompositeType) String() string {
	if i < 0 || i >= CompositeType(len(_CompositeType_index)-1) {
		return "CompositeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompositeType_name[_CompositeType_index[i]:_CompositeType_index[i+1]]
}
