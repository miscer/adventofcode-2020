package parser

import (
	"reflect"
	"testing"
)

func TestParseBag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		bag   ParsedBag
		err   bool
	}{
		{
			name:  "light red",
			input: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			bag: ParsedBag{
				Color: "light red",
				Contains: map[string]int{
					"bright white": 1,
					"muted yellow": 2,
				},
			},
			err: false,
		},
		{
			name:  "bright white",
			input: "bright white bags contain 1 shiny gold bag.",
			bag: ParsedBag{
				Color: "bright white",
				Contains: map[string]int{
					"shiny gold": 1,
				},
			},
			err: false,
		},
		{
			name:  "drab tan",
			input: "drab tan bags contain 5 drab maroon bags, 5 bright silver bags, 2 dim tan bags.",
			bag: ParsedBag{
				Color: "drab tan",
				Contains: map[string]int{
					"drab maroon":   5,
					"bright silver": 5,
					"dim tan":       2,
				},
			},
			err: false,
		},
		{
			name:  "faded blue",
			input: "faded blue bags contain no other bags.",
			bag: ParsedBag{
				Color:    "faded blue",
				Contains: map[string]int{},
			},
			err: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBag(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("ParseBag() error = %v, expected %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.bag) {
				t.Errorf("ParseBag() got = %v, want %v", got, tt.bag)
			}
		})
	}
}
