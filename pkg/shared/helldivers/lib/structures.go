package lib

import "time"

type WarSeasons struct {
	Current string   `json:"current"`
	Seasons []string `json:"seasons"`
}

type Message struct {
	Es string `json:"es"`
	Fr string `json:"fr"`
	De string `json:"de"`
	En string `json:"en"`
	It string `json:"it"`
	Pl string `json:"pl"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}

type FeedMessage struct {
	Id      int           `json:"id"`
	Message Message       `json:"message"`
	TagIds  []interface{} `json:"tag_ids"`
	Type    int           `json:"type"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Planet struct {
	Disabled     bool     `json:"disabled"`
	Hash         int64    `json:"hash"`
	Index        int      `json:"index"`
	InitialOwner string   `json:"initial_owner"`
	MaxHealth    int      `json:"max_health"`
	Name         string   `json:"name"`
	Position     Position `json:"position"`
	Sector       string   `json:"sector"`
	Waypoints    []int    `json:"waypoints"`
}

type HomeWorld struct {
	Planets []Planet `json:"planets"`
	Race    string   `json:"race"`
}

type Campaign struct {
	Count  int    `json:"count"`
	Id     int    `json:"id"`
	Planet Planet `json:"planet"`
	Type   int    `json:"type"`
}

type PlanetAttack struct {
	Source Planet `json:"source"`
	Target Planet `json:"target"`
}

type PlanetStatus struct {
	Health         int     `json:"health"`
	Liberation     float64 `json:"liberation"`
	Owner          string  `json:"owner"`
	Planet         Planet  `json:"planet"`
	Players        int     `json:"players"`
	RegenPerSecond float64 `json:"regen_per_second"`
}

type Info struct {
	Capitals               []interface{} `json:"capitals"`
	EndDate                time.Time     `json:"end_date"`
	HomeWorlds             []HomeWorld   `json:"home_worlds"`
	MinimumClientVersion   string        `json:"minimum_client_version"`
	PlanetPermanentEffects []interface{} `json:"planet_permanent_effects"`
	Planets                []Planet      `json:"planets"`
	StartDate              time.Time     `json:"start_date"`
	WarId                  int           `json:"war_id"`
}

type GlobalEvent struct {
	AssignmentId32 int      `json:"assignment_id_32"`
	Effects        []string `json:"effects"`
	Flag           int      `json:"flag"`
	Id             int      `json:"id"`
	Id32           int      `json:"id_32"`
	Message        Message  `json:"message"`
	MessageId32    int      `json:"message_id_32"`
	Planets        []Planet `json:"planets"`
	PortraitId32   int      `json:"portrait_id_32"`
	Race           string   `json:"race"`
	Title          string   `json:"title"`
	Title32        int      `json:"title_32"`
}

type JointOperation struct {
	HqNodeIndex int    `json:"hq_node_index"`
	Id          int    `json:"id"`
	Planet      Planet `json:"planet"`
}

type PlanetEvent struct {
	Campaign        Campaign         `json:"campaign"`
	EventType       int              `json:"event_type"`
	ExpireTime      time.Time        `json:"expire_time"`
	Health          int              `json:"health"`
	Id              int              `json:"id"`
	JointOperations []JointOperation `json:"joint_operations"`
	MaxHealth       int              `json:"max_health"`
	Planet          Planet           `json:"planet"`
	Race            string           `json:"race"`
	StartTime       time.Time        `json:"start_time"`
}

type Status struct {
	ActiveElectionPolicyEffects []int            `json:"active_election_policy_effects"`
	Campaigns                   []Campaign       `json:"campaigns"`
	CommunityTargets            []int            `json:"community_targets"`
	GlobalEvents                []GlobalEvent    `json:"global_events"`
	ImpactMultiplier            int              `json:"impact_multiplier"`
	JointOperations             []JointOperation `json:"joint_operations"`
	PlanetActiveEffects         []int            `json:"planet_active_effects"`
	PlanetAttacks               []PlanetAttack   `json:"planet_attacks"`
	PlanetEvents                []PlanetEvent    `json:"planet_events"`
	PlanetStatus                []PlanetStatus   `json:"planet_status"`
	SnapshotAt                  time.Time        `json:"snapshot_at"`
	StartedAt                   time.Time        `json:"started_at"`
	WarId                       int              `json:"war_id"`
}
