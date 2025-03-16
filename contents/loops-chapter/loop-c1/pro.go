package loopc1

import "fmt"

func countConnections(groupSize int) int {
	conn := 0
	for i := 1; i < groupSize; i++ {
		conn += i
	}
	fmt.Println(conn)
	return conn
}
