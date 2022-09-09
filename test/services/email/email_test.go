package email

import (
	"btcApp/internal/services/email"
	"btcApp/test/utils"
	"fmt"
	"github.com/antihax/optional"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
	"golang.org/x/net/context"
	"testing"
)

func GetMailClient() (*mailslurp.APIClient, context.Context) {
	emailContext := context.WithValue(
		context.Background(),
		mailslurp.ContextAPIKey,
		mailslurp.APIKey{Key: utils.ServiceForTestingMailingApiKey},
	)

	config := mailslurp.NewConfiguration()
	client := mailslurp.NewAPIClient(config)

	return client, emailContext
}

func CreateInbox(t *testing.T, client *mailslurp.APIClient, emailContext context.Context) *mailslurp.InboxDto {
	inbox, _, err := client.InboxControllerApi.CreateInbox(emailContext, nil)

	if err != nil {
		utils.Failure(t, "Error during creating an inbox: %s", err)
	}
	return &inbox
}

func TestSendEmail(t *testing.T) {
	client, emailContext := GetMailClient()
	inbox := CreateInbox(t, client, emailContext)

	testData := 10000
	msg := email.InitializeMessage(testData)
	dialer := email.InitializeDialer()
	email.SendEmail(inbox.EmailAddress, msg, dialer)

	waitOpts := &mailslurp.WaitForLatestEmailOpts{
		InboxId:    optional.NewInterface(inbox.Id),
		Timeout:    optional.NewInt64(30000),
		UnreadOnly: optional.NewBool(true),
	}
	receivedEmail, _, _ := client.WaitForControllerApi.WaitForLatestEmail(emailContext, waitOpts)

	expectedMessage := fmt.Sprintf("BTC rate in UAH: %d\r\n", testData)
	if *receivedEmail.Body == expectedMessage {
		utils.Success(t, "Message received")
	} else {
		utils.Failure(t, "Received message is wrong. Actual: %s, expected: %s", *receivedEmail.Body, expectedMessage)
	}
}
