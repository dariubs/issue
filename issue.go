package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dariubs/s2i"
)

func main() {
	f := os.Args[1]

	if f == "init" {
		initIssue()
	} else {
		checkInit()
	}

	switch f {
	case "add":
		addIssue()
	case "fix":
		id := s2i.ParseInt(os.Args[2], -1)
		fixIssue(id)
	case "show":
		id := s2i.ParseInt(os.Args[2], -1)
		showIssue(id)
	case "list":
		listIssue()
	}
}

func getuserinput(label string) (string, error) {
	fmt.Print(label)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return "", err
	}
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	return text, nil
}

func getindex() int {
	rawindex, err := ioutil.ReadFile(".issue/index/last")
	if err != nil {
		log.Println(err)
		return 0
	}

	index, err := strconv.ParseInt(string(rawindex), 10, 64)
	if err != nil {
		log.Println(err)
	}
	return int(index)
}

func incrindex() error {
	index := getindex() + 1

	f, err := os.OpenFile(".issue/index/last",
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d", index))

	return nil
}

func initIssue() {
	err := os.Mkdir(".issue", 0700)
	if err != nil {
		log.Println(err)
		return
	}

	err = os.MkdirAll(".issue/index", 0700)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.Create(".issue/index/last")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	f.WriteString("1")

	l, err := os.Create(".issue/index/list")
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("issue inited")
}

func addIssue() {
	var content string
	var title string

	text, err := getuserinput("title: ")
	content += ":title:\n"
	content += text
	content += "\n\n"
	title = text

	text, err = getuserinput("description: ")
	content += ":description:\n"
	content += text
	content += "\n\n"

	index := getindex()

	f, err := os.Create(fmt.Sprintf(".issue/%d", index))
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	f.WriteString(content)

	l, err := os.OpenFile(".issue/index/list",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	l.WriteString(fmt.Sprintf("%d - %s\n", index, title))

	incrindex()

	fmt.Println("issue added")
}

func fixIssue(id int) {
	os.Rename(fmt.Sprintf(".issue/%d", id), fmt.Sprintf(".issue/%d.done", id))
}

func showIssue(id int) {
	rawissue, err := ioutil.ReadFile(fmt.Sprintf(".issue/%d", id))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(rawissue))
}

func checkInit() {
	if _, err := os.Stat(".issue/"); os.IsNotExist(err) {
		fmt.Println("Issue not inited here")
		os.Exit(0)
	}
}

func listIssue() {
	rawlist, err := ioutil.ReadFile(".issue/index/list")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(rawlist))
}
