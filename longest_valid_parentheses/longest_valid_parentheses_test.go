package longest_valid_parentheses

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_longestValidParentheses(t *testing.T) {
	assert.Equal(t, 2, longestValidParentheses("(()"))
	assert.Equal(t, 6, longestValidParentheses("()(())"))
}
