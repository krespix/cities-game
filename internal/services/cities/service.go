package cities

import (
	"bufio"
	"fmt"
	"os"
)

type Service interface {
	InitCitiesMap()
}

type service struct {
}

func readCsvFile(filePath string) string {
	f, err := os.Open("cities.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return ""
}

func (s *service) InitCitiesMap() {
	records := readCsvFile("./city.csv")
	fmt.Println(records)
}

func New() Service {
	return &service{}
}
