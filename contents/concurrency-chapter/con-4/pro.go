package con4

func addEmailsToQueue(emails []string) chan string {
	bufferedChan := make(chan string, len(emails))
	for _, email := range emails {
		bufferedChan <- email
	}
	return bufferedChan
}
