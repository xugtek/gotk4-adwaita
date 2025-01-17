// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeBin = coreglib.Type(C.adw_bin_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBin, F: marshalBin},
	})
}

// BinOverrides contains methods that are overridable.
type BinOverrides struct {
}

func defaultBinOverrides(v *Bin) BinOverrides {
	return BinOverrides{}
}

// Bin: widget with one child.
//
// <picture> <source srcset="bin-dark.png" media="(prefers-color-scheme: dark)">
// <img src="bin.png" alt="bin"> </picture>
//
// The AdwBin widget has only one child, set with the bin:child property.
//
// It is useful for deriving subclasses, since it provides common code needed
// for handling a single child widget.
type Bin struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*Bin)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Bin, *BinClass, BinOverrides](
		GTypeBin,
		initBinClass,
		wrapBin,
		defaultBinOverrides,
	)
}

func initBinClass(gclass unsafe.Pointer, overrides BinOverrides, classInitFunc func(*BinClass)) {
	if classInitFunc != nil {
		class := (*BinClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBin(obj *coreglib.Object) *Bin {
	return &Bin{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalBin(p uintptr) (interface{}, error) {
	return wrapBin(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBin creates a new AdwBin.
//
// The function returns the following values:
//
//   - bin: new created AdwBin.
func NewBin() *Bin {
	var _cret *C.GtkWidget // in

	_cret = C.adw_bin_new()

	var _bin *Bin // out

	_bin = wrapBin(coreglib.Take(unsafe.Pointer(_cret)))

	return _bin
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//   - widget (optional): child widget of self.
func (self *Bin) Child() gtk.Widgetter {
	var _arg0 *C.AdwBin    // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwBin)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_bin_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//   - child (optional) widget.
func (self *Bin) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwBin    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwBin)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.adw_bin_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// BinClass: instance of this type is always passed by reference.
type BinClass struct {
	*binClass
}

// binClass is the struct that's finalized.
type binClass struct {
	native *C.AdwBinClass
}

func (b *BinClass) ParentClass() *gtk.WidgetClass {
	valptr := &b.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
