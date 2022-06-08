package raws

type RawProfileData struct {
	PlayerInfo struct {
		Nickname          string `json:"nickname"`
		Level             int    `json:"level"`
		Signature         string `json:"signature"`
		WorldLevel        int    `json:"worldLevel"`
		NameCardID        int    `json:"nameCardId"`
		FinishAchievement int    `json:"finishAchievementNum"`
		AbyssFloor        int    `json:"towerFloorIndex"`
		AbyssChamber      int    `json:"towerLevelIndex"`
		ProfileIcon       struct {
			AvatarId int `json:"avatarId"`
		} `json:"profilePicture"`
	} `json:"playerInfo"`
}
