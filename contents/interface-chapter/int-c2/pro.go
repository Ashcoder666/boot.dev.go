package intc2

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	} else {
		return d.priorityLevel
	}
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch val := n.(type) {
	case directMessage:
		return val.senderUsername, val.importance()
	case groupMessage:
		return val.groupName, val.importance()
	case systemAlert:
		return val.alertCode, val.importance()
	default:
		return "", 0
	}
}
