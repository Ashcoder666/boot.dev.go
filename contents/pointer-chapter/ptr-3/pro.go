package ptr3

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analytics *Analytics, msg Message) {
	if msg.Success {
		analytics.MessagesTotal += 1
		analytics.MessagesFailed += 0
		analytics.MessagesSucceeded += 1
	} else {
		analytics.MessagesTotal += 1
		analytics.MessagesFailed += 1
		analytics.MessagesSucceeded += 0
	}
}
