package main

import (
	"bufio"
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

func ToMap(lines []string) map[int][]string {
	var lenMap map[int][]string
	lenMap = make(map[int][]string)

	for _, line := range lines {
		var re = regexp.MustCompile("[,.?! ~\n]")
		lineSub := re.ReplaceAllString(line, "")
		// println(lineSub)
		_, exists := lenMap[len(lineSub)]
		if !exists {
			lenMap[len(lineSub)] = []string{line}
		} else {
			lenMap[len(lineSub)] = append(lenMap[len(lineSub)], line)
		}
	}
	return lenMap
}

func FileOpen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	return f
}

func main() {
	wPunct := ToSlice(`C:\Project\20200804.-방송DB후처리\broadcast_w_punct_u.txt`)
	wPunctMap := ToMap(wPunct)
	f, err := os.Open(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.trn`)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	newF, err := os.Create(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.punct.trn`)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	notFound, err := os.Create(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_not_found.trn`)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var line string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line = scanner.Text()
		fileName := strings.Split(line, " :: ")[0]
		orgTxt := strings.Split(line, " :: ")[1]
		orgTxtSub := strings.ReplaceAll(orgTxt, " ", "")

		cmpLen := len(orgTxtSub)
		point := cmpLen + 5

		splitOrgTxt := strings.Split(orgTxt, " ")
		firstWord := splitOrgTxt[0]
		lastWord := splitOrgTxt[len(splitOrgTxt)-1]

		flag := true

		for flag {
			cmpList, exists := wPunctMap[cmpLen]
			if !exists || cmpLen > point {

				flag = false
				break
			}
			var re = regexp.MustCompile("[,.?! ~\n]")
			for _, cmpTxt := range cmpList {
				cmpTxtSub := re.ReplaceAllString(cmpTxt, "")
				if strings.Contains(cmpTxtSub, orgTxtSub) {
					start := strings.Index(cmpTxt, firstWord)
					check := strings.Index(cmpTxt[start+len(firstWord):], lastWord)
					check = check + start + len(firstWord)
					end := strings.Index(cmpTxt[check+len(lastWord):], " ")
					if end == -1 {

					} else {
						end = end + check + len(lastWord)
					}
					flag = false
					break
				}
			}
			cmpLen++
		}
	}

	// print(f, newF, notFound)
	// for key, _ := range wPunctMap {
	// 	println(key)
	// }
	// print(wPunctMap[53][0])
	// print(wPunctMap[40][0])
	// print(wPunctMap[30])
}
