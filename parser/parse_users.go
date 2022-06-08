package parser

import (
	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
)

func ParseUsers(data *raws.RawUsersData) *structs.Users {
	var rarity int64
	var offering structs.Offerings
	var chsl []structs.SimpleCharacters
	var wsl []structs.Worlds
	var rsl []structs.Realms

	characters := make(chan []structs.SimpleCharacters)
	worlds := make(chan []structs.Worlds)
	realms := make(chan []structs.Realms)
	user := data.Data

	go func(chars chan []structs.SimpleCharacters) {
		for _, c := range user.Avatars {
			if c.Rarity < 100 {
				rarity = c.Rarity
			} else {
				rarity = c.Rarity - 100
			}
			chsl = append(chsl, structs.SimpleCharacters{
				ID:         c.ID,
				Image:      c.Image,
				Name:       c.Name,
				Element:    c.Element,
				Friendship: c.Fetter,
				Level:      c.Level,
				Rarity:     rarity,
				Cons:       c.Cons,
			})
		}
		chars <- chsl
	}(characters)

	go func(wchan chan []structs.Worlds) {
		for _, w := range user.World {
			if len(w.Offerings) != 0 {
				offering = structs.Offerings(w.Offerings[0])
			} else {
				offering = structs.Offerings{}
			}
			wsl = append(wsl, structs.Worlds{
				ID:        w.ID,
				Type:      w.Type,
				Name:      w.Name,
				Icon:      w.Icon,
				Level:     w.Level,
				Explored:  w.Explored / 10,
				Offerings: offering,
			})
		}
		wchan <- wsl
	}(worlds)

	go func(rchan chan []structs.Realms) {
		for _, t := range user.Homes {
			rsl = append(rsl, structs.Realms{
				Name: t.Name,
				Icon: t.Icon,
			})
		}
		rchan <- rsl
	}(realms)

	return &structs.Users{
		Stats: structs.Stats{
			ActiveDay:      user.Stats.ActiveDay,
			Achievement:    user.Stats.Achievement,
			Characters:     user.Stats.Avatar,
			WayPoint:       user.Stats.WayPoint,
			Domain:         user.Stats.Domain,
			SpiralAbyss:    user.Stats.SpiralAbyss,
			Anemoculus:     user.Stats.Anemoculus,
			Geoculus:       user.Stats.Geoculus,
			Electroculus:   user.Stats.Electroculus,
			PreciousChest:  user.Stats.PreciousChest,
			LuxuriousChest: user.Stats.LuxuriousChest,
			ExquisiteChest: user.Stats.ExquisiteChest,
			CommonChest:    user.Stats.CommonChest,
			RemarkChest:    user.Stats.MagicChest,
		},
		Worlds: <-worlds,
		Teapots: structs.Teapots{
			Realms:      <-realms,
			Level:       user.Homes[0].Level,
			Visitors:    user.Homes[0].Visitors,
			Comforts:    user.Homes[0].Comforts,
			Items:       user.Homes[0].Items,
			ComfortName: user.Homes[0].ComfortName,
			ComfortIcon: user.Homes[0].ComfortIcon,
		},
		Characters: <-characters,
	}
}
