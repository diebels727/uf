package uf

import (
  "reflect"
  "fmt"
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

type UnionFind struct{
  Clusters []int
  ClusterMap ClusterMap
}

func (u *UnionFind) Init(objects Objects) {
  u.Clusters = make([]int,objects.Len())
  u.ClusterMap = make(ClusterMap)

  switch reflect.TypeOf(objects).Kind() {
  case reflect.Slice:
    o := reflect.ValueOf(objects)
    for i := 0; i < o.Len(); i++ {
      object := o.Index(i).Interface().(Object)
      u.ClusterMap[object] = i
      u.Clusters[i] = object.ID()
    }
  }
}

func (u *UnionFind) Find(object Object) int {
  objectPositionInClusters := u.ClusterMap[object]
  return u.Clusters[objectPositionInClusters]
}

func (u *UnionFind) Connected(p Object,q Object) bool {
  pPositionInClusters := u.ClusterMap[p]
  qPositionInClusters := u.ClusterMap[q]
  return u.Clusters[pPositionInClusters] == u.Clusters[qPositionInClusters]
}

func (u *UnionFind) Union(p Object,q Object) {
  if u.Connected(p,q) {
    return
  }
  //find all q's and switch all q clusters to p's
  pClusterID := u.Find(p)
  qClusterID := u.Find(q)
  for i,clusterID := range u.Clusters {
    if clusterID == qClusterID {
      u.Clusters[i] = pClusterID
    }
  }
}

func (u *UnionFind) PrettyPrintClusterIDs() {
  for index,id := range u.Clusters {
    fmt.Println("Index in Clusters: ",index," Cluster ID: ",id)
  }
}