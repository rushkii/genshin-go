package structs

type Users struct {
	Characters []SimpleCharacters `json:"characters"`
	Stats      Stats              `json:"stats"`
	Worlds     []Worlds           `json:"explorations"`
	Teapots    Teapots            `json:"teapots"`
}

type SimpleCharacters struct {
	ID         int64  `json:"id"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Element    string `json:"element"`
	Friendship int64  `json:"friendship"`
	Level      int64  `json:"level"`
	Rarity     int64  `json:"rarity"`
	Cons       int64  `json:"constellations"`
}

type Stats struct {
	ActiveDay      int64  `json:"active_day"`
	Achievement    int64  `json:"achievement"`
	Characters     int64  `json:"characters"`
	WayPoint       int64  `json:"unlocked_waypoint"`
	Domain         int64  `json:"unlocked_domain"`
	SpiralAbyss    string `json:"spiral_abyss"`
	Anemoculus     int64  `json:"anemoculus"`
	Geoculus       int64  `json:"geoculus"`
	Electroculus   int64  `json:"electroculus"`
	PreciousChest  int64  `json:"precious_chest"`
	LuxuriousChest int64  `json:"luxurious_chest"`
	ExquisiteChest int64  `json:"exquisite_chest"`
	CommonChest    int64  `json:"common_chest"`
	MagicChest     int64  `json:"magic_chest"`
}

type Worlds struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Level     int64     `json:"level"`
	Explored  int64     `json:"progress"`
	Offerings Offerings `json:"offerings"`
}

type Teapots struct {
	Realms      []Realms `json:"realms"`
	Level       int64    `json:"level"`
	Visitors    int64    `json:"visitors"`
	Comforts    int64    `json:"comforts"`
	Items       int64    `json:"items"`
	ComfortName string   `json:"comfort_name"`
	ComfortIcon string   `json:"comfort_icon"`
}

type Realms struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Offerings struct {
	Name  string `json:"name"`
	Level int64  `json:"level"`
}
