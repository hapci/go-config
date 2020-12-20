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
	"sync"

	"github.com/Netflix/go-env"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	readerLock = &sync.RWMutex{}
	reader     = viper.New()
)

// UnmarshalFromFile decodes the file and assigns the decoded values
// into the out value.
func UnmarshalFromFile(filename string, v interface{}) error {
	readerLock.Lock()
	defer readerLock.Unlock()

	reader.SetConfigFile(filename)

	if err := reader.ReadInConfig(); err != nil {
		return errors.Wrap(err, "reading configuration")
	}

	return errors.Wrap(reader.Unmarshal(v), "decoding configuration")
}

func UnmarshalFromEnv(v interface{}) error {
	_, err := env.UnmarshalFromEnviron(v)
	return err
}
