package three
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at:
// 2017-08-08 16:38:51.987402183 +0100 BST
//
// using the following cmd:
// object3d_method_generator -typeName DirectionalLight -typeSlug directional_light

import "github.com/gopherjs/gopherjs/js"
	
func (obj *DirectionalLight) ApplyMatrix(matrix *Matrix4) {
	obj.Call("applyMatrix", matrix)
}

func (obj *DirectionalLight) Add(m Object3D) {
	obj.Object.Call("add", m.getInternalObject())
}

// func (obj *DirectionalLight) Copy() *DirectionalLight {
// 	return &DirectionalLight{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *DirectionalLight) ToJSON() interface{} {
	return obj.Object.Call("toJSON").Interface()
}

func (obj *DirectionalLight) getInternalObject() *js.Object {
	return obj.Object
}

