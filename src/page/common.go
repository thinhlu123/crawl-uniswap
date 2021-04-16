package page

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"thinhlu123/crawl-uniswap/src/model"
)

func Render(w http.ResponseWriter, tmpl string, pageVars model.PageVars) {

	tmpl = fmt.Sprintf("public/html/%s", tmpl) // prefix the name passed
	t, err := template.ParseFiles(tmpl)        //parse the template file

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, pageVars) //execute the template and pass in the variables to fill the gaps

	if err != nil { // if there is an error
		log.Print("template executing error: ", err)
	}
}
