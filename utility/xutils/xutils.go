package xutils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"strings"
	"time"
)

// CodeLanguageToLower
// 将 code language 转换为小写，并将 go 转换为 golang
func CodeLanguageToLower(codeLanguage string) string {
	codeLanguage = strings.ToLower(codeLanguage)
	if codeLanguage == "go" {
		codeLanguage = "golang"
	}
	return codeLanguage
}

func CodeLanguageAnnotationFlagGet(codeLanguage string) string {
	var codeLanguageAnnotationFlagMap = map[string]string{
		"golang":     "//",
		"java":       "//",
		"python":     "#",
		"javascript": "//",
		"typescript": "//",
		"c":          "//",
		"cpp":        "//",
		"csharp":     "//",
		"php":        "//",
		"ruby":       "#",
		"rust":       "//",
		"scala":      "//",
		"swift":      "//",
		"kotlin":     "//",
		"groovy":     "//",
		"dart":       "//",
		"elixir":     "#",
		"haskell":    "--",
		"erlang":     "%%",
		"fortran":    "!",
	}

	if val, ok := codeLanguageAnnotationFlagMap[codeLanguage]; ok {
		return val
	} else {
		return "//"
	}
}

// GetEnv
// 获取环境变量，如果没有则返回默认值
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// StringIsInSmallSlice 判断字符串是否在字符串数组中
func StringIsInSmallSlice(target string, elements []string) bool {
	for _, s := range elements {
		if s == target {
			return true
		}
	}
	return false
}

// JudgeSliceIsEq
// 判断字符串数组是否相等
func JudgeSliceIsEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// FirstNoEmptyLineIdx
// 获取字符串中不是空行的第一行索引
func FirstNoEmptyLineIdx(lines []string) int {
	var idx int
	for i, line := range lines {
		if strings.TrimSpace(line) != "" {
			idx = i
			break
		}
	}

	return idx
}

// TrimSlices
// 去掉字符串数组中，每行首尾的 \t 或空格
func TrimSlices(lines []string) []string {
	var res []string
	for _, line := range lines {
		res = append(res, strings.Trim(line, "\t "))
	}

	return res
}

func findESIdx(ctx context.Context, gLines []string, sLines []string) int {
	gIdx := len(gLines) - 1
	if len(sLines) == 0 {
		return gIdx + 1
	}

	//multiLinesNeedCheckMoreFlags := xapollo.ShortCutParamsConfigInfo.MultiLineNeedCheckMoreFlags
	multiLinesNeedCheckMoreFlags := []string{"}", "]"}

	// 获取去掉 \t 或空格后的字符串数组
	gCLines := TrimSlices(gLines)
	sCLines := TrimSlices(sLines)

	// 获取不是空行的第一行，作为开始比较的行
	gSIdx := FirstNoEmptyLineIdx(gCLines)
	sSIdx := FirstNoEmptyLineIdx(sCLines)

	firstNoEmptyLine := sCLines[sSIdx]
	g.Log().Infof(ctx, "firstNoEmptyLine: %s", firstNoEmptyLine)
	for i := gSIdx; i <= len(gCLines)-1; i++ {
		if gCLines[i] == firstNoEmptyLine {
			// 白名单里，当前行相等，存在下一行不相等的话，eIdx 取全；否则，返回当前行的上一行
			if StringIsInSmallSlice(firstNoEmptyLine, multiLinesNeedCheckMoreFlags) {
				needCheckLines := 1
				if len(gCLines) >= i+needCheckLines+1 && len(sCLines) >= sSIdx+needCheckLines+1 {
					if !JudgeSliceIsEq(gCLines[i+1:i+needCheckLines+1], sCLines[sSIdx+1:sSIdx+needCheckLines+1]) {
						gIdx = gIdx
					} else {
						// 返回当前行的上一行
						gIdx = i - 1
						break
					}
				}
				// 相等的符号在最后一行，直接返回当前行的上一行
				if len(gCLines) == i+1 {
					gIdx = i - 1
					break
				}
			} else {
				// 非白名单里，当前行相等，直接返回当前行的上一行
				gIdx = i - 1
				break
			}
		}
	}

	return gIdx + 1
}

func HandleRemoveDuplicateCodeForMultiLine(ctx context.Context, generateCode string, suffixCode string) string {
	g.Log().Infof(ctx, "RemoveDuplicateCodeForMultiLine start ...")
	start := time.Now()

	gLines := strings.Split(generateCode, "\n")
	sLines := strings.Split(suffixCode, "\n")

	eIdx := findESIdx(ctx, gLines, sLines)

	duration := time.Since(start).Milliseconds()
	resp := strings.Join(gLines[:eIdx], "\n")

	g.Log().Infof(ctx, "RemoveDuplicateCodeForMultiLine end duration: %d ms, code: %s", duration, resp)

	return resp
}

func HandleGenerateTextWithLengthStop(ctx context.Context, s string) string {
	g.Log().Infof(ctx, "handleGenerateTextWithLengthStop with string: %s", s)
	lines := strings.Split(s, "\n")
	if len(lines) > 1 {
		lines = lines[:len(lines)-1]
		return strings.Join(lines, "\n")
	}
	return s
}

func HandleRemoveTailSpaceIfExistForSingleLine(ctx context.Context, s string) string {
	g.Log().Infof(ctx, "handleRemoveTailSpaceIfExistForSingleLine with string: %s", s)
	return strings.TrimRight(s, " ")
}
