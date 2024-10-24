package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sergi/go-diff/diffmatchpatch"
	"gopkg.in/guregu/null.v3"
	"math/rand"
	"strconv"
	"strings"
)

// StringIsInSmallSlice 判断字符串是否在字符串数组中
func StringIsInSmallSlice(target string, elements []string) bool {
	for _, s := range elements {
		if s == target {
			return true
		}
	}
	return false
}

func countLeadingSpaces(s string) int {
	trimmed := strings.TrimLeft(s, " ")
	return len(s) - len(trimmed)
}

func replaceTabsWithSpaces(input string, spaceCount int) string {
	// 使用 strings.Replace 将制表符 \t 替换为指定数量的空格
	// 注意：这里假设每个制表符都会替换为 spaceCount 个空格
	result := strings.Replace(input, "\t", strings.Repeat(" ", spaceCount), -1)
	return result
}

func getLastLineOfCodeWithIndent(input string) string {
	lines := strings.Split(input, "\n")
	// 遍历字符串的每一行，找到最后一个非空白行
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if strings.TrimSpace(line) != "" {
			return line
		}
	}
	return "" // 如果没有找到非空行，返回空字符串
}

func JudgeCodeIsMultiLine(prefix string, suffix string) bool {
	lastPrefixCode := getLastLineOfCodeWithIndent(prefix)
	suffixCodes := strings.SplitAfter(suffix, "\n")

	var firstSuffixCode string
	for _, s := range suffixCodes {
		if s != "\n" {
			firstSuffixCode = s
			break
		}
	}

	g.Log().Infof(context.Background(), "lastPrefixCode: %s", lastPrefixCode)
	g.Log().Infof(context.Background(), "firstSuffixCode: %s", firstSuffixCode)

	lastPrefixCodeSpaceNums := countLeadingSpaces(replaceTabsWithSpaces(lastPrefixCode, 4))
	firstSuffixCodeSpaceNums := countLeadingSpaces(replaceTabsWithSpaces(firstSuffixCode, 4))

	g.Log().Infof(context.Background(), "lastPrefixCodeSpaceNums: %d", lastPrefixCodeSpaceNums)
	g.Log().Infof(context.Background(), "firstSuffixCodeSpaceNums: %d", firstSuffixCodeSpaceNums)

	return lastPrefixCodeSpaceNums == firstSuffixCodeSpaceNums
}

// RemoveDuplicateElementSlice slice 元素去重
func RemoveDuplicateElementSlice[T comparable](rawSlice []T) (uniqSlice []T) {
	uniqSlice = make([]T, 0, len(rawSlice))
	uMap := map[T]struct{}{}
	for _, item := range rawSlice {
		if _, ok := uMap[item]; !ok {
			uMap[item] = struct{}{}
			uniqSlice = append(uniqSlice, item)
		}
	}
	return uniqSlice
}

// RandSliceValue 随机获取 slice 中的一个元素
func RandSliceValue(xs []string) string {
	return xs[rand.Intn(len(xs))]
}

// IsEqualIgnoreSliceOrder 判断两个 interface{} 是否相等，若为 slice，则忽略 slice 中元素的顺序
func IsEqualIgnoreSliceOrder(a, b interface{}) bool {
	return cmp.Equal(a, b, cmpopts.SortSlices(func(x, y interface{}) bool {
		return ToString(x) < ToString(y)
	}))
}

// GetDiffOfTwoSlice 获取两个 slice 的差集
func GetDiffOfTwoSlice[T comparable](a, b []T) (onlyInA []T, onlyInB []T) {
	if IsEqualIgnoreSliceOrder(a, b) {
		return nil, nil
	}
	for _, itemA := range a {
		found := false
		for _, itemB := range b {
			if IsEqualIgnoreSliceOrder(itemA, itemB) {
				found = true
				break
			}
		}
		if !found {
			onlyInA = append(onlyInA, itemA)
		}
	}

	for _, itemB := range b {
		found := false
		for _, itemA := range a {
			if IsEqualIgnoreSliceOrder(itemB, itemA) {
				found = true
				break
			}
		}
		if !found {
			onlyInB = append(onlyInB, itemB)
		}
	}

	return onlyInA, onlyInB
}

// GetIntersectionOfTwoSlice 获取两个 slice 的交集
func GetIntersectionOfTwoSlice[T comparable](a, b []T) []T {
	elementMap := make(map[T]bool)
	for _, v := range a {
		elementMap[v] = true
	}

	var intersection []T
	for _, v := range b {
		if elementMap[v] {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

// ToString 将 interface{} 转为 string
func ToString(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case null.String:
		if value.Valid {
			return value.String
		}
		return "null"
	case null.Int:
		if value.Valid {
			return strconv.FormatInt(value.Int64, 10)
		}
		return "null"
	case null.Float:
		if value.Valid {
			return strconv.FormatFloat(value.Float64, 'f', -1, 64)
		}
		return "null"
	case null.Bool:
		if value.Valid {
			return strconv.FormatBool(value.Bool)
		}
		return "null"
	case null.Time:
		if value.Valid {
			return value.Time.String()
		}
		return "null"
	default:
		return gconv.String(any)
	}
}

// GetDiffOfTwoCode
// 比较两段代码的增减行数
func GetDiffOfTwoCode(codeOne string, codeTwo string) int {
	//log.Infof(context.Background(), "codeOne: %s", codeOne)
	//log.Infof(context.Background(), "codeTwo: %s", codeTwo)

	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(codeOne, codeTwo, false)

	modifiedLines := 0

	for _, diff := range diffs {
		if diff.Type == diffmatchpatch.DiffInsert {
			modifiedLines++
		}
	}
	g.Log().Infof(context.Background(), "modifiedLines: %d", modifiedLines)
	return modifiedLines
}

// GenerateMD5
// 给定字符串，生成MD5
func GenerateMD5(s string) string {
	hash := md5.Sum([]byte(s))
	md5Hash := hex.EncodeToString(hash[:])
	return md5Hash
}
