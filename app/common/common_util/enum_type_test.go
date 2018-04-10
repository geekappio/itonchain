package common_util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	YES_A := DefEnumType("YES")
	YES_B := DefEnumType("YES")
	var NO = DefEnumType("NO")

	// Compare YES enumeration
	assert.EqualValues(t, YES_A, YES_B)

	// Compare YES enumeration
	assert.True(t, YES_A == YES_B)

	// Compare valueOf method.
	assert.EqualValues(t, YES_A, YES_A.ValueOf("YES"))
	assert.EqualValues(t, YES_A, NO.ValueOf("YES"))

	// Compare value of YES enumeration
	assert.True(t, YES_A.Value == YES_B.Value)

}