module github.com/eukarya-inc/plateau-spec/cmd

go 1.24.0

require (
	github.com/eukarya-inc/plateau-spec/plateaudocsearch v0.0.0
	github.com/stretchr/testify v1.11.1
)

replace github.com/eukarya-inc/plateau-spec/plateaudocsearch => ../plateaudocsearch

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/reearth/mokuji/go v0.2.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
