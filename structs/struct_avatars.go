package structs

type AvatarIcon struct {
	Avatars []struct {
		ID    int    `json:"id"`
		Icon  string `json:"icon"`
		Cover string `json:"image"`
	} `json:"avatars"`
}
