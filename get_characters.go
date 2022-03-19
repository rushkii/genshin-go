package genshin

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/rushkii/genshin-go/parser"
	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
	"github.com/rushkii/genshin-go/utils"
)

func (c *Client) GetCharacters(uid int) (*[]structs.Characters, error) {
	var raw_data raws.RawCharsData
	var err error

	data, err := json.Marshal(map[string]string{
		"role_id": strconv.Itoa(uid),
		"server":  utils.GetServer(uid),
	})
	if err != nil {
		log.Println(err)
	}

	body := bytes.NewBuffer(c.POST("/genshin/api/character", data))

	if err = json.NewDecoder(body).Decode(&raw_data); err != nil {
		return &[]structs.Characters{}, err
	}

	if raw_data.Retcode == 10102 {
		return &[]structs.Characters{}, DataNotPublic("the owner of %d settings is private, please go to hoyolab profile and set to public", uid)
	}

	return parser.ParseCharacters(&raw_data), nil
}
