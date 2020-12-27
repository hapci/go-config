/*
   Copyright The HAPCI Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package config_test

import (
	"fmt"
	"os"

	"github.com/hapci/go-config"
)

func ExampleUnmarshalFromFile() {
	cfg := struct {
		String string `mapstructure:"string"`
		Map    map[string]struct {
			Array []string `mapstructure:"array"`
			Int   int      `mapstructure:"int"`
			Float float64  `mapstructure:"float"`
			Bool  bool     `mapstructure:"bool"`
		} `mapstructure:"map"`
	}{}

	err := config.UnmarshalFromFile("filename.yml", &cfg)
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleUnmarshalFromEnv() {
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
