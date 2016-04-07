package interior

type InteriorTreePolicy interface {
	GetScore(numTrials, numWins, parentTrials int64) float64
}