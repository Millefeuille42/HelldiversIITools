package commands

import (
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func buildOrderComponents(order lib.MajorOrder) []discordgo.MessageComponent {
	var planets []lib.PlanetName

	for _, task := range order.Tasks {
		if task.Progress <= 0 {
			planets = append(planets, task.Target)
		}
	}

	if len(planets) <= 0 {
		return nil
	}

	var buttons []discordgo.MessageComponent

	for _, planet := range planets {
		buttons = append(buttons, discordgo.Button{
			Label:    planet.Name,
			Style:    0,
			Disabled: false,
			Emoji: discordgo.ComponentEmoji{
				Name: "🌎",
			},
			CustomID: fmt.Sprintf("planet_button-%d", planet.Index),
		})
	}

	return []discordgo.MessageComponent{discordgo.ActionsRow{Components: buttons}}
}

func buildReward(reward lib.Reward) string {
	rType := ""
	switch reward.Type {
	case lib.MedalRewardType:
		rType = "medal"
	}
	if reward.Amount > 1 {
		rType += "s"
	}
	return fmt.Sprintf("%d %s", reward.Amount, rType)
}

func buildTaskProgress(task lib.Task) string {
	taskType := ""
	switch task.Type {
	case lib.LiberateTaskType:
		taskType = "Liberate"
	}

	progress := "⚠️"
	if task.Progress == 1 {
		progress = "✅"
	}

	return fmt.Sprintf("%s %s: %s", taskType, task.Target.Name, progress)
}

func buildOrderEmbed(order lib.MajorOrder) *discordgo.MessageEmbed {
	ret := &discordgo.MessageEmbed{
		Type: "rich",
		// TODO add reward image
		Title:       order.Title,
		Description: order.Description,
		Fields:      make([]*discordgo.MessageEmbedField, 0),
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Briefing",
		Value:  order.Briefing,
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Rewards",
		Value:  buildReward(order.Reward),
		Inline: false,
	})

	tasksValue := ""
	for _, task := range order.Tasks {
		if tasksValue != "" {
			tasksValue += "\n"
		}
		tasksValue += buildTaskProgress(task)
	}
	if tasksValue != "" {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Tasks",
			Value:  tasksValue,
			Inline: false,
		})
	}

	return ret
}

func orderCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting order", 0)
		return
	}

	order, err := helldivers.GoDiversClient.GetMajorOrder()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting order", 0)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    "",
			Components: buildOrderComponents(order),
			Embeds: []*discordgo.MessageEmbed{
				buildOrderEmbed(order),
			},
			AllowedMentions: nil,
			Choices:         nil,
			CustomID:        "",
			Title:           "",
		},
	})

	if err != nil {
		log.Println(err)
	}
}
