package struct7

//the order of fields in a struct can have a big impact on memory usage.
// just most menory to least

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}
