package parser

import (
	"sync"

	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
)

var wg sync.WaitGroup

func ParseUsers(data *raws.RawUsersData) *structs.Users {
	var characters []structs.SimpleCharacters
	var worlds []structs.Worlds
	var realms []structs.Realms

	user := data.Data
	wg.Add(3)

	go func() {
		defer wg.Done()
		var rarity int64

		for _, c := range user.Avatars {
			if c.Rarity < 100 {
				rarity = c.Rarity
			} else {
				rarity = c.Rarity - 100
			}
			characters = append(characters, structs.SimpleCharacters{
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
	}()

	go func() {
		defer wg.Done()
		for _, w := range user.World {
			var offering structs.Offerings

			if len(w.Offerings) != 0 {
				offering = structs.Offerings(w.Offerings[0])
			} else {
				offering = structs.Offerings{}
			}
			worlds = append(worlds, structs.Worlds{
				ID:        w.ID,
				Type:      w.Type,
				Name:      w.Name,
				Icon:      w.Icon,
				Level:     w.Level,
				Explored:  w.Explored / 10,
				Offerings: offering,
			})
		}
	}()

	go func() {
		defer wg.Done()
		for _, t := range user.Homes {
			realms = append(realms, structs.Realms{
				Name: t.Name,
				Icon: t.Icon,
			})
		}
	}()

	wg.Wait()

	return &structs.Users{
		Characters: characters,
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
			MagicChest:     user.Stats.MagicChest,
		},
		Worlds: worlds,
		Teapots: structs.Teapots{
			Realms:      realms,
			Level:       user.Homes[0].Level,
			Visitors:    user.Homes[0].Visitors,
			Comforts:    user.Homes[0].Comforts,
			Items:       user.Homes[0].Items,
			ComfortName: user.Homes[0].ComfortName,
			ComfortIcon: user.Homes[0].ComfortIcon,
		},
	}
}
