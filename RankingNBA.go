package main

import (
	"regexp"
	"strconv"
	"strings"
)

type ScoreMeta struct {
	Wins     int
	Draws    int
	Losses   int
	Scored   int
	Conceded int
	Points   int
}

func NbaCup(resultSheet, toFind string) string {
	if len(toFind) == 0 {
		return ""
	}
	scoreList := strings.Split(resultSheet, ",")
	meta := ScoreMeta{
		Wins:     0,
		Draws:    0,
		Losses:   0,
		Scored:   0,
		Conceded: 0,
		Points:   0,
	}
	occurences := 0

	for _, e := range scoreList {
		matched, _ := regexp.MatchString("\\b"+toFind+"\\b", e)
		if matched {
			occurences += 1
		}
		if strings.Contains(e, toFind) {
			start := strings.Index(e, toFind)
			scores := parseScores(e)
			if len(scores) == 1 {
				return "Error(float number):" + e
			}
			incrementGameMeta(&meta, scores, start)
		}
	}
	incrementSummary(&meta)
	if occurences > 0 {
		return formatResult(&meta, toFind)
	}
	return toFind + ":This team didn't play!"
}
func incrementSummary(meta *ScoreMeta) {
	meta.Points += meta.Wins * 3
	meta.Points += meta.Draws * 1
	meta.Points += meta.Losses * 0
}

func formatResult(meta *ScoreMeta, toFind string) string {
	wins := strconv.Itoa(meta.Wins)
	draws := strconv.Itoa(meta.Draws)
	losses := strconv.Itoa(meta.Losses)
	scored := strconv.Itoa(meta.Scored)
	conceded := strconv.Itoa(meta.Conceded)
	points := strconv.Itoa(meta.Points)
	return toFind + ":W=" + wins + ";D=" + draws + ";L=" + losses + ";Scored=" + scored + ";Conceded=" + conceded + ";Points=" + points
}

func incrementGameMeta(meta *ScoreMeta, scores []int, index int) {
	if index <= 1 {
		meta.Scored += scores[0]
		meta.Conceded += scores[1]
		if scores[0] > scores[1] {
			meta.Wins += 1
		} else if scores[0] == scores[1] {
			meta.Draws += 1
		} else {
			meta.Losses += 1
		}
	} else {
		meta.Scored += scores[1]
		meta.Conceded += scores[0]
		if scores[1] > scores[0] {
			meta.Wins += 1
		} else if scores[1] == scores[0] {
			meta.Draws += 1
		} else {
			meta.Losses += 1
		}
	}
}

func parseScores(entry string) []int {
	scoreItem := strings.Split(entry, " ")
	scores := make([]int, 0)
	for _, e := range scoreItem {
		s, err := strconv.Atoi(e)
		if err == nil {
			scores = append(scores, s)
		}
	}
	return scores
}
