package common_util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	YES_A := DefEnumType("YES")
	YES_B := DefEnumType("YES")

	// Compare YES enumeration
	assert.EqualValues(t, YES_A, YES_B)

	// Compare YES enumeration
	assert.True(t, YES_A == YES_B)

	// Compare valueOf method.
	assert.EqualValues(t, YES_A, EnumValueOf("YES"))
	assert.EqualValues(t, YES_A, EnumValueOf("YES"))

	// Compare value of YES enumeration
	assert.True(t, YES_A.Value == YES_B.Value)

}