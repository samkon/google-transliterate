# Samkon Google Transliterate

This source is got powered to enable transliteration over the forked library. Thanks to `Conight`. Useful for speacially arabic transliteration.

## Download from Github
```shell script
GO111MODULE=on go get github.com/samkon/google-transliterate
```

## Quick Start Examples

### Simple translate
```go
package main

import (
	"fmt"
	"github.com/samkon/google-transliterate"
)

func main() {
	t := translator.New()
	result, err := t.Transliterate("كامل الحارونى متفرع من احمد فخرى امام حديقه الطفل مدينه نصر القاهره الدور السابع شقه ١٦", "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
```

### You can look into forked library for details

https://github.com/Conight/go-googletrans

## License
This SDK is distributed under the [The MIT License](https://opensource.org/licenses/MIT), see [LICENSE](./LICENSE) for more information.
