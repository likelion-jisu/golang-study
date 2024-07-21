package main

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main6() {
	var s Speaker
	s = Dog{}

	if _, ok := s.(Dog); ok {
		fmt.Println("s is a Dog")
	}

	parseTree, err := parse("5*10+20")
	if err != nil {
		panic(err)
	}
	result, err := walkTree(parseTree)
	fmt.Println(result, err)
}

// 선택적 인터페이스의 사용
func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	// reader가 WriteTo 메서드를 구현하는지 확인
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}

	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

	// 함수 하단 구현부...
	return 0, nil
}

func ctxDriverStmtExec(ctx context.Context, si driver.Stmt, nvdargs []driver.NamedValue) (driver.Result, error) {
	if siCtx, is := si.(driver.StmtExecContext); is {
		return siCtx.ExecContext(ctx, nvdargs)
	}
	// 대비책 코드...
	return nil, nil
}

// 래핑된 구현중 하나에 의해 구현된 선택적 인터페이스는 타입 단언, 타입 스위치로 검출이 되지 않는다.

type Speaker interface {
	Speak() string
	// Talk() string -> 주석 해제하면 컴파일 타임에서 타입 단언, 타입 스위치 불가
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type treeNode struct {
	val    treeVal
	lchild *treeNode
	rchild *treeNode
}

// treeVal defines an unexported marker interface that makes it clear
// which types can be assigned to val in treeNode
type treeVal interface {
	isToken()
}

type number int

func (number) isToken() {}

type operator func(int, int) int

func (operator) isToken() {}

func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

var operators = map[string]operator{
	"+": func(n1, n2 int) int {
		return n1 + n2
	},
	"-": func(n1, n2 int) int {
		return n1 - n2
	},
	"*": func(n1, n2 int) int {
		return n1 * n2
	},
	"/": func(n1, n2 int) int {
		return n1 / n2
	},
}

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("invalid expression")
	case number:
		// we know that t.val is of type number, so return the
		// int value
		return int(val), nil
	case operator:
		// we know that t.val is of type operator, so
		// find the values of the left and right children, then
		// call the process() method on operator to return the
		// result of processing their values.
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		// if a new treeVal type is defined, but walkTree wasn't updated
		// to process it, this detects it
		return 0, errors.New("unknown node type")
	}
}

func parse(s string) (*treeNode, error) {
	// not important for our example, so return something hard-coded
	return &treeNode{
		val: operators["+"],
		lchild: &treeNode{
			val:    operators["*"],
			lchild: &treeNode{val: number(5)},
			rchild: &treeNode{val: number(10)},
		},
		rchild: &treeNode{val: number(20)},
	}, nil
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
