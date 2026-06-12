# CSMT

### About library

Counter Strike Marketplace Tools - go library for job with CS marketplace. Currently, the main work is on improving the technical component of the library and implementing unfinished API methods.

### Quick start

Firstly, you need to install library into your project

```bash
go get github.com/cyrillemad/csmt
```

After that, you need to select the API you'll be working with and create a suitable client. For free work without authorization, we'll use `NoAuthorizeClient`. This is a convenient set of all APIs from the library that don't require authorization.

```go
package main

import (
	"github.com/cyrillemad/csmt"
)

func main() {
	c := csmt.NewNoAuthClient()
}
```

Now let's get `MarketHash`. This type is recommended to be filled only in the results of other API methods, such as `SearchHash`, but nothing prevents you from simply substituting your string as this type. Once we have this universal type, we can substitute it in any methods that work with `MarketHash`, regardless of the API.

```go
package main

import (
	"fmt"

	"github.com/cyrillemad/csmt"
)

func main() {
	c := csmt.NewNoAuthClient()

	hash, err := c.Community.SearchHash("DragonLore")
	if err != nil {
		fmt.Println(err)
	}

	price, err := c.Community.PriceOverview(hash)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash, price.MedianPrice)
}
```

Let's check it out!

```
Sticker Slab | Dragon Lore (Foil) 781.66
Process finished with the exit code 0
```

So, we have a short program, soon there will be a link to the full documentation here.

### Special thanks

Many thanks to all users who contribute to the development, use the library, or simply explore it!
