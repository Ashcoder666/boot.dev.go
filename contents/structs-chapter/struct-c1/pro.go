package structc1

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	var messageLimit int
	if membershipType == "premium" {
		messageLimit = 1000
	} else {
		messageLimit = 100
	}
	newUser := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: messageLimit,
		},
	}

	return newUser
}
