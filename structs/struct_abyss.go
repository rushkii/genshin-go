package structs

type Abyss struct {
	Season      int64   `json:"season"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	MaxFloor    string  `json:"max_floor"`
	Battles     int64   `json:"battles"`
	Wins        int64   `json:"wins"`
	Stars       int64   `json:"stars"`
	MostPlayed  []Rank  `json:"most_played"`
	MostDefeats []Rank  `json:"most_defeats"`
	HighDamage  []Rank  `json:"highest_damage"`
	DamageTaken []Rank  `json:"damage_taken"`
	BurstsUsed  []Rank  `json:"bursts_used"`
	SkillsUsed  []Rank  `json:"skills_used"`
	Floors      []Floor `json:"floors"`
}

type Rank struct {
	ID     int64  `json:"id"`
	Icon   string `json:"icon"`
	Rarity int64  `json:"rarity"`
	Value  int64  `json:"value"`
}

type Floor struct {
	Floor      int64     `json:"floor"`
	Icon       string    `json:"icon"`
	MaxStar    int64     `json:"max_stars"`
	Stars      int64     `json:"stars_earned"`
	SettleTime string    `json:"settle_time"`
	Chambers   []Chamber `json:"chambers"`
}

type Chamber struct {
	Chamber int64    `json:"chamber"`
	Stars   int64    `json:"stars_earned"`
	MaxStar int64    `json:"max_stars"`
	Battles []Battle `json:"battles"`
}

type Battle struct {
	Party      int64       `json:"party"`
	Timestamp  string      `json:"timestamp"`
	Characters []Character `json:"characters"`
}

type Character struct {
	ID     int64  `json:"id"`
	Icon   string `json:"icon"`
	Level  int64  `json:"level"`
	Rarity int64  `json:"rarity"`
}
