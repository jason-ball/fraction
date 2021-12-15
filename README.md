# fractions [![Go Reference](https://pkg.go.dev/badge/github.com/jason-ball/fraction.svg)](https://pkg.go.dev/github.com/jason-ball/fraction)

A simple fraction library for Go. Created to teach myself Go modules.

## How to use:

```go
import "github.com/jason-ball/fractions"

func main() {
    // Create two new fractions
    a := fraction.Fraction{3, 6}
    b := fraction.Fraction{25, 100}

    // Add them
    result := fraction.Add(a, b)

    // Result is automatically simplifed and formatted
    println(result.String())
    // Prints: 3/4
}
```

## How to run tests:

```bash
go test
```

## License:

```
MIT License

Copyright (c) 2021 Jason Ball

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```