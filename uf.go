package uf

import (
  "reflect"
  // "fmt"
)

type Pair interface {
  P() Object
  Q() Object
}

type Object interface {
  ID() int
}

type Objects interface {
  Len() int
}

type ClusterMap map[Object]int

// func (objects *Objects) Len() int {
//   return len([]Object(*objects));
// }

type UnionFind struct{
  Clusters []Object
  ClusterMap ClusterMap
}

func (u *UnionFind) Init(objects Objects) {
  u.Clusters = make([]Object,objects.Len())
  u.ClusterMap = make(ClusterMap)

  switch reflect.TypeOf(objects).Kind() {
  case reflect.Slice:
    o := reflect.ValueOf(objects)
    for i := 0; i < o.Len(); i++ {
      object := o.Index(i).Interface().(Object)
      u.ClusterMap[object] = object.ID()
    }
  }

  // for id,object := range(objects) {
  //   u.ClusterMap[&object] = id
  // }
}

func (u *UnionFind) find(object Object) Object {
  return u.Clusters[(object).ID()]
}

// func connected(pair Pair,ids []int) bool {
//   return ids[pair.P] == ids[pair.Q]
// }
//
// func union(pair Pair,ids []int) []int {
//   logger.Printf("[union] called")
//   if (connected(pair,ids)) { return ids };
//   pid := ids[pair.P];
//   for i:=0;i<len(ids);i++ {
//       if (ids[i] == pid) { ids[i] = ids[pair.Q] };
//   }
//   logger.Printf("[union] finished")
//   return ids
// }