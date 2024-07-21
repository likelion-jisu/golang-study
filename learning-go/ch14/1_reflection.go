package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 타입, 종류

	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name(), xt.Kind())

	f := Foo{}
	ft := reflect.TypeOf(f)
	// Foo struct
	fmt.Println(ft.Name(), ft.Kind())

	// 슬라이스, 포인터 - 이름을 갖지 않는다.
	xpt := reflect.TypeOf(&x)
	//  ptr
	fmt.Println(xpt.Name(), xpt.Kind())
	// int int
	fmt.Println(xpt.Elem().Name(), xpt.Elem().Kind())

	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(),
			curField.Tag.Get("myTag"))
	}

	// 값

	vValue := reflect.ValueOf(x)
	fmt.Println(vValue.Int())

	s := []string{"a", "b", "c"}
	// reflect.Value 타입
	sv := reflect.ValueOf(s)
	// Interface : 비어있는 인터페이스(interface{})로 변수 값을 반환한다.
	// 타입 단언을 통해 확인한다.
	s2 := sv.Interface().([]string)
	fmt.Println(s2)

	i := 10
	// 포인터를 나타내는 reflect.Value
	iv := reflect.ValueOf(&i)
	// 포인터가 가리키는 값
	ivv := iv.Elem()
	ivv.SetInt(20)
	fmt.Println(i) // 20

	// 새로운 값 만들기
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	// slice []string
	fmt.Println(ssv.Kind(), ssv.Type())

	sv2 := reflect.New(stringType).Elem()
	// string string
	fmt.Println(sv2.Kind(), sv2.Type())
	sv2.SetString("hello")

	ssv = reflect.Append(ssv, sv2)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)

	var a interface{}
	fmt.Println(a == nil, hasNoValue(a)) // prints true true

	var b *int
	fmt.Println(b == nil, hasNoValue(b)) // prints true true

	var c interface{} = b
	fmt.Println(c == nil, hasNoValue(c)) // prints false true

	var d int
	fmt.Println(hasNoValue(d)) // prints false

	var e interface{} = d
	fmt.Println(e == nil, hasNoValue(e)) // prints false false

	// Marshal, Unmarshal

	data := `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith",42,false
Martha,37,"true"
`
	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var entries []MyData
	Unmarshal(allData, &entries)
	fmt.Println(entries)

	//now to turn entries into output
	out, err := Marshal(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)

	timed := MakeTimedFunction(timeMe).(func())
	timed()
	timedToo := MakeTimedFunction(timeMeToo).(func(int) int)
	fmt.Println(timedToo(2))
}

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

// changeInt, changeIntReflect는 같은 처리를 한다.
func changeInt(i *int) {
	*i = 20
}

func changeIntReflect(i *int) {
	iv := reflect.ValueOf(i)
	iv.Elem().SetInt(20)
}

// 문자열을 나타내는 reflect.Type
var stringType = reflect.TypeOf((*string)(nil)).Elem()

// 문자열 슬라이스를 나타내는 reflect.Type
var stringSliceType = reflect.TypeOf([]string(nil))

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	// reflect.Kind가 nil이 될 수 있는 것만 IsNil()을 호출할 수 있다.
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

// Marshal maps all of structs in a slice of structs to a slice of slice of strings.
// The first row written is the header with the column names.
// 슬라이스를 수정하지 않고 읽기만 하므로 구조체의 슬라이스를 가리키는 포인터가 아니다.
func Marshal(v interface{}) ([][]string, error) {
	sliceVal := reflect.ValueOf(v)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := sliceVal.Type().Elem() // 슬라이스 요소의 reflect.Type을 얻는다.
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}
	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)
	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshalHeader(vt reflect.Type) []string {
	var row []string
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

func marshalOne(vv reflect.Value) ([]string, error) {
	var row []string
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			row = append(row, strconv.FormatUint(fieldVal.Uint(), 10))
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
		}
	}
	return row, nil
}

// Unmarshal maps all of the rows of data in slice of slice of strings into a slice of structs.
// The first row is assumed to be the header with the column names.
// 파라미터에 저장된 값을 수정하므로 구조체 슬라이스를 가리키는 포인터를 전달해야 한다.
func Unmarshal(data [][]string, v interface{}) error {
	// 구조체 슬라이스 포인터를 reflect.Value로 변환한다.
	sliceValPtr := reflect.ValueOf(v)
	if sliceValPtr.Kind() != reflect.Ptr {
		return errors.New("must be a pointer to a slice of structs")
	}
	sliceVal := sliceValPtr.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("must be a pointer to a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("must be a pointer to a slice of structs")
	}

	// assume the first row is a header
	header := data[0]
	namePos := make(map[string]int, len(header))
	for k, v := range header {
		namePos[v] = k
	}

	for _, row := range data[1:] {
		newVal := reflect.New(structType).Elem()
		err := unmarshalOne(row, namePos, newVal)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}
	return nil
}

func unmarshalOne(row []string, namePos map[string]int, vv reflect.Value) error {
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		typeField := vt.Field(i)
		pos, ok := namePos[typeField.Tag.Get("csv")]
		if !ok {
			continue
		}
		val := row[pos]
		field := vv.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetUint(i)
		case reflect.String:
			field.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(b)
		default:
			return fmt.Errorf("cannot handle field of kind %v", field.Kind())
		}
	}
	return nil
}

type MyData struct {
	Name   string `csv:"name"`
	HasPet bool   `csv:"has_pet"`
	Age    int    `csv:"age"`
}

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func timeMe() {
	time.Sleep(1 * time.Second)
}

func timeMeToo(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}
