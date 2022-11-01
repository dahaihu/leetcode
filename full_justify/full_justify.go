package full_justify

import (
	"bytes"
	"strings"
)

func justifyLine(words []string, maxWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-len(words[0]))
	}
	extraCount := maxWidth
	for _, word := range words {
		extraCount -= len(word)
	}
	extraCount -= len(words) - 1
	var line bytes.Buffer
	gap, remainder := extraCount/(len(words)-1), extraCount%(len(words)-1)
	for idx, word := range words[:len(words)-1] {
		line.WriteString(word + strings.Repeat(" ", gap+1))
		if idx < remainder {
			line.WriteString(" ")
		}
	}
	line.WriteString(words[len(words)-1])
	return line.String()
}

func lineWords(words []string, maxWidth int) ([]string, string) {
	wordsLen := 0
	for idx, word := range words {
		wordsLen += len(word)
		switch {
		case wordsLen == maxWidth:
			return words[idx+1:], strings.Join(words[:idx+1], " ")
		case wordsLen < maxWidth:
			wordsLen += 1
		case wordsLen > maxWidth:
			return words[idx:], justifyLine(words[:idx], maxWidth)
		}
	}
	line := strings.Join(words, " ")
	return nil, line + strings.Repeat(" ", maxWidth-len(line))
}

func fullJustify(words []string, maxWidth int) []string {
	lines := []string{}
	var line string
	for len(words) != 0 {
		words, line = lineWords(words, maxWidth)
		lines = append(lines, line)
	}
	return lines
}
