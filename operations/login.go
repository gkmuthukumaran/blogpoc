package operations

import "github.com/blogpoc/interfaces/db"

func IsValidUser(username, password string) bool {

	userDetails, _ := db.GetUserDetails(username)

	if password == userDetails.Password {
		return true
	}
	return false
}
