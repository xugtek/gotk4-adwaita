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
	GTypeOverlaySplitView = coreglib.Type(C.adw_overlay_split_view_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOverlaySplitView, F: marshalOverlaySplitView},
	})
}

// OverlaySplitViewOverrides contains methods that are overridable.
type OverlaySplitViewOverrides struct {
}

func defaultOverlaySplitViewOverrides(v *OverlaySplitView) OverlaySplitViewOverrides {
	return OverlaySplitViewOverrides{}
}

// OverlaySplitView: widget presenting sidebar and content side by side or as an
// overlay.
//
// <picture> <source srcset="overlay-split-view-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="overlay-split-view.png"
// alt="overlay-split-view"> </picture> <picture> <source
// srcset="overlay-split-view-collapsed-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="overlay-split-view-collapsed.png"
// alt="overlay-split-view-collapsed"> </picture>
//
// AdwOverlaySplitView has two children: sidebar and content, and displays them
// side by side.
//
// When overlaysplitview:collapsed is set to TRUE, the sidebar is instead shown
// as an overlay above the content widget.
//
// The sidebar can be hidden or shown using the overlaysplitview:show-sidebar
// property.
//
// Sidebar can be displayed before or after the content, this can be controlled
// with the overlaysplitview:sidebar-position property.
//
// Collapsing the split view automatically hides the sidebar widget,
// and uncollapsing it shows the sidebar. If this behavior is not desired,
// the overlaysplitview:pin-sidebar property can be used to override it.
//
// AdwOverlaySplitView supports an edge swipe gesture for showing the sidebar,
// and a swipe from the sidebar for hiding it. Gestures are only supported
// on touchscreen, but not touchpad. Gestures can be controlled with the
// overlaysplitview:enable-show-gesture and overlaysplitview:enable-hide-gesture
// properties.
//
// See also navigationsplitview.
//
// AdwOverlaySplitView is typically used together with an breakpoint setting the
// collapsed property to TRUE on small widths, as follows:
//
//	<object class="AdwWindow">
//	  <property name="width-request">360</property>
//	  <property name="height-request">200</property>
//	  <property name="default-width">800</property>
//	  <property name="default-height">800</property>
//	  <child>
//	    <object class="AdwBreakpoint">
//	      <condition>max-width: 400sp</condition>
//	      <setter object="split_view" property="collapsed">True</setter>
//	    </object>
//	  </child>
//	  <property name="content">
//	    <object class="AdwOverlaySplitView" id="split_view">
//	      <property name="sidebar">
//	        <!-- ... -->
//	      </property>
//	      <property name="content">
//	        <!-- ... -->
//	      </property>
//	    </object>
//	  </property>
//	</object>
//
// AdwOverlaySplitView is often used for implementing the utility pane
// (https://developer.gnome.org/hig/patterns/containers/utility-panes.html)
// pattern.
//
// # Sizing
//
// When not collapsed, AdwOverlaySplitView changes the sidebar width depending
// on its own width.
//
// If possible, it tries to allocate a fraction of the total width, controlled
// with the overlaysplitview:sidebar-width-fraction property.
//
// The sidebar also has minimum and maximum sizes, controlled with the
// overlaysplitview:min-sidebar-width and overlaysplitview:max-sidebar-width
// properties.
//
// The minimum and maximum sizes are using the length unit specified with the
// overlaysplitview:sidebar-width-unit.
//
// By default, sidebar is using 25% of the total width, with 180sp as the
// minimum size and 280sp as the maximum size.
//
// When collapsed, the preferred width fraction is ignored and the sidebar uses
// overlaysplitview:max-sidebar-width when possible.
//
// # Header Bar Integration
//
// When used inside AdwOverlaySplitView, headerbar will automatically hide the
// window buttons in the middle.
//
// # AdwOverlaySplitView as GtkBuildable
//
// The AdwOverlaySplitView implementation of the gtk.Buildable interface
// supports setting the sidebar widget by specifying “sidebar” as the “type”
// attribute of a <child> element, Specifying “content” child type or omitting
// it results in setting the content widget.
//
// # CSS nodes
//
// AdwOverlaySplitView has a single CSS node with the name overlay-split-view.
//
// It contains two nodes with the name widget, containing the sidebar and
// content children.
//
// When not collapsed, they have the .sidebar-view and .content-view style
// classes respectively.
//
//	overlay-split-view
//	├── widget.sidebar-pane
//	│   ╰── [sidebar child]
//	╰── widget.content-pane
//	    ╰── [content child]
//
// When collapsed, the one containing the sidebar child has the .background
// style class and the other one has no style classes.
//
//	overlay-split-view
//	├── widget.background
//	│   ╰── [sidebar child]
//	╰── widget
//	    ╰── [content child]
//
// # Accessibility
//
// AdwOverlaySplitView uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type OverlaySplitView struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	Swipeable
}

var (
	_ gtk.Widgetter     = (*OverlaySplitView)(nil)
	_ coreglib.Objector = (*OverlaySplitView)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*OverlaySplitView, *OverlaySplitViewClass, OverlaySplitViewOverrides](
		GTypeOverlaySplitView,
		initOverlaySplitViewClass,
		wrapOverlaySplitView,
		defaultOverlaySplitViewOverrides,
	)
}

func initOverlaySplitViewClass(gclass unsafe.Pointer, overrides OverlaySplitViewOverrides, classInitFunc func(*OverlaySplitViewClass)) {
	if classInitFunc != nil {
		class := (*OverlaySplitViewClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapOverlaySplitView(obj *coreglib.Object) *OverlaySplitView {
	return &OverlaySplitView{
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
		Swipeable: Swipeable{
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
	}
}

func marshalOverlaySplitView(p uintptr) (interface{}, error) {
	return wrapOverlaySplitView(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewOverlaySplitView creates a new AdwOverlaySplitView.
//
// The function returns the following values:
//
//   - overlaySplitView: newly created AdwOverlaySplitView.
func NewOverlaySplitView() *OverlaySplitView {
	var _cret *C.GtkWidget // in

	_cret = C.adw_overlay_split_view_new()

	var _overlaySplitView *OverlaySplitView // out

	_overlaySplitView = wrapOverlaySplitView(coreglib.Take(unsafe.Pointer(_cret)))

	return _overlaySplitView
}

// Collapsed gets whether self is collapsed.
//
// The function returns the following values:
//
//   - ok: whether self is collapsed.
func (self *OverlaySplitView) Collapsed() bool {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.gboolean             // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_collapsed(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Content gets the content widget for self.
//
// The function returns the following values:
//
//   - widget (optional): content widget for self.
func (self *OverlaySplitView) Content() gtk.Widgetter {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_content(_arg0)
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

// EnableHideGesture gets whether self can be closed with a swipe gesture.
//
// The function returns the following values:
//
//   - ok: TRUE if self can be closed with a swipe gesture.
func (self *OverlaySplitView) EnableHideGesture() bool {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.gboolean             // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_enable_hide_gesture(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EnableShowGesture gets whether self can be opened with an edge swipe gesture.
//
// The function returns the following values:
//
//   - ok: TRUE if self can be opened with a swipe gesture.
func (self *OverlaySplitView) EnableShowGesture() bool {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.gboolean             // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_enable_show_gesture(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxSidebarWidth gets the maximum sidebar width for self.
//
// The function returns the following values:
//
//   - gdouble: maximum width.
func (self *OverlaySplitView) MaxSidebarWidth() float64 {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.double               // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_max_sidebar_width(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinSidebarWidth gets the minimum sidebar width for self.
//
// The function returns the following values:
//
//   - gdouble: minimum width.
func (self *OverlaySplitView) MinSidebarWidth() float64 {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.double               // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_min_sidebar_width(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PINSidebar gets whether the sidebar widget is pinned for self.
//
// The function returns the following values:
//
//   - ok: whether if the sidebar widget is pinned.
func (self *OverlaySplitView) PINSidebar() bool {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.gboolean             // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_pin_sidebar(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowSidebar gets whether the sidebar widget is shown for self.
//
// The function returns the following values:
//
//   - ok: TRUE if the sidebar widget is shown.
func (self *OverlaySplitView) ShowSidebar() bool {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.gboolean             // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_show_sidebar(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sidebar gets the sidebar widget for self.
//
// The function returns the following values:
//
//   - widget (optional): sidebar widget for self.
func (self *OverlaySplitView) Sidebar() gtk.Widgetter {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_sidebar(_arg0)
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

// SidebarPosition gets the sidebar position for self.
//
// The function returns the following values:
//
//   - packType: sidebar position for self.
func (self *OverlaySplitView) SidebarPosition() gtk.PackType {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.GtkPackType          // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_sidebar_position(_arg0)
	runtime.KeepAlive(self)

	var _packType gtk.PackType // out

	_packType = gtk.PackType(_cret)

	return _packType
}

// SidebarWidthFraction gets the preferred sidebar width fraction for self.
//
// The function returns the following values:
//
//   - gdouble: preferred width fraction.
func (self *OverlaySplitView) SidebarWidthFraction() float64 {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.double               // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_sidebar_width_fraction(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SidebarWidthUnit gets the length unit for minimum and maximum sidebar widths.
//
// The function returns the following values:
//
//   - lengthUnit: length unit.
func (self *OverlaySplitView) SidebarWidthUnit() LengthUnit {
	var _arg0 *C.AdwOverlaySplitView // out
	var _cret C.AdwLengthUnit        // in

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_overlay_split_view_get_sidebar_width_unit(_arg0)
	runtime.KeepAlive(self)

	var _lengthUnit LengthUnit // out

	_lengthUnit = LengthUnit(_cret)

	return _lengthUnit
}

// SetCollapsed sets whether self view is collapsed.
//
// When collapsed, the sidebar widget is presented as an overlay above the
// content widget, otherwise they are displayed side by side.
//
// The function takes the following parameters:
//
//   - collapsed: whether self is collapsed.
func (self *OverlaySplitView) SetCollapsed(collapsed bool) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if collapsed {
		_arg1 = C.TRUE
	}

	C.adw_overlay_split_view_set_collapsed(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(collapsed)
}

// SetContent sets the content widget for self.
//
// The function takes the following parameters:
//
//   - content (optional) widget.
func (self *OverlaySplitView) SetContent(content gtk.Widgetter) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if content != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(content).Native()))
	}

	C.adw_overlay_split_view_set_content(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(content)
}

// SetEnableHideGesture sets whether self can be closed with a swipe gesture.
//
// Only touchscreen swipes are supported.
//
// The function takes the following parameters:
//
//   - enableHideGesture: whether self can be closed with a swipe gesture.
func (self *OverlaySplitView) SetEnableHideGesture(enableHideGesture bool) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enableHideGesture {
		_arg1 = C.TRUE
	}

	C.adw_overlay_split_view_set_enable_hide_gesture(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableHideGesture)
}

// SetEnableShowGesture sets whether self can be opened with an edge swipe
// gesture.
//
// Only touchscreen swipes are supported.
//
// The function takes the following parameters:
//
//   - enableShowGesture: whether self can be opened with a swipe gesture.
func (self *OverlaySplitView) SetEnableShowGesture(enableShowGesture bool) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enableShowGesture {
		_arg1 = C.TRUE
	}

	C.adw_overlay_split_view_set_enable_show_gesture(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableShowGesture)
}

// SetMaxSidebarWidth sets the maximum sidebar width for self.
//
// Maximum width is affected by overlaysplitview:sidebar-width-unit.
//
// The sidebar widget can still be allocated with larger width if its own
// minimum width exceeds it.
//
// The function takes the following parameters:
//
//   - width: maximum width.
func (self *OverlaySplitView) SetMaxSidebarWidth(width float64) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.double               // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(width)

	C.adw_overlay_split_view_set_max_sidebar_width(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(width)
}

// SetMinSidebarWidth sets the minimum sidebar width for self.
//
// Minimum width is affected by overlaysplitview:sidebar-width-unit.
//
// The sidebar widget can still be allocated with larger width if its own
// minimum width exceeds it.
//
// The function takes the following parameters:
//
//   - width: minimum width.
func (self *OverlaySplitView) SetMinSidebarWidth(width float64) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.double               // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(width)

	C.adw_overlay_split_view_set_min_sidebar_width(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(width)
}

// SetPINSidebar sets whether the sidebar widget is pinned for self.
//
// By default, collapsing self automatically hides the sidebar widget, and
// uncollapsing it shows the sidebar. If set to TRUE, sidebar visibility never
// changes on its own.
//
// The function takes the following parameters:
//
//   - pinSidebar: whether to pin the sidebar widget.
func (self *OverlaySplitView) SetPINSidebar(pinSidebar bool) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if pinSidebar {
		_arg1 = C.TRUE
	}

	C.adw_overlay_split_view_set_pin_sidebar(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(pinSidebar)
}

// SetShowSidebar sets whether the sidebar widget is shown for self.
//
// The function takes the following parameters:
//
//   - showSidebar: whether to show the sidebar widget.
func (self *OverlaySplitView) SetShowSidebar(showSidebar bool) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if showSidebar {
		_arg1 = C.TRUE
	}

	C.adw_overlay_split_view_set_show_sidebar(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showSidebar)
}

// SetSidebar sets the sidebar widget for self.
//
// The function takes the following parameters:
//
//   - sidebar (optional) widget.
func (self *OverlaySplitView) SetSidebar(sidebar gtk.Widgetter) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if sidebar != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(sidebar).Native()))
	}

	C.adw_overlay_split_view_set_sidebar(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sidebar)
}

// SetSidebarPosition sets the sidebar position for self.
//
// If it's set to GTK_PACK_START, the sidebar is displayed before the content,
// if GTK_PACK_END, it's displayed after the content.
//
// The function takes the following parameters:
//
//   - position: new position.
func (self *OverlaySplitView) SetSidebarPosition(position gtk.PackType) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.GtkPackType          // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkPackType(position)

	C.adw_overlay_split_view_set_sidebar_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}

// SetSidebarWidthFraction sets the preferred sidebar width as a fraction of the
// total width of self.
//
// The preferred width is additionally limited by
// overlaysplitview:min-sidebar-width and overlaysplitview:max-sidebar-width.
//
// The sidebar widget can be allocated with larger width if its own minimum
// width exceeds the preferred width.
//
// The function takes the following parameters:
//
//   - fraction: preferred width fraction.
func (self *OverlaySplitView) SetSidebarWidthFraction(fraction float64) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.double               // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(fraction)

	C.adw_overlay_split_view_set_sidebar_width_fraction(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(fraction)
}

// SetSidebarWidthUnit sets the length unit for minimum and maximum sidebar
// widths.
//
// See overlaysplitview:min-sidebar-width and
// overlaysplitview:max-sidebar-width.
//
// The function takes the following parameters:
//
//   - unit: length unit.
func (self *OverlaySplitView) SetSidebarWidthUnit(unit LengthUnit) {
	var _arg0 *C.AdwOverlaySplitView // out
	var _arg1 C.AdwLengthUnit        // out

	_arg0 = (*C.AdwOverlaySplitView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwLengthUnit(unit)

	C.adw_overlay_split_view_set_sidebar_width_unit(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(unit)
}
