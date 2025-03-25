package slic5

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for index, message := range messages {
		messages[index].tags = tagger(message)
	}
}

func tagger(msg sms) []string {
	tags := []string{}
	// ?
}
