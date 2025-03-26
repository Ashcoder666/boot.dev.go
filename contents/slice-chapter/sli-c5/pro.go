package slic5

import (
	"regexp"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for index, message := range messages {
		messages[index].tags = tagger(message)
	}

	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	words := strings.Split(msg.content, " ")

	re := regexp.MustCompile("[^a-zA-Z]+")
	for _, word := range words {

		word = re.ReplaceAllString(word, "")

		if strings.ToLower(word) == "urgent" {
			tags = append(tags, "Urgent")
		}

	}

	for _, word := range words {

		word = re.ReplaceAllString(word, "")

		if strings.ToLower(word) == "sale" {
			tags = append(tags, "Promo")
		}
	}

	return tags

}
