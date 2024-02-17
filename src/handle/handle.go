package handle

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanInt(input *int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	*input = n
}

func Scan(input *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	*input = line
}

func ToRP(amount int) string {
	var Rp string
	for amount >= 1000 {
		Rp = fmt.Sprintf(".%03d%s", amount%1000, Rp)
		amount /= 1000
	}
	Rp = fmt.Sprintf("%d%s", amount, Rp)
	return Rp
}
