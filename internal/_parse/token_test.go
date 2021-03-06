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

package parse_test

import (
	"fmt"

	. "github.com/hajimehoshi/goc/internal/parse"
	"github.com/hajimehoshi/goc/internal/preprocess"
)

func outputTokens(path string, srcs map[string]string) {
	files := map[string]preprocess.PPTokenReader{}
	for path, src := range srcs {
		files[path] = preprocess.Tokenize([]byte(src), "")
	}

	tokens := Tokenize(preprocess.Preprocess(path, files))
	for {
		t, err := tokens.NextToken()
		if err != nil {
			fmt.Println("error")
			return
		}
		if t.Type == EOF {
			break
		}
		fmt.Println(t)
	}
}

func ExampleEmpty() {
	outputTokens("main.c", map[string]string{
		"main.c": `#`,
	})
	// Output:
}

func ExampleCalc() {
	outputTokens("main.c", map[string]string{
		"main.c": `1+1=2`,
	})
	// Output:
	// integer: 1 (int)
	// +
	// integer: 1 (int)
	// =
	// integer: 2 (int)
}

func ExampleHelloWorld() {
	outputTokens("main.c", map[string]string{
		"stdio.h": `foo bar`,
		"main.c": `#include <stdio.h>

int main() {
  printf("Hello, World!\n");
  return 0;
}`,
	})
	// Output:
	// identifier: foo
	// identifier: bar
	// int
	// identifier: main
	// (
	// )
	// {
	// identifier: printf
	// (
	// string: "Hello, World!\n"
	// )
	// ;
	// return
	// integer: 0 (int)
	// ;
	// }
}
