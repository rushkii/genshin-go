package genshin

import (
	"encoding/json"
	"fmt"

	"github.com/rushkii/genshin-go/parser"
	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
	"github.com/rushkii/genshin-go/utils"
)

func (c *Client) GetUserStats(uid int) (*structs.Users, error) {
	var raw_data raws.RawUsersData
	var err error

	// users, err := c.Database.DBGetUserStats(uid)
	// if err == nil {
	// 	log.Panicln("Error DBGetuserStats: ", err)
	// }
	// if (time.Now().Unix() - users.UpdateAt) >= 3600 {

	// }

	body := c.GET(fmt.Sprintf("/genshin/api/index?server=%s&role_id=%d", utils.GetServer(uid), uid))

	if err = json.Unmarshal(body, &raw_data); err != nil {
		return &structs.Users{}, err
	}

	if raw_data.Retcode == 10102 {
		return &structs.Users{}, DataNotPublic("the owner of %d settings is private, please go to hoyolab profile and set to public", uid)
	}

	return parser.ParseUsers(&raw_data), nil
}
