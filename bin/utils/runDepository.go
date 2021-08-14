package utils

import (
	"fmt"
)

func RunDepository(URL string) {
	Cmd("git", fmt.Sprintln("clone ", URL))
}
