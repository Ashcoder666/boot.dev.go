package struct6

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (a authenticationInfo) getBasicAuth() string {
	res := fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
	return res
}
