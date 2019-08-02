package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dariubs/s2i"
)

func main() {
	f := os.Args[1]
	fmt.Println("command: ", f)

	switch f {
	case "init":
		initIssue()
	case "add":
		addIssue()
	case "fix":
		id := s2i.ParseInt(os.Args[2], -1)
		fixIssue(id)
	}
}

func getuserinput(label string) (string, error) {
	fmt.Print(label)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	return text, nil
}

func getindex() int {
	rawindex, err := ioutil.ReadFile(".issue/index/last")
	if err != nil {
		return 0
	}
	index := s2i.ParseInt(string(rawindex), 0)
	return index
}

func incrindex() error {
	index := getindex()
	index++
	f, err := os.Open(".issue/index/last")
	if err != nil {

	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d", index))

	return nil
}

func initIssue() {
	err := os.Mkdir(".issue", 0700)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll(".issue/index", 0700)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(".issue/index/last")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString("1\n")
	fmt.Println("issue inited")
}

func addIssue() {
	var content string

	text, err := getuserinput("title: ")
	content += ":title:\n"
	content += text
	content += "\n\n"

	text, err = getuserinput("description: ")
	content += ":description:\n"
	content += text
	content += "\n\n"

	index := getindex()

	f, err := os.Create(fmt.Sprintf(".issue/%d", index))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.WriteString(content)

	incrindex()

	fmt.Println("issue added")
}

func fixIssue(id int) {
	os.Rename(fmt.Sprintf(".issue/%d", id), fmt.Sprintf(".issue/%d.done", id))
}
