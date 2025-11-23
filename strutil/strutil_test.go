package strutil

import (
	"reflect"
	"testing"
)

func TestMaxLen(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "nil_slice",
			input: nil,
			want:  0,
		},
		{
			name:  "empty_slice",
			input: []string{},
			want:  0,
		},
		{
			name:  "single_empty_string",
			input: []string{""},
			want:  0,
		},
		{
			name:  "single_string",
			input: []string{"hello"},
			want:  5,
		},
		{
			name:  "mixed_lengths_longest_first",
			input: []string{"longest", "short", "tiny"},
			want:  7,
		},
		{
			name:  "mixed_lengths_longest_last",
			input: []string{"tiny", "short", "longest"},
			want:  7,
		},
		{
			name:  "mixed_lengths_longest_middle",
			input: []string{"tiny", "longest", "short"},
			want:  7,
		},
		{
			name:  "multibyte_characters",
			input: []string{"a", "£", "😊"},
			want:  4, // 😊 is 4 bytes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxLen(tt.input); got != tt.want {
				t.Errorf("MaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRows(t *testing.T) {
	tests := []struct {
		name    string
		rows    []string
		height  int
		width   int
		noValue string
		want    []string
	}{
		{
			name:    "basic padding",
			rows:    []string{"foo", "bar"},
			height:  3,
			width:   5,
			noValue: "",
			want:    []string{"foo  ", "bar  ", "     "},
		},
		{
			name:    "truncation",
			rows:    []string{"foo", "bar", "baz"},
			height:  2,
			width:   3,
			noValue: "",
			want:    []string{"foo", "bar"},
		},
		{
			name:    "empty rows zero dimensions",
			rows:    []string{},
			height:  0,
			width:   0,
			noValue: "N/A",
			want:    nil,
		},
		{
			name:    "empty rows default height",
			rows:    []string{},
			height:  0,
			width:   5,
			noValue: "-",
			want:    []string{"-    "},
		},
		{
			name:    "empty rows default width",
			rows:    []string{},
			height:  2,
			width:   0,
			noValue: "---",
			want:    []string{"---", "   "},
		},
		{
			name:    "width expansion",
			rows:    []string{"a"},
			height:  1,
			width:   3,
			noValue: "",
			want:    []string{"a  "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRows(tt.rows, tt.height, tt.width, tt.noValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PadRows() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPadBlock(t *testing.T) {
	tests := []struct {
		name    string
		block   string
		height  int
		width   int
		noValue string
		want    []string
	}{
		{
			name:    "single line block",
			block:   "hello",
			height:  2,
			width:   6,
			noValue: "",
			want:    []string{"hello ", "      "},
		},
		{
			name:    "multi line block",
			block:   "line1\nline2",
			height:  3,
			width:   5,
			noValue: "",
			want:    []string{"line1", "line2", "     "},
		},
		{
			name:    "empty block with noValue",
			block:   "",
			height:  2,
			width:   3,
			noValue: "N/A",
			want:    []string{"   ", "   "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadBlock(tt.block, tt.height, tt.width, tt.noValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PadBlock() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSideBySide(t *testing.T) {
	tests := []struct {
		name    string
		spacing int
		blocks  []string
		want    string
	}{
		{
			name:    "nil blocks",
			spacing: 1,
			blocks:  nil,
			want:    "",
		},
		{
			name:    "empty blocks slice",
			spacing: 1,
			blocks:  []string{},
			want:    "",
		},
		{
			name:    "single block",
			spacing: 1,
			blocks:  []string{"foo\nbar"},
			want:    "foo\nbar\n",
		},
		{
			name:    "two blocks single line",
			spacing: 1,
			blocks:  []string{"A", "B"},
			want:    "A B\n",
		},
		{
			name:    "two blocks varied height first taller",
			spacing: 1,
			blocks:  []string{"A\nB", "C"},
			want:    "A C\nB \n",
		},
		{
			name:    "two blocks varied height second taller",
			spacing: 1,
			blocks:  []string{"A", "B\nC"},
			want:    "A B\n  C\n",
		},
		{
			name:    "alignment with variable width",
			spacing: 2,
			blocks:  []string{"long\nshort", "R"},
			want:    "long   R\nshort  \n",
		},
		{
			name:    "multiple blocks mixed",
			spacing: 1,
			blocks:  []string{"A", "B\nB", "C"},
			want:    "A B C\n  B \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SideBySide(tt.spacing, tt.blocks...); got != tt.want {
				t.Errorf("SideBySide() = \n%q\nwant \n%q", got, tt.want)
			}
		})
	}
}
