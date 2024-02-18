// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/datumforge/fgax/entfga/_examples/basic/ent/organization"
	"github.com/datumforge/fgax/entfga/_examples/basic/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	orgmembershipFields := schema.OrgMembership{}.Fields()
	_ = orgmembershipFields
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[1].Descriptor()
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = organizationDescName.Validators[0].(func(string) error)
}
