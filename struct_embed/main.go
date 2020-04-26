package main

import "fmt"

type KeyInfo struct {
	Id       string
	ParentId string
}
type Master struct {
	KeyInfo
	Name string
}

func (m *Master) SetParentId(id string) {
	m.ParentId = id
}

func NewMaster(id, name string) *Master {
	return &Master{
		KeyInfo: KeyInfo{
			Id: id,
		},
		Name: name,
	}
}

func main() {
	m := NewMaster("a123zz", "user1")
	fmt.Println(m)

	m.SetParentId("parent_123abc")
	fmt.Println(m)
}
