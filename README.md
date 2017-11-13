[![GoDoc](https://godoc.org/github.com/uudashr/go-expo?status.svg)](https://godoc.org/github.com/uudashr/go-expo)
# go-expo
Package expo provides expand expression library.

Common usage is on the REST API that provide `expand` key on the URL:

```
https://api.some-blog-service.com/blog?expand=author,publisher.contact
```

## Installation

```
go get github.com/uudashr/go-expo
```

## Example
```go
opts := expo.ExpandOptions{"author", "publisher.contact"}

expand := opts.Expand("author")
fmt.Println("Expand author:", expand)

expand = opts.Expand("category")
fmt.Println("Expand category:", expand)

expand = opts.Expand("publisher")
fmt.Println("Expand publisher:", expand)

pubOpts := opts.Sub("publisher")
expand = pubOpts.Expand("contact")
fmt.Println("Expand publisher.contact:", expand)

// Output:
// Expand author: true
// Expand category: false
// Expand publisher: true
// Expand publisher.contact: true
```
