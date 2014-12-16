# Pine.js Go Library

This is a simple Go library for interacting with [pine.js][pine].

## Usage

```go
var data foo

if err := pinejs.Get(&data); err != nil {
	log.Fatalln(err)
}
```

[pine]:https://bitbucket.org/rulemotion/pinejs/overview
