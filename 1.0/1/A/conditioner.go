package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func problem(reader io.Reader) int {
	result := 0
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	temperature := scanner.Text()

	temperatureValues := strings.Split(temperature, " ")
	tRoom, _ := strconv.Atoi(temperatureValues[0])
	tCond, _ := strconv.Atoi(temperatureValues[1])

	scanner.Scan()
	mode := scanner.Text()
	switch mode {
	case "freeze":
		if tRoom <= tCond {
			return tRoom
		} else {
			return tCond
		}
	case "heat":
		if tRoom >= tCond {
			return tRoom
		} else {
			return tCond
		}
	case "auto":
		return tCond
	case "fan":
		return tRoom
	}

	return result
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
