package uf

import (
  "testing"
  // "fmt"
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

func TestInit(t *testing.T) {
  var testObjects []TestObject
  testObjects = make([]TestObject,2)
  u := UnionFind{}
  u.Init(TestObjects(testObjects))
  if len(u.Clusters) != len(testObjects) {
    t.Error("Expected ",len(testObjects)," got ",len(u.Clusters))
  }
}

func TestFind(t *testing.T) {
  var testObjects []TestObject
  var objectUnderTest TestObject

  testObjects = make([]TestObject,0)
  testObjects = append(testObjects,TestObject{10})
  testObjects = append(testObjects,TestObject{6})
  testObjects = append(testObjects,TestObject{1})
  u := UnionFind{}
  u.Init(TestObjects(testObjects))
  objectUnderTest = TestObject{6}
  clusterID := u.Find(objectUnderTest)
  if clusterID != objectUnderTest.ID() {
    t.Error("FAIL")
  }
}

func TestFindIdentityShouldRetrieveTheSameObject(t *testing.T) {
  var testObjects []TestObject
  p := TestObject{1}
  testObjects = make([]TestObject,0)
  testObjects = append(testObjects,p)
  u := UnionFind{}
  u.Init(TestObjects(testObjects))
  if p.ID() != u.Find(p) {
    t.Error("FAIL")
  }
}

func TestClusterInitializationShouldBeDistinctOnInit(t *testing.T) {

}

func TestTrivialUnionShouldConnectBothObjects(t *testing.T) {
  var testObjects []TestObject
  p := TestObject{10}
  q := TestObject{1}
  testObjects = make([]TestObject,0)
  testObjects = append(testObjects,p)
  testObjects = append(testObjects,q)
  u := UnionFind{}
  u.Init(TestObjects(testObjects))
  u.Union(p,q)
  if u.Find(p) != u.Find(q) {
    t.Error("Expected ",u.Find(p)," to equal ",u.Find(q))
  }
}

func TestComplexUnionShouldConnectObjects(t *testing.T) {
  var testObjects []TestObject
  p := TestObject{10}
  q := TestObject{1}
  r := TestObject{7}
  s := TestObject{9}
  v := TestObject{11}
  w := TestObject{14}
  testObjects = make([]TestObject,0)
  testObjects = append(testObjects,p)
  testObjects = append(testObjects,q)
  testObjects = append(testObjects,r)
  testObjects = append(testObjects,s)
  testObjects = append(testObjects,v)
  testObjects = append(testObjects,w)
  u := UnionFind{}
  u.Init(TestObjects(testObjects))
  u.Union(p,q)
  if u.Find(p) != u.Find(q) {
    t.Error("Expected ",u.Find(p)," to equal ",u.Find(q))
  }
  u.Union(r,s)
  if u.Find(r) != u.Find(s) {
    t.Error("Expected ",u.Find(r)," to equal ",u.Find(s))
  }
  u.Union(v,w)
  if u.Find(v) != u.Find(w) {
    t.Error("Expected ",u.Find(v)," to equal ",u.Find(w))
  }


}




