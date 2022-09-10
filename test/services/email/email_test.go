package email

import (
	"btcApp/internal/services/email"
	"btcApp/test/utils"
	"fmt"
	"github.com/antihax/optional"
	"github.com/joho/godotenv"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
	"testing"
)

func TestSendEmail(t *testing.T) {
	client, emailContext := GetMailClient(t)
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

func GetMailClient(t *testing.T) (*mailslurp.APIClient, context.Context) {
	loadEnvironment(t)

	emailContext := context.WithValue(
		context.Background(),
		mailslurp.ContextAPIKey,
		mailslurp.APIKey{Key: os.Getenv("MAILSLURP_API_KEY")},
	)

	config := mailslurp.NewConfiguration()
	client := mailslurp.NewAPIClient(config)

	return client, emailContext
}

func loadEnvironment(t *testing.T) {
	projectDirectory := findProjectDirectoryWithEnvFile()

	loadingError := godotenv.Load(projectDirectory + "/.env")
	if loadingError != nil {
		utils.Failure(t, "Environment not loaded")
	}
}

func findProjectDirectoryWithEnvFile() string {
	currentDirectory, _ := os.Getwd()
	var projectDirectory string
	for i := 0; i < 3; i++ {
		projectDirectory = filepath.Dir(currentDirectory)
		currentDirectory = projectDirectory
	}
	return projectDirectory
}

func CreateInbox(t *testing.T, client *mailslurp.APIClient, emailContext context.Context) *mailslurp.InboxDto {
	inbox, _, err := client.InboxControllerApi.CreateInbox(emailContext, nil)

	if err != nil {
		utils.Failure(t, "Error during creating an inbox: %s", err)
	}
	return &inbox
}
