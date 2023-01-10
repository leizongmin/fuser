[![Go](https://github.com/leizongmin/fuser/actions/workflows/go.yml/badge.svg)](https://github.com/leizongmin/fuser/actions/workflows/go.yml)
[![Node.js CI](https://github.com/leizongmin/fuser/actions/workflows/node.js.yml/badge.svg)](https://github.com/leizongmin/fuser/actions/workflows/node.js.yml)
![NPM](https://img.shields.io/npm/l/@leizm/fuser)
![npm (scoped)](https://img.shields.io/npm/v/@leizm/fuser)

# fuser

Find the Process That is Using a File in Linux

## Go version

### Installation

```bash
go get github.com/leizongmin/fuser
```

#### Build open files and pid map data

```go
package main

import (
	"fmt"

	"github.com/leizongmin/fuser"
)

func main() {
	data, err := fuser.BuildMap(nil)
	if err != nil {
		panic(err)
	}
	for k, v := range data {
		fmt.Println(k, v)
	}
}
```

Outputs like this:

```
socket:[90732] [3004]
pipe:[99713] [4117 4117]
/home/codespace/.vscode-remote/data/logs/20230110T093419/exthost1/output_logging_20230110T093428/12-LLDB.log [3004 31348 31405 4619 4629 6359 7158]
pipe:[732477] [31405 31405]
...
```

#### Gets a list of Pids for which the file is currently being opened

```go
err := fuser.Update(nil)
if err != nil {
   panic(err)
}
fmt.Println(fuser.GetPath("/dev/null2"))
```

Outputs like this:

```
[1 2542 2568 2577 2590 2962 3004 31348 31405 3523 3565 4117 4251 4257 4270 4309 4619 4629 5215 5237 7 7158]
```

or:

```
[]
```

#### Filter file path when buildMap or update

```go
fuser.Update(&fuser.Options{
   Filter: func(s string) bool {
      return true
   },
})

fuser.Update(&fuser.Options{
   Filter: func(s string) bool {
      return false
   },
})
```

## Node.js version

### Installation

```bash
npm i @leizm/fuser -S
```

### Usage

#### Build open files and pid map data

```js
const fuser = require("@leizm/fuser");

const data = await fuser.buildMap();
console.log(data);
```

Outputs like this:

```
{
 '/dev/null': [
       1, 11547,  1163,  1163,
    1181, 11856, 12283, 12290,
   12298, 12431, 13220, 13245,
   21131, 22466, 22525, 23318,
   31545, 31546, 31563,  3196,
    3222,  6370,     7
 ],
 'pipe:[104532]': [ 1, 7 ],
 'pipe:[104533]': [ 1, 7 ],
 'pipe:[334229]': [ 10976 ],
 'pipe:[334230]': [ 10976 ],
 'pipe:[334231]': [ 10976 ],
 ...
}
```

#### Gets a list of Pids for which the file is currently being opened

```js
const fuser = require("@leizm/fuser");

// update the cache firstly
await fuser.update();

const pids = fuser.getPath("/dev/null");
console.log(pids);
```

Outputs like this:

```
[ 1, 7 ]
```

or:

```
null
```

#### Filter file path when buildMap or update

```js
await fuser.update({ filter: (path) => true });
await fuser.buildMap({ filter: (path) => true });
```

## License

```
MIT License

Copyright (c) 2023 LEI Zongmin <leizongmin@gmail.com>

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
