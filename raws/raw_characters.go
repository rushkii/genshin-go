package raws

type RawCharsData struct {
	Retcode int64  `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Chars []RawCharacters `json:"avatars"`
	} `json:"data"`
}

type RawCharacters struct {
	ID         int64               `json:"id"`
	Name       string              `json:"name"`
	Element    string              `json:"element"`
	Level      int64               `json:"level"`
	Rarity     int64               `json:"rarity"`
	Friendship int64               `json:"fetter"`
	Icon       string              `json:"icon"`
	Image      string              `json:"image"`
	ConsNum    int64               `json:"actived_constellation_num"`
	Weapon     RawWeapon           `json:"weapon"`
	Artifacts  []RawArtifact       `json:"reliquaries"`
	Cons       []RawConstellations `json:"constellations"`
	Skins      []RawSkin           `json:"costumes"`
}

type RawWeapon struct {
	ID         int64  `json:"id"`
	Type       string `json:"type_name"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Rarity     int64  `json:"rarity"`
	Level      int64  `json:"level"`
	Ascension  int64  `json:"promote_level"`
	Refinement int64  `json:"affix_level"`
	Icon       string `json:"icon"`
}

type RawArtifact struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	ArtName string    `json:"pos_name"`
	Pos     int64     `json:"pos"`
	Rarity  int64     `json:"rarity"`
	Level   int64     `json:"level"`
	Icon    string    `json:"icon"`
	Set     RawArtSet `json:"set"`
}

type RawArtSet struct {
	ID      int64       `json:"id"`
	Name    string      `json:"name"`
	Effects []RawEffect `json:"affixes"`
}

type RawEffect struct {
	Pieces int64  `json:"activation_number"`
	Fx     string `json:"effect"`
}

type RawConstellations struct {
	ID        int64  `json:"id"`
	Pos       int64  `json:"pos"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Effect    string `json:"effect"`
	IsActived bool   `json:"is_actived"`
}

type RawSkin struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
