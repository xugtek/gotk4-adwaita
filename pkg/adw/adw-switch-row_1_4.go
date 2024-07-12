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
	GTypeSwitchRow = coreglib.Type(C.adw_switch_row_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSwitchRow, F: marshalSwitchRow},
	})
}

// SwitchRowOverrides contains methods that are overridable.
type SwitchRowOverrides struct {
}

func defaultSwitchRowOverrides(v *SwitchRow) SwitchRowOverrides {
	return SwitchRowOverrides{}
}

// SwitchRow: gtk.ListBoxRow used to represent two states.
//
// <picture> <source srcset="switch-row-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="switch-row.png" alt="switch-row"> </picture>
//
// The AdwSwitchRow widget contains a gtk.Switch that allows the user to select
// between two states: "on" or "off". When activated, the row will invert its
// active state.
//
// The user can control the switch by activating the row or by dragging on the
// switch handle.
//
// See gtk.Switch for details.
//
// Example of an AdwSwitchRow UI definition:
//
//	<object class="AdwSwitchRow">
//	  <property name="title" translatable="yes">Switch Row</property>
//	  <signal name="notify::active" handler="switch_row_notify_active_cb"/>
//	</object>
//
// The switchrow:active property should be connected to in order to monitor
// changes to the active state.
type SwitchRow struct {
	_ [0]func() // equal guard
	ActionRow
}

var (
	_ gtk.Widgetter     = (*SwitchRow)(nil)
	_ coreglib.Objector = (*SwitchRow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SwitchRow, *SwitchRowClass, SwitchRowOverrides](
		GTypeSwitchRow,
		initSwitchRowClass,
		wrapSwitchRow,
		defaultSwitchRowOverrides,
	)
}

func initSwitchRowClass(gclass unsafe.Pointer, overrides SwitchRowOverrides, classInitFunc func(*SwitchRowClass)) {
	if classInitFunc != nil {
		class := (*SwitchRowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSwitchRow(obj *coreglib.Object) *SwitchRow {
	return &SwitchRow{
		ActionRow: ActionRow{
			PreferencesRow: PreferencesRow{
				ListBoxRow: gtk.ListBoxRow{
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
					Object: obj,
					Actionable: gtk.Actionable{
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
					},
				},
			},
		},
	}
}

func marshalSwitchRow(p uintptr) (interface{}, error) {
	return wrapSwitchRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSwitchRow creates a new AdwSwitchRow.
//
// The function returns the following values:
//
//   - switchRow: newly created AdwSwitchRow.
func NewSwitchRow() *SwitchRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_switch_row_new()

	var _switchRow *SwitchRow // out

	_switchRow = wrapSwitchRow(coreglib.Take(unsafe.Pointer(_cret)))

	return _switchRow
}

// Active gets whether self is in its "on" or "off" position.
//
// The function returns the following values:
//
//   - ok: whether self is active or not.
func (self *SwitchRow) Active() bool {
	var _arg0 *C.AdwSwitchRow // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwSwitchRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_switch_row_get_active(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets whether self is in its "on" or "off" position.
//
// The function takes the following parameters:
//
//   - isActive: whether self should be active.
func (self *SwitchRow) SetActive(isActive bool) {
	var _arg0 *C.AdwSwitchRow // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwSwitchRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.adw_switch_row_set_active(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(isActive)
}
