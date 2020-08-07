package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func MakeSubMap(entireMap map[int][]string, size int) map[int][]string {
	var subMap map[int][]string
	subMap = make(map[int][]string)
	for k, v := range entireMap {
		if k >= size {
			subMap[k] = v
		}
	}
	return subMap
}

func SearchKeyIndex(keys []int, size int) int {
	for i, key := range keys {
		if key >= size {
			return i
		}
	}
	return len(keys) - 1
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

func run(fileNumber int) {
	wPunct := ToSlice(`C:\Project\20200804.-방송DB후처리\broadcast_2017_01_03.captions_u.txt`)
	wPunctMap := ToMap(wPunct)
	keys := make([]int, 0, len(wPunctMap))
	for k := range wPunctMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
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
	f, err := os.Open(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list_` + strconv.Itoa(fileNumber) + `.trn`)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	newF, err := os.OpenFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_01_03_pcm.list.punct_`+strconv.Itoa(fileNumber)+`.trn`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	notFound, err := os.OpenFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_not_found_`+strconv.Itoa(fileNumber)+`.trn`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	errFound, err := os.OpenFile(`C:\Project\20200804.-방송DB후처리\SubtTV_2017_err_found_`+strconv.Itoa(fileNumber)+`.trn`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var line string
	scanner := bufio.NewScanner(f)
	// i := 0
	var re = regexp.MustCompile("[,.?! ~\n]")
	fmt.Println("start")
Loop:
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
		keyIndex := SearchKeyIndex(keys, cmpLen)
		// point := cmpLen + 20
		// subMap := MakeSubMap(wPunctMap, cmpLen)

		splitOrgTxt := strings.Split(orgTxt, " ")
		firstWord := splitOrgTxt[0]
		lastWord := splitOrgTxt[len(splitOrgTxt)-1]
		wordCount := len(splitOrgTxt)
		defer func() {
			recover()
			s := line
			errFound.WriteString(s)
		}()
		for _, mapIndex := range keys[keyIndex:] {
			cmpList := wPunctMap[mapIndex]
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
					}
					goto Loop
				}
			}
		}
		s := line + "\n"
		notFound.WriteString(s)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(7)
	go run(1)
	go run(2)
	go run(3)
	go run(4)
	go run(5)
	go run(6)
	go run(7)
	go run(8)
	wg.Wait()
}
