package parser

import (
	"github.com/rushkii/genshin-go/raws"
	"github.com/rushkii/genshin-go/structs"
)

func ParseAbyss(data *raws.RawAbyssData) *structs.Abyss {
	abyss := data.Data
	var r raws.Rank
	played := make(chan []structs.Rank)
	defeats := make(chan []structs.Rank)
	highdmg := make(chan []structs.Rank)
	dmgtkn := make(chan []structs.Rank)
	bursts := make(chan []structs.Rank)
	skills := make(chan []structs.Rank)
	floors := make(chan []structs.Floor)
	chambers := make(chan []structs.Chamber)
	battles := make(chan []structs.Battle)
	characters := make(chan []structs.Character)

	go func(plchan chan []structs.Rank) {
		var plksl []structs.Rank
		for _, r = range abyss.RevealRank {
			plksl = append(plksl, structs.Rank{
				ID:     r.AvatarID,
				Icon:   r.AvatarIcon,
				Rarity: r.Rarity,
				Value:  r.Value,
			})
		}
		plchan <- plksl
	}(played)

	go func(mdchan chan []structs.Rank) {
		var mdksl []structs.Rank
		for _, r = range abyss.DefeatRank {
			mdksl = append(mdksl, structs.Rank{
				ID:     r.AvatarID,
				Icon:   r.AvatarIcon,
				Rarity: r.Rarity,
				Value:  r.Value,
			})
		}
		mdchan <- mdksl
	}(defeats)

	go func(hdchan chan []structs.Rank) {
		var hdksl []structs.Rank
		for _, r = range abyss.DamageRank {
			hdksl = append(hdksl, structs.Rank{
				ID:     r.AvatarID,
				Icon:   r.AvatarIcon,
				Rarity: r.Rarity,
				Value:  r.Value,
			})
		}
		hdchan <- hdksl
	}(highdmg)

	go func(dtchan chan []structs.Rank) {
		var dtksl []structs.Rank
		for _, r = range abyss.TakeDamageRank {
			dtksl = append(dtksl, structs.Rank{
				ID:     r.AvatarID,
				Icon:   r.AvatarIcon,
				Rarity: r.Rarity,
				Value:  r.Value,
			})
		}
		dtchan <- dtksl
	}(dmgtkn)

	go func(brchan chan []structs.Rank) {
		var brksl []structs.Rank
		for _, r = range abyss.EnergySkillRank {
			brksl = append(brksl, structs.Rank{
				ID:     r.AvatarID,
				Icon:   r.AvatarIcon,
				Rarity: r.Rarity,
				Value:  r.Value,
			})
		}
		brchan <- brksl
	}(bursts)

	go func(skchan chan []structs.Rank) {
		var skksl []structs.Rank
		for _, r = range abyss.NormalSkillRank {
			skksl = append(skksl, structs.Rank{
				ID:     r.AvatarID,
				Icon:   r.AvatarIcon,
				Rarity: r.Rarity,
				Value:  r.Value,
			})
		}
		skchan <- skksl
	}(skills)

	go func(flchan chan []structs.Floor) {
		var flksl []structs.Floor
		for _, f := range abyss.Floors {
			go func(cmbchan chan []structs.Chamber) {
				var cmbsl []structs.Chamber
				for _, c := range f.Levels {
					go func(btlchan chan []structs.Battle) {
						var btlsl []structs.Battle
						for _, b := range c.Battles {
							go func(chchan chan []structs.Character) {
								var chsl []structs.Character
								for _, ch := range b.Avatars {
									chsl = append(chsl, structs.Character{
										ID:     ch.ID,
										Icon:   ch.Icon,
										Level:  ch.Level,
										Rarity: ch.Rarity,
									})
								}
								chchan <- chsl
							}(characters)
							btlsl = append(btlsl, structs.Battle{
								Party:      b.Index,
								Timestamp:  b.Timestamp,
								Characters: <-characters,
							})
						}
						btlchan <- btlsl
					}(battles)
					cmbsl = append(cmbsl, structs.Chamber{
						Chamber: c.Index,
						Stars:   c.Star,
						MaxStar: c.MaxStar,
						Battles: <-battles,
					})
				}
				cmbchan <- cmbsl
			}(chambers)
			flksl = append(flksl, structs.Floor{
				Floor:      f.Index,
				Icon:       f.Icon,
				MaxStar:    f.MaxStar,
				Stars:      f.Star,
				SettleTime: f.SettleTime,
				Chambers:   <-chambers,
			})
		}
		flchan <- flksl
	}(floors)

	return &structs.Abyss{
		Season:      abyss.ScheduleID,
		StartTime:   abyss.StartTime,
		EndTime:     abyss.EndTime,
		MaxFloor:    abyss.MaxFloor,
		Battles:     abyss.TotalBattleTimes,
		Wins:        abyss.TotalWinTimes,
		MostPlayed:  <-played,
		MostDefeats: <-defeats,
		HighDamage:  <-highdmg,
		DamageTaken: <-dmgtkn,
		BurstsUsed:  <-bursts,
		SkillsUsed:  <-skills,
		Floors:      <-floors,
	}
}
