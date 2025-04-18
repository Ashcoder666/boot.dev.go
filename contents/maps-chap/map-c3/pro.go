package mapc3

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	user, ok := users[name]
	if !ok {
		// delete(users, name)
		return logNotFound
	}
	if user.admin {
		delete(users, name)
		return logAdmin
	}
	defer delete(users, name)
	return logDeleted
}
