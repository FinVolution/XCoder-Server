// FILEPATH: /core/retrieval/retrieval_test.go

package retrieval

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatcher_Retrieval(t *testing.T) {
	matcher := NewMatcher("golang", "/path/to/code", 3, 2, 0.5, "query document content")

	tests := []struct {
		name     string
		ctxDoc   string
		expected string
		score    float64
		wantErr  bool
	}{
		{
			name:     "basic functionality",
			ctxDoc:   `package main\nfunc main() {\nprintln("Hello, World!")\n}`,
			expected: `package main\nfunc main() {\nprintln("Hello, World!")\n}`,
			score:    0.5,
			wantErr:  false,
		},
		{
			name:     "empty document",
			ctxDoc:   "",
			expected: "",
			score:    0,
			wantErr:  false,
		},
		{
			name:     "no matching snippet",
			ctxDoc:   `package main\nfunc main() {\nprintln("No Match")\n}`,
			expected: "",
			score:    0,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snippet, score, err := matcher.Retrieval(context.Background(), tt.ctxDoc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matcher.Retrieval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.expected, snippet)
			assert.Equal(t, tt.score, score)
		})
	}
}

func TestMatcher_calJaccardSimilarityScore(t *testing.T) {
	tests := []struct {
		name   string
		words1 []string
		set2   map[string]bool
		want   float64
	}{
		{
			name:   "basic functionality",
			words1: []string{"hello", "world"},
			set2:   map[string]bool{"hello": true, "world": true},
			want:   1.0,
		},
		{
			name:   "partial match",
			words1: []string{"hello", "world"},
			set2:   map[string]bool{"hello": true, "golang": true},
			want:   0.3333333333333333,
		},
		{
			name:   "no match",
			words1: []string{"hello", "world"},
			set2:   map[string]bool{"golang": true, "programming": true},
			want:   0.0,
		},
		{
			name:   "empty words1",
			words1: []string{},
			set2:   map[string]bool{"hello": true, "world": true},
			want:   0.0,
		},
		{
			name:   "empty set2",
			words1: []string{"hello", "world"},
			set2:   map[string]bool{},
			want:   0.0,
		},
		{
			name:   "both empty",
			words1: []string{},
			set2:   map[string]bool{},
			want:   0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matcher{}
			got := m.calJaccardSimilarityScore(tt.words1, tt.set2)
			assert.Equal(t, tt.want, got)
		})
	}
}
