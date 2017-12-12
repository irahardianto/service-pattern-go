package interfaces

type IPlayerService interface {
	GetScores(player1Name string, player2Name string) (string, error)
}
