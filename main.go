package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)



var domain = "s"
var folder_path = "00f38d96-b252-4d12-961a-ff8468323815"

type FileType struct {
	FileName string
	Date     time.Time
}

func GetDir() []FileType {
	entries, err := os.ReadDir(folder_path)
	if err != nil {
		log.Fatal(err)
	}

	var listOfFiles = make([]FileType, 0)

	for _, e := range entries {
		//layout := "20060102150405000"
		tN := e.Name()[37:51]
		fmt.Println(tN)
		t, err := time.Parse("20060102150405", tN)
		if err != nil {
			fmt.Println(err)
		}
		listOfFiles = append(listOfFiles, FileType{FileName: e.Name(), Date: t})
		// fmt.Println(e.Name()[37:54])
	}
	sort.Slice(listOfFiles, func(i, j int) bool { return listOfFiles[i].Date.Before(listOfFiles[j].Date) })
	return listOfFiles
}
func PostToServer(fileName FileType) {

	now := time.Now()
	formatted2 := now.Format("2006-01-02T15:04:05")

	url := "http://" + domain + "/" + formatted2 + "/4"
	fmt.Println("URL:>", url)

}

func main() {
	fmt.Println(domain)
	listOfFils := GetDir()
	ticker := time.NewTicker(4 * time.Second)

	var i = 0
	go func() {

		for range ticker.C {
			if i > len(listOfFils)-1 {
				i = 0
			}

			//fmt.Println(listOfFils[i].Date)
			PostToServer(listOfFils[i])
			i++
		}

	}()

	time.Sleep(100 * time.Second)
	ticker.Stop()

}
