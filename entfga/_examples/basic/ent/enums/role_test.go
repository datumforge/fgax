package enums_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/ent/enums"
)

func TestEnum(t *testing.T) {
	testCases := []struct {
		name     string
		role     string
		expected enums.Role
	}{
		{
			name:     "admin",
			role:     "admin",
			expected: enums.RoleAdmin,
		},
		{
			name:     "member",
			role:     "member",
			expected: enums.RoleMember,
		},
		{
			name:     "owner",
			role:     "owner",
			expected: enums.RoleOwner,
		},
		{
			name:     "invalid role",
			role:     "cattypist",
			expected: enums.Invalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := enums.ToRole(tc.role)
			assert.Equal(t, tc.expected, res)
		})
	}
}
