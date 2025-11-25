package strutil_test

import (
	"reflect"
	"testing"

	"github.com/jamieyoung5/gostrc/strutil"
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
			if got := strutil.MaxLen(tt.input); got != tt.want {
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
			if got := strutil.PadRows(tt.rows, tt.height, tt.width, tt.noValue); !reflect.DeepEqual(got, tt.want) {
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
			if got := strutil.PadBlock(tt.block, tt.height, tt.width, tt.noValue); !reflect.DeepEqual(got, tt.want) {
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
			if got := strutil.SideBySide(tt.spacing, tt.blocks...); got != tt.want {
				t.Errorf("SideBySide() = \n%q\nwant \n%q", got, tt.want)
			}
		})
	}
}

func TestVisibleLen(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "empty string",
			input: "",
			want:  0,
		},
		{
			name:  "plain text",
			input: "hello",
			want:  5,
		},
		{
			name:  "simple ansi color",
			input: "\x1b[31mred\x1b[0m",
			want:  3,
		},
		{
			name:  "complex ansi sequences",
			input: "\x1b[31;1;4mboldredunder\x1b[0m",
			want:  12,
		},
		{
			name:  "mixed ansi and text",
			input: "pre\x1b[32mmid\x1b[0mpost",
			want:  10,
		},
		{
			name:  "only ansi",
			input: "\x1b[H\x1b[2J",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strutil.VisibleLen(tt.input); got != tt.want {
				t.Errorf("VisibleLen(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaxVisibleLen(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "plain strings",
			input: []string{"a", "bb", "ccc"},
			want:  3,
		},
		{
			name:  "ansi shorter than plain visible",
			input: []string{"\x1b[31ma\x1b[0m", "bb"},
			want:  2, // "a" (1) vs "bb" (2)
		},
		{
			name:  "ansi longer than plain visible",
			input: []string{"\x1b[31mhello\x1b[0m", "hi"},
			want:  5, // "hello" (5) vs "hi" (2)
		},
		{
			name:  "mixed ansi strings",
			input: []string{"\x1b[31mshort\x1b[0m", "\x1b[32mlonger\x1b[0m"},
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strutil.MaxVisibleLen(tt.input); got != tt.want {
				t.Errorf("MaxVisibleLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadVisibleRows(t *testing.T) {
	tests := []struct {
		name    string
		rows    []string
		height  int
		width   int
		noValue string
		want    []string
	}{
		{
			name:    "pad ansi string",
			rows:    []string{"\x1b[31mred\x1b[0m"},
			height:  1,
			width:   5,
			noValue: "",
			// Should add 2 spaces to the raw string
			want: []string{"\x1b[31mred\x1b[0m  "},
		},
		{
			name:    "ansi string exact width",
			rows:    []string{"\x1b[31mred\x1b[0m"},
			height:  1,
			width:   3,
			noValue: "",
			want:    []string{"\x1b[31mred\x1b[0m"},
		},
		{
			name:    "ansi string truncation (not supported but ensures no panic)",
			rows:    []string{"\x1b[31mred\x1b[0m"},
			height:  1,
			width:   2,
			noValue: "",
			// Current implementation does not truncate, just returns as is
			want: []string{"\x1b[31mred\x1b[0m"},
		},
		{
			name:    "mixed plain and ansi alignment",
			rows:    []string{"foo", "\x1b[32mbar\x1b[0m"},
			height:  2,
			width:   5,
			noValue: "",
			want:    []string{"foo  ", "\x1b[32mbar\x1b[0m  "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strutil.PadVisibleRows(tt.rows, tt.height, tt.width, tt.noValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PadVisibleRows() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPadVisibleBlock(t *testing.T) {
	tests := []struct {
		name    string
		block   string
		height  int
		width   int
		noValue string
		want    []string
	}{
		{
			name:    "multiline ansi block",
			block:   "\x1b[31mline1\x1b[0m\nline2",
			height:  2,
			width:   6,
			noValue: "",
			want:    []string{"\x1b[31mline1\x1b[0m ", "line2 "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strutil.PadVisibleBlock(tt.block, tt.height, tt.width, tt.noValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PadVisibleBlock() = %q, want %q", got, tt.want)
			}
		})
	}
}
