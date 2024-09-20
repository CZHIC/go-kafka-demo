package util

import (
	"reflect"
	"sort"
)

//求并集
func Union(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func Intersect(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	nn := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func Difference(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	nn := make([]int64, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//通过反射将interface数据转换成数组
func InterfaceToArr(obj interface{}) []interface{} {
	val := reflect.ValueOf(obj) //获取reflect.Type类型
	kd := val.Kind()            //获取到a对应的类别

	if kd != reflect.Slice {
		return nil

	}
	//获取到该结构体有几个字段
	var ret []interface{}
	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		s := reflect.ValueOf(obj)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			ret = append(ret, ele.Interface())
		}
	}
	return ret
}

// 查找两个数组的异同(add,delete)
func FindDifferentList(old []int64, new []int64) ([]int64, []int64) {
	msrc := make(map[int64]byte) //按源数组建索引
	mall := make(map[int64]byte) //源+目所有元素建索引

	var set []int64 //交集

	//1.源数组建立map
	for _, v := range old {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range new {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) { //长度变化，即可以存
			l = len(mall)
		} else { //存不了，进并集
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除的，找不到的必定能在目数组中找到，即新加的
	var added, deleted []int64
	for v, _ := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}

	return added, deleted
}

// 查找两个数组的异同(add,delete)
func FindDifferentStringList(old []string, new []string) ([]string, []string) {
	msrc := make(map[string]byte) //按源数组建索引
	mall := make(map[string]byte) //源+目所有元素建索引

	var set []string //交集

	//1.源数组建立map
	for _, v := range old {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range new {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) { //长度变化，即可以存
			l = len(mall)
		} else { //存不了，进并集
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除的，找不到的必定能在目数组中找到，即新加的
	var added, deleted []string
	for v, _ := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}

	return added, deleted
}

//判断判断元素是否存在数组中[]int
func IsContainInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 数组去重
func RemoveDuplicateElement(element []int) []int {
	result := make([]int, 0, len(element))
	temp := map[int]struct{}{}
	for _, item := range element {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// map value 跑 排序
type Pair struct {
	Key   string
	Value int64
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func SortMapByValue(m map[string]int64) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

// map value 跑 排序
type MapInt64String struct {
	Key   int64
	Value string
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type MapInt64StringList []MapInt64String

func (p MapInt64StringList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p MapInt64StringList) Len() int           { return len(p) }
func (p MapInt64StringList) Less(i, j int) bool { return p[i].Key < p[j].Key }

// A function to turn a map into a PairList, then sort and return it.
func SortMapByValueMapInt64String(m map[int64]string) MapInt64StringList {
	p := make(MapInt64StringList, len(m))
	i := 0
	for k, v := range m {
		p[i] = MapInt64String{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
