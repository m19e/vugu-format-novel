package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

import (
	"strings"
	"syscall/js"
)

type RootData struct {
	Text   string
	Result string
}

const checkList = "　「『#"

func (data *RootData) HandleChange(event *vugu.DOMEvent) {
	data.Text = event.JSEvent().Get("target").Get("value").String()
}

func (data *RootData) Copy() {
	js.Global().Get("document").Call("getElementById", "result").Call("select")
	js.Global().Get("document").Call("execCommand", "Copy")
	js.Global().Get("window").Call("getSelection").Call("removeAllRanges")
}

func checkBOL(s string) bool {
	if s == "" {
		return false
	}
	if strings.Contains(checkList, string([]rune(s)[:1])) {
		return false
	}
	return true
}

func (data *RootData) Format() {
	var result string
	slice := strings.Split(data.Text, "\n")
	for _, str := range slice {
		if checkBOL(str) {
			str = "　" + str
		}
		result += str + "\n"
	}
	data.Result = result
}

var _ vugu.ComponentType = (*Root)(nil)

func (comp *Root) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*RootData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	css = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
	css.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    .my-first-vugu-comp {\n        background: #eee;\n    }\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)})
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "my-first-vugu-comp"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "container-fluid p-5"}}}
		parent.AppendChild(n)
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
			parent.AppendChild(n)
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "textarea", DataAtom: vugu.VGAtom(217608), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "text"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-control"}, vugu.VGAttribute{Namespace: "", Key: "placeholder", Val: "変換したいテキストを入力"}, vugu.VGAttribute{Namespace: "", Key: "style", Val: "height:35vh; resize:none;"}}}
				parent.AppendChild(n)
				// @change = { data.HandleChange(event) }
				{
					var i_ interface{} = data
					idat_ := reflect.ValueOf(&i_).Elem().InterfaceData()
					var i2_ interface{} = data.HandleChange
					i2dat_ := reflect.ValueOf(&i2_).Elem().InterfaceData()
					n.SetDOMEventHandler("change", vugu.DOMEventHandler{
						ReceiverAndMethodHash: uint64(idat_[0]) ^ uint64(idat_[1]) ^ uint64(i2dat_[0]) ^ uint64(i2dat_[1]),
						Method:                reflect.ValueOf(data).MethodByName("HandleChange"),
						Args:                  []interface{}{event},
					})
				}
				if false {
					// force compiler to check arguments for type safety
					data.HandleChange(event)
				}
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", DataAtom: vugu.VGAtom(102662), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn btn-primary btn-lg btn-block"}}}
			parent.AppendChild(n)
			// @click = { data.Format() }
			{
				var i_ interface{} = data
				idat_ := reflect.ValueOf(&i_).Elem().InterfaceData()
				var i2_ interface{} = data.Format
				i2dat_ := reflect.ValueOf(&i2_).Elem().InterfaceData()
				n.SetDOMEventHandler("click", vugu.DOMEventHandler{
					ReceiverAndMethodHash: uint64(idat_[0]) ^ uint64(idat_[1]) ^ uint64(i2dat_[0]) ^ uint64(i2dat_[1]),
					Method:                reflect.ValueOf(data).MethodByName("Format"),
					Args:                  []interface{}{},
				})
			}
			if false {
				// force compiler to check arguments for type safety
				data.Format()
			}
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            Format\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group my-3"}}}
			parent.AppendChild(n)
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "textarea", DataAtom: vugu.VGAtom(217608), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "result"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-control"}, vugu.VGAttribute{Namespace: "", Key: "style", Val: "height:40vh; resize:none;"}, vugu.VGAttribute{Namespace: "", Key: "readonly", Val: ""}}}
				parent.AppendChild(n)
				n.InnerHTML = fmt.Sprint(data.Result)
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", DataAtom: vugu.VGAtom(102662), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "type", Val: "button"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn btn-primary"}}}
			parent.AppendChild(n)
			// @click = { data.Copy() }
			{
				var i_ interface{} = data
				idat_ := reflect.ValueOf(&i_).Elem().InterfaceData()
				var i2_ interface{} = data.Copy
				i2dat_ := reflect.ValueOf(&i2_).Elem().InterfaceData()
				n.SetDOMEventHandler("click", vugu.DOMEventHandler{
					ReceiverAndMethodHash: uint64(idat_[0]) ^ uint64(idat_[1]) ^ uint64(i2dat_[0]) ^ uint64(i2dat_[1]),
					Method:                reflect.ValueOf(data).MethodByName("Copy"),
					Args:                  []interface{}{},
				})
			}
			if false {
				// force compiler to check arguments for type safety
				data.Copy()
			}
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            Copy\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "link", DataAtom: vugu.VGAtom(95236), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "rel", Val: "stylesheet"}, vugu.VGAttribute{Namespace: "", Key: "href", Val: "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type Root struct {}

func (ct *Root) NewData(props vugu.Props) (interface{}, error) { return &RootData{}, nil }

func init() { vugu.RegisterComponentType("root", &Root{}) }
