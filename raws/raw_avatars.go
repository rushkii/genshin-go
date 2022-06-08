package raws

type RawAvatarData struct {
	Data struct {
		Avatars []struct {
			ID    int    `json:"id"`
			Icon  string `json:"icon"`
			Cover string `json:"image"`
		} `json:"avatars"`
	} `json:"data"`
	Message string `json:"message"`
	Retcode int    `json:"retcode"`
}
