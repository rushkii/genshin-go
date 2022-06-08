package structs

type Profile struct {
	Info *Info `json:"info"`
}

type Info struct {
	Nickname          string `json:"nickname"`
	Level             int    `json:"level"`
	Signature         string `json:"signature"`
	WorldLevel        int    `json:"world_level"`
	NameCardID        int    `json:"name_card_id"`
	FinishAchievement int    `json:"finish_achievement_num"`
	AbyssFloor        int    `json:"abyss_floor"`
	AbyssChamber      int    `json:"abyss_chamber"`
	ProfileIcon       string `json:"profile_icon"`
}
