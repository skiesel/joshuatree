package players

import (
	"bufio"
	"fmt"
	"github.com/skiesel/joshuatree/domains"
	"os"
	"strconv"
	"strings"
)

type InteractivePlayer struct {
}

func NewInteractivePlayer() *InteractivePlayer {
	return &InteractivePlayer{}
}

func (player InteractivePlayer) GetAction(domain domains.Domain, state domains.State, opposingActions []domains.Action) domains.Action {
	actions := domain.GetAvailableActions(state)
	for {

		fmt.Printf("Please choose a move:\n")
		for i, action := range actions {
			fmt.Printf("%d) %s\n", i, domain.GetString(action))
		}

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		selection, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("failed to parse selection")
			continue
		}
		if selection < 0 || selection >= int64(len(actions)) {
			fmt.Println("selection is not a valid move")
			continue
		}
		return actions[selection]
	}
}
