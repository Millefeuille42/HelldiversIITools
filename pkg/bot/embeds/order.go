package embeds

import (
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

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
	case lib.ControlTaskType:
		taskType = "Control"
	}

	planet, err := helldivers.GoDiversClient.GetPlanet(task.Target.Index)
	if err != nil {
		log.Println(err)
	} else {
		for _, event := range planet.Events {
			switch lib.EventType(event.EventType) {
			case lib.DefenseEventType:
				taskType = "Defend"
			}
		}
	}

	progress := "⚠️"
	if task.Progress == 1 {
		progress = "✅"
	}

	return fmt.Sprintf("%s %s: %s", taskType, task.Target.Name, progress)
}

func BuildOrderEmbed(order lib.MajorOrder) *discordgo.MessageEmbed {
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

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("Ends <t:%d:R>", order.EndsAt.Unix()),
		Inline: false,
	})

	return ret
}
