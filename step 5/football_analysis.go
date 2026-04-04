package main

import (
	"cmp"
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Mises   int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() float64 {
	if p.Mises == 0 {
		p.Rating = float64(p.Goals+p.Assists / 2)
		return p.Rating
	}
	p.Rating = ((float64(p.Goals+p.Assists / 2)) / float64(p.Mises))
	return p.Rating
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{
		Name:    name,
		Goals:   goals,
		Mises:   misses,
		Assists: assists,
	}
	p.calculateRating()
	return p
}

func goalsSort(players []Player) []Player {
	sliceGoalSort := make([]Player, len(players))
	copy(sliceGoalSort, players)

	slices.SortFunc(sliceGoalSort, func(a, b Player) int {
		if n := cmp.Compare(b.Goals, a.Goals); n != 0 {
			return n
		}
		return cmp.Compare(a.Name, b.Name)
	})
	return sliceGoalSort
}

func ratingSort(players []Player) []Player {
	sliceRatingSort := make([]Player, len(players))
	copy(sliceRatingSort, players)

	slices.SortFunc(sliceRatingSort, func(a, b Player) int {
		if n := cmp.Compare(b.Rating, a.Rating); n != 0 {
			return n
		}
		return cmp.Compare(a.Name, b.Name)
	})
	return sliceRatingSort
}

func gmSort(players []Player) []Player {
	sliceGmSort := make([]Player, len(players))
	copy(sliceGmSort, players)

	slices.SortFunc(sliceGmSort, func(a, b Player) int {
		rationA := float64(a.Goals)
		if a.Mises != 0 {
			rationA = float64(a.Goals) / float64(a.Mises)
		}

		rationB := float64(b.Goals)
		if b.Mises != 0 {
			rationB = float64(b.Goals) / float64(b.Mises)
		}

		if n := cmp.Compare(rationB, rationA); n != 0 {
			return n
		}
		return cmp.Compare(a.Name, b.Name)
	})
	return sliceGmSort
}

func main() {
	plr1 := NewPlayer("Федор",  10, 0, 4)
	plr2 := NewPlayer("Артём", 8, 2, 4)
	plr3 := NewPlayer("Михаил", 5, 1, 0)
	plr4 := NewPlayer("Стёпа", 12, 3, 6)

	fmt.Println(goalsSort([]Player{plr1, plr2, plr3, plr4}))
	fmt.Println(ratingSort([]Player{plr1, plr2, plr3, plr4}))
	fmt.Println(gmSort([]Player{plr1, plr2, plr3, plr4}))
}
