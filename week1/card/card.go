package card

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// a credit card has a checksum built into them
// we use Luhn's algorithm to determine if a credit card is syntactically correct
// if checksum is 0 is valid
// else is going to return 1
func checksumOk(s string) int {
	sum := 0
	for i := len(s) - 2; i >= 0; i -= 2 {
		// if number *2 is less than 10 we add the number to the sum
		// detract 48 according to ASCII table -- 48 is 0
		if int((s[i]-48)*2) < 10 {
			sum += int((s[i] - 48) * 2)
		} else {
			// else if the number *2 is bigger than 10 we have to separate the number
			// like this : 12 ---> 1+2=3
			// what i do in this case , ex : 12 - 10 = 2 , 2+1 =3
			sum += int((s[i]-48)*2 - 10 + 1)
		}
	}
	for j := len(s) - 1; j >= 0; j -= 2 {
		sum += int(s[j] - 48)
	}
	return sum % 10
}

// function that checks whether the input is valid
func validInput(s string) (bool, error) {
	if len(s) > 16 {
		return false, errors.New("input not valid: too long")
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		return false, errors.New("input not valid: input not numerical")
	}
	return true, nil
}

// AMEXs start with 34 , 37
// Mastercards start with 51,52,53,54,55
// Visas start with 4
func Card(s string) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "-", "")
	_, err := validInput(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if checksumOk(s) == 0 {
		if len(s) == 15 && (s[0:2] == "34" || s[0:2] == "37") {
			fmt.Println(s, "\nAmerican Express")
		} else if len(s) >= 13 && len(s) <= 16 && s[0] == '4' {
			fmt.Println(s, "\nVisa")
		} else if len(s) == 16 && (s[0:2] == "51" || s[0:2] == "52" || s[0:2] == "53" || s[0:2] == "54" || s[0:2] == "55") {
			fmt.Println(s, "\nMastercard")
		} else {
			fmt.Println("Card unknown")
		}
	} else {
		fmt.Println("Card not valid")
	}
}
