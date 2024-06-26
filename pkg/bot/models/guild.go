package models

import (
	"Helldivers2Tools/pkg/bot/database"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type GuildModel struct {
	Id                  string `json:"id"`
	GuildId             string `json:"guild_id"`
	Name                string `json:"name"`
	AnnouncementChannel string
	Location            string `json:"location"`
}

func (guild *GuildModel) CreateTable() error {
	_, err := database.Database.Exec(`
		CREATE TABLE IF NOT EXISTS guilds (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			guild_id TEXT NOT NULL UNIQUE,
			name TEXT NOT NULL,
			announcement_channel TEXT,
			location TEXT
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func appendGuildsToList(list []GuildModel, rows *sql.Rows) ([]GuildModel, error) {
	var guild GuildModel
	err := rows.Scan(
		&guild.Id,
		&guild.GuildId,
		&guild.Name,
		&guild.AnnouncementChannel,
		&guild.Location,
	)
	if err != nil {
		return nil, err
	}

	list = append(list, guild)
	return list, nil
}

func getGuildsListFromRows(rows *sql.Rows) ([]GuildModel, error) {
	var list []GuildModel
	var err error
	list, err = appendGuildsToList(list, rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list, err = appendGuildsToList(list, rows)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func (guild *GuildModel) getGuildByQuery(query squirrel.SelectBuilder) error {
	rows, err := database.SelectHelper(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = rows.Scan(
		&guild.Id,
		&guild.GuildId,
		&guild.Name,
		&guild.AnnouncementChannel,
		&guild.Location,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetGuilds() ([]GuildModel, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("*").
			From("guilds"),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	guilds, err := getGuildsListFromRows(rows)
	return guilds, err
}

func (guild *GuildModel) GetGuildById(id string) (*GuildModel, error) {
	return guild, guild.getGuildByQuery(
		squirrel.
			Select("*").
			From("guilds").
			Where(squirrel.Eq{"id": id}),
	)
}

func (guild *GuildModel) GetGuildByGuildId(guildId string) (*GuildModel, error) {
	return guild, guild.getGuildByQuery(
		squirrel.
			Select("*").
			From("guilds").
			Where(squirrel.Eq{"guild_id": guildId}),
	)
}

func (guild *GuildModel) CreateGuild() error {
	result, err := squirrel.
		Insert("guilds").
		Columns(
			"guild_id",
			"name",
			"announcement_channel",
			"location",
		).
		Values(
			guild.GuildId,
			guild.Name,
			guild.AnnouncementChannel,
			guild.Location,
		).
		RunWith(database.Database).Exec()
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	guild.Id = fmt.Sprintf("%d", id)
	return nil
}

func (guild *GuildModel) DeleteGuild() error {
	_, err := squirrel.
		Delete("guilds").
		Where(squirrel.Eq{"id": guild.Id}).
		RunWith(database.Database).Exec()
	return err
}

func (guild *GuildModel) UpdateGuild() error {
	guildQuery := squirrel.Update("guilds")

	if guild.GuildId != "" {
		guildQuery = guildQuery.Set("guild_id", guild.GuildId)
	}

	if guild.Name != "" {
		guildQuery = guildQuery.Set("name", guild.Name)
	}

	if guild.AnnouncementChannel != "" {
		guildQuery = guildQuery.Set("announcement_channel", guild.AnnouncementChannel)
	}

	if guild.Location != "" {
		guildQuery = guildQuery.Set("location", guild.Location)
	}

	_, err := guildQuery.
		Where(squirrel.Eq{"id": guild.Id}).
		RunWith(database.Database).Exec()
	return err
}
