package search

import (
	"fmt"
	"github.com/skiesel/joshuatree/domains"
	"github.com/skiesel/joshuatree/policies/exterior"
	"github.com/skiesel/joshuatree/policies/interior"
	"time"
)

type MonteCarlo struct {
	root           *decisionNode
	current        *decisionNode
	MoveTimeout    float64
	PlayerId       int64
	OtherPlayerIds []int64
	InteriorPolicy interior.InteriorTreePolicy
	ExteriorPolicy exterior.ExteriorTreePolicy
}

type decisionNode struct {
	state       domains.State
	arcs        map[int64]*decisionArc
	numTrials   int64
	initialized bool
}

type decisionArc struct {
	next      map[int64]*decisionNode
	action    domains.Action
	numTrials int64
	numWins   int64
}

func NewMonteCarlo(playerId int64, otherPlayerIds []int64, moveTimeout, uctCValue float64, domain domains.Domain) *MonteCarlo {
	startNode := buildStartNode(domain)
	return &MonteCarlo{
		root:           startNode,
		current:        startNode,
		MoveTimeout:    moveTimeout,
		PlayerId:       playerId,
		OtherPlayerIds: otherPlayerIds,
		InteriorPolicy: interior.NewUCTPolicy(uctCValue),
		ExteriorPolicy: exterior.NewRandomPolicy(),
	}
}

func buildStartNode(domain domains.Domain) *decisionNode {
	startState := domain.GetStartState()
	return &decisionNode{
		state:       startState,
		arcs:        map[int64]*decisionArc{},
		numTrials:   1,
		initialized: false,
	}
}

func (monteCarlo *MonteCarlo) UpdateForOpposingAction(domain domains.Domain, state domains.State, opposingActions []domains.Action) {
	//in case we get the first move...
	if len(opposingActions) == 0 {
		return
	}

	if len(monteCarlo.current.arcs) == 0 {
		monteCarlo.current.state = state
		return
	}

	actionHash := domain.HashAction(opposingActions[0])
	arc, found := monteCarlo.current.arcs[actionHash]
	if !found {
		arc = &decisionArc{}
		monteCarlo.current.arcs[actionHash] = arc
	}

	stateHash := domain.HashState(state)
	node, found := arc.next[stateHash]
	if !found {
		node = &decisionNode{
			state:       state,
			arcs:        map[int64]*decisionArc{},
			numTrials:   1,
			initialized: false,
		}
		arc.next[stateHash] = node
	}

	monteCarlo.current = node
}

func (monteCarlo MonteCarlo) GetAction(domain domains.Domain) domains.Action {

	moveStartTime := time.Now()

	for time.Since(moveStartTime).Seconds() < monteCarlo.MoveTimeout {

		currentNode := monteCarlo.current

		arcs := []*decisionArc{}
		for currentNode.initialized {
			bestArc := monteCarlo.getBestArc(currentNode)
			currentNode.numTrials++
			arcs = append(arcs, bestArc)
			currentNode = monteCarlo.sampleOutcome(domain, currentNode.state, bestArc)
		}

		if !domain.IsTerminal(currentNode.state) {
			currentNode.arcs = monteCarlo.createdecisionArcs(domain, currentNode.state)
			currentNode.initialized = true

			win := monteCarlo.doFullRollOut(domain, currentNode.state)
			for _, arc := range arcs {
				if win {
					arc.numWins++
				}
				arc.numTrials++
			}
		}
	}

	bestArc := monteCarlo.getBestArc(monteCarlo.current)

	for _, arc := range monteCarlo.current.arcs {
		fmt.Printf("%s : %d / %d\n", domain.GetString(arc.action), arc.numWins, arc.numTrials)
	}

	return bestArc.action
}

func (monteCarlo MonteCarlo) sampleOutcome(domain domains.Domain, state domains.State, arc *decisionArc) *decisionNode {
	nextState := domain.ApplyAction(state, arc.action, monteCarlo.PlayerId)

	for _, id := range monteCarlo.OtherPlayerIds {
		nextState = monteCarlo.doSimulationSingleStep(domain, state, id)
		if domain.IsTerminal(nextState) {
			break
		}
	}

	hash := domain.HashState(nextState)
	node, found := arc.next[hash]
	if !found {
		node = &decisionNode{
			state:       nextState,
			arcs:        map[int64]*decisionArc{},
			numTrials:   1,
			initialized: false,
		}
		arc.next[hash] = node
	}

	return node
}

func (monteCarlo MonteCarlo) getBestArc(node *decisionNode) *decisionArc {
	bestScore := -1.
	var bestArc *decisionArc
	for _, innerArc := range node.arcs {
		score := monteCarlo.InteriorPolicy.GetScore(innerArc.numTrials, innerArc.numWins, node.numTrials)
		if score > bestScore {
			bestArc = innerArc
		}
	}
	return bestArc
}

func (monteCarlo MonteCarlo) doFullRollOut(domain domains.Domain, state domains.State) bool {
	currentState := state
	for !domain.IsTerminal(currentState) {
		currentState = monteCarlo.doSimulationFullRound(domain, currentState)
	}

	// domain.Draw(currentState)
	// if domain.DidWin(currentState, monteCarlo.PlayerId) {
	// 	fmt.Println("win")
	// } else {
	// 	fmt.Println("!win")
	// }

	return domain.DidWin(currentState, monteCarlo.PlayerId)
}

func (monteCarlo MonteCarlo) doSimulationFullRound(domain domains.Domain, state domains.State) domains.State {
	nextState := monteCarlo.doSimulationSingleStep(domain, state, monteCarlo.PlayerId)

	if domain.IsTerminal(nextState) {
		return nextState
	}

	for _, id := range monteCarlo.OtherPlayerIds {
		nextState = monteCarlo.doSimulationSingleStep(domain, nextState, id)
		if domain.IsTerminal(nextState) {
			return nextState
		}
	}
	return nextState
}

func (monteCarlo MonteCarlo) doSimulationSingleStep(domain domains.Domain, state domains.State, playerId int64) domains.State {
	action := monteCarlo.ExteriorPolicy.GetAction(domain, state)
	return domain.ApplyAction(state, action, playerId)
}

func (monteCarlo MonteCarlo) createdecisionArcs(domain domains.Domain, state domains.State) map[int64]*decisionArc {
	actions := domain.GetAvailableActions(state)
	nodes := map[int64]*decisionArc{}
	for _, action := range actions {
		hash := domain.HashAction(action)
		nodes[hash] = &decisionArc{
			next:      map[int64]*decisionNode{},
			action:    action,
			numTrials: 1,
			numWins:   1,
		}
	}
	return nodes
}
