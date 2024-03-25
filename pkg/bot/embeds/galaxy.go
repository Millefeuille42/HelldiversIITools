package embeds

import (
	"Helldivers2Tools/pkg/bot/globals"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"github.com/bwmarrin/discordgo"
)

func BuildGalaxyEmbed(stats lib.GalaxyStats) *discordgo.MessageEmbed {
	ret := &discordgo.MessageEmbed{
		Type:   "rich",
		Title:  "Galaxy",
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Kills",
		Value:  globals.NumberPrinter.Sprintf("%d terminids", stats.BugKills),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Value:  globals.NumberPrinter.Sprintf("%d automatons", stats.AutomatonKills),
		Inline: true,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Value:  globals.NumberPrinter.Sprintf("%d illuminates", stats.IlluminateKills),
		Inline: true,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Missions",
		Value:  globals.NumberPrinter.Sprintf("%d won, %d lost", stats.MissionsWon, stats.MissionsLost),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Bullets",
		Value:  globals.NumberPrinter.Sprintf("%d fired, %d hits", stats.BulletsFired, stats.BulletsHit),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Accidentals",
		Value:  globals.NumberPrinter.Sprintf("%d accidental deaths", stats.Friendlies),
		Inline: false,
	})

	return ret
}
