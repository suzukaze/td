package main

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"regexp"
	"strconv"
	"strings"
)

type Todo struct {
	Id       int64  `json:"id"`
	Desc     string `json:"desc"`
	Status   string `json:"status"`
	Modified string `json:"modified"`
}

func (t *Todo) MakeOutput() {
	var symbole string
	var color ct.Color

	if t.Status == "done" {
		color = ct.Green
		symbole = "✓"
	} else {
		color = ct.Red
		symbole = "✕"
	}

	hashtagReg := regexp.MustCompile("#[^\\s]*")

	spaceCount := 6 - len(strconv.FormatInt(t.Id, 10))

	fmt.Print(strings.Repeat(" ", spaceCount)," ", t.Id, " | ")
	ct.ChangeColor(color, false, ct.None, false)
	fmt.Print(symbole)
	ct.ResetColor()
	fmt.Print(" ")
	pos := 0
	for _, token := range hashtagReg.FindAllStringIndex(t.Desc, -1) {
		fmt.Print(t.Desc[pos:token[0]])
		ct.ChangeColor(ct.Yellow, false, ct.None, false)
		fmt.Print(t.Desc[token[0]:token[1]])
		ct.ResetColor()
		pos = token[1]
	}
	fmt.Println(t.Desc[pos:])
}
