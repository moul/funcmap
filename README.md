# funcmap

[![GuardRails badge](https://badges.production.guardrails.io/moul/funcmap.svg)](https://www.guardrails.io)

:neckbeard: funcmap for go text/template

This repo extends the builtin helpers of Golang's [text/template](https://golang.org/pkg/text/template) library with new helpers.

## Usage

```go
import "github.com/moul/funcmap"

tmpl, _ := template.New("").Funcs(FuncMap).Parse(`Hello {{"woRld" | rev}} !`)
tmpl.Execute(os.Stdout, nil)
// Output: Hello dlRow !
```

## Helpers

* `json`
* `prettyjson`
* `indent`
* `split`
* `join`
* `title`
* `unexport`
* `add`
* `trimspace`
* `lower`
* `upper`
* `rev`
* `int`
