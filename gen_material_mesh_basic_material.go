package three
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at:
// 2017-08-08 16:38:52.20564721 +0100 BST
//
// using the following cmd:
// material_method_generator -materialName MeshBasicMaterial -materialSlug mesh_basic_material

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)
	
func (m MeshBasicMaterial) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m MeshBasicMaterial) SetValues(values MaterialParameters) {
	m.Call("setValues", structs.Map(values))
}

func (m MeshBasicMaterial) ToJSON(meta interface{}) interface{} {
	return m.Call("toJSON", meta)
}

func (m MeshBasicMaterial) Clone() {
	m.Call("clone")
}

func (m MeshBasicMaterial) Copy(source Object3D) {
	m.Call("copy", source)
}

func (m MeshBasicMaterial) Dispose() {
	m.Call("dispose")
}

func (m MeshBasicMaterial) getInternalObject() *js.Object {
	return m.Object
}

