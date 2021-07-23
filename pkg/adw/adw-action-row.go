// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_action_row_get_type()), F: marshalActionRower},
	})
}

// ActionRowOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionRowOverrider interface {
	Activate()
}

type ActionRow struct {
	PreferencesRow
}

func wrapActionRow(obj *externglib.Object) *ActionRow {
	return &ActionRow{
		PreferencesRow: PreferencesRow{
			ListBoxRow: gtk.ListBoxRow{
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
				Actionable: gtk.Actionable{
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
				},
				Object: obj,
			},
		},
	}
}

func marshalActionRower(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActionRow(obj), nil
}

// NewActionRow creates a new ActionRow.
func NewActionRow() *ActionRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_action_row_new()

	var _actionRow *ActionRow // out

	_actionRow = wrapActionRow(externglib.Take(unsafe.Pointer(_cret)))

	return _actionRow
}

func (self *ActionRow) Activate() {
	var _arg0 *C.AdwActionRow // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	C.adw_action_row_activate(_arg0)
}

// AddPrefix adds a prefix widget to self.
func (self *ActionRow) AddPrefix(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.adw_action_row_add_prefix(_arg0, _arg1)
}

// AddSuffix adds a suffix widget to self.
func (self *ActionRow) AddSuffix(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.adw_action_row_add_suffix(_arg0, _arg1)
}

// ActivatableWidget gets the widget activated when self is activated.
func (self *ActionRow) ActivatableWidget() gtk.Widgetter {
	var _arg0 *C.AdwActionRow // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_activatable_widget(_arg0)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// IconName gets the icon name for self.
func (self *ActionRow) IconName() string {
	var _arg0 *C.AdwActionRow // out
	var _cret *C.char         // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Subtitle gets the subtitle for self.
func (self *ActionRow) Subtitle() string {
	var _arg0 *C.AdwActionRow // out
	var _cret *C.char         // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_subtitle(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SubtitleLines gets the number of lines at the end of which the subtitle label
// will be ellipsized. If the value is 0, the number of lines won't be limited.
func (self *ActionRow) SubtitleLines() int {
	var _arg0 *C.AdwActionRow // out
	var _cret C.int           // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_subtitle_lines(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TitleLines gets the number of lines at the end of which the title label will
// be ellipsized. If the value is 0, the number of lines won't be limited.
func (self *ActionRow) TitleLines() int {
	var _arg0 *C.AdwActionRow // out
	var _cret C.int           // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_title_lines(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UseUnderline gets whether an embedded underline in the text of the title and
// subtitle labels indicates a mnemonic. See adw_action_row_set_use_underline().
func (self *ActionRow) UseUnderline() bool {
	var _arg0 *C.AdwActionRow // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remove removes a child from self.
func (self *ActionRow) Remove(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.adw_action_row_remove(_arg0, _arg1)
}

// SetActivatableWidget sets the widget to activate when self is activated,
// either by clicking on it, by calling adw_action_row_activate(), or via
// mnemonics in the title or the subtitle. See the “use_underline” property to
// enable mnemonics.
//
// The target widget will be activated by emitting the
// GtkWidget::mnemonic-activate signal on it.
func (self *ActionRow) SetActivatableWidget(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.adw_action_row_set_activatable_widget(_arg0, _arg1)
}

// SetIconName sets the icon name for self.
func (self *ActionRow) SetIconName(iconName string) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.char         // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_action_row_set_icon_name(_arg0, _arg1)
}

// SetSubtitle sets the subtitle for self.
func (self *ActionRow) SetSubtitle(subtitle string) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.char         // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	if subtitle != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_action_row_set_subtitle(_arg0, _arg1)
}

// SetSubtitleLines sets the number of lines at the end of which the subtitle
// label will be ellipsized. If the value is 0, the number of lines won't be
// limited.
func (self *ActionRow) SetSubtitleLines(subtitleLines int) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 C.int           // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(subtitleLines)

	C.adw_action_row_set_subtitle_lines(_arg0, _arg1)
}

// SetTitleLines sets the number of lines at the end of which the title label
// will be ellipsized. If the value is 0, the number of lines won't be limited.
func (self *ActionRow) SetTitleLines(titleLines int) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 C.int           // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(titleLines)

	C.adw_action_row_set_title_lines(_arg0, _arg1)
}

// SetUseUnderline: if true, an underline in the text of the title and subtitle
// labels indicates the next character should be used for the mnemonic
// accelerator key.
func (self *ActionRow) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_action_row_set_use_underline(_arg0, _arg1)
}
