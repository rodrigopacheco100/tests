package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Event struct {
	Type  string
	Value int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entriesLength, _ := strconv.Atoi(scanner.Text())

	events := make([]Event, entriesLength)
	for i := range entriesLength {
		scanner.Scan()
		entry := scanner.Text()
		parts := strings.Split(entry, " ")
		value, _ := strconv.Atoi(parts[1])
		events[i] = Event{Type: parts[0], Value: value}
	}

	result := CalculateResponseTimes(events)
	printFormattedResult(result)
}

func CalculateResponseTimes(events []Event) map[int]int {
	friendsWaitingAnswer := map[int]int{}
	eventsTimesMap := map[int]int{}
	lastEntryType := ""

	for _, ev := range events {
		if lastEntryType != "T" && ev.Type != "T" {
			incrementAllWaitingFriends(&friendsWaitingAnswer, 1)
		}

		switch ev.Type {
		case "T":
			incrementAllWaitingFriends(&friendsWaitingAnswer, ev.Value)
		case "R":
			friendsWaitingAnswer[ev.Value] = 0
		case "E":
			if awaitedTime, ok := friendsWaitingAnswer[ev.Value]; ok {
				eventsTimesMap[ev.Value] += awaitedTime
				delete(friendsWaitingAnswer, ev.Value)
			}
		}

		lastEntryType = ev.Type
	}

	for friend := range friendsWaitingAnswer {
		eventsTimesMap[friend] = -1
	}

	return eventsTimesMap
}

func incrementAllWaitingFriends(friendsWaitingAnswer *map[int]int, value int) {
	for friend := range *friendsWaitingAnswer {
		(*friendsWaitingAnswer)[friend] += value
	}
}

func printFormattedResult(eventsTimesMap map[int]int) {
	sortedFriends := make([]int, 0, len(eventsTimesMap))
	for friend := range eventsTimesMap {
		sortedFriends = append(sortedFriends, friend)
	}

	sort.Ints(sortedFriends)

	for _, friend := range sortedFriends {
		fmt.Printf("%d %d\n", friend, eventsTimesMap[friend])
	}
}
