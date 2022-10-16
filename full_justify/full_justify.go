package full_justify

import (
	"bytes"
	"fmt"
	"strings"
)

func lineFull(words []string, maxWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-len(words[0]))
	}
	wordsLen := 0
	for i := 0; i < len(words); i++ {
		wordsLen += len(words[i])
	}
	extraCount := maxWidth - wordsLen - (len(words) - 1)
	var line bytes.Buffer
	divided, left := extraCount/(len(words)-1), extraCount%(len(words)-1)
	for i, word := range words {
		var wrappedWord string
		if i < len(words)-1 {
			if i < left {
				wrappedWord = fmt.Sprintf("%s %s", word, strings.Repeat(" ", divided+1))
			} else {
				wrappedWord = fmt.Sprintf("%s %s", word, strings.Repeat(" ", divided))
			}
		} else {
			wrappedWord = word
		}
		line.WriteString(wrappedWord)
	}
	return line.String()
}

func lineDivide(words []string, maxWidth int) (int, string) {
	var length int
	for i, word := range words {
		wordLength := len(word)
		length += wordLength
		switch {
		case length == maxWidth:
			return i + 1, strings.Join((words)[:i+1], " ")
		case length < maxWidth:
			length += 1
		case length > maxWidth:
			return i, lineFull((words)[:i], maxWidth)
		}
	}
	lastLine := strings.Join(words, " ")
	return len(words), lastLine + strings.Repeat(" ", maxWidth-len(lastLine))
}

func fullJustify(words []string, maxWidth int) []string {
	lines := make([]string, 0)
	for len(words) != 0 {
		used, line := lineDivide(words, maxWidth)
		lines = append(lines, line)
		words = words[used:]
	}
	return lines
}
