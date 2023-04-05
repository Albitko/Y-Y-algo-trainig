package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func problem(reader io.Reader) string {
	var a, b, c, d, e int
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	a, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	d, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	e, _ = strconv.Atoi(scanner.Text())

	// Проверяем, можно ли вставить кирпич в отверстие
	if (a <= d && b <= e) || (a <= e && b <= d) ||
		(a <= d && c <= e) || (a <= e && c <= d) ||
		(b <= d && c <= e) || (b <= e && c <= d) {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
