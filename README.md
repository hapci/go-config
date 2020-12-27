<pre>
                                            ____________         
    _______ ______       ______________________  __/__(_)______ _
    __  __ `/  __ \_______  ___/  __ \_  __ \_  /_ __  /__  __ `/
    _  /_/ // /_/ //_____/ /__ / /_/ /  / / /  __/ _  / _  /_/ / 
    _\__, / \____/       \___/ \____//_/ /_//_/    /_/  _\__, /  
    /____/                                              /____/
</pre>

__Simple and idiomatic parsing of configuration files and environment variables.__

### Installation

```
$ go get github.com/hapci/go-config
```

### Getting started

#### Unmarshal files

example.yml

```yaml
string: "text"
map:
  key1:
    array: [ "item10", "item20", "item30" ]
    int: 40
    float: 50.6
    bool: true
```

code:

```go
package main

import (
	"fmt"

	"github.com/hapci/go-config"
)

func main() {
	cfg := struct {
		String string `mapstructure:"string"`
		Map    map[string]struct {
			Array []string `mapstructure:"array"`
			Int   int      `mapstructure:"int"`
			Float float64  `mapstructure:"float"`
			Bool  bool     `mapstructure:"bool"`
		} `mapstructure:"map"`
	}{}

	err := config.UnmarshalFromFile("example.yml", &cfg)
	if err != nil {
		fmt.Println(err)
	}
}
```

Supported formats: JSON, TOML, YAML, HCL, INI, envfile and Java properties.

#### Unmarshal environment variables

```go
package main

import (
	"fmt"
	"os"

	"github.com/hapci/go-config"
)

func main() {
	err := os.Setenv("int", "1")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Setenv("string", "text")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Setenv("BOOL", "true")
	if err != nil {
		fmt.Println(err)
	}

	cfg := struct {
		Int    int     `env:"int"`
		String string  `env:"string,required=true"`
		Float  float64 `env:"float,default=3.14"`
		Bool   bool    `env:"bool,BOOL"`
	}{}

	err = config.UnmarshalFromEnv(&cfg)
	if err != nil {
		fmt.Println(err)
	}
}
```