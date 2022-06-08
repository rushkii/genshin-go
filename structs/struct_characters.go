package structs

type Characters struct {
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Element    string           `json:"element"`
	Level      int64            `json:"level"`
	Rarity     int64            `json:"rarity"`
	Friendship int64            `json:"friendship"`
	Icon       string           `json:"icon"`
	Image      string           `json:"image"`
	ConsNum    int64            `json:"cons_num"`
	Weapon     Weapon           `json:"weapon"`
	Artifacts  []Artifact       `json:"artifacts"`
	Cons       []Constellations `json:"constellations"`
	Skins      []Skin           `json:"skins"`
}

type Weapon struct {
	ID         int64  `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Desc       string `json:"description"`
	Rarity     int64  `json:"rarity"`
	Level      int64  `json:"level"`
	Ascension  int64  `json:"ascension"`
	Refinement int64  `json:"refinement"`
	Icon       string `json:"icon"`
}

type Artifact struct {
	ID      int64  `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	ArtName string `json:"item_name"`
	Rarity  int64  `json:"rarity"`
	Level   int64  `json:"level"`
	Icon    string `json:"icon"`
	Set     ArtSet `json:"set"`
}

type ArtSet struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Effects []Effect `json:"effects"`
}

type Effect struct {
	Pieces int64  `json:"pieces"`
	Fx     string `json:"effect"`
}

type Constellations struct {
	ID        int64  `json:"id"`
	Pos       int64  `json:"pos"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Effect    string `json:"effect"`
	IsActived bool   `json:"is_actived"`
}

type Skin struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
