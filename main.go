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
	fiveNumbers()
}

func fiveNumbers() {
	// adding 5 numbers
	number1 := askUser("Tell me a number")
	fmt.Println(number1)

	number2 := askUser("Tell me a number")
	fmt.Println(number2)

	number3 := askUser("Tell me a number")
	fmt.Println(number3)

	number4 := askUser("Tell me a number")
	fmt.Println(number4)

	number5 := askUser("Tell me a number")
	fmt.Println(number5)

	i, _ := strconv.Atoi(number1)
	i2, _ := strconv.Atoi(number2)
	i3, _ := strconv.Atoi(number3)
	i4, _ := strconv.Atoi(number4)
	i5, _ := strconv.Atoi(number5)

	fmt.Println(i + i2 + i3 + i4 + i5)
}

func favouriteColour() {
	// fav colour
	colour := askUser("What is your favourite colour?")
	if strings.EqualFold("yellow", colour) {
		fmt.Println("Wow you have amazing taste!")
	} else {
		fmt.Println("You lack taste")
	}
}

func inputValidation() {
	// Input validation
	for {
		numberValidation := askUser("Please enter either 1,2 or 3")
		if (numberValidation == "1") || (numberValidation == "2") || (numberValidation == "3") {
			fmt.Println("You entered:", numberValidation)
			break
		} else {
			fmt.Println("Please enter a valid number")
		}
	}
}

func expenseFormChecklist() {
	// checklist expenses
	receiptCheck := askUser("Do you have a receipt?")
	if receiptCheck != "yes" {
		fmt.Println("Resubmit with a receipt!")
	} else {
		vatReceieptCheck := askUser("Does your receipt have a VAT Registration Number, Date, Amount?")
		if vatReceieptCheck != "yes" {
			fmt.Println("Resubmit with a valid VAT invoice")
		} else {
			vatReceiptUpload := askUser("Vat receipt uploaded?")
			if vatReceiptUpload != "yes" {
				fmt.Println("Upload and resubmit")
			} else {
				vatReceiptClaimDate := askUser("Date on Vat receiept used on claim?")
				if vatReceiptClaimDate != "yes" {
					fmt.Println("Correct or resubmit")
				} else {
					vatReceiptCorrectClaimAmount := askUser("Does your claimed amount in the form match the VAT receipts total?")
					if vatReceiptCorrectClaimAmount != "yes" {
						fmt.Println("Correct or resubmit")
					} else {
						vatReceiptAttendees := askUser("Is the attendees field filled out with at least one person?")
						if vatReceiptAttendees != "yes" {
							fmt.Println("Correct or resubmit")
						} else {
							vatReceiptProvidedServiceClaim := askUser("Did you claim something already provided by BJSS?")
							if vatReceiptProvidedServiceClaim == "yes" {
								fmt.Println("Remove item from claims")
							} else {
								vatReceiptAnyMoreClaims := askUser("Do you have more expenses to claim?")
								if vatReceiptAnyMoreClaims == "yes" {
									fmt.Println("Use add expense to add them to the same claim form!")
								} else {
									fmt.Println("Thanks for completing the form!")
								}
							}
						}
					}
				}
			}

		}
	}
}

func guessMyNumber() {
	fmt.Println("Game: Guess player 1's secret number")
	secretNumber := askUser("Please enter secret number:")
	secretNum, _ := strconv.Atoi(secretNumber)
	fmt.Println(secretNumber)

	clearScreen()
	// convert into int
	for numOfGuesses := 1; numOfGuesses <= 5; numOfGuesses++ {
		numberGuess := askUser("Guess the secret number:")
		numGuess, _ := strconv.Atoi(numberGuess)
		if numGuess < secretNum {
			fmt.Println("Number is too low")
		}
		if numberGuess > secretNumber {
			fmt.Println("Number is too high")
		}
		if numberGuess == secretNumber {
			fmt.Println("You have won!")
			return
		}
		if numOfGuesses == 5 {
			fmt.Println("You have lost!")
			return
		}
	}
}

func clearScreen() {
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
}

func acmeChatBot() {
	fmt.Println("Welcome to the Acme One Store - We sell only one product!")
	userHelp := askUser("What would you like help with?")
	helpString := "help"
	priceString := "price"
	if userHelp == "help" {
		fmt.Println("Please read the manual")
	}
	if userHelp == "price" {
		fmt.Println("The product price is £99")
	}
	if (helpString != userHelp) && (priceString != userHelp) {
		fmt.Println("Not understood, please call: 0845 555 5555")
	}
}

// make while loop
func addingMachine() {
	n := 0
	for {
		fmt.Println("If you input 99999 program will end...")
		addingMachineInput := askUser("Please input whole number:")
		MachineInput, _ := strconv.Atoi(addingMachineInput)
		if MachineInput == 99999 {
			// fmt.Println(n)
			return
		}
		n += MachineInput
		fmt.Println(n)
	}
}

func nationalAnthems() {

	// ukAnthem := strings.ContainsAny("God save our gracious king")
	// germanAnthem := "Germany, Germany above all Above all in the world"
	// usAnthem := "Oh say can you see, by the dawn's early light"

	nationalAnthemLyric := askUser("Please enter the first line of a national anthem:")

	ukAnthem := "God save our gracious king"
	germanAnthem := "Germany, Germany above all Above all in the world"
	usAnthem := "Oh say can you see, by the dawn's early light"

	if strings.EqualFold(germanAnthem, nationalAnthemLyric) || strings.ContainsAny("germany", nationalAnthemLyric) {
		fmt.Println("This is the national anthem of Germany")
	}
	if strings.EqualFold(usAnthem, nationalAnthemLyric) || strings.HasPrefix("oh say", nationalAnthemLyric) {
		fmt.Println("This is the national anthem of the United States")
	}
	if strings.EqualFold(ukAnthem, nationalAnthemLyric) {
		fmt.Println("This is the national anthem of the United Kingdom")
	}

}

func favePetFinder() {
	for {
		favePetAsker := askUser("Do you like meows, barks, or bubbles?")
		if favePetAsker == "bubbles" {
			fmt.Println("Your favourite pet is a fish")
			return
		}
		if favePetAsker == "meows" {
			meowsType := askUser("Short or long tails?")
			if meowsType == "short" {
				fmt.Println("Your favourite pet is a Manx cat")
				return
			} else if meowsType == "long" {
				meowsLongType := askUser("Do you like solid colour fur or mixed colours?")
				if meowsLongType == "mixed" {
					fmt.Println("You like tabby cats")
					return
				} else if meowsLongType == "solid" {
					fmt.Println("You like black cats!?")
					MeowsLongTypeSolid := askUser("It is a black cat right?")
					if MeowsLongTypeSolid == "yes" {
						fmt.Println("See I knew I couldn't be wrong haha!")
						return
					} else {
						fmt.Println("My guess is that you like white cats!")
						return
					}

				}

			}
		}
		if favePetAsker == "barks" {
			barksType := askUser("Do you like tiny or big dogs?")
			if barksType == "tiny" {
				fmt.Println("You like chihuahuas")
				return
			} else {
				fmt.Println("Wow there are lots of dogs you might like!")
				return
			}
		}
	}
}

// NEW GO TASKS - FRIDAY

func threeIntLargestNum() {

	for {
		number1 := askUser("First number between 1-10:")
		intNum1, _ := strconv.Atoi(number1)
		if intNum1 > 10 || intNum1 <= 0 {
			fmt.Println("Error incorrect value")
			continue

		} else {

			for {
				number2 := askUser("Second number between 1-10:")
				intNum2, _ := strconv.Atoi(number2)
				if intNum2 > 10 || intNum2 <= 0 {
					fmt.Println("Error incorrect value")
					continue

				} else {

					for {
						number3 := askUser("Third number between 1-10:")
						intNum3, _ := strconv.Atoi(number3)
						if intNum3 > 10 || intNum3 <= 0 {
							fmt.Println("Error incorrect value")
							continue

						} else {

							firstValue := (intNum1 * (intNum2 + intNum3))
							secondValue := (intNum1 * intNum2 * intNum3)
							thirdValue := (intNum1 + intNum2*intNum3)
							fourthValue := ((intNum1 + intNum2) * intNum3)

							numbers := []int{firstValue, secondValue, fourthValue, thirdValue}
							max := max(numbers)
							fmt.Println("The largest value calculated:", max)
							return

						}
					}

				}
			}
		}

	}
}

func max(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}

func palindromeStrCheck() {
	palindromeChecker := askUser("Palindrome check: ")
	fmt.Println(palindromeChecker)

}

func isPalindrome(str string) {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func askUser(prompt string) string {
	fmt.Println(prompt)
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	return scanner.Text()
}
