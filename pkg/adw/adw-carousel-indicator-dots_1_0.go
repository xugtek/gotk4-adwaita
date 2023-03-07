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
	GTypeCarouselIndicatorDots = coreglib.Type(C.adw_carousel_indicator_dots_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCarouselIndicatorDots, F: marshalCarouselIndicatorDots},
	})
}

// CarouselIndicatorDotsOverrides contains methods that are overridable.
type CarouselIndicatorDotsOverrides struct {
}

func defaultCarouselIndicatorDotsOverrides(v *CarouselIndicatorDots) CarouselIndicatorDotsOverrides {
	return CarouselIndicatorDotsOverrides{}
}

// CarouselIndicatorDots dots indicator for carousel.
//
// <picture> <source srcset="carousel-indicator-dots-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="carousel-indicator-dots.png"
// alt="carousel-indicator-dots"> </picture>
//
// The AdwCarouselIndicatorDots widget shows a set of dots for each page of a
// given carousel. The dot representing the carousel's active page is larger and
// more opaque than the others, the transition to the active and inactive state
// is gradual to match the carousel's position.
//
// See also carouselindicatorlines.
//
//
// CSS nodes
//
// AdwCarouselIndicatorDots has a single CSS node with name
// carouselindicatordots.
type CarouselIndicatorDots struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	gtk.Orientable
}

var (
	_ gtk.Widgetter     = (*CarouselIndicatorDots)(nil)
	_ coreglib.Objector = (*CarouselIndicatorDots)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CarouselIndicatorDots, *CarouselIndicatorDotsClass, CarouselIndicatorDotsOverrides](
		GTypeCarouselIndicatorDots,
		initCarouselIndicatorDotsClass,
		wrapCarouselIndicatorDots,
		defaultCarouselIndicatorDotsOverrides,
	)
}

func initCarouselIndicatorDotsClass(gclass unsafe.Pointer, overrides CarouselIndicatorDotsOverrides, classInitFunc func(*CarouselIndicatorDotsClass)) {
	if classInitFunc != nil {
		class := (*CarouselIndicatorDotsClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCarouselIndicatorDots(obj *coreglib.Object) *CarouselIndicatorDots {
	return &CarouselIndicatorDots{
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
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalCarouselIndicatorDots(p uintptr) (interface{}, error) {
	return wrapCarouselIndicatorDots(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCarouselIndicatorDots creates a new AdwCarouselIndicatorDots.
//
// The function returns the following values:
//
//    - carouselIndicatorDots: newly created AdwCarouselIndicatorDots.
//
func NewCarouselIndicatorDots() *CarouselIndicatorDots {
	var _cret *C.GtkWidget // in

	_cret = C.adw_carousel_indicator_dots_new()

	var _carouselIndicatorDots *CarouselIndicatorDots // out

	_carouselIndicatorDots = wrapCarouselIndicatorDots(coreglib.Take(unsafe.Pointer(_cret)))

	return _carouselIndicatorDots
}

// Carousel gets the displayed carousel.
//
// The function returns the following values:
//
//    - carousel (optional): displayed carousel.
//
func (self *CarouselIndicatorDots) Carousel() *Carousel {
	var _arg0 *C.AdwCarouselIndicatorDots // out
	var _cret *C.AdwCarousel              // in

	_arg0 = (*C.AdwCarouselIndicatorDots)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_carousel_indicator_dots_get_carousel(_arg0)
	runtime.KeepAlive(self)

	var _carousel *Carousel // out

	if _cret != nil {
		_carousel = wrapCarousel(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _carousel
}

// SetCarousel sets the displayed carousel.
//
// The function takes the following parameters:
//
//    - carousel (optional): carousel.
//
func (self *CarouselIndicatorDots) SetCarousel(carousel *Carousel) {
	var _arg0 *C.AdwCarouselIndicatorDots // out
	var _arg1 *C.AdwCarousel              // out

	_arg0 = (*C.AdwCarouselIndicatorDots)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if carousel != nil {
		_arg1 = (*C.AdwCarousel)(unsafe.Pointer(coreglib.InternObject(carousel).Native()))
	}

	C.adw_carousel_indicator_dots_set_carousel(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(carousel)
}