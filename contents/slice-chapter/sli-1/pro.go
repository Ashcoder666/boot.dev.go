package sli1

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	resArr := [3]string{primary, secondary, tertiary}

	resNumArr := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}

	return resArr, resNumArr
}
