package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	draw   string
	points int
}

var cardHierarchyWithoutJoker = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

var cardHierarchyWithJoker = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

type ByPowerWithJoker []Hand

func (h ByPowerWithJoker) Len() int      { return len(h) }
func (h ByPowerWithJoker) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByPowerWithJoker) Less(i, j int) bool {
	first := h[i]
	second := h[j]

	firstTypePower := GetHandPower(first, true)
	secondTypePower := GetHandPower(second, true)

	if firstTypePower != secondTypePower {
		return firstTypePower < secondTypePower
	}

	return LessBySecondOrderingRule(first, second, cardHierarchyWithJoker)
}

type ByPower []Hand

func (h ByPower) Len() int      { return len(h) }
func (h ByPower) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByPower) Less(i, j int) bool {
	first := h[i]
	second := h[j]

	firstTypePower := GetHandPower(first, false)
	secondTypePower := GetHandPower(second, false)

	if firstTypePower != secondTypePower {
		return firstTypePower < secondTypePower
	}

	return LessBySecondOrderingRule(first, second, cardHierarchyWithoutJoker)
}

func LessBySecondOrderingRule(first Hand, second Hand, cardHierarchy map[rune]int) bool {
	secondAsRunes := []rune(second.draw)
	for i, card := range first.draw {
		if card != secondAsRunes[i] {
			return cardHierarchy[card] < cardHierarchy[secondAsRunes[i]]
		}
	}

	return false
}

func GetHandPower(hand Hand, withJoker bool) int {
	switch {
	case IsFiveOfAKind(hand, withJoker):
		return 7
	case IsFourOfAKind(hand, withJoker):
		return 6
	case IsFullHouse(hand, withJoker):
		return 5
	case IsThreeOfAKind(hand, withJoker):
		return 4
	case IsTwoPair(hand, withJoker):
		return 3
	case IsPair(hand, withJoker):
		return 2
	default:
		return 1
	}
}

func IsTwoPair(hand Hand, withJoker bool) bool {
	if withJoker {
		return IsTwoPairWithJoker(hand)
	}

	return IsTwoPairNoJoker(hand)
}

func IsTwoPairNoJoker(hand Hand) bool {
	remainder := ""

	for _, card := range hand.draw {
		if !strings.ContainsRune(remainder, card) {
			remainder += string(card)
		}
	}

	return len(remainder) == 3
}

func IsTwoPairWithJoker(hand Hand) bool {
	remainder := ""
	jokerCount := 0

	for _, card := range hand.draw {
		if card == 'J' {
			jokerCount++
		} else {
			remainder += string(card)
		}
	}

	switch {
	case jokerCount >= 2:
		return true
	case jokerCount == 1:
		return HasCountOfSame(remainder, 2, false)
	default:
		return IsTwoPairNoJoker(hand)
	}
}

func IsFullHouse(hand Hand, withJoker bool) bool {
	if withJoker {
		return IsFullHouseWithJoker(hand)
	}
	return IsFullHouseNoJoker(hand)
}

func IsFullHouseNoJoker(hand Hand) bool {
	test := []rune(hand.draw)[0]
	remainder := ""

	for _, card := range hand.draw {
		if card != test {
			remainder += string(card)
		}
	}

	if len(remainder) == 3 && HasCountOfSame(remainder, 3, false) {
		return true
	}

	if len(remainder) == 2 && HasCountOfSame(remainder, 2, false) {
		return true
	}

	return false
}

func IsFullHouseWithJoker(hand Hand) bool {
	storage := ""

	for _, card := range hand.draw {
		if card != 'J' {
			storage += string(card)
		}
	}

	switch {
	case len(storage) == 5:
		return IsFullHouseNoJoker(hand)
	case len(storage) == 4:
		return IsTwoPairNoJoker(hand) || HasCountOfSameNoJokers(storage, 3)
	case len(storage) == 3:
		return HasCountOfSameNoJokers(storage, 2) || HasCountOfSameNoJokers(storage, 3)
	default:
		return true
	}
}

func IsPair(hand Hand, withJoker bool) bool {
	return HasCountOfSame(hand.draw, 2, withJoker)
}

func IsThreeOfAKind(hand Hand, withJoker bool) bool {
	return HasCountOfSame(hand.draw, 3, withJoker)
}

func IsFourOfAKind(hand Hand, withJoker bool) bool {
	return HasCountOfSame(hand.draw, 4, withJoker)
}

func IsFiveOfAKind(hand Hand, withJoker bool) bool {
	return HasCountOfSame(hand.draw, 5, withJoker)
}

func HasCountOfSame(draw string, count int, withJoker bool) bool {
	if withJoker {
		return HasCountOfSameWithJokers(draw, count)
	}

	return HasCountOfSameNoJokers(draw, count)
}

func HasCountOfSameWithJokers(draw string, count int) bool {
	highestCount := 0
	jokerCount := 0
	for _, outerCard := range draw {
		if outerCard == 'J' {
			jokerCount++
			continue
		}
		currentCount := 0
		for _, innerCard := range draw {
			if outerCard == innerCard {
				currentCount++
			}
		}
		if highestCount < currentCount {
			highestCount = currentCount
		}
		if highestCount+jokerCount >= count {
			return true
		}
	}

	return highestCount+jokerCount >= count
}

func HasCountOfSameNoJokers(draw string, count int) bool {
	highestCount := 0
	for _, outerCard := range draw {
		currentCount := 0
		for _, innerCard := range draw {
			if outerCard == innerCard {
				currentCount++
			}
		}
		if highestCount < currentCount {
			highestCount = currentCount
		}
		if highestCount >= count {
			return true
		}
	}

	return false
}

func main() {
	inputLines, _ := utils.GetProblemLines()

	hands := []Hand{}

	for _, line := range inputLines {
		split := strings.Split(line, " ")
		points, _ := strconv.Atoi(split[1])
		hands = append(hands, Hand{draw: split[0], points: points})
	}

	sort.Sort(ByPower(hands))

	sum := 0

	for i, hand := range hands {
		sum += hand.points * (i + 1)
	}

	fmt.Println(sum)

	sort.Sort(ByPowerWithJoker(hands))

	sum = 0

	for i, hand := range hands {
		sum += hand.points * (i + 1)
	}

	fmt.Println(sum)
}
