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

func (c *Client) GetCharacters(uid string) (*[]structs.Characters, error) {
	var raw_data raws.RawCharsData
	data := []byte(fmt.Sprintf(`{
		"role_id": "%s",
		"server": "%s"
	}`, uid, utils.GuessServer(uid)))

	body := bytes.NewBuffer(c.POST("/genshin/api/character", data))

	if err := json.NewDecoder(body).Decode(&raw_data); err != nil {
		return &[]structs.Characters{}, err
	}

	if raw_data.Retcode == 10102 {
		return &[]structs.Characters{}, DataNotPublic("the owner of %d settings is private, please go to hoyolab profile and set to public", uid)
	}

	return parser.ParseCharacters(&raw_data), nil
}
