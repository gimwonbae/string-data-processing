package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"unicode/utf8"
)

func toSlice(path string) []string {
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

func viewList(line string, path string, fileType string) []string {
	orgPath := strings.Split(line, " :: ")[0]
	forFileName := strings.Split(orgPath, "/")
	fileName := forFileName[len(forFileName)-2]
	if fileType == "text" {
		return toSlice(filepath.Join("..", path, fileName, "data", fileName+".text"))
	} else if fileType == "tok" {
		return toSlice(filepath.Join("..", path, fileName, "data", fileName+".tok"))
	} else {
		print("type = text or tok")
		return nil
	}
}

func match(path string, dirPath string) {
	var b bytes.Buffer
	wNum, _ := os.OpenFile(filepath.Join(`..`, `w_num`), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	woNum, _ := os.OpenFile(filepath.Join(`..`, `wo_num`), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fail, _ := os.OpenFile(filepath.Join(`..`, `fail`), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	lines := toSlice(filepath.Join(`..`, path))
	var re = regexp.MustCompile("[^가-힣0-9a-zA-Z]")
	var re2 = regexp.MustCompile("[0-9]")
	for _, line := range lines {
		textList := viewList(line, dirPath, "text")
		tokList := viewList(line, dirPath, "tok")

		fileName := strings.Split(line, " :: ")[0]
		lineText := strings.Split(line, " :: ")[1]
		lineSub := strings.ToLower(re.ReplaceAllString(lineText, ""))

		index := 0
		for _, tok := range tokList {
			tokSub := strings.ToLower(re.ReplaceAllString(tok, ""))
			if strings.Contains(tokSub, lineSub) {
				break
			}
			index++
		}
		if index >= len(textList) {
			b.WriteString(line)
			b.WriteString("\n")
			fail.WriteString(b.String())
			b.Reset()
			// fmt.Println("index")
			continue
		}
		text := textList[index]

		lineSplit := strings.Split(lineText, " ")
		firstWord := lineSplit[0]
		lastWord := lineSplit[len(lineSplit)-1]
		start := strings.Index(text, firstWord)
		if start == -1 {
			b.WriteString(line)
			b.WriteString("\n")
			fail.WriteString(b.String())
			b.Reset()
			continue
		}
		if len(lineSplit) == 1 {
			end := strings.Index(text[start+len(firstWord):], " ")
			b.WriteString(fileName)
			b.WriteString(" :: ")
			if end == -1 {
				b.WriteString(text[start:])
				b.WriteString("\n")
				matched := re2.MatchString(text[start:])
				if matched {
					wNum.WriteString(b.String())
				} else {
					woNum.WriteString(b.String())
				}
				continue
			} else {
				end = end + start + len(lastWord)
				b.WriteString(text[start:end])
				b.WriteString("\n")
				matched := re2.MatchString(text[start:end])
				if matched {
					wNum.WriteString(b.String())
				} else {
					woNum.WriteString(b.String())
				}
				continue
			}
		}
		check := strings.Index(text[start+len(firstWord):], lastWord)
		if check == -1 {
			b.WriteString(line)
			b.WriteString("\n")
			fail.WriteString(b.String())
			b.Reset()
			continue
		}
		check = check + start + len(firstWord)
		end := strings.Index(text[check+len(lastWord):], " ")

		b.WriteString(fileName)
		b.WriteString(" :: ")
		if end == -1 {
			b.WriteString(text[start:])
			b.WriteString("\n")
			matched := re2.MatchString(text[start:])
			if matched {
				wNum.WriteString(b.String())
			} else {
				woNum.WriteString(b.String())
			}
		} else {
			end = end + check + len(lastWord)
			b.WriteString(text[start:end])
			b.WriteString("\n")
			matched := re2.MatchString(text[start:])
			if matched {
				wNum.WriteString(b.String())
			} else {
				woNum.WriteString(b.String())
			}
		}
		b.Reset()
	}
}

func woNum(checkInput string, orgInput string) {
	var b bytes.Buffer
	checkList := toSlice(filepath.Join(`..`, checkInput))
	orgList := toSlice(filepath.Join(`..`, orgInput))

	match, _ := os.OpenFile(filepath.Join(`..`, `match_wo_num`), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	miss, _ := os.OpenFile(filepath.Join(`..`, `miss_wo_num`), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

func wNum(checkInput string, orgInput string) {
	var b bytes.Buffer
	checkList := toSlice(filepath.Join(`..`, checkInput))
	orgList := toSlice(filepath.Join(`..`, orgInput))

	match, _ := os.OpenFile(filepath.Join(`..`, `match_w_num`), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	miss, _ := os.OpenFile(filepath.Join(`..`, `miss_w_num`), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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
				// fmt.Println(checkSub)
				// fmt.Println(orgSub)
				// fmt.Println(strings.ToLower(checkSub))
				// fmt.Println(strings.ToLower(orgSub))
				// fmt.Println(utf8.RuneCountInString(strings.ToLower(checkSub)))
				// fmt.Println(utf8.RuneCountInString(strings.ToLower(orgSub)))
				condition := utf8.RuneCountInString(strings.ToLower(checkSub)) - utf8.RuneCountInString(strings.ToLower(orgSub))
				// fmt.Println(condition)
				if -2 <= condition && condition <= 2 {
					match.WriteString(b.String())
				} else {
					miss.WriteString(b.String())
				}
				b.Reset()
				orgList = orgList[index+1:]
				break
			}
		}
	}
}

func main() {
	source := flag.String("source", "", "Source File")
	ref := flag.String("ref", "", "Reference File")
	goal := flag.String("goal", "", "Matching or Checking")

	flag.Parse()
	if *goal == "matching" {
		match(*source, *ref)
	} else if *goal == "checking" {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			woNum(`wo_num`, *source)
		}()
		go func() {
			defer wg.Done()
			wNum(`w_num`, *source)
		}()
		wg.Wait()
	} else {
		fmt.Println("Wrong Usage")
	}
	// match("SubtTV_2017_01_03_pcm.list.trn", `broadcast_text\KOR`)
	// woNum(`wo_num`, `SubtTV_2017_01_03_pcm.list.trn`)
	// wNum(`w_num`, `SubtTV_2017_01_03_pcm.list.trn`)
}
