package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/twip-link/bashir"
)

// go test -run=.*/line_zero -v
func TestElim(t *testing.T) {
	var debug bool

	df1 := "Mon 02 January 2006 03:04:05 PM MST"
	df2 := "2006-01-02"

	var input1 = []string{`From the Poetry Foundation on 2023-05-26:`}

	var input2 = []string{`From the Poetry Foundation on 2023-05-26:
https://www.poetryfoundation.org/poems/44272/the-road-not-taken`}

	var input0 = []string{`From the Poetry Foundation on 2023-05-26:
https://www.poetryfoundation.org/poems/44272/the-road-not-taken
The Road Not Taken
BY ROBERT FROST
Two roads diverged in a yellow wood,
And sorry I could not travel both
And be one traveler, long I stood
And looked down one as far as I could
To where it bent in the undergrowth;
Then took the other, as just as fair,
And having perhaps the better claim,
Because it was grassy and wanted wear;
Though as for that the passing there
Had worn them really about the same,
And both that morning equally lay
In leaves no step had trodden black.
Oh, I kept the first for another day!
Yet knowing how way leads on to way,
I doubted if I should ever come back.
I shall be telling this with a sigh
Somewhere ages and ages hence:
Two roads diverged in a wood, and I—
I took the one less traveled by,
And that has made all the difference.`}

	call := func(l int) []string {
		path := "/home/jeff/tl/examples/elim/elim"
		return []string{fmt.Sprintf(`%v < input.txt -l %d`, path, l)}
	}

	tests := map[string]struct {
		input []string
		want  []string
	}{
		"line one":  {call(1), input1},
		"line two":  {call(2), input2},
		"line zero": {call(0), input0},
		"multi": {
			input: []string{
				"date",
				"date -u",
				`date +"%Y-%m-%d"`,
			},
			want: []string{
				df(time.Now(), df1),
				df(time.Now().UTC(), df1),
				df(time.Now(), df2),
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := bashir.Bash(tc.input...)

			if debug {
				for i := 0; i < len(got); i++ {
					fmt.Printf("%d : %s\n", i, got[i])
				}
			}

			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func df(d time.Time, f string) string {
	return d.Format(f)
}
