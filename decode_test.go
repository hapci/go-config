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

package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Map TestMap `mapstructure:"map"`
}

type TestMap map[string]TestStruct

type TestStruct struct {
	Array []string `mapstructure:"array"`
	Int   int      `mapstructure:"int" env:"int"`
	Float float64  `mapstructure:"float" env:"float"`
	Bool  bool     `mapstructure:"bool" env:"bool,BOOL"`
}

func TestUnmarshalFromFile(t *testing.T) {
	t.Parallel()

	type args struct {
		filename string
		v        interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantValue interface{}
	}{
		{
			name: "decode yaml",
			args: args{
				filename: filepath.Join("config", "test.yaml"),
				v:        &TestConfig{},
			},
			wantErr: false,
			wantValue: &TestConfig{
				Map: map[string]TestStruct{
					"key1": {
						Array: []string{"item1", "item2", "item3"},
						Int:   4,
						Float: 5.6,
						Bool:  true,
					},
				},
			},
		},
		{
			name: "decode yml",
			args: args{
				filename: filepath.Join("config", "test.yml"),
				v:        &TestConfig{},
			},
			wantErr: false,
			wantValue: &TestConfig{
				Map: map[string]TestStruct{
					"key1": {
						Array: []string{"item10", "item20", "item30"},
						Int:   40,
						Float: 50.6,
						Bool:  true,
					},
				},
			},
		},
		{
			name: "decode toml",
			args: args{
				filename: filepath.Join("config", "test.toml"),
				v:        &TestConfig{},
			},
			wantErr: false,
			wantValue: &TestConfig{
				Map: map[string]TestStruct{
					"key1": {
						Array: []string{"item1", "item2", "item3"},
						Int:   4,
						Float: 5.6,
						Bool:  true,
					},
				},
			},
		},
		{
			name: "decode json",
			args: args{
				filename: filepath.Join("config", "test.json"),
				v:        &TestConfig{},
			},
			wantErr: false,
			wantValue: &TestConfig{
				Map: map[string]TestStruct{
					"key1": {
						Array: []string{"item1", "item2", "item3"},
						Int:   4,
						Float: 5.6,
						Bool:  true,
					},
				},
			},
		},
		{
			name: "no file",
			args: args{
				filename: filepath.Join("config", "wrong_file.json"),
				v:        &TestConfig{},
			},
			wantErr:   true,
			wantValue: &TestConfig{},
		},
		{
			name: "wrong format",
			args: args{
				filename: filepath.Join("config", "wrong_format.yaml"),
				v:        &TestConfig{},
			},
			wantErr:   true,
			wantValue: &TestConfig{Map: map[string]TestStruct{}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := UnmarshalFromFile(tt.args.filename, tt.args.v)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantValue, tt.args.v)
		})
	}
}

func TestUnmarshalFromEnv(t *testing.T) {
	cfg := &TestStruct{}

	err := os.Setenv("int", "1")
	assert.NoError(t, err)

	err = os.Setenv("float", "2.3")
	assert.NoError(t, err)

	err = os.Setenv("BOOL", "true")
	assert.NoError(t, err)

	wantCfg := &TestStruct{
		Int:   1,
		Float: 2.3,
		Bool:  true,
	}

	err = UnmarshalFromEnv(cfg)
	assert.NoError(t, err)
	assert.Equal(t, wantCfg, cfg)
}
