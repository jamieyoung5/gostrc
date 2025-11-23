package strutil

import (
	"fmt"
	"strings"
)

// MaxLen finds the length of the longest string in a slice.
func MaxLen(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	return maxWidth
}

// PadBlock pads a text block to a target height and width.
// It right-pads with spaces and bottom-pads with empty lines.
func PadBlock(block string, height, width int, noValue string) []string {
	rows := strings.Split(block, "\n")
	return PadRows(rows, height, width, noValue)
}

// PadRows pads lines of text to a target height and width.
// It right-pads with spaces and bottom-pads with empty lines.
func PadRows(rows []string, height, width int, noValue string) []string {
	if len(rows) == 0 {
		if height == 0 && width == 0 {
			return nil
		}
		if height == 0 {
			height = 1
		}
		if width == 0 {
			width = len(noValue)
		}
		rows = []string{noValue}
	}

	padded := make([]string, height)
	emptyLine := strings.Repeat(" ", width)

	for i := 0; i < height; i++ {
		if i < len(rows) {
			padded[i] = fmt.Sprintf("%-*s", width, rows[i])
		} else {
			padded[i] = emptyLine
		}
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
