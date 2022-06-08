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

func (c *Client) GetSpiralAbyss(uid string, previous bool) (*structs.Abyss, error) {
	var raw_data raws.RawAbyssData

	sched := 1
	if previous {
		sched = 2
	}

	body := bytes.NewBuffer(c.GET(fmt.Sprintf(
		"/genshin/api/spiralAbyss?server=%s&role_id=%s&schedule_type=%d", utils.GuessServer(uid), uid, sched,
	)))

	if err := json.NewDecoder(body).Decode(&raw_data); err != nil {
		return &structs.Abyss{}, err
	}

	if raw_data.Retcode == 10102 {
		return &structs.Abyss{}, DataNotPublic("the owner of %d settings is private, please go to hoyolab profile and set to public", uid)
	}

	return parser.ParseAbyss(&raw_data), nil
}
