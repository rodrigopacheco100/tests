package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entriesLength, _ := strconv.Atoi(scanner.Text())

	friendsWaitingAnswer := map[int]int{}
	eventsTimesMap := map[int]int{}

	lastEntryType := ""

	for range entriesLength {
		scanner.Scan()
		entry := scanner.Text()

		parts := strings.Split(entry, " ")
		entryType := parts[0]
		value, _ := strconv.Atoi(parts[1])

		if lastEntryType != "T" && entryType != "T" {
			incrementAllWaitingFriends(&friendsWaitingAnswer, 1)
		}

		switch entryType {
		case "T":
			incrementAllWaitingFriends(&friendsWaitingAnswer, value)

		case "R":
			friendsWaitingAnswer[value] = 0

		case "E":
			if awaitedTime, ok := friendsWaitingAnswer[value]; ok {
				eventsTimesMap[value] += awaitedTime
				delete(friendsWaitingAnswer, value)
			}
		}

		lastEntryType = entryType
	}

	for friendMissedAnswer := range friendsWaitingAnswer {
		eventsTimesMap[friendMissedAnswer] = -1
	}

	printFormattedResult(eventsTimesMap)
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
