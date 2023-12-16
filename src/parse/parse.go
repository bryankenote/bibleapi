package parse

import (
	"bibleapi/src/codegen/sqlc"
	"bibleapi/src/db"
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Verse struct {
	Book    string
	Chapter int
	Verse   int
	Content string
}

func ParseBibleFile(fileName string, processVerse func(verse Verse)) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error openning file: %s\n", err.Error())
		return
	}
	defer file.Close()

	referenceRE := regexp.MustCompile("([1-3] )*[a-zA-Z]+ [0-9]+:[0-9]+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		referenceValueArr := referenceRE.FindAllString(line, -1)
		if len(referenceValueArr) != 0 {
			splitReference := strings.Split(referenceValueArr[0], " ")

			var book string
			var splitNumbers []string

			if slices.Contains([]string{"1", "2", "3"}, splitReference[0]) {
				book = strings.Join(splitReference[0:2], " ")
				splitNumbers = strings.Split(splitReference[2], ":")
			} else {
				book = splitReference[0]
				splitNumbers = strings.Split(splitReference[1], ":")
			}

			chapterNumber, err := strconv.Atoi(splitNumbers[0])
			if err != nil {
				fmt.Printf("Error casting string to number: %s\n", err.Error())
				return
			}
			verseNumber, err := strconv.Atoi(splitNumbers[1])
			if err != nil {
				fmt.Printf("Error casting string to number: %s\n", err.Error())
				return
			}
			referenceString := fmt.Sprintf("%s %d:%d\n", book, chapterNumber, verseNumber)
			verseContent := strings.Trim(strings.Trim(line[len(referenceString):], " "), "\t")

			verse := Verse{Book: book, Chapter: chapterNumber, Verse: verseNumber, Content: verseContent}
			processVerse(verse)
		}
	}
}

func ImportBibleFile(ctx context.Context) {
	currentBook := ""
	ParseBibleFile("../bsb/bible.txt", func(verse Verse) {
		if currentBook != verse.Book {
			currentBook = verse.Book
			fmt.Printf("Importing %s ...\n", currentBook)
		}
		db.Instance.CreateVerse(ctx, sqlc.CreateVerseParams{
			Translation: "BSB",
			Book:        verse.Book,
			Chapter:     int64(verse.Chapter),
			Verse:       int64(verse.Verse),
			Content:     verse.Content,
		})
	})
	fmt.Println("Done.")
}
