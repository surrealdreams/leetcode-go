package main

import "testing"

func BenchmarkGenerateParenthesis5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(5)
	}
}

func BenchmarkGenerateParenthesis10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(10)
	}
}
