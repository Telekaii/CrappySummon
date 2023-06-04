package main

import (
	"fmt"
	"log"
	"time"

	ui "github.com/manifoldco/promptui"
)

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	isGameOver := false
	inventory := []string{"Stick", "Computer", "Pen", "Mug", "Wallet", "Apple", "Sunglasses", "Notebook", "Vase", "Spoon"}
	tempInventory := inventory

	fmt.Println("Welcome to ̶ ̶M̶a̶g̶i̶c̶S̶u̶m̶m̶o̶n̶  CrappySummon!\nThis game is all about mixing items, to summon random things! The main objective is summoning the.. 😳")
	fmt.Println("Would you like to start now?")
	startGame := yesNo()

	if startGame == "Yes" {
		fmt.Println("Starting game.. Might take a few seconds..")
		time.Sleep(2e+9)
		fmt.Print("\033[H\033[2J")
		for !isGameOver {
			itemDeleted := mixing(inventory)
			inventory = remove(inventory, itemDeleted)
			actualMixing(inventory, tempInventory)
			fmt.Println(inventory)
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

func actualMixing(inventory []string, tempInventory []string) {
	fmt.Println(".")
}
