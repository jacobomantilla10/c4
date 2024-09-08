package game

type Player struct {
	playerId  int
	character rune
}

func MakePlayer(id int, symbol rune) Player {
	return Player{id, symbol}
}

func (p *Player) GetId() int {
	return p.playerId
}

func (p *Player) GetChar() rune {
	return p.character
}
