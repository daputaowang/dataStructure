package sstring

import (
	"fmt"
	"testing"
)

func TestIndexBF(t *testing.T) {
	// 引入表组测试
	testCases := []struct {
		name   string
		s      SString
		m      SString
		expect int
	}{
		{"substring not exist", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'c'}, 5}, 6},
		{"match successful", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'a'}, 5}, 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := IndexBF(testCase.s, testCase.m)
			t.Log(res)
			if res != testCase.expect {
				t.Errorf("received elem %d, expect elem %d", res, testCase.expect)
			}
		})
	}
}

func BenchmarkIndexBF(b *testing.B) {
	testCases := []struct {
		name   string
		s      SString
		m      SString
		expect int
	}{
		{"substring not exist", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'c'}, 5}, 6},
		{"match successful", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'a'}, 5}, 0},
	}

	for _, testCase := range testCases {
		b.Run(testCase.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				IndexBF(testCase.s, testCase.m)
			}
		})
	}
}

func TestIndexKMP(t *testing.T) {
	// 引入表组测试
	testCases := []struct {
		name   string
		s      SString
		m      SString
		expect int
	}{
		{"substring not exist", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'c'}, 5}, 6},
		{"match successful", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'a'}, 5}, 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			next := getNext(testCase.m)
			res := IndexKMP(testCase.s, testCase.m, next)
			t.Log(res)
			if res != testCase.expect {
				t.Errorf("received elem %d, expect elem %d", res, testCase.expect)
			}
		})
	}
}

func BenchmarkIndexKMP(b *testing.B) {
	testCases := []struct {
		name   string
		s      SString
		m      SString
		expect int
	}{
		{"substring not exist", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'c'}, 5}, 6},
		{"match successful", SString{[]byte{' ', 'a', 'b', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'c', 'b', 'a', 'b'}, 13}, SString{[]byte{' ', 'a', 'b', 'c', 'a', 'a'}, 5}, 0},
	}

	for _, testCase := range testCases {
		b.Run(testCase.name, func(b *testing.B) {
			b.ResetTimer()
			next := getNext(testCase.m)
			for i := 0; i < b.N; i++ {
				IndexKMP(testCase.s, testCase.m, next)
			}
		})
	}
}

func TestGetNext(t *testing.T) {
	s := SString{[]byte{' ', 'a', 'b', 'a', 'a', 'b', 'c', 'a', 'c'}, 8}
	nextS := getNextval(s)
	fmt.Println(nextS[1:])
}
