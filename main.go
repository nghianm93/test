package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	errENV := godotenv.Load()
	if errENV != nil {
		log.Fatal(errENV)
	}
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")

	// Create bot and enable debugging info
	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	// (more on configuration in examples/configuration/main.go)
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe (https://core.telegram.org/bots/api#getme)
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print Bot information
	fmt.Printf("Bot user: %+v\n", botUser)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage.
			// Send a message to sender with the same text (echo bot).
			// (https://core.telegram.org/bots/api#sendmessage)
			sentMessage, err := bot.SendDice(
				//tu.Message(
				//	tu.ID(chatID),
				//	"No",
				//),
				tu.Dice(tu.ID(chatID), "🎲"),
			)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
