// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_leaflet_transition_type_get_type()), F: marshalLeafletTransitionType},
		{T: externglib.Type(C.adw_leaflet_get_type()), F: marshalLeafletter},
		{T: externglib.Type(C.adw_leaflet_page_get_type()), F: marshalLeafletPager},
	})
}

// LeafletTransitionType describes the possible transitions in a adw.Leaflet
// widget.
//
// New values may be added to this enumeration over time.
type LeafletTransitionType C.gint

const (
	// LeafletTransitionTypeOver: cover the old page or uncover the new page,
	// sliding from or towards the end according to orientation, text direction
	// and children order.
	LeafletTransitionTypeOver LeafletTransitionType = iota
	// LeafletTransitionTypeUnder: uncover the new page or cover the old page,
	// sliding from or towards the start according to orientation, text
	// direction and children order.
	LeafletTransitionTypeUnder
	// LeafletTransitionTypeSlide: slide from left, right, up or down according
	// to the orientation, text direction and the children order.
	LeafletTransitionTypeSlide
)

func marshalLeafletTransitionType(p uintptr) (interface{}, error) {
	return LeafletTransitionType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for LeafletTransitionType.
func (l LeafletTransitionType) String() string {
	switch l {
	case LeafletTransitionTypeOver:
		return "Over"
	case LeafletTransitionTypeUnder:
		return "Under"
	case LeafletTransitionTypeSlide:
		return "Slide"
	default:
		return fmt.Sprintf("LeafletTransitionType(%d)", l)
	}
}

// Leaflet: adaptive container acting like a box or a stack.
//
// The AdwLeaflet widget can display its children like a gtk.Box does or like a
// gtk.Stack does, adapting to size changes by switching between the two modes.
//
// When there is enough space the children are displayed side by side, otherwise
// only one is displayed and the leaflet is said to be “folded”. The threshold
// is dictated by the preferred minimum sizes of the children. When a leaflet is
// folded, the children can be navigated using swipe gestures.
//
// The “over” and “under” transition types stack the children one on top of the
// other, while the “slide” transition puts the children side by side. While
// navigating to a child on the side or below can be performed by swiping the
// current child away, navigating to an upper child requires dragging it from
// the edge where it resides. This doesn't affect non-dragging swipes.
//
//
// CSS nodes
//
// AdwLeaflet has a single CSS node with name leaflet. The node will get the
// style classes .folded when it is folded, .unfolded when it's not, or none if
// it hasn't computed its fold yet.
type Leaflet struct {
	gtk.Widget

	Swipeable
	gtk.Orientable
	*externglib.Object
}

var (
	_ gtk.Widgetter       = (*Leaflet)(nil)
	_ externglib.Objector = (*Leaflet)(nil)
)

func wrapLeaflet(obj *externglib.Object) *Leaflet {
	return &Leaflet{
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
		Swipeable: Swipeable{
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
		Orientable: gtk.Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalLeafletter(p uintptr) (interface{}, error) {
	return wrapLeaflet(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLeaflet creates a new AdwLeaflet.
func NewLeaflet() *Leaflet {
	var _cret *C.GtkWidget // in

	_cret = C.adw_leaflet_new()

	var _leaflet *Leaflet // out

	_leaflet = wrapLeaflet(externglib.Take(unsafe.Pointer(_cret)))

	return _leaflet
}

// Append adds a child to self.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
func (self *Leaflet) Append(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_leaflet_append(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// AdjacentChild finds the previous or next navigatable child.
//
// This will be the same child adw.Leaflet.Navigate() or swipe gestures will
// navigate to.
//
// If there's no child to navigate to, NULL will be returned instead.
//
// See adw.LeafletPage:navigatable.
//
// The function takes the following parameters:
//
//    - direction: direction.
//
func (self *Leaflet) AdjacentChild(direction NavigationDirection) gtk.Widgetter {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwNavigationDirection // out
	var _cret *C.GtkWidget             // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwNavigationDirection(direction)

	_cret = C.adw_leaflet_get_adjacent_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(direction)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// CanSwipeBack gets whether a swipe gesture can be used to navigate to the
// previous child.
func (self *Leaflet) CanSwipeBack() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_can_swipe_back(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanSwipeForward gets whether a swipe gesture can be used to navigate to the
// next child.
func (self *Leaflet) CanSwipeForward() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_can_swipe_forward(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUnfold gets whether self can unfold.
func (self *Leaflet) CanUnfold() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_can_unfold(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildByName finds the child of self with name.
//
// Returns NULL if there is no child with this name.
//
// See adw.LeafletPage:name.
//
// The function takes the following parameters:
//
//    - name of the child to find.
//
func (self *Leaflet) ChildByName(name string) gtk.Widgetter {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.char       // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_leaflet_get_child_by_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// ChildTransitionDuration gets the child transition animation duration for
// self.
func (self *Leaflet) ChildTransitionDuration() uint {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.guint       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_child_transition_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ChildTransitionRunning gets whether a child transition is currently running
// for self.
func (self *Leaflet) ChildTransitionRunning() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_child_transition_running(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FoldThresholdPolicy gets the fold threshold policy for self.
func (self *Leaflet) FoldThresholdPolicy() FoldThresholdPolicy {
	var _arg0 *C.AdwLeaflet            // out
	var _cret C.AdwFoldThresholdPolicy // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_fold_threshold_policy(_arg0)
	runtime.KeepAlive(self)

	var _foldThresholdPolicy FoldThresholdPolicy // out

	_foldThresholdPolicy = FoldThresholdPolicy(_cret)

	return _foldThresholdPolicy
}

// Folded gets whether self is folded.
func (self *Leaflet) Folded() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_folded(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous gets whether self is homogeneous for the given fold and
// orientation.
//
// See adw.Leaflet:hhomogeneous-folded, adw.Leaflet:vhomogeneous-folded,
// adw.Leaflet:hhomogeneous-unfolded, adw.Leaflet:vhomogeneous-unfolded.
//
// The function takes the following parameters:
//
//    - folded: fold.
//    - orientation: orientation.
//
func (self *Leaflet) Homogeneous(folded bool, orientation gtk.Orientation) bool {
	var _arg0 *C.AdwLeaflet    // out
	var _arg1 C.gboolean       // out
	var _arg2 C.GtkOrientation // out
	var _cret C.gboolean       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if folded {
		_arg1 = C.TRUE
	}
	_arg2 = C.GtkOrientation(orientation)

	_cret = C.adw_leaflet_get_homogeneous(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(folded)
	runtime.KeepAlive(orientation)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize gets whether the leaflet interpolates its size when changing
// the visible child.
func (self *Leaflet) InterpolateSize() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_interpolate_size(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ModeTransitionDuration gets the mode transition animation duration for self.
func (self *Leaflet) ModeTransitionDuration() uint {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.guint       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_mode_transition_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Page returns the adw.LeafletPage object for child.
//
// The function takes the following parameters:
//
//    - child of self.
//
func (self *Leaflet) Page(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_leaflet_get_page(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Pages returns a GListModel that contains the pages of the leaflet.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track and change the visible page.
func (self *Leaflet) Pages() gtk.SelectionModeller {
	var _arg0 *C.AdwLeaflet        // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_pages(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel gtk.SelectionModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.SelectionModeller is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(gtk.SelectionModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.SelectionModeller")
		}
		_selectionModel = rv
	}

	return _selectionModel
}

// TransitionType gets the type of animation used for transitions between modes
// and children.
func (self *Leaflet) TransitionType() LeafletTransitionType {
	var _arg0 *C.AdwLeaflet              // out
	var _cret C.AdwLeafletTransitionType // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_transition_type(_arg0)
	runtime.KeepAlive(self)

	var _leafletTransitionType LeafletTransitionType // out

	_leafletTransitionType = LeafletTransitionType(_cret)

	return _leafletTransitionType
}

// VisibleChild gets the widget currently visible when the leaflet is folded.
func (self *Leaflet) VisibleChild() gtk.Widgetter {
	var _arg0 *C.AdwLeaflet // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_visible_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// VisibleChildName gets the name of the currently visible child widget.
func (self *Leaflet) VisibleChildName() string {
	var _arg0 *C.AdwLeaflet // out
	var _cret *C.char       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_visible_child_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// InsertChildAfter inserts child in the position after sibling in the list of
// children.
//
// If sibling is NULL, inserts child at the first position.
//
// The function takes the following parameters:
//
//    - child: widget to insert.
//    - sibling after which to insert child.
//
func (self *Leaflet) InsertChildAfter(child, sibling gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if sibling != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	}

	_cret = C.adw_leaflet_insert_child_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Navigate navigates to the previous or next child.
//
// The child must have the adw.LeafletPage:navigatable property set to TRUE,
// otherwise it will be skipped.
//
// This will be the same child as returned by adw.Leaflet.GetAdjacentChild() or
// navigated to via swipe gestures.
//
// The function takes the following parameters:
//
//    - direction: direction.
//
func (self *Leaflet) Navigate(direction NavigationDirection) bool {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwNavigationDirection // out
	var _cret C.gboolean               // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwNavigationDirection(direction)

	_cret = C.adw_leaflet_navigate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(direction)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Prepend inserts child at the first position in self.
//
// The function takes the following parameters:
//
//    - child: widget to prepend.
//
func (self *Leaflet) Prepend(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_leaflet_prepend(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Remove removes a child widget from self.
//
// The function takes the following parameters:
//
//    - child to remove.
//
func (self *Leaflet) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_leaflet_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// ReorderChildAfter moves child to the position after sibling in the list of
// children.
//
// If sibling is NULL, moves child to the first position.
//
// The function takes the following parameters:
//
//    - child: widget to move, must be a child of self.
//    - sibling to move child after.
//
func (self *Leaflet) ReorderChildAfter(child, sibling gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if sibling != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	}

	C.adw_leaflet_reorder_child_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
}

// SetCanSwipeBack sets whether a swipe gesture can be used to navigate to the
// previous child.
//
// The function takes the following parameters:
//
//    - canSwipeBack: new value.
//
func (self *Leaflet) SetCanSwipeBack(canSwipeBack bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if canSwipeBack {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_swipe_back(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canSwipeBack)
}

// SetCanSwipeForward sets whether a swipe gesture can be used to navigate to
// the next child.
//
// The function takes the following parameters:
//
//    - canSwipeForward: new value.
//
func (self *Leaflet) SetCanSwipeForward(canSwipeForward bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if canSwipeForward {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_swipe_forward(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canSwipeForward)
}

// SetCanUnfold sets whether self can unfold.
//
// The function takes the following parameters:
//
//    - canUnfold: whether self can unfold.
//
func (self *Leaflet) SetCanUnfold(canUnfold bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if canUnfold {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_unfold(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canUnfold)
}

// SetChildTransitionDuration sets the child transition animation duration for
// self.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (self *Leaflet) SetChildTransitionDuration(duration uint) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.guint       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(duration)

	C.adw_leaflet_set_child_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetFoldThresholdPolicy sets the fold threshold policy for self.
//
// The function takes the following parameters:
//
//    - policy to use.
//
func (self *Leaflet) SetFoldThresholdPolicy(policy FoldThresholdPolicy) {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwFoldThresholdPolicy // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwFoldThresholdPolicy(policy)

	C.adw_leaflet_set_fold_threshold_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetHomogeneous sets self to be homogeneous or not for the given fold and
// orientation.
//
// If it is homogeneous, self will request the same width or height for all its
// children depending on the orientation. If it isn't and it is folded, the
// leaflet may change width or height when a different child becomes visible.
//
// See adw.Leaflet:hhomogeneous-folded, adw.Leaflet:vhomogeneous-folded,
// adw.Leaflet:hhomogeneous-unfolded, adw.Leaflet:vhomogeneous-unfolded.
//
// The function takes the following parameters:
//
//    - folded: fold.
//    - orientation: orientation.
//    - homogeneous: TRUE to make self homogeneous.
//
func (self *Leaflet) SetHomogeneous(folded bool, orientation gtk.Orientation, homogeneous bool) {
	var _arg0 *C.AdwLeaflet    // out
	var _arg1 C.gboolean       // out
	var _arg2 C.GtkOrientation // out
	var _arg3 C.gboolean       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if folded {
		_arg1 = C.TRUE
	}
	_arg2 = C.GtkOrientation(orientation)
	if homogeneous {
		_arg3 = C.TRUE
	}

	C.adw_leaflet_set_homogeneous(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(folded)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(homogeneous)
}

// SetInterpolateSize sets whether the leaflet interpolates its size when
// changing the visible child.
//
// The function takes the following parameters:
//
//    - interpolateSize: new value.
//
func (self *Leaflet) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_interpolate_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(interpolateSize)
}

// SetModeTransitionDuration sets the mode transition animation duration for
// self.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (self *Leaflet) SetModeTransitionDuration(duration uint) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.guint       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(duration)

	C.adw_leaflet_set_mode_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetTransitionType sets the type of animation used for transitions between
// modes and children.
//
// The function takes the following parameters:
//
//    - transition: new transition type.
//
func (self *Leaflet) SetTransitionType(transition LeafletTransitionType) {
	var _arg0 *C.AdwLeaflet              // out
	var _arg1 C.AdwLeafletTransitionType // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwLeafletTransitionType(transition)

	C.adw_leaflet_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(transition)
}

// SetVisibleChild sets the widget currently visible when the leaflet is folded.
//
// The function takes the following parameters:
//
//    - visibleChild: new child.
//
func (self *Leaflet) SetVisibleChild(visibleChild gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(visibleChild.Native()))

	C.adw_leaflet_set_visible_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(visibleChild)
}

// SetVisibleChildName makes the child with the name name visible.
//
// See adw_leaflet_set_visible_child() for more details.
//
// The function takes the following parameters:
//
//    - name of a child.
//
func (self *Leaflet) SetVisibleChildName(name string) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.char       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_leaflet_set_visible_child_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// LeafletPage: auxiliary class used by adw.Leaflet.
type LeafletPage struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*LeafletPage)(nil)
)

func wrapLeafletPage(obj *externglib.Object) *LeafletPage {
	return &LeafletPage{
		Object: obj,
	}
}

func marshalLeafletPager(p uintptr) (interface{}, error) {
	return wrapLeafletPage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child gets the leaflet child th which self belongs.
func (self *LeafletPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwLeafletPage // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_page_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Name gets the name of self.
func (self *LeafletPage) Name() string {
	var _arg0 *C.AdwLeafletPage // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_page_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Navigatable gets whether the child can be navigated to when folded.
func (self *LeafletPage) Navigatable() bool {
	var _arg0 *C.AdwLeafletPage // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_page_get_navigatable(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetName sets the name of the self.
//
// The function takes the following parameters:
//
//    - name: new value to set.
//
func (self *LeafletPage) SetName(name string) {
	var _arg0 *C.AdwLeafletPage // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_leaflet_page_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetNavigatable sets whether self can be navigated to when folded.
//
// The function takes the following parameters:
//
//    - navigatable: whether self can be navigated to when folded.
//
func (self *LeafletPage) SetNavigatable(navigatable bool) {
	var _arg0 *C.AdwLeafletPage // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))
	if navigatable {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_page_set_navigatable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(navigatable)
}
