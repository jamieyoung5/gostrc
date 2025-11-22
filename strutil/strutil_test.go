package strutil

import (
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
