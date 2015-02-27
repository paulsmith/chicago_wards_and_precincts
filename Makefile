index.html: index.html.tmpl gen-index.go
	go run gen-index.go > $@
