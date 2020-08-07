package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

func max(numbers map[int][]string) int {
	var maxNumber int
	for maxNumber = range numbers {
		break
	}
	for n := range numbers {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

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
	wPunct := ToSlice(`C:\Project\20200804.-방송DB후처리\broadcast_2017_01_03.captions_u.txt`)
	wPunctMap := ToMap(wPunct)
	// keys := make([]int, 0, len(wPunctMap))
	// for k := range wPunctMap {
	// 	keys = append(keys, k)
	// }
	// sort.Ints(keys)
	// fmt.Print(keys)
	// sort.Sort(sort.IntSlice(keys))
	// fmt.Println("max:", max(wPunctMap))
	// for i, k := range keys {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if keys[i-1]+1 != k {
	// 		fmt.Print(k, " , ")
	// 	}
	// }
	// fmt.Print(wPunctMap[8649])
	f, err := os.Open(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.trn`)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	newF, err := os.OpenFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.punct.trn`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	notFound, err := os.OpenFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_not_found.trn`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var line string
	scanner := bufio.NewScanner(f)
	// i := 0
	fmt.Print("start")
	for scanner.Scan() {
		// i++
		// if i > 10 {
		// 	break
		// }
		// startTime := time.Now()

		line = scanner.Text()
		fileName := strings.Split(line, " :: ")[0]
		orgTxt := strings.Split(line, " :: ")[1]
		orgTxtSub := strings.ReplaceAll(orgTxt, " ", "")

		cmpLen := len(orgTxtSub)
		// point := cmpLen + 20

		splitOrgTxt := strings.Split(orgTxt, " ")
		firstWord := splitOrgTxt[0]
		lastWord := splitOrgTxt[len(splitOrgTxt)-1]
		wordCount := len(splitOrgTxt)

		flag := true

		for flag {
			cmpList, exists := wPunctMap[cmpLen]
			if !exists {
				if cmpLen > 1000 {
					s := line + "\n"
					notFound.WriteString(s)
					flag = false
					break
				}
				cmpLen++
				continue
			}
			var re = regexp.MustCompile("[,.?! ~\n]")
			for _, cmpTxt := range cmpList {
				cmpTxtSub := re.ReplaceAllString(cmpTxt, "")
				if strings.Contains(cmpTxtSub, orgTxtSub) {
					if wordCount == 1 {
						start := strings.Index(cmpTxt, firstWord)
						end := strings.Index(cmpTxt[start+len(firstWord):], " ")
						if end == -1 {
							s := fileName + " :: " + cmpTxt[start:] + "\n"
							newF.WriteString(s)
							// ioutil.WriteFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.punct.trn`, []byte(s), 0644)
						} else {
							end = end + start + len(firstWord)
							s := fileName + " :: " + cmpTxt[start:end] + "\n"
							newF.WriteString(s)
							// ioutil.WriteFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.punct.trn`, []byte(s), 0644)
						}
					} else {
						start := strings.Index(cmpTxt, firstWord)
						check := strings.Index(cmpTxt[start+len(firstWord):], lastWord)
						var end int

						check = check + start + len(firstWord)
						end = strings.Index(cmpTxt[check+len(lastWord):], " ")

						if end == -1 {
							s := fileName + " :: " + cmpTxt[start:] + "\n"
							newF.WriteString(s)
							// ioutil.WriteFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.punct.trn`, []byte(s), 0644)
						} else {
							end = end + check + len(lastWord)
							s := fileName + " :: " + cmpTxt[start:end] + "\n"
							newF.WriteString(s)
							// ioutil.WriteFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.punct.trn`, []byte(s), 0644)
						}
						// elapsedTime := time.Since(startTime)
						// fmt.Printf("done: %s\n", elapsedTime)
					}
					flag = false
					break
				}
			}
			cmpLen++
		}
	}
}
