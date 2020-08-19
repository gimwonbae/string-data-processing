package main

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strings"
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

func run() {
	var b bytes.Buffer
	checkList := ToSlice(`..\..\Set\SubtTV_2017_01_03_pcm.list.punct.trn`)
	orgList := ToSlice(`..\..\Set\SubtTV_2017_01_03_pcm.list.trn`)

	match, _ := os.OpenFile(`match`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	miss, _ := os.OpenFile(`miss`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	var re = regexp.MustCompile("[^가-힣0-9a-zA-Z]")

	for _, checkLine := range checkList {
		checkLineSplit := strings.Split(checkLine, " :: ")
		checkFile := checkLineSplit[0]
		checkText := checkLineSplit[1]
		checkSub := re.ReplaceAllString(checkText, "")

		for index, orgLine := range orgList {
			orgLineSplit := strings.Split(orgLine, " :: ")
			orgFile := orgLineSplit[0]
			if orgFile == checkFile {
				orgText := orgLineSplit[1]
				orgSub := re.ReplaceAllString(orgText, "")

				b.WriteString(checkLine)
				b.WriteString("\n")
				if strings.ToLower(checkSub) != strings.ToLower(orgSub) {
					miss.WriteString(b.String())
				} else {
					match.WriteString(b.String())
				}
				b.Reset()
				orgList = orgList[index+1:]
				break
			}
		}
	}
}

func main() {
	run()
}
