package mapc2

func includes[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
func findSuggestedFriends(username string, friendships map[string][]string) []string {
	// friendsCounter := make(map[string]int)
	mutualFriendsUnique := []string{}
	currentFriends := friendships[username]

	for _, friend := range currentFriends {
		friendsofCurrentFriends := friendships[friend]
		for _, friendcurrent := range friendsofCurrentFriends {
			if friendcurrent == username || includes(currentFriends, friendcurrent) || includes(mutualFriendsUnique, friendcurrent) {
				// fmt.Println(friendcurrent, "skipped")
				continue
			}

			mutualFriendsUnique = append(mutualFriendsUnique, friendcurrent)
			// fmt.Println(mutualFriendsUnique, "im result array")
		}
	}

	if len(mutualFriendsUnique) == 0 {
		return nil
	}

	return mutualFriendsUnique

}
