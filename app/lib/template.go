package lib

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"strings"
)

const VIEW_PATH = "view"

func TemplatePath() string {
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		return "system/error.html"
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		return "system/error.html"
	}

	// ambil nama file
	file, _ := caller.FileLine(fpcs[0] - 1)
	file_arr := strings.Split(file, "/")
	file = file_arr[len(file_arr)-1]
	file = strings.Replace(file, ".go", "", -1)

	// ambil nama method
	fname := caller.Name()
	f := strings.Split(fname, ".")
	return file + "/" + strings.ToLower(f[len(f)-1]) + ".html"
} //end func

/** FOR SINGLE TEMPLATE FILE **/
func TemplateRender(w http.ResponseWriter, templatePath string, templateData map[string]interface{}) {
	var tpl *template.Template
	var err error

	tpl, err = template.ParseFiles(VIEW_PATH + "/" + templatePath)

	if err != nil {
		fmt.Fprintf(w, "%v\n", err.Error())
		return
	}

	tpl.Execute(w, templateData)
} //end func

/** FOR MULTI TEMPLATE FILES **/
func TemplateRenders(w http.ResponseWriter, templatePath string, templateData map[string]interface{}, vbase string, isComplete bool) {
	var tpl *template.Template
	var err error

	if isComplete == true {
		tpl, err = template.ParseFiles(VIEW_PATH+"/"+vbase+".html", VIEW_PATH+"/"+vbase+"_header.html", VIEW_PATH+"/"+vbase+"_footer.html", VIEW_PATH+"/"+templatePath)
	} else {
		tpl, err = template.ParseFiles(VIEW_PATH+"/"+vbase+".html", VIEW_PATH+"/"+templatePath)
	}

	if err != nil {
		fmt.Fprintf(w, "%v\n", err.Error())
		return
	}

	tpl.ExecuteTemplate(w, "layout", templateData)
} //end func
