// Copyright 2020 Eryx <evorui аt gmail dοt com>, All rights reserved.
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

package hmsg

import (
	"testing"

	"golang.org/x/text/language"
)

func Test_MailTemplateLang_Valid(t *testing.T) {

	for _, v := range []string{
		"zh-cn",
		"zh_cn",
		"zh-CN",
		"zh_CN",
	} {
		if _, err := language.Parse(v); err != nil {
			t.Fatal(err)
		}
	}
}

func Benchmark_LangParse_HIT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		language.Parse("zh-cn")
	}
}

func Benchmark_LangParse_MIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		language.Parse("zh-00")
	}
}
