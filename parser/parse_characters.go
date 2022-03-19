package parser

import (
	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
	"github.com/rushkii/genshin-go/utils"
)

func ParseCharacters(data *raws.RawCharsData) *[]structs.Characters {
	var characters []structs.Characters
	Artifacts := make(chan []structs.Artifact)
	Cons := make(chan []structs.Constellations)
	Skins := make(chan []structs.Skin)
	Fx := make(chan []structs.Effect)

	for _, chars := range data.Data.Chars {
		go func(arch chan []structs.Artifact) {
			var arsl []structs.Artifact
			for i := 0; i < len(chars.Artifacts); i++ {
				go func(fxch chan []structs.Effect, ind int) {
					var fxsl []structs.Effect
					for a := 0; a < len(chars.Artifacts[ind].Set.Effects); a++ {
						fxsl = append(fxsl, structs.Effect{
							Pieces: chars.Artifacts[ind].Set.Effects[a].Pieces,
							Fx:     chars.Artifacts[ind].Set.Effects[a].Fx,
						})
					}
					fxch <- fxsl
				}(Fx, i)

				arsl = append(arsl, structs.Artifact{
					ID:      chars.Artifacts[i].ID,
					Type:    utils.GetArtifactsPos(chars.Artifacts[i].Pos),
					Name:    chars.Artifacts[i].Name,
					ArtName: chars.Artifacts[i].ArtName,
					Rarity:  chars.Artifacts[i].Rarity,
					Level:   chars.Artifacts[i].Level,
					Icon:    chars.Artifacts[i].Icon,
					Set: structs.ArtSet{
						ID:      chars.Artifacts[i].Set.ID,
						Name:    chars.Artifacts[i].Set.Name,
						Effects: <-Fx,
					},
				})
			}
			arch <- arsl
		}(Artifacts)

		go func(conch chan []structs.Constellations) {
			var consl []structs.Constellations
			for i := 0; i < len(chars.Cons); i++ {
				consl = append(consl, structs.Constellations{
					ID:        chars.Cons[i].ID,
					Pos:       chars.Cons[i].Pos,
					Name:      chars.Cons[i].Name,
					Icon:      chars.Cons[i].Icon,
					Effect:    chars.Cons[i].Effect,
					IsActived: chars.Cons[i].IsActived,
				})
			}
			conch <- consl
		}(Cons)

		go func(skinch chan []structs.Skin) {
			var skinsl []structs.Skin
			for i := 0; i < len(chars.Skins); i++ {
				skinsl = append(skinsl, structs.Skin{
					ID:   chars.Skins[i].ID,
					Name: chars.Skins[i].Name,
					Icon: chars.Skins[i].Icon,
				})
			}
			skinch <- skinsl
		}(Skins)

		characters = append(characters, structs.Characters{
			ID:         chars.ID,
			Name:       chars.Name,
			Element:    chars.Element,
			Level:      chars.Level,
			Rarity:     chars.Rarity,
			Friendship: chars.Friendship,
			Icon:       chars.Icon,
			Image:      chars.Image,
			ConsNum:    chars.ConsNum,
			Weapon: structs.Weapon{
				ID:         chars.Weapon.ID,
				Type:       chars.Weapon.Type,
				Name:       chars.Weapon.Name,
				Desc:       chars.Weapon.Desc,
				Rarity:     chars.Weapon.Rarity,
				Level:      chars.Weapon.Level,
				Ascension:  chars.Weapon.Ascension,
				Refinement: chars.Weapon.Refinement,
				Icon:       chars.Weapon.Icon,
			},
			Artifacts: <-Artifacts,
			Cons:      <-Cons,
			Skins:     <-Skins,
		})
	}
	return &characters
}
