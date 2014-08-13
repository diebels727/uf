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
  return len(t)
}

func (t TestObject) ID() int {
  return t.id
}

// func TestInit(t *testing.T) {
//   var testObjects []TestObject
//   testObjects = make([]TestObject,2)
//   u := UnionFind{}
//   u.Init(TestObjects(testObjects))
//   if len(u.Clusters) != len(testObjects) {
//     t.Error("Expected ",len(testObjects)," got ",len(u.Clusters))
//   }
// }

func TestFind(t *testing.T) {
  var testObjects []TestObject
  var objectUnderTest TestObject

  testObjects = make([]TestObject,0)
  testObjects = append(testObjects,TestObject{10})
  testObjects = append(testObjects,TestObject{6})
  testObjects = append(testObjects,TestObject{1})
  u := UnionFind{}
  u.Init(TestObjects(testObjects))
  fmt.Println("Length of Clusters: ",len(u.Clusters))
  objectUnderTest = TestObject{6}
  fmt.Println("Pre find")
  object := u.find(objectUnderTest)
  fmt.Println("Post find")
  if object.ID() != objectUnderTest.ID() {
    t.Error("FAIL")
  }
}


