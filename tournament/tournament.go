package tournament

import (
	"bufio"
	//"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name         string
	matcheswon   int
	matchesdrawn int
	matcheslost  int
	matchesnum   int
	points       int
}

func Tally(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	//scanner.Scan()
	teams := make(map[string]Team)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		data := strings.Split(line, ";")
		if len(data) != 3 {
			return fmt.Errorf("input is wrron")
		}

		leftTeam, rightTeam := data[0], data[1]
		tl := teams[leftTeam]
		tl.name = leftTeam
		tr := teams[rightTeam]
		tr.name = rightTeam

		switch data[2] { // win or loss or draw
		case "win":
			tl.matchesnum++
			tl.matcheswon++
			tl.points += 3
			tr.matcheslost++
			tr.matchesnum++

		case "loss":
			tl.matcheslost++
			tl.matchesnum++
			tr.matcheswon++
			tr.matchesnum++
			tr.points += 3
		case "draw":
			tl.matchesdrawn++
			tl.matchesnum++
			tl.points++
			tr.matchesdrawn++
			tr.matchesnum++
			tr.points++
		default:
			return fmt.Errorf("unknown match result %q in line %q (expected 'win', 'loss', or 'draw')", data, line)
		}
		teams[leftTeam] = tl
		teams[rightTeam] = tr

	}
	results := make([]Team, 0)
	for _, r := range teams {
		results = append(results, r)
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].points == results[j].points {
			return results[i].name < results[j].name
		}
		return results[i].points > results[j].points
	})
	fmt.Fprintf(output, "Team                           | MP |  W |  D |  L |  P\n")
	for _, r := range results {
		fmt.Fprintf(output, "%-30s | %2d | %2d | %2d | %2d | %2d\n", r.name, r.matchesnum, r.matcheswon, r.matchesdrawn, r.matcheslost, r.points)
	}

	return nil
}
