// Code generated by "stringer -type TypeKind enum.go"; DO NOT EDIT.

package ir

import "fmt"

const _TypeKind_name = "Int8Int16Int32Int64Uint8Uint16Uint32Uint64Float32Float64Float128Complex64Complex128Complex256ArrayUnionStructPointerFunction"

var _TypeKind_index = [...]uint8{0, 4, 9, 14, 19, 24, 30, 36, 42, 49, 56, 64, 73, 83, 93, 98, 103, 109, 116, 124}

func (i TypeKind) String() string {
	i -= 1
	if i < 0 || i >= TypeKind(len(_TypeKind_index)-1) {
		return fmt.Sprintf("TypeKind(%d)", i+1)
	}
	return _TypeKind_name[_TypeKind_index[i]:_TypeKind_index[i+1]]
}
