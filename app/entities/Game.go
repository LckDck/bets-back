package entities

// Game represents one game
type Game struct {
	ID int
	ShemeID int
	bets []Bet 
	Answer Bet 
	Results []Bet
	EndDate int32
}

// AddBet adds a bet to the game
func (g Game) AddBet(b Bet) {
	g.bets = append(g.bets, b)
}

// AddAnswer adds an correct answer to the game
func (g Game) AddAnswer(answer Bet) {
	g.Answer = answer
	g.calculate()
}


func (g Game) calculate() {
	//считаем, сколько баллов набрала каждая ставка и записываем в результат
	g.Results = nil
}