# go-talib
A Go(lang) wrapper for TA-Lib(Techinal Analysis Library) which is often used for stock/financial analysis.

http://ta-lib.org/

[![Build Status](https://travis-ci.org/abulo/talib.svg?branch=master)](https://travis-ci.org/abulo/talib)
[![GoDoc](https://godoc.org/github.com/abulo/talib?status.svg)](https://godoc.org/github.com/abulo/talib)

To use the library you need TA-Lib installed.


===

每个函数增加两个返回值，分别代表ta-lib的outBegIdx和outNBElement参数，以便处理需要ignore的值

===

## Example
```go
package main

import (
	"fmt"
	"math"

	"github.com/abulo/talib"
)

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
}
```

## Installing

To compile first download [ta-lib-0.6.4-src.tar.gz](https://github.com/TA-Lib/ta-lib/releases/download/v0.6.4/ta-lib-0.6.4-src.tar.gz) and:
```
./configure --prefix=/usr/local LDFLAGS="-lm" && \
make -j && \
make install && \
ldconfig && \
cd /usr/local/lib && \
ln -s libta-lib.so libta_lib.so
```