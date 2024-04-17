package entities

import "iGaming/database/models"

type Player struct {
	UserID   string
	Nickname string
	JPKey    string
	Balance  *Balance
}

func (p *Player) ConvertPlayerToEntity(player *models.Player) {
	p.UserID = player.UserID
	p.Nickname = player.UserNick
	p.JPKey = player.JPKey
}
