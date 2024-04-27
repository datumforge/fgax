package fgax

import (
	"testing"

	openfga "github.com/openfga/go-sdk"
	typesystem "github.com/openfga/openfga/pkg/typesystem"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewDirectRelation(t *testing.T) {
	testCases := []struct {
		name        string
		role        string
		expectedRes openfga.Userset
	}{
		{
			name: "new admin role",
			role: "admin",
			expectedRes: openfga.Userset{
				This: &map[string]interface{}{
					"admin": typesystem.This(),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userset := newDirectRelation(tc.role)

			assert.Equal(t, tc.expectedRes, userset)
		})
	}
}

func TestNewComputedUsersetRelation(t *testing.T) {
	testCases := []struct {
		name        string
		relation    string
		expectedRes openfga.Userset
	}{
		{
			name:     "new relation",
			relation: "meow",
			expectedRes: openfga.Userset{
				ComputedUserset: &openfga.ObjectRelation{
					Relation: lo.ToPtr("meow"),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userset := newComputedUsersetRelation(tc.relation)

			assert.Equal(t, tc.expectedRes, userset)
		})
	}
}
func TestNewTupleUsersetRelation(t *testing.T) {
	testCases := []struct {
		name         string
		relation     string
		fromRelation string
		expectedRes  openfga.Userset
	}{
		{
			name:         "new tuple userset relation",
			relation:     "relation",
			fromRelation: "fromRelation",
			expectedRes: openfga.Userset{
				TupleToUserset: &openfga.TupleToUserset{
					Tupleset: openfga.ObjectRelation{
						Relation: lo.ToPtr("fromRelation"),
					},
					ComputedUserset: openfga.ObjectRelation{
						Relation: lo.ToPtr("relation"),
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userset := newTupleUsersetRelation(tc.relation, tc.fromRelation)

			assert.Equal(t, tc.expectedRes, userset)
		})
	}
}
func TestCreateNewMetadata(t *testing.T) {
	testCases := []struct {
		name       string
		relation   string
		userType   string
		expectedRD map[string]openfga.RelationMetadata
	}{
		{
			name:     "empty user type, not direct",
			relation: "relation",
			userType: "",
			expectedRD: map[string]openfga.RelationMetadata{
				"relation": {
					DirectlyRelatedUserTypes: &[]openfga.RelationReference{},
				},
			},
		},
		{
			name:     "non-empty user type, direct relation",
			relation: "relation",
			userType: "userType",
			expectedRD: map[string]openfga.RelationMetadata{
				"relation": {
					DirectlyRelatedUserTypes: &[]openfga.RelationReference{
						{
							Type: "userType",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rd := createNewMetadata(tc.relation, tc.userType)

			assert.Equal(t, tc.expectedRD, rd)
		})
	}
}
