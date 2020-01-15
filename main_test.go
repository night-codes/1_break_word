package main

import (
	"testing"
)

var tcs = []struct {
	str    string
	tokens []string
	result bool
}{{
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
	[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
	false,
}, {
	"bb",
	[]string{"a", "b", "bbb", "bbbb"},
	true,
}, {
	"a",
	[]string{},
	false,
}, {
	"",
	[]string{},
	true,
}, {
	"aaaaaaa",
	[]string{"aaaa", "aaa"},
	true,
}, {
	"leetc",
	[]string{"leet", "code"},
	false,
}, {
	"leetcode",
	[]string{"leet", "code"},
	true,
}, {
	"leetleet",
	[]string{"leet", "code"},
	true,
}, {
	"aaabbbccc",
	[]string{"aaa", "bbb", "ccc", "ddd"},
	true,
}, {
	"",
	[]string{"aaa", "bbb", "ccc", "ddd"},
	true,
}, {
	"aaabbbcccddd",
	[]string{"aaa", "bbb", "ccc", "ddd"},
	true,
}, {
	"aaabbbcccdddcccbbbaaa",
	[]string{"aaa", "bbb", "ccc", "ddd"},
	true,
}, {
	"aaabbbfff",
	[]string{"aaa", "bbb", "ccc", "ddd"},
	false,
}}

func TestValidator(t *testing.T) {
	for _, v := range tcs {
		result := preprocess(v.tokens)(v.str)
		if result != v.result {
			t.Errorf("TestValidator fail for '%v': got <%v> want <%v>", v.str, result, v.result)
		}
	}
}

func BenchmarkSampleShort(b *testing.B) {
	validator := preprocess([]string{"aaa", "bbb", "ccc", "ddd", "aa", "bb", "cc", "dd", "ab", "cd", "ef", "gh"})
	for i := 0; i < b.N; i++ {
		validator("aaabbbccc")
	}
}

func BenchmarkSampleMedium(b *testing.B) {
	validator := preprocess([]string{"aaa", "bbb", "ccc", "ddd", "aa", "bb", "cc", "dd", "ab", "cd", "ef", "gh"})
	for i := 0; i < b.N; i++ {
		validator("aaabbbcccaaabbbccc")
	}
}

func BenchmarkSampleLong(b *testing.B) {
	validator := preprocess([]string{"aaa", "bbb", "ccc", "ddd", "aa", "bb", "cc", "dd", "ab", "bc", "cd", "ef", "gh"})
	for i := 0; i < b.N; i++ {
		validator("aaaabbbbbccccaaabbbcccaaaabbbbbccccaaabbbccc")
	}
}
