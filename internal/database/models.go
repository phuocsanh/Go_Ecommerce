// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

// Account
type PreGoCrmUserC struct {
	// Account ID
	UserID uint32
	// Email
	UserEmail string
	// Phone Number
	UserPhone string
	// Username
	UserUsername string
	// Password
	UserPassword string
	// Creation Time
	UserCreatedAt int32
	// Update Time
	UserUpdatedAt int32
	// Creation IP
	UserCreateIpAt string
	// Last Login Time
	UserLastLoginAt int32
	// Last Login IP
	UserLastLoginIpAt string
	// Login Times
	UserLoginTimes int32
	// Status 1:enable, 0:disable
	UserStatus bool
}