package retrieval

import (
	"context"
	"math"
	"regexp"
	"sort"
	"strings"
	"xcoder/internal/model/input/codein"
)

var (
	golangKeyWords     = []string{"return", "continue", "break", "if", "else", "func", "interface", "struct", "select", "case", "package", "defer", "go", "map", "struct", "chan", "range", "goto", "switch", "const", "fallthrough", "if", "type", "default", "for", "import", "var"}
	pythonKeywords     = []string{"class", "except", "import", "as", "if", "with", "elif", "else", "for", "while", "def", "try", "from", "finally", "in", "is", "lambda", "yield", "del", "assert", "raise", "break", "continue", "pass", "nonlocal", "return"}
	javaKeywords       = []string{"case", "boolean", "break", "finally", "abstract", "char", "class", "catch", "const", "case", "double", "continue", "default", "do", "byte", "goto", "else", "implements", "extends", "final", "finally", "float", "for", "assert", "enum", "if", "import"}
	javascriptKeywords = []string{"class", "catch", "delete", "await", "break", "byte", "float", "case", "boolean", "char", "class", "abstract", "const", "continue", "debugger", "default", "delete", "extends", "do", "arguments", "double", "final", "else", "enum", "eval", "export", "false", "finally"}
	defaultKeywords    = []string{"for", "break", "it", "class", "import"}
	re                 = regexp.MustCompile(`[^a-zA-Z0-9]+`)
)

type Matcher struct {
	codeLanguage    string
	codePath        string
	chunkSize       int
	slideWindowSize int
	minScore        float64
	queryDoc        string
}

func NewMatcher(codeLanguage, codePath string, chunkSize int, slideWindowSize int, minScore float64, queryDoc string) *Matcher {
	return &Matcher{
		codeLanguage:    codeLanguage,
		codePath:        codePath,
		chunkSize:       chunkSize,
		slideWindowSize: slideWindowSize,
		minScore:        minScore,
		queryDoc:        queryDoc,
	}
}

type RetrievalResult struct {
	Snippet string
	Score   float64
	Type    string
	Path    string
}

func (m *Matcher) RetrieveAllSnippets(ctx context.Context, ctxDoc []*codein.Context) ([]*RetrievalResult, error) {
	allSnippets := make([]*RetrievalResult, 0)
	for _, doc := range ctxDoc {
		snippet, score, err := m.Retrieval(ctx, doc.Content)
		if err != nil {
			return nil, err
		}

		if score > m.minScore {
			res := &RetrievalResult{
				Snippet: snippet,
				Score:   score,
				Type:    doc.Type,
				Path:    doc.Path,
			}
			allSnippets = append(allSnippets, res)
		}
	}

	// 将 allSnippets 按照 Score 从小到大排序
	sort.Slice(allSnippets, func(i, j int) bool {
		return allSnippets[i].Score < allSnippets[j].Score
	})

	return allSnippets, nil
}

func (m *Matcher) Retrieval(ctx context.Context, ctxDoc string) (string, float64, error) {
	lines := strings.Split(ctxDoc, "\n")

	tokenizeLines := make([][]string, len(lines))
	for idx, line := range lines {
		i, err := m.tokenize(line, m.codeLanguage)
		if err != nil {
			return "", 0, err
		}
		tokenizeLines[idx] = i
	}

	var data []map[string]interface{}
	delineations := m.getWindowDelineationsWithNext(lines)

	queryDocTokensSet := m.GetQueryDocTokens()
	for _, delineation := range delineations {
		start, end, sEnd := delineation[0], delineation[1], delineation[2]
		var lineTokens []string
		for _, tokenLine := range tokenizeLines[start:end] {
			lineTokens = append(lineTokens, tokenLine...)
		}
		score := m.calJaccardSimilarityScore(lineTokens, queryDocTokensSet)
		data = append(data, map[string]interface{}{
			"score": score,
			"start": start,
			"end":   end,
			"sEnd":  sEnd,
		})
	}

	snippet, score := m.getMaxScoreSnippet(ctx, data, lines)
	return snippet, score, nil
}

func (m *Matcher) GetQueryDocTokens() map[string]bool {
	var queryDocTokens []string
	queryDocTokensSet := make(map[string]bool)

	lines := strings.Split(m.queryDoc, "\n")
	for _, line := range lines {
		tokens, _ := m.tokenize(line, m.codeLanguage)
		queryDocTokens = append(queryDocTokens, tokens...)
	}

	// tokens 去重
	for _, word := range queryDocTokens {
		queryDocTokensSet[word] = true
	}
	return queryDocTokensSet
}

func (m *Matcher) getMaxScoreSnippet(ctx context.Context, data []map[string]interface{}, rawLines []string) (string, float64) {
	maxScore := math.Inf(-1)
	var maxIndex int

	for idx, item := range data {
		score := item["score"].(float64)
		if score > maxScore {
			maxScore = score
			maxIndex = idx
		}
	}

	start := data[maxIndex]["start"].(int)
	_ = data[maxIndex]["end"].(int)
	sEnd := data[maxIndex]["sEnd"].(int)

	if sEnd >= len(rawLines) {
		sEnd = len(rawLines)
	}

	return strings.Join(rawLines[start:sEnd], "\n"), maxScore
}

func (m *Matcher) calJaccardSimilarityScore(words1 []string, set2 map[string]bool) float64 {
	set1 := make(map[string]bool, len(words1))
	for _, word := range words1 {
		set1[word] = true
	}

	intersection := 0
	for word := range set1 {
		if set2[word] {
			intersection++
		}
	}

	return float64(intersection) / float64(len(set1)+len(set2)-intersection)
}

func (m *Matcher) getWindowDelineationsWithNext(lines []string) [][3]int {
	delineations := make([][3]int, 0)
	arrayLength := len(lines)

	if arrayLength == 0 {
		return delineations
	}
	if arrayLength < m.chunkSize {
		return [][3]int{{0, arrayLength, arrayLength}}
	}
	for i := 0; i < arrayLength-m.chunkSize+1; i++ {
		delineations = append(delineations, [3]int{i, i + m.chunkSize - 1, i + 2*m.chunkSize - 1})
	}
	return delineations
}

func (m *Matcher) tokenize(code string, language string) ([]string, error) {
	words := re.Split(code, -1)

	var keywords []string
	switch language {
	case "golang":
		keywords = golangKeyWords
	case "python":
		keywords = pythonKeywords
	case "java":
		keywords = javaKeywords
	case "kotlin":
		keywords = append(javaKeywords, "static")
	case "groovy":
		keywords = append(javaKeywords, "static")
	case "javascript":
		keywords = javascriptKeywords
	case "typescript":
		keywords = append(javascriptKeywords, "var")
	default:
		keywords = defaultKeywords
	}

	resp := filterWords(words, keywords)

	return resp, nil
}

func filterWords(words []string, keywords []string) []string {
	filteredWords := []string{}
	for _, word := range words {
		if !contains(keywords, word) {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords
}

func contains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
