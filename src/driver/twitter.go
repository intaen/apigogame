package driver

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
)

func InitiateClient() (*twitter.Client, error) {
	// Declare config client
	config := oauth1.NewConfig(viper.GetString("cred.twitter.consumer-key"), viper.GetString("cred.twitter.consumer-secret"))
	token := oauth1.NewToken(viper.GetString("cred.twitter.access-token"), viper.GetString("cred.twitter.access-secret"))
	// Initiate client
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, err
}
