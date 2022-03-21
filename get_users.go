package genshin

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rushkii/genshin-go/parser"
	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
	"github.com/rushkii/genshin-go/utils"
)

func (c *Client) GetUserStats(uid string) (*structs.Users, error) {
	var raw_data raws.RawUsersData

	// users, err := c.Database.DBGetUserStats(uid)
	// if err == nil {
	// 	log.Panicln("Error DBGetuserStats: ", err)
	// }
	// if (time.Now().Unix() - users.UpdateAt) >= 3600 {

	// }

	body := bytes.NewBuffer(c.GET(fmt.Sprintf(
		"/genshin/api/index?server=%s&role_id=%s", utils.GuessServer(uid), uid,
	)))

	if err := json.NewDecoder(body).Decode(&raw_data); err != nil {
		return &structs.Users{}, err
	}

	if raw_data.Retcode == 10102 {
		return &structs.Users{}, DataNotPublic("the owner of %d settings is private, please go to hoyolab profile and set to public", uid)
	}

	return parser.ParseUsers(&raw_data), nil
}
