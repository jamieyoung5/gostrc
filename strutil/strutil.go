package strutil

import (
	"regexp"
	"strings"
)

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

func VisibleLen(str string) int {
	return len(ansiEscape.ReplaceAllString(str, ""))
}

// MaxLen finds the length of the longest string in a slice.
func MaxLen(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		if l := len(line); l > maxWidth {
			maxWidth = l
		}
	}
	return maxWidth
}

// MaxVisibleLen finds the visible length of the longest string in a slice.
func MaxVisibleLen(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		if l := VisibleLen(line); l > maxWidth {
			maxWidth = l
		}
	}
	return maxWidth
}

// PadBlock splits a block into rows and pads them using standard byte length.
func PadBlock(block string, height, width int, noValue string) []string {
	rows := strings.Split(block, "\n")
	return PadRows(rows, height, width, noValue)
}

// PadVisibleBlock splits a block into rows and pads them using visible length (ANSI-aware).
func PadVisibleBlock(block string, height, width int, noValue string) []string {
	rows := strings.Split(block, "\n")
	return PadVisibleRows(rows, height, width, noValue)
}

// PadRows pads lines of text to a target height and width using byte length.
func PadRows(rows []string, height, width int, noValue string) []string {
	return padRows(rows, height, width, noValue, func(s string) int { return len(s) })
}

// PadVisibleRows pads lines of text to a target height and width using visible length.
func PadVisibleRows(rows []string, height, width int, noValue string) []string {
	return padRows(rows, height, width, noValue, VisibleLen)
}

// padRows is the internal worker that applies the specific length strategy.
func padRows(
	rows []string,
	height, width int,
	noValue string,
	lenFunc func(string) int,
) []string {
	if len(rows) == 0 {
		if height == 0 && width == 0 {
			return nil
		}
		if height == 0 {
			height = 1
		}
		if width == 0 {
			width = lenFunc(noValue)
		}
		rows = []string{noValue}
	}

	padded := make([]string, height)

	for i := 0; i < height; i++ {
		var line string
		if i < len(rows) {
			line = rows[i]
		}

		currentLen := lenFunc(line)
		if currentLen < width {
			line += strings.Repeat(" ", width-currentLen)
		}

		padded[i] = line
	}

	return padded
}

// SideBySide joins text blocks horizontally with specified spacing, padding lines to maintain column alignment.
func SideBySide(spacing int, blocks ...string) string {
	if len(blocks) == 0 {
		return ""
	}

	numBlocks := len(blocks)
	grid := make([][]string, numBlocks)
	widths := make([]int, numBlocks)
	height, maxW := 0, 0

	for i, b := range blocks {
		lines := strings.Split(b, "\n")
		grid[i] = lines
		if len(lines) > height {
			height = len(lines)
		}
		for _, l := range lines {
			if len(l) > widths[i] {
				widths[i] = len(l)
			}
		}
		if widths[i] > maxW {
			maxW = widths[i]
		}
	}

	var sb strings.Builder
	sep := strings.Repeat(" ", spacing)
	pads := strings.Repeat(" ", maxW)

	for y := 0; y < height; y++ {
		for x, lines := range grid {
			s := ""
			if y < len(lines) {
				s = lines[y]
			}
			sb.WriteString(s)

			if x < numBlocks-1 {
				sb.WriteString(pads[:widths[x]-len(s)])
				sb.WriteString(sep)
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
