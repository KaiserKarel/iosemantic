# iosemantic

[![GoDoc](https://godoc.org/godoc.org/github.com/KaiserKarel/iosemantic?status.svg)](https://godoc.org/github.com/KaiserKarel/iosemantic)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaiserkarel/iosemantic)](https://goreportcard.com/report/github.com/kaiserkarel/iosemantic)

A testing library containing helper function to verify that io.Readers and io.Writers implement their respective specifications.

An example usage is [DRFS](https://github.com/kaiserkarel/drfs), where I implemented a file abstraction and use this library to ensure git remote add origin git@github.com:KaiserKarel/iosemantic.git correctness.

## Example

```go 
func TestMyCustomFileBackendSemantics(t *testing.T) {
    var file = NewCustomFileBackend()
    iosemantic.ImplementsReader(t, file)
    iosemantic.ImplementsWriter(t, file)
    iosemantic.ImplementsWriterAt(t, file)
}
```

## Caveats

`iosemantic` only verifies that the interfaces match their specifications, not that the input and output buffers remain
consistent. You will still need to write tests to verify your business logic.

## Stability

The current API will remain consistent. Functions accepting respective option structs may be expanded on by adding options to these structs, where unset fields are set to sane defaults.


