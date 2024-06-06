package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/term"
)

func main() {
	printIntro()
	loop := true

	for loop {
		fmt.Println("Enter your choice to try prompts:")
		fmt.Println("1. Text input CLI prompt")
		fmt.Println("2. Password input CLI prompt")
		fmt.Println("3. Yes/No input CLI prompt")
		fmt.Println("4. CheckBox CLI prompt")
		fmt.Println("0. Exit")
		fmt.Print("Your choice: ")

		choice := getUserInput()

		switch choice {
		case "1":
			breakfast := TextPrompt("What would you like to have as breakfast, sir?")
			fmt.Printf("Here's your breakfast, sir: %s\n", breakfast)
		case "2":
			password := PasswordPrompt("Don't type any password here!!!")
			fmt.Printf("Oh, I can see your password: %s\n", password)
		case "3":
			ok := YesNoPrompt("Pineapple pizza is the best..", true)
			if ok {
				fmt.Println("No bruvvh, you are are gay")
			} else {
				fmt.Println("Huh?")
			}
		case "4":
			answers := Checkboxes(
				"Which are your favourite programming languages?",
				[]string{
					"C",
					"Python",
					"Java",
					"C++",
					"C#",
					"Visual Basic",
					"JavaScript",
					"PHP",
					"Assembly Language",
					"SQL",
					"Groovy",
					"Classic Visual Basic",
					"Fortran",
					"R",
					"Ruby",
					"Swift",
					"MATLAB",
					"Go",
					"Prolog",
					"Perl",
				},
			)
			s := strings.Join(answers, ", ")
			fmt.Println("Oh, I see! You like", s)
		case "0":
			fmt.Println("Exited")
			loop = false
		default:
			fmt.Println("Invalid choice, please enter an appropriate choice from the menu.")
		}
	}
}

func getUserInput() string {
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func TextPrompt(label string) string {
	fmt.Print(label + " ")
	return getUserInput()
}

func PasswordPrompt(label string) string {
	var input string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		p, _ := term.ReadPassword(int(syscall.Stdin))
		input = string(p)
		if p != nil {
			break
		}
		fmt.Println("Password cannot be empty. Please try again.")
	}
	fmt.Println() // Print a newline after password input
	return input
}

func YesNoPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)
		s := getUserInput()
		if s == "" {
			return def
		}
		s = strings.ToLower(s)
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}

func Checkboxes(label string, options []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: options,
	}
	survey.AskOne(prompt, &res)

	return res
}

func printIntro() {
	fmt.Println(`
 ▄████▄   ▒█████   ███▄ ▄███▓ ███▄ ▄███▓ ▄▄▄       ███▄    █ ▓█████▄  █    ██ 
▒██▀ ▀█  ▒██▒  ██▒▓██▒▀█▀ ██▒▓██▒▀█▀ ██▒▒████▄     ██ ▀█   █ ▒██▀ ██▌ ██  ▓██▒
▒▓█    ▄ ▒██░  ██▒▓██    ▓██░▓██    ▓██░▒██  ▀█▄  ▓██  ▀█ ██▒░██   █▌▓██  ▒██░
▒▓▓▄ ▄██▒▒██   ██░▒██    ▒██ ▒██    ▒██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█▄   ▌▓▓█  ░██░
▒ ▓███▀ ░░ ████▓▒░▒██▒   ░██▒▒██▒   ░██▒ ▓█   ▓██▒▒██░   ▓██░░▒████▓ ▒▒█████▓ 
░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ░  ░░ ▒░   ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ▒▒▓  ▒ ░▒▓▒ ▒ ▒ 
  ░  ▒     ░ ▒ ▒░ ░  ░      ░░  ░      ░  ▒   ▒▒ ░░ ░░   ░ ▒░ ░ ▒  ▒ ░░▒░ ░ ░ 
░        ░ ░ ░ ▒  ░      ░   ░      ░     ░   ▒      ░   ░ ░  ░ ░  ░  ░░░ ░ ░ 
░ ░          ░ ░         ░          ░         ░  ░         ░    ░       ░     
░                                                             ░               
`)
}
