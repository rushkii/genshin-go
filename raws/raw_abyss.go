package raws

type RawAbyssData struct {
	Retcode int64    `json:"retcode"`
	Message string   `json:"message"`
	Data    RawAbyss `json:"data"`
}

type RawAbyss struct {
	ScheduleID       int64   `json:"schedule_id"`
	StartTime        string  `json:"start_time"`
	EndTime          string  `json:"end_time"`
	TotalBattleTimes int64   `json:"total_battle_times"`
	TotalWinTimes    int64   `json:"total_win_times"`
	MaxFloor         string  `json:"max_floor"`
	RevealRank       []Rank  `json:"reveal_rank"`
	DefeatRank       []Rank  `json:"defeat_rank"`
	DamageRank       []Rank  `json:"damage_rank"`
	TakeDamageRank   []Rank  `json:"take_damage_rank"`
	NormalSkillRank  []Rank  `json:"normal_skill_rank"`
	EnergySkillRank  []Rank  `json:"energy_skill_rank"`
	Floors           []Floor `json:"floors"`
	TotalStar        int64   `json:"total_star"`
	IsUnlock         bool    `json:"is_unlock"`
}

type Rank struct {
	AvatarID   int64  `json:"avatar_id"`
	AvatarIcon string `json:"avatar_icon"`
	Value      int64  `json:"value"`
	Rarity     int64  `json:"rarity"`
}

type Floor struct {
	Index      int64   `json:"index"`
	Icon       string  `json:"icon"`
	IsUnlock   bool    `json:"is_unlock"`
	SettleTime string  `json:"settle_time"`
	Star       int64   `json:"star"`
	MaxStar    int64   `json:"max_star"`
	Levels     []Level `json:"levels"`
}

type Level struct {
	Index   int64    `json:"index"`
	Star    int64    `json:"star"`
	MaxStar int64    `json:"max_star"`
	Battles []Battle `json:"battles"`
}

type Battle struct {
	Index     int64    `json:"index"`
	Timestamp string   `json:"timestamp"`
	Avatars   []Avatar `json:"avatars"`
}

type Avatar struct {
	ID     int64  `json:"id"`
	Icon   string `json:"icon"`
	Level  int64  `json:"level"`
	Rarity int64  `json:"rarity"`
}
