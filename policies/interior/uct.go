package interior

import (
	"math"
)

type UCTPolicy struct {
	c float64
}

func NewUCTPolicy(c float64) *UCTPolicy {
	return &UCTPolicy{
		c: c,
	}
}

func (uct UCTPolicy) GetScore(numTrials, numWins, parentTrials int64) float64 {
	return float64(numWins) / float64(numTrials) * uct.c * math.Sqrt(math.Log(float64(parentTrials))/float64(numTrials))
}
