package lib

const (
	EventsRoute       = "/%s/events"            // warSeasonId
	LatestEventRoute  = "/%s/events/latest"     // warSeasonId
	FeedRoute         = "/%s/feed"              // warSeasonId
	InfoRoute         = "/%s/info"              // warSeasonId
	PlanetsRoute      = "/%s/planets"           // warSeasonId
	PlanetRoute       = "/%s/planets/%d"        // warSeasonId, PlanetId
	PlanetStatusRoute = "/%s/planets/%d/status" // warSeasonId, PlanetId
	StatusRoute       = "/%s/events"            // warSeasonId
)
