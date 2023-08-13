package utils

import (
	"fmt"
	"go-spotify-cli/constants"
)

func PrintError(message string, err error) {
	fmt.Println(constants.Red + message + fmt.Sprint(err) + constants.Reset)
}
