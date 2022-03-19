package raws

type RawUsersData struct {
	Retcode int64    `json:"retcode"`
	Message string   `json:"message"`
	Data    RawUsers `json:"data"`
}

type RawUsers struct {
	Avatars []RawUsersAvatar `json:"avatars"`
	Stats   RawUsersStats    `json:"stats"`
	World   []RawUsersWorld  `json:"world_explorations"`
	Homes   []RawUsersHome   `json:"homes"`
}

type RawUsersAvatar struct {
	ID        int64  `json:"id"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Element   string `json:"element"`
	Fetter    int64  `json:"fetter"`
	Level     int64  `json:"level"`
	Rarity    int64  `json:"rarity"`
	Cons      int64  `json:"actived_constellation_num"`
	CardImage string `json:"card_image"`
	IsChosen  bool   `json:"is_chosen"`
}

type RawUsersHome struct {
	Level       int64  `json:"level"`
	Visitors    int64  `json:"visit_num"`
	Comforts    int64  `json:"comfort_num"`
	Items       int64  `json:"item_num"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	ComfortName string `json:"comfort_level_name"`
	ComfortIcon string `json:"comfort_level_icon"`
}

type RawUsersStats struct {
	ActiveDay      int64  `json:"active_day_number"`
	Achievement    int64  `json:"achievement_number"`
	WinRate        int64  `json:"win_rate"`
	Anemoculus     int64  `json:"anemoculus_number"`
	Geoculus       int64  `json:"geoculus_number"`
	Avatar         int64  `json:"avatar_number"`
	WayPoint       int64  `json:"way_point_number"`
	Domain         int64  `json:"domain_number"`
	SpiralAbyss    string `json:"spiral_abyss"`
	PreciousChest  int64  `json:"precious_chest_number"`
	LuxuriousChest int64  `json:"luxurious_chest_number"`
	ExquisiteChest int64  `json:"exquisite_chest_number"`
	CommonChest    int64  `json:"common_chest_number"`
	Electroculus   int64  `json:"electroculus_number"`
	MagicChest     int64  `json:"magic_chest_number"`
}

type RawUsersWorld struct {
	Level     int64              `json:"level"`
	Explored  int64              `json:"exploration_percentage"`
	Icon      string             `json:"icon"`
	Name      string             `json:"name"`
	Type      string             `json:"type"`
	Offerings []RawUsersOffering `json:"offerings"`
	ID        int64              `json:"id"`
}

type RawUsersOffering struct {
	Name  string `json:"name"`
	Level int64  `json:"level"`
}
