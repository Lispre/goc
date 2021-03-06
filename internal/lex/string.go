// Copyright 2018 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lex

import (
	"fmt"
)

func ReadString(src ByteReadPeeker) (string, error) {
	if err := shouldRead(src, '"'); err != nil {
		return "", err
	}

	bs := []byte{}
loop:
	for {
		b, err := shouldPeekByte(src)
		if err != nil {
			return "", err
		}
		switch b {
		case '"':
			mustDiscard(src, 1)
			break loop
		case '\\':
			b, err := ReadEscapedChar(src)
			if err != nil {
				return "", err
			}
			if b > 255 {
				return "", fmt.Errorf("lex: not implemented")
			}
			bs = append(bs, byte(b))
			continue loop
		case '\r', '\n':
			return "", fmt.Errorf("lex: newline in string")
		}
		mustDiscard(src, 1)
		bs = append(bs, b)
	}

	return string(bs), nil
}
