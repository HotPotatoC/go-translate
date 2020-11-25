package cli

import (
	"sync"
	"testing"
)

type TestCase struct {
	sl  string
	tl  string
	st  string
	out string
}

var testCases = []TestCase{
	{"en", "id", "How are you?", "Apa kabar?"},
	{"en", "id", "Have you done today's assignment?", "Sudahkah Anda menyelesaikan tugas hari ini?"},
	{"en", "ja", "Good morning friend", "おはよう友達"},
	{"en", "no", "I am going to buy a cup of coffee", "Jeg skal kjøpe en kopp kaffe"},
	{"en", "ru", "This is the documentation for our API", "Это документация для нашего API"},
	{"id", "fr", "Aku menggunakan bahasa pemrograman Javascript", "J'utilise le langage de programmation Javascript"},
}

func TestRequestTranslate(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan string, len(testCases))

	for _, tc := range testCases {
		wg.Add(1)
		go RequestTranslate(&RequestBody{
			tc.sl, tc.tl, tc.st,
		}, ch, &wg)
		res := <-ch
		if res == "You have been rate limited, Try again later." {
			t.Log(res)
			break
		}

		if res != tc.out {
			t.Errorf("Expected \"%s\" got \"%s\"", tc.out, res)
		}
	}
	close(ch)
	wg.Wait()
}
