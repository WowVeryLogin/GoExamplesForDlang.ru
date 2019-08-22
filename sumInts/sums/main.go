package main

import (
	"log"
	"os"
	"text/template"
)

var t = template.New("SumFunc")
var tmpl = template.Must(t.Parse(`//Package sums is generated. Do not edit it.
package main

func GeneratedSumInts(a interface{}) interface{} {
	switch v := a.(type) {
{{range $}}	case []{{.}}:
		var res {{.}}
		for i := range v {
			res += v[i]
		}
		return res
{{end}}	default:
		return nil
	}
}
`))

func main() {
	f, err := os.Create("sums.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	tmpl.Execute(f, []string{"int", "int16", "int32", "int64"})
}
