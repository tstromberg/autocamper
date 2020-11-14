// The "main" package contains the command-line utility functions.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"text/template"

	"github.com/tstromberg/campwiz/pkg/mixer"
	"github.com/tstromberg/campwiz/pkg/engine"
	"k8s.io/klog/v2"
)

type formValues struct {
	Dates    string
	Nights   int
	Distance int
	Standard bool
	Group    bool
	WalkIn   bool
	BoatIn   bool
}

type TemplateContext struct {
	Query engine.Query
	Results  []engine.Result
	Form     formValues
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Incoming request: %+v", r)
	klog.Infof("Incoming request: %+v", r)
	crit := engine.Query{}
	results, err := engine.Search(crit)
	klog.V(1).Infof("RESULTS: %+v", results)
	if err != nil {
		klog.Errorf("Search error: %s", err)
	}

	outTmpl, err := data.Read("http.tmpl")
	if err != nil {
		klog.Errorf("Failed to read template: %v", err)
	}
	tmpl := template.Must(template.New("http").Parse(string(outTmpl)))
	ctx := TemplateContext{
		Query: crit,
		Results:  results,
		Form: formValues{
			Dates: "2018-09-20",
		},
	}
	err = tmpl.ExecuteTemplate(w, "http", ctx)
	if err != nil {
		klog.Errorf("template: %v", err)
	}
}

func init() {
	flag.Parse()
}

func main() {
	http.HandleFunc("/", handler)
	klog.Fatal(http.ListenAndServe(":8080", nil))
}
