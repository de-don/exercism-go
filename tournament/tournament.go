package tournament

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

/* ============================================= */
/* ========== Class for storing info =========== */
/* ============================================= */

type Info struct {
	team                            string
	played, won, drawn, lost, score int
}

func (info *Info) newWin() {
	info.won++
	info.played++
	info.score += 3
}

func (info *Info) newLose() {
	info.lost++
	info.played++
}

func (info *Info) newDraw() {
	info.played++
	info.drawn++
	info.score += 1
}

/* ============================================= */
/* ========== Class for sort info ============== */
/* ============================================= */

type byScore []*Info

func (s byScore) Len() int {
	return len(s)
}

func (s byScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byScore) Less(i, j int) bool {
	if s[i].score < s[j].score {
		return true
	}
	if s[i].score == s[j].score {
		return s[i].team > s[j].team
	}
	return false
}

/* ============================================= */
/* ============= Tally function ================ */
/* ============================================= */

func Tally(reader io.Reader, buffer *bytes.Buffer) error {
	teams := map[string]*Info{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		// skip newlines and comments
		if strings.TrimSpace(line) == "" || line[0] == '#' {
			continue
		}
		data := strings.Split(line, ";")
		if len(data) != 3 {
			return errors.New("failed read line")
		}

		if _, ok := teams[data[0]]; !ok {
			teams[data[0]] = &Info{team: data[0]}
		}
		team1 := teams[data[0]]

		if _, ok := teams[data[1]]; !ok {
			teams[data[1]] = &Info{team: data[1]}
		}
		team2 := teams[data[1]]

		switch data[2] {
		case "win":
			team1.newWin()
			team2.newLose()
		case "loss":
			team1.newLose()
			team2.newWin()
		case "draw":
			team1.newDraw()
			team2.newDraw()
		default:
			return errors.New("undefined result")
		}
	}

	// convert map to array
	data := make([]*Info, 0, len(teams))
	for _, g := range teams {
		data = append(data, g)
	}
	sort.Sort(byScore(data))

	// print header
	buffer.WriteString(fmt.Sprintf(
		"%-31s| %2s | %2s | %2s | %2s | %2s\n",
		"Team", "MP", "W", "D", "L", "P"))
	// print lines
	for i := range data {
		info := data[len(data)-i-1]
		buffer.WriteString(fmt.Sprintf(
			"%-31s| %2d | %2d | %2d | %2d | %2d\n",
			info.team, info.played, info.won, info.drawn, info.lost, info.score))
	}

	return nil
}
