package factory

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/pop/v6"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

// BuildAdminUser creates an AdminUser
// Also creates, if not provided
//   - User
//
// Params:
// - customs is a slice that will be modified by the factory
// - db can be set to nil to create a stubbed model that is not stored in DB.
func BuildAdminUser(db *pop.Connection, customs []Customization, traits []Trait) models.AdminUser {
	customs = setupCustomizations(customs, traits)

	// Find adminuser assertion and convert to models adminuser
	var cAdminUser models.AdminUser
	if result := findValidCustomization(customs, AdminUser); result != nil {
		cAdminUser = result.Model.(models.AdminUser)
		if result.LinkOnly {
			return cAdminUser
		}
	}

	// Create the associated user model
	user := BuildUser(db, customs, nil)

	// create adminuser
	adminUser := models.AdminUser{
		UserID:    &user.ID,
		User:      user,
		FirstName: "Leo",
		LastName:  "Spaceman",
		Email:     "leo_spaceman_admin@example.com",
		Role:      "SYSTEM_ADMIN",
	}
	// Overwrite values with those from assertions
	testdatagen.MergeModels(&adminUser, cAdminUser)

	// If db is false, it's a stub. No need to create in database
	if db != nil {
		mustCreate(db, &adminUser)
	}

	return adminUser
}

// BuildDefaultAdminUser returns an admin user with appropriate email
// Also creates
//   - User
func BuildDefaultAdminUser(db *pop.Connection) models.AdminUser {
	return BuildAdminUser(db, nil, []Trait{GetTraitAdminUserEmail})
}

// ------------------------
//        TRAITS
// ------------------------

// GetTraitAdminUserEmail helps comply with the uniqueness constraint on emails
func GetTraitAdminUserEmail() []Customization {
	// There's a uniqueness constraint on admin user emails so add some randomness
	email := strings.ToLower(fmt.Sprintf("leo_spaceman_admin_%s@example.com", makeRandomString(5)))
	return []Customization{
		{
			Model: models.User{
				LoginGovEmail: email,
			},
			Type: &User,
		},
		{
			Model: models.AdminUser{
				Email: email,
			},
			Type: &AdminUser,
		},
	}
}

// GetTraitActiveAdminUser sets the User and AdminUser as Active
func GetTraitActiveAdminUser() []Customization {
	return []Customization{
		{
			Model: models.AdminUser{
				Active: true,
			},
		},
		{
			Model: models.User{
				Active: true,
			},
		},
	}
}
