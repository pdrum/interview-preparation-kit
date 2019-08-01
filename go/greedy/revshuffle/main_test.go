package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	run(t, []TC{
		{
			In:     "jjcddjggcdjd",
			Wanted: "cgddjj",
		},
		{
			In:     "beaa",
			Wanted: "aeb",
		},
		{
			In:     "djjcddjggbiigjhfghehhbgdigjicafgjcehhfgifadihiajgciagicdahcbajjbhifjiaajigdgdfhdiijjgaiejgegbbiigida",
			Wanted: "aaaaabccigicgjihidfiejfijgidgbhhehgfhjgiibggjddjjd",
		},
		{
			In:     "bdabaceadaedaaaeaecdeadababdbeaeeacacaba",
			Wanted: "aaaaaabaaceededecbdb",
		},
		{
			In:     "aahaxhxhxxahxx",
			Wanted: "axhxhxa",
		},
		{
			In:     "eggegg",
			Wanted: "egg",
		},
		{
			In:     "abcdefgabcdefg",
			Wanted: "agfedcb",
		},
		{
			In:     "aeiouuoiea",
			Wanted: "aeiou",
		},
	})
}

type TC struct {
	In     string
	Wanted string
}

func run(t *testing.T, tcs []TC) {
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("for input %s", tc.In), func(t *testing.T) {
			have := reverseShuffleMerge(tc.In)
			if reverseShuffleMerge(tc.In) != tc.Wanted {
				fmt.Printf("want %s have %s\n", tc.Wanted, have)
				t.Fail()
			}
		})
	}
}
