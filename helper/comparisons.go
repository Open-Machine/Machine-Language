package helper

import (
	"github.com/open-machine/assembler/data"
	"reflect"
)

func SafeIsEqualStrPointer(a *string, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return *a == *b
}

func SafeIsEqualProgramPointer(a *data.Program, b *data.Program) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return reflect.DeepEqual(*a, *b)
}

func SafeIsEqualInstructionPointer(a *data.Instruction, b *data.Instruction) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return *a == *b
}

func SafeIsEqualInstructionParamPointer(a *data.InstructionParameter, b *data.InstructionParameter) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return *a == *b
}
