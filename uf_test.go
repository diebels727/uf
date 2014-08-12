package uf

import (
  "testing"
  "fmt"
)


type TestObject struct {
  id int
}

type TestObjects []TestObject

func (t TestObjects) Len() int {
  return 1
}

func (t *TestObject) ID() int {
  return t.id
}

func TestGeneralUnion(t *testing.T) {
  var testObjects []TestObject
  testObjects = make([]TestObject,1)
  u := UnionFind{}
  u.Init(TestObjects(testObjects))
  // o := &TestObject{1}

  fmt.Println("done.")
}

