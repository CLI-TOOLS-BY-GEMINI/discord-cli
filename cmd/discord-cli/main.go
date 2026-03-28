package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/minatosingull/discord-cli/pkg/discord"
)

func main() {
	token := flag.String("token", os.Getenv("DISCORD_TOKEN"), "Discord Bot Token")
	flag.Parse()

	if *token == "" {
		fmt.Println("Error: DISCORD_TOKEN environment variable or -token flag is required")
		os.Exit(1)
	}

	client := discord.NewClient(*token)

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		return
	}

	command := args[0]
	switch command {
	case "me":
		user, err := client.GetMe()
		if err != nil {
			log.Fatalf("Failed to get me: %v", err)
		}
		fmt.Printf("Bot User: %s#%s (ID: %s)\n", user.Username, user.Discriminator, user.ID)

	case "channel":
		if len(args) < 2 {
			fmt.Println("Usage: discord-cli channel <channel_id>")
			os.Exit(1)
		}
		channelID := args[1]
		channel, err := client.GetChannel(channelID)
		if err != nil {
			log.Fatalf("Failed to get channel: %v", err)
		}
		fmt.Printf("Channel: %s (ID: %s, Type: %d)\n", channel.Name, channel.ID, channel.Type)

	case "message":
		if len(args) < 3 {
			fmt.Println("Usage: discord-cli message <channel_id> <content>")
			os.Exit(1)
		}
		channelID := args[1]
		content := args[2]
		msg, err := client.CreateMessage(channelID, content)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		fmt.Printf("Message Sent! (ID: %s, Author: %s)\n", msg.ID, msg.Author.Username)

	case "user":
		if len(args) < 2 {
			fmt.Println("Usage: discord-cli user <user_id>")
			os.Exit(1)
		}
		user, err := client.GetUser(args[1])
		if err != nil {
			log.Fatalf("Failed to get user: %v", err)
		}
		fmt.Printf("User: %s#%s (ID: %s, Bot: %v)\n", user.Username, user.Discriminator, user.ID, user.Bot)

	case "guild":
		if len(args) < 2 {
			fmt.Println("Usage: discord-cli guild <guild_id>")
			os.Exit(1)
		}
		guild, err := client.GetGuild(args[1])
		if err != nil {
			log.Fatalf("Failed to get guild: %v", err)
		}
		fmt.Printf("Guild: %s (ID: %s, Owner: %s)\n", guild.Name, guild.ID, guild.OwnerID)

	case "get-message":
		if len(args) < 3 {
			fmt.Println("Usage: discord-cli get-message <channel_id> <message_id>")
			os.Exit(1)
		}
		msg, err := client.GetMessage(args[1], args[2])
		if err != nil {
			log.Fatalf("Failed to get message: %v", err)
		}
		fmt.Printf("Message: %s (ID: %s, Author: %s)\n", msg.Content, msg.ID, msg.Author.Username)

	case "delete-message":
		if len(args) < 3 {
			fmt.Println("Usage: discord-cli delete-message <channel_id> <message_id>")
			os.Exit(1)
		}
		err := client.DeleteMessage(args[1], args[2])
		if err != nil {
			log.Fatalf("Failed to delete message: %v", err)
		}
		fmt.Println("Message deleted successfully")

	case "modify-channel":
		if len(args) < 3 {
			fmt.Println("Usage: discord-cli modify-channel <channel_id> <new_name>")
			os.Exit(1)
		}
		channel, err := client.ModifyChannel(args[1], args[2])
		if err != nil {
			log.Fatalf("Failed to modify channel: %v", err)
		}
		fmt.Printf("Channel modified: %s (ID: %s)\n", channel.Name, channel.ID)

	case "delete-channel":
		if len(args) < 2 {
			fmt.Println("Usage: discord-cli delete-channel <channel_id>")
			os.Exit(1)
		}
		err := client.DeleteChannel(args[1])
		if err != nil {
			log.Fatalf("Failed to delete channel: %v", err)
		}
		fmt.Println("Channel deleted successfully")

	case "edit-message":
		if len(args) < 4 {
			fmt.Println("Usage: discord-cli edit-message <channel_id> <message_id> <new_content>")
			os.Exit(1)
		}
		msg, err := client.EditMessage(args[1], args[2], args[3])
		if err != nil {
			log.Fatalf("Failed to edit message: %v", err)
		}
		fmt.Printf("Message edited: %s (ID: %s)\n", msg.Content, msg.ID)

	case "guild-channels":
		if len(args) < 2 {
			fmt.Println("Usage: discord-cli guild-channels <guild_id>")
			os.Exit(1)
		}
		channels, err := client.GetGuildChannels(args[1])
		if err != nil {
			log.Fatalf("Failed to get guild channels: %v", err)
		}
		for _, ch := range channels {
			fmt.Printf("- %s (ID: %s, Type: %d)\n", ch.Name, ch.ID, ch.Type)
		}

	case "me-guilds":
		guilds, err := client.GetMeGuilds()
		if err != nil {
			log.Fatalf("Failed to get my guilds: %v", err)
		}
		for _, g := range guilds {
			fmt.Printf("- %s (ID: %s)\n", g.Name, g.ID)
		}

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Discord API CLI Tool")
	fmt.Println("Usage:")
	fmt.Println("  discord-cli me")
	fmt.Println("  discord-cli me-guilds")
	fmt.Println("  discord-cli user <user_id>")
	fmt.Println("  discord-cli guild <guild_id>")
	fmt.Println("  discord-cli guild-channels <guild_id>")
	fmt.Println("  discord-cli channel <channel_id>")
	fmt.Println("  discord-cli modify-channel <channel_id> <new_name>")
	fmt.Println("  discord-cli delete-channel <channel_id>")
	fmt.Println("  discord-cli message <channel_id> <content>")
	fmt.Println("  discord-cli get-message <channel_id> <message_id>")
	fmt.Println("  discord-cli edit-message <channel_id> <message_id> <new_content>")
	fmt.Println("  discord-cli delete-message <channel_id> <message_id>")
}
