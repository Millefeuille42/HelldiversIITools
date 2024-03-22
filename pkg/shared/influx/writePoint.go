package influx

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"context"
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

func WritePlanet(planet lib.Planet) error {
	if planet.LibPercent == 0 {
		return nil
	}

	point := influxdb3.NewPointWithMeasurement("planet").
		SetTag("name", planet.PlanetName).
		SetStringField("CurrentOwner", planet.CurrentOwner).
		SetDoubleField("LibPercent", planet.LibPercent).
		SetIntegerField("AutomatonKills", planet.AutomatonKills).
		SetIntegerField("BugKills", planet.BugKills).
		SetIntegerField("IlluminateKills", planet.IlluminateKills).
		SetIntegerField("MissionsLost", planet.MissionsLost).
		SetIntegerField("MissionsWon", planet.MissionsWon).
		SetIntegerField("BulletsFired", planet.BulletsFired).
		SetIntegerField("BulletsHit", planet.BulletsHit).
		SetIntegerField("Friendlies", planet.Friendlies).
		SetIntegerField("Players", planet.Players).
		SetIntegerField("Deaths", planet.Deaths).
		SetIntegerField("TimePlayed", planet.TimePlayed)

	var points = []*influxdb3.Point{point}
	return Client.WritePoints(context.Background(), points,
		influxdb3.WithDatabase("helldivers"))
}

func WriteGalaxy(planet lib.GalaxyStats) error {
	point := influxdb3.NewPointWithMeasurement("galaxy").
		SetIntegerField("AutomatonKills", planet.AutomatonKills).
		SetIntegerField("BugKills", planet.BugKills).
		SetIntegerField("IlluminateKills", planet.IlluminateKills).
		SetIntegerField("MissionsLost", planet.MissionsLost).
		SetIntegerField("MissionsWon", planet.MissionsWon).
		SetIntegerField("BulletsFired", planet.BulletsFired).
		SetIntegerField("BulletsHit", planet.BulletsHit).
		SetIntegerField("Friendlies", planet.Friendlies).
		SetIntegerField("Deaths", planet.Deaths).
		SetIntegerField("TimePlayed", planet.TimePlayed)

	var points = []*influxdb3.Point{point}
	return Client.WritePoints(context.Background(), points,
		influxdb3.WithDatabase("helldivers"))
}
