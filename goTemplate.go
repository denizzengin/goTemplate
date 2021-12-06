package template

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
)

// ErrTemplateCodeNotFound 
var ErrTemplateCodeNotFound = errors.New("template code cannot found")
const (
	EmailTemplate string = "emailTemplate.html"
)

var EmailTemplates map[string]*template.Template

type WrappedOrder struct {
	Description string 
	Orders []Order
}

type Order struct {
	OrderId      int
	OrderDesc    string
	CustomerId   int
	CustomerName string
}

func ParseTemplate(templateCode string, obj interface{}) string {
	t, ok := EmailTemplates[templateCode] 
	if !ok {
		checkErr(ErrTemplateCodeNotFound)
	}
	buffer := bytes.NewBuffer([]byte(""))
	err := t.Execute(buffer, obj)
	checkErr(err)	
	// Speed up performance	
	return buffer.String()
}

func GetAllTemplate() {
	EmailTemplates = make(map[string]*template.Template)
	commonTemplate, err := template.ParseFiles("./" + EmailTemplate)
	checkErr(err)
	for _, v := range commonTemplate.Templates() {
		EmailTemplates[v.Name()] = v
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("unknown error ", err.Error())
		os.Exit(1)
	}
}
