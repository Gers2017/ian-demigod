package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Ian struct {
	languages []string
	modLevel  int
	isAChad   bool
}

func newIan(languages []string, modLevel int, isAChad bool) *Ian {
	ian := Ian{languages: languages, modLevel: modLevel, isAChad: isAChad}
	return &ian
}

func (ian *Ian) computePower() int {
	power := len(ian.languages) + ian.modLevel
	level := 10
	if !ian.isAChad {
		level = -10
	}
	power += level
	return power
}

func (ian *Ian) printIan() {
	fmt.Println("---Ian results---")
	fmt.Println("Languages: ", strings.Join(ian.languages, ", ")) // expected: [js, python, java]
	fmt.Println("ModLevel: ", ian.modLevel)
	if ian.isAChad {
		fmt.Println("Absolute chad")
	} else {
		fmt.Println("Not an absolute chad")
	}
	fmt.Println("Power: ", ian.computePower())
}

func (ian *Ian) isMostPowerfulThan(secondIan *Ian) bool {
	return ian.computePower() > secondIan.computePower()
}

func getLanguages() []string {
	var ianLanguages []string
	fmt.Println("So, what programming languages does your Ian know? like: java, python, lua")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		text = strings.ReplaceAll(text, " ", "")
		ianLanguages = append(strings.Split(text, ","))
	}
	return ianLanguages
}

func getModLevel() (int, error) {
	var modLevel int
	fmt.Println("What's his mod level? from 0 to 100")
	_, err := fmt.Scanln(&modLevel)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	if modLevel < 0 || modLevel > 100 {
		return 0, errors.New("mod level has to be in range(0, 100)")
	}

	return modLevel, nil
}

func getIsAChad() (*bool, error) {
	fmt.Println("Is your Ian an absolute chad? [Y/n]")
	var ianChadInput string
	_, err := fmt.Scanln(&ianChadInput)
	if err != nil {
		return nil, err
	}
	var lowerInput = strings.ToLower(ianChadInput)
	var isAChad = lowerInput == "yes" || lowerInput == "y" || lowerInput == "ye"

	return &isAChad, nil
}

func main() {
	// create a list of languages
	var ianLanguages = getLanguages()

	// get the mod level from the user from 0 to 100
	modLevel, err := getModLevel()
	if err != nil {
		fmt.Println(err)
		return
	}

	// get the chad boolean from the user
	isAChad, err := getIsAChad()
	if err != nil {
		fmt.Println(err)
		return
	}

	// generate the Ian
	ian := newIan(ianLanguages, modLevel, *isAChad)
	ian.printIan()
}
