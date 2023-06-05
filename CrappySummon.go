package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	ui "github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
)

func main() {
	items := []string{"Stick", "Additives", "Clackers", "Scissors", "Solvents", "Paint", "Pole", "Money", "Buckets", "Cows"}
	mixedItems := make([]string, 0)
	solution := []string{"Stick-Paint", "Additives-Solvents", "Clackers-Scissors", "Pole-Money", "Buckets-Cows"}
	solutionReverse := []string{"Paint-Stick", "Solvents-Additives", "Scissors-Clackers", "Money-Pole", "Cows-Buckets"}
	yesNo := []string{"YES", "NO"}

	fmt.Println("Welcome to ̶ ̶M̶a̶g̶i̶c̶S̶u̶m̶m̶o̶n̶  CrappySummon!\nThis game is all about mixing items, to summon your worst nightmare. Which is the.. 😳")
	fmt.Println("Would you like to start now?")
	startGame := prompt("Please select [YES/NO]", yesNo)

	if startGame == "YES" {
		fmt.Println("Starting game.. Might take a few seconds..")
		loadingBar(3)
		clearScreen()
		for len(items) > 0 {
			choiceOne := prompt("Mixing Items", items)
			items = remove(items, choiceOne)
			choiceTwo := prompt("Mixing Items", items)
			items = remove(items, choiceTwo)
			clearScreen()
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

func prompt(label string, items []string) string {
	prompt := ui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func generateNewItem(item1, item2 string) string {
	return strings.Join([]string{item1, item2}, "-")
}

func endGame(mixedItems, solution, solutionReverse []string) {
	var dump string

	fmt.Println("Ran out of items to mix. Now, let's see if your worst nightmare spawns..")
	loadingBar(5)

	isCorrect := false
	for _, item := range mixedItems {
		if contains(solution, item) || contains(solutionReverse, item) {
			isCorrect = true
			break
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

func loadingBar(second int) {
	b := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		b.Add(1)
		time.Sleep(50 * time.Millisecond)
	}
}

func contains(arr []string, item string) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}
	return false
}

func clearScreen() {
	if os.Getenv("OS") == "Windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
