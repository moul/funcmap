package funcmap

import (
	"bytes"
	"os"
	"testing"
	"text/template"

	. "github.com/smartystreets/goconvey/convey"
)

func test(templateString string, data interface{}) (string, error) {
	tmpl, err := template.New("").Funcs(FuncMap).Parse(templateString)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	err = tmpl.Execute(buffer, data)
	return buffer.String(), err
}

func TestFuncMap(t *testing.T) {
	Convey("Testing FuncMap", t, func() {
		output, err := test(`Hello {{"world" | title}} !`, nil)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "Hello World !")

		output, err = test(`Hello {{"world" | upper}} !`, nil)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "Hello WORLD !")

		output, err = test(`Hello {{"woRld" | lower}} !`, nil)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "Hello world !")
	})
}

func ExampleFuncMap_title() {
	tmpl, _ := template.New("").Funcs(FuncMap).Parse(`Hello {{"world" | title}} !`)
	tmpl.Execute(os.Stdout, nil)
	// Output: Hello World !
}

func ExampleFuncMap_unexport() {
	tmpl, _ := template.New("").Funcs(FuncMap).Parse(`Hello {{"WORLD" | unexport}} !`)
	tmpl.Execute(os.Stdout, nil)
	// Output: Hello wORLD !
}

func ExampleFuncMap_upper() {
	tmpl, _ := template.New("").Funcs(FuncMap).Parse(`Hello {{"world" | upper}} !`)
	tmpl.Execute(os.Stdout, nil)
	// Output: Hello WORLD !
}

func ExampleFuncMap_lower() {
	tmpl, _ := template.New("").Funcs(FuncMap).Parse(`Hello {{"woRld" | lower}} !`)
	tmpl.Execute(os.Stdout, nil)
	// Output: Hello world !
}

func ExampleFuncMap_rev() {
	tmpl, _ := template.New("").Funcs(FuncMap).Parse(`Hello {{"woRld" | rev}} !`)
	tmpl.Execute(os.Stdout, nil)
	// Output: Hello dlRow !
}

func ExampleFuncMap_indent() {
	tmpl, _ := template.New("").Funcs(FuncMap).Parse(`{{indent "Hello World!" ">>> "}}`)
	tmpl.Execute(os.Stdout, nil)
	// Output: >>> Hello World!
}

func ExampleFuncMap_trimspace() {
	tmpl, _ := template.New("").Funcs(FuncMap).Parse(`Hello {{"      World     " | trimspace}}!`)
	tmpl.Execute(os.Stdout, nil)
	// Output: Hello World!
}
