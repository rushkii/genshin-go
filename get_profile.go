package genshin

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
)

func (c *Client) GetProfile(uid string) (*structs.Profile, error) {
	var enka_data raws.RawProfileData
	var hoyo_data raws.RawAvatarData

	body := bytes.NewBuffer(EnkaReq(uid))
	if err := json.NewDecoder(body).Decode(&enka_data); err != nil {
		return &structs.Profile{}, err
	}

	data := []byte(fmt.Sprintf(`{
		"character_ids": [%d]
	}`, enka_data.PlayerInfo.ProfileIcon.AvatarId))
	body = bytes.NewBuffer(c.POST("genshin/api/avatarBasicInfo", data))
	if err := json.NewDecoder(body).Decode(&hoyo_data); err != nil {
		return &structs.Profile{}, err
	}

	if hoyo_data.Retcode == 10102 {
		return &structs.Profile{}, DataNotPublic("the owner of %d settings is private, please go to hoyolab profile and set to public", uid)
	}

	return &structs.Profile{
		Info: &structs.Info{
			Nickname:          enka_data.PlayerInfo.Nickname,
			Level:             enka_data.PlayerInfo.Level,
			Signature:         enka_data.PlayerInfo.Signature,
			WorldLevel:        enka_data.PlayerInfo.WorldLevel,
			NameCardID:        enka_data.PlayerInfo.NameCardID,
			FinishAchievement: enka_data.PlayerInfo.FinishAchievement,
			AbyssFloor:        enka_data.PlayerInfo.AbyssFloor,
			AbyssChamber:      enka_data.PlayerInfo.AbyssChamber,
			ProfileIcon:       hoyo_data.Data.Avatars[0].Icon,
		},
	}, nil
}
