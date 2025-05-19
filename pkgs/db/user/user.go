package user

import "go-api/pkgs/db"

func GetUser() {
	// do something
	db.ConnectOracleDb("local")
}
// Compare this snippet from go-api-db-integration/pkgs/db/user/user.go:
// package user	