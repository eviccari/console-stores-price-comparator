package domain

type GameInfo struct {
	Name     string
	Price    float64
	Platform string
}

func NewGameInfo(name, platform string, price float64) GameInfo {
	return GameInfo{
		Name:     name,
		Price:    price,
		Platform: platform,
	}
}
