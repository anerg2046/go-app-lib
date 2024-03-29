package util

import (
	"fmt"
	"reflect"

	"github.com/goccy/go-json"
)

// 检查obj是否包含在target里
func Contain[T comparable](obj T, target []T) bool {
	for _, v := range target {
		if obj == v {
			return true
		}
	}
	return false
}

// 使用反射检查obj是否包含在target里
func DeepContain[T any](obj T, target []T) bool {
	for _, v := range target {
		if reflect.DeepEqual(obj, v) {
			return true
		}
	}
	return false
}

// Pretty 友好显示控制台输出数据
func Pretty(data any) {
	src, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(src))
}

// 判断指针是否为空
func IsNil(i any) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// 返回任意数据的指针
func Ptr[T any](i T) *T {
	return &i
}

// 从切片中删除某个元素
func RemoveItem[T comparable](a []T, elem T) []T {
	ret := make([]T, 0, len(a))
	for _, val := range a {
		if val != elem {
			ret = append(ret, val)
		}
	}
	return ret
}

// 切片去重
func RemoveDuplicates[T comparable](a []T) []T {
	ret := []T{}
	m := make(map[T]struct{}) //map的值不重要
	for _, v := range a {
		if _, ok := m[v]; !ok {
			ret = append(ret, v)
			m[v] = struct{}{}
		}
	}
	return ret
}

// 切片转为interface{}切片
func ToInterfaceSlice(slice any) []any {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]any, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// 切分切片
func SplitSlice[T any](slice []T, size int) [][]T {
	var chunk [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunk = append(chunk, slice[i:end])
	}
	return chunk
}

// 切分map
func ChunkMap[K comparable, V any](m map[K]V, size int) [][]V {
	var chunk [][]V
	for _, v := range m {
		chunk = append(chunk, []V{v})
		if len(chunk) == size {
			chunk = append(chunk, []V{})
		}
	}
	return chunk
}
