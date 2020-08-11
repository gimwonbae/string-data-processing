package main

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func ToSlice(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// println(lines)
	return lines
}

func lcs(a, b string) int {
	arunes := []rune(a)
	brunes := []rune(b)
	aLen := len(arunes)
	bLen := len(brunes)
	lengths := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lengths[i] = make([]int, bLen+1)
	}
	// row 0 and column 0 are initialized to 0 already

	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if arunes[i] == brunes[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]rune, 0, lengths[aLen][bLen])
	for x, y := aLen, bLen; x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, arunes[x-1])
			x--
			y--
		}
	}
	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s)
}

func run(fileNumber int) {
	var b bytes.Buffer
	cpLines := ToSlice(`C:\Project\20200804.-방송DB후처리\broadcast_2017_01_03.captions_u.txt`)
	nfLines := ToSlice(`..\..\testSet\SubtTV_2017_not_found_` + strconv.Itoa(fileNumber) + `.trn`)
	lcsMatch, err := os.OpenFile(`..\..\testSet\SubtTV_2017_lcs_match_`+strconv.Itoa(fileNumber)+`.trn`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// var re = regexp.MustCompile("[,.?! ~\n]")
	if err != nil {
		panic(err)
	}
	defer lcsMatch.Close()

	for _, nfLine := range nfLines {
		// fileName := strings.Split(nfLine, " :: ")[0]
		nfTxt := strings.Split(nfLine, " :: ")[1]
		// nfTxtSub := strings.ReplaceAll(nfTxt, " ", "")
		maxValue := 0
		lcsValue := 0
		var line string
		for _, cpLine := range cpLines {
			// cpLineSub := re.ReplaceAllString(cpLine, "")
			// lcsValue = lcs(cpLineSub, nfTxtSub)
			lcsValue = lcs(cpLine, nfTxt)
			if lcsValue > maxValue {
				line = cpLine
				maxValue = lcsValue
			}
		}
		b.WriteString(nfTxt)
		b.WriteString(" => ")
		b.WriteString(line)
		b.WriteString("\n")
		lcsMatch.WriteString(b.String())
		b.Reset()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(4)
	start := 1
	for i := start; i <= start+3; i++ {
		go func(i int) {
			defer wg.Done()
			run(i)
		}(i)
	}
	wg.Wait()
}
