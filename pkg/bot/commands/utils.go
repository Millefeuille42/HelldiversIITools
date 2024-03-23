package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func interactionSendFollowupError(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
	_, _ = s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{Content: message, Flags: flags})
}

func interactionSendDefer(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: nil,
	})
}

func interactionSendResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   flags,
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func interactionSendError(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
	interactionSendResponse(s, i, message, flags)
}

func parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	return optionMap
}

func handleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate, val string, choices []*discordgo.ApplicationCommandOptionChoice) {
	if i.Type != discordgo.InteractionApplicationCommandAutocomplete {
		return
	}

	choices = filterChoices(choices, val)
	choices = rankChoices(choices, val)
	maxResults := 7
	if len(choices) < maxResults {
		maxResults = len(choices)
	}
	choices = choices[:maxResults]

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
