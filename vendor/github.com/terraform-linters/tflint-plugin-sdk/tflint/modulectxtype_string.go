// Code generated by "stringer -type=ModuleCtxType"; DO NOT EDIT.

package tflint

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SelfModuleCtxType-0]
	_ = x[RootModuleCtxType-1]
}

const _ModuleCtxType_name = "SelfModuleCtxTypeRootModuleCtxType"

var _ModuleCtxType_index = [...]uint8{0, 17, 34}

func (i ModuleCtxType) String() string {
	if i < 0 || i >= ModuleCtxType(len(_ModuleCtxType_index)-1) {
		return "ModuleCtxType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ModuleCtxType_name[_ModuleCtxType_index[i]:_ModuleCtxType_index[i+1]]
}
