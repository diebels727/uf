package uf

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

// func (objects *Objects) Len() int {
//   return len([]Object(*objects));
// }


type UnionFind struct{
  objects []Object
}

func (u *UnionFind) Init(objects Objects) {
  u.objects = make([]Object,objects.Len())
  // (*u).objects = make([]Object,0)
}

func (u *UnionFind) find(object Object) Object {
  return u.objects[object.ID()]
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