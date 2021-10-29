// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_view_switcher_policy_get_type()), F: marshalViewSwitcherPolicy},
		{T: externglib.Type(C.adw_view_switcher_get_type()), F: marshalViewSwitcherer},
	})
}

// ViewSwitcherPolicy describes the adaptive modes of adw.ViewSwitcher.
type ViewSwitcherPolicy C.gint

const (
	// ViewSwitcherPolicyAuto: automatically adapt to the best fitting mode.
	ViewSwitcherPolicyAuto ViewSwitcherPolicy = iota
	// ViewSwitcherPolicyNarrow: force the narrow mode.
	ViewSwitcherPolicyNarrow
	// ViewSwitcherPolicyWide: force the wide mode.
	ViewSwitcherPolicyWide
)

func marshalViewSwitcherPolicy(p uintptr) (interface{}, error) {
	return ViewSwitcherPolicy(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ViewSwitcherPolicy.
func (v ViewSwitcherPolicy) String() string {
	switch v {
	case ViewSwitcherPolicyAuto:
		return "Auto"
	case ViewSwitcherPolicyNarrow:
		return "Narrow"
	case ViewSwitcherPolicyWide:
		return "Wide"
	default:
		return fmt.Sprintf("ViewSwitcherPolicy(%d)", v)
	}
}

// ViewSwitcher: adaptive view switcher.
//
// An adaptive view switcher designed to switch between multiple views contained
// in a adw.ViewStack in a similar fashion to gtk.StackSwitcher.
//
// Depending on the available width, the view switcher can adapt from a wide
// mode showing the view's icon and title side by side, to a narrow mode showing
// the view's icon and title one on top of the other, in a more compact way.
// This can be controlled via the adw.ViewSwitcher:policy property.
//
//
// CSS nodes
//
// AdwViewSwitcher has a single CSS node with name viewswitcher.
//
//
// Accessibility
//
// AdwViewSwitcher uses the GTK_ACCESSIBLE_ROLE_TAB_LIST role and uses the
// GTK_ACCESSIBLE_ROLE_TAB for its buttons.
type ViewSwitcher struct {
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*ViewSwitcher)(nil)
)

func wrapViewSwitcher(obj *externglib.Object) *ViewSwitcher {
	return &ViewSwitcher{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalViewSwitcherer(p uintptr) (interface{}, error) {
	return wrapViewSwitcher(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewViewSwitcher creates a new AdwViewSwitcher.
func NewViewSwitcher() *ViewSwitcher {
	var _cret *C.GtkWidget // in

	_cret = C.adw_view_switcher_new()

	var _viewSwitcher *ViewSwitcher // out

	_viewSwitcher = wrapViewSwitcher(externglib.Take(unsafe.Pointer(_cret)))

	return _viewSwitcher
}

// NarrowEllipsize gets the ellipsizing position for the titles.
func (self *ViewSwitcher) NarrowEllipsize() pango.EllipsizeMode {
	var _arg0 *C.AdwViewSwitcher   // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_get_narrow_ellipsize(_arg0)
	runtime.KeepAlive(self)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// Policy gets the policy of self.
func (self *ViewSwitcher) Policy() ViewSwitcherPolicy {
	var _arg0 *C.AdwViewSwitcher      // out
	var _cret C.AdwViewSwitcherPolicy // in

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_get_policy(_arg0)
	runtime.KeepAlive(self)

	var _viewSwitcherPolicy ViewSwitcherPolicy // out

	_viewSwitcherPolicy = ViewSwitcherPolicy(_cret)

	return _viewSwitcherPolicy
}

// Stack gets the stack controlled by self.
func (self *ViewSwitcher) Stack() *ViewStack {
	var _arg0 *C.AdwViewSwitcher // out
	var _cret *C.AdwViewStack    // in

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_get_stack(_arg0)
	runtime.KeepAlive(self)

	var _viewStack *ViewStack // out

	if _cret != nil {
		_viewStack = wrapViewStack(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _viewStack
}

// SetNarrowEllipsize sets the ellipsizing position for the titles.
//
// The function takes the following parameters:
//
//    - mode: new value.
//
func (self *ViewSwitcher) SetNarrowEllipsize(mode pango.EllipsizeMode) {
	var _arg0 *C.AdwViewSwitcher   // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.adw_view_switcher_set_narrow_ellipsize(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(mode)
}

// SetPolicy sets the policy of self.
//
// The function takes the following parameters:
//
//    - policy: new policy.
//
func (self *ViewSwitcher) SetPolicy(policy ViewSwitcherPolicy) {
	var _arg0 *C.AdwViewSwitcher      // out
	var _arg1 C.AdwViewSwitcherPolicy // out

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwViewSwitcherPolicy(policy)

	C.adw_view_switcher_set_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetStack sets the stack controlled by self.
//
// The function takes the following parameters:
//
//    - stack: stack.
//
func (self *ViewSwitcher) SetStack(stack *ViewStack) {
	var _arg0 *C.AdwViewSwitcher // out
	var _arg1 *C.AdwViewStack    // out

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))
	if stack != nil {
		_arg1 = (*C.AdwViewStack)(unsafe.Pointer(stack.Native()))
	}

	C.adw_view_switcher_set_stack(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(stack)
}
