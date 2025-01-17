// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeStatusPage = coreglib.Type(C.adw_status_page_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStatusPage, F: marshalStatusPage},
	})
}

// StatusPageOverrides contains methods that are overridable.
type StatusPageOverrides struct {
}

func defaultStatusPageOverrides(v *StatusPage) StatusPageOverrides {
	return StatusPageOverrides{}
}

// StatusPage: page used for empty/error states and similar use-cases.
//
// <picture> <source srcset="status-page-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="status-page.png" alt="status-page"> </picture>
//
// The AdwStatusPage widget can have an icon, a title, a description and a
// custom widget which is displayed below them.
//
// # CSS nodes
//
// AdwStatusPage has a main CSS node with name statuspage.
//
// AdwStatusPage can use the .compact (style-classes.html#compact-status-page)
// style class for when it needs to fit into a small space such a sidebar or a
// popover.
type StatusPage struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*StatusPage)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StatusPage, *StatusPageClass, StatusPageOverrides](
		GTypeStatusPage,
		initStatusPageClass,
		wrapStatusPage,
		defaultStatusPageOverrides,
	)
}

func initStatusPageClass(gclass unsafe.Pointer, overrides StatusPageOverrides, classInitFunc func(*StatusPageClass)) {
	if classInitFunc != nil {
		class := (*StatusPageClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStatusPage(obj *coreglib.Object) *StatusPage {
	return &StatusPage{
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

func marshalStatusPage(p uintptr) (interface{}, error) {
	return wrapStatusPage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStatusPage creates a new AdwStatusPage.
//
// The function returns the following values:
//
//   - statusPage: newly created AdwStatusPage.
func NewStatusPage() *StatusPage {
	var _cret *C.GtkWidget // in

	_cret = C.adw_status_page_new()

	var _statusPage *StatusPage // out

	_statusPage = wrapStatusPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _statusPage
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//   - widget (optional): child widget of self.
func (self *StatusPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_status_page_get_child(_arg0)
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

// Description gets the description markup for self.
//
// The function returns the following values:
//
//   - utf8 (optional): description.
func (self *StatusPage) Description() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_status_page_get_description(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IconName gets the icon name for self.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name.
func (self *StatusPage) IconName() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_status_page_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Paintable gets the paintable for self.
//
// The function returns the following values:
//
//   - paintable (optional): paintable.
func (self *StatusPage) Paintable() *gdk.Paintable {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.GdkPaintable  // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_status_page_get_paintable(_arg0)
	runtime.KeepAlive(self)

	var _paintable *gdk.Paintable // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_paintable = &gdk.Paintable{
				Object: obj,
			}
		}
	}

	return _paintable
}

// Title gets the title for self.
//
// The function returns the following values:
//
//   - utf8: title.
func (self *StatusPage) Title() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_status_page_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//   - child (optional) widget.
func (self *StatusPage) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.adw_status_page_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetDescription sets the description markup for self.
//
// The description is displayed below the title. It is parsed as Pango markup.
//
// The function takes the following parameters:
//
//   - description (optional): description.
func (self *StatusPage) SetDescription(description string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if description != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(description)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_status_page_set_description(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(description)
}

// SetIconName sets the icon name for self.
//
// Changing this will set statuspage:paintable to NULL.
//
// The function takes the following parameters:
//
//   - iconName (optional): icon name.
func (self *StatusPage) SetIconName(iconName string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_status_page_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetPaintable sets the paintable for self.
//
// Changing this will set statuspage:icon-name to NULL.
//
// The function takes the following parameters:
//
//   - paintable (optional): paintable.
func (self *StatusPage) SetPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.GdkPaintable  // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	}

	C.adw_status_page_set_paintable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(paintable)
}

// SetTitle sets the title for self.
//
// The title is displayed below the icon. It is not parsed as Pango markup.
//
// The function takes the following parameters:
//
//   - title: title.
func (self *StatusPage) SetTitle(title string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_status_page_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// StatusPageClass: instance of this type is always passed by reference.
type StatusPageClass struct {
	*statusPageClass
}

// statusPageClass is the struct that's finalized.
type statusPageClass struct {
	native *C.AdwStatusPageClass
}

func (s *StatusPageClass) ParentClass() *gtk.WidgetClass {
	valptr := &s.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
