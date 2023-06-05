package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	ui "github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
)

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

//func areArraysEqual(arr1, arr2 []string) bool {
	//if len(arr1) != len(arr2) {
		//return false
	//}

	//freq := make(map[string]int)

	//for _, item := range arr1 {
		//freq[item]++
	//}

	//for _, item := range arr2 {
		//freq[item]--
	//}

	//for _, count := range freq {
		//if count != 0 {
			//return false
		//}
	//}

	//return true
//}

//func isSolutionCorrect(mixedItems, solution, solutionReverse []string) bool {
	//return areArraysEqual(mixedItems, solution) || areArraysEqual(mixedItems, solutionReverse)
//}

func main() {
	items := []string{"Stick", "Additives", "Clackers", "Scissors", "Solvents", "Paint", "Pole", "Money", "Buckets", "Cows"}
	mixedItems := make([]string, 0)
	solution := []string{"Stick-Paint", "Additives-Solvents", "Clackers-Scissors", "Pole-Money", "Buckets-Cows"}
	solutionReverse := []string{"Paint-Stick", "Solvents-Additives", "Scissors-Clackers", "Money-Pole", "Cows-Buckets"}

	fmt.Println("Welcome to ̶ ̶M̶a̶g̶i̶c̶S̶u̶m̶m̶o̶n̶  CrappySummon!\nThis game is all about mixing items, to summon your worst nightmare. Which is the.. 😳")
	fmt.Println("Would you like to start now?")
	startGame := yesNo()

	if startGame == "Yes" {
		fmt.Println("Starting game.. Might take a few seconds..")
		bar1 := progressbar.Default(100)
		for i := 0; i < 100; i++ {
			bar1.Add(1)
			time.Sleep(30 * time.Millisecond)
		}
		fmt.Print("\033[H\033[2J")
		for len(items) > 0 {
			choiceOne := mixing(items)
			items = remove(items, choiceOne)
			choiceTwo := mixing(items)
			items = remove(items, choiceTwo)
			fmt.Print("\033[H\033[2J")

			if len(items) >= 2 {
				newItem := generateNewItem(choiceOne, choiceTwo)
				mixedItems = append(mixedItems, newItem)
			}

			if len(items) == 0 {
				newItem := generateNewItem(choiceOne, choiceTwo)
				mixedItems = append(mixedItems, newItem)
				endGame(mixedItems, solution, solutionReverse)
			}
		}
	} else {
		time.Sleep(1e+9)
		fmt.Println("Quitting the game.")
		return
	}
}

func yesNo() string {
	prompt := ui.Select{
		Label: "Select Yes or No",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}

func mixing(items []string) string {
	prompt := ui.Select{
		Label: "Mix Items",
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}

func generateNewItem(item1, item2 string) string {
	return strings.Join([]string{item1, item2}, "-")
}

func endGame(mixedItems, solution, solutionReverse []string) {
	var dump string

	fmt.Println("Ran out of items to mix. Now, let's see if your worst nightmare spawns..")
	bar2 := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar2.Add(1)
		time.Sleep(50 * time.Millisecond)
	}

	isCorrect := true
	for i := 0; i < len(mixedItems); i++ {
		if i < len(solution) && mixedItems[i] != solution[i] {
			if i < len(solutionReverse) && mixedItems[i] != solutionReverse[i] {
				isCorrect = false
				break
			}
		}
	}
	if isCorrect {
		fmt.Println("CORRECT SOLUTION!")
		fmt.Scanln(&dump)
	} else {
		fmt.Println("Noob")
		fmt.Scanln(&dump)
	}

}
