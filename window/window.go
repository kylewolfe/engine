// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 Package window abstracts the OpenGL Window manager
 Currently only "glfw" is supported
*/
package window

import (
	"github.com/g3n/engine/core"
	"github.com/go-gl/glfw/v3.2/glfw"
)

//
// Interface for all window managers
//
type IWindow interface {
	core.IDispatcher
	GetScreenResolution(interface{}) (width, height int)
	SwapInterval(interval int)
	MakeContextCurrent()
	GetSize() (width int, height int)
	SetSize(width int, height int)
	GetPos() (xpos, ypos int)
	SetPos(xpos, ypos int)
	SetTitle(title string)
	SetStandardCursor(cursor StandardCursor)
	SwapBuffers()
	ShouldClose() bool
	SetShouldClose(bool)
	Destroy()
	PollEvents()
	GetTime() float64
}

// Key corresponds to a keyboard key.
type Key int

//
// Keycodes (from glfw)
//
const (
	KeyUnknown      = Key(glfw.KeyUnknown)
	KeySpace        = Key(glfw.KeySpace)
	KeyApostrophe   = Key(glfw.KeyApostrophe)
	KeyComma        = Key(glfw.KeyComma)
	KeyMinus        = Key(glfw.KeyMinus)
	KeyPeriod       = Key(glfw.KeyPeriod)
	KeySlash        = Key(glfw.KeySlash)
	Key0            = Key(glfw.Key0)
	Key1            = Key(glfw.Key1)
	Key2            = Key(glfw.Key2)
	Key3            = Key(glfw.Key3)
	Key4            = Key(glfw.Key4)
	Key5            = Key(glfw.Key5)
	Key6            = Key(glfw.Key6)
	Key7            = Key(glfw.Key7)
	Key8            = Key(glfw.Key8)
	Key9            = Key(glfw.Key8)
	KeySemicolon    = Key(glfw.KeySemicolon)
	KeyEqual        = Key(glfw.KeyEqual)
	KeyA            = Key(glfw.KeyA)
	KeyB            = Key(glfw.KeyB)
	KeyC            = Key(glfw.KeyC)
	KeyD            = Key(glfw.KeyD)
	KeyE            = Key(glfw.KeyE)
	KeyF            = Key(glfw.KeyF)
	KeyG            = Key(glfw.KeyG)
	KeyH            = Key(glfw.KeyH)
	KeyI            = Key(glfw.KeyI)
	KeyJ            = Key(glfw.KeyJ)
	KeyK            = Key(glfw.KeyK)
	KeyL            = Key(glfw.KeyL)
	KeyM            = Key(glfw.KeyM)
	KeyN            = Key(glfw.KeyN)
	KeyO            = Key(glfw.KeyO)
	KeyP            = Key(glfw.KeyP)
	KeyQ            = Key(glfw.KeyQ)
	KeyR            = Key(glfw.KeyR)
	KeyS            = Key(glfw.KeyS)
	KeyT            = Key(glfw.KeyT)
	KeyU            = Key(glfw.KeyU)
	KeyV            = Key(glfw.KeyV)
	KeyW            = Key(glfw.KeyW)
	KeyX            = Key(glfw.KeyX)
	KeyY            = Key(glfw.KeyY)
	KeyZ            = Key(glfw.KeyZ)
	KeyLeftBracket  = Key(glfw.KeyLeftBracket)
	KeyBackslash    = Key(glfw.KeyBackslash)
	KeyRightBracket = Key(glfw.KeyRightBracket)
	KeyGraveAccent  = Key(glfw.KeyGraveAccent)
	KeyWorld1       = Key(glfw.KeyWorld1)
	KeyWorld2       = Key(glfw.KeyWorld2)
	KeyEscape       = Key(glfw.KeyEscape)
	KeyEnter        = Key(glfw.KeyEnter)
	KeyTab          = Key(glfw.KeyTab)
	KeyBackspace    = Key(glfw.KeyBackspace)
	KeyInsert       = Key(glfw.KeyInsert)
	KeyDelete       = Key(glfw.KeyDelete)
	KeyRight        = Key(glfw.KeyRight)
	KeyLeft         = Key(glfw.KeyLeft)
	KeyDown         = Key(glfw.KeyDown)
	KeyUp           = Key(glfw.KeyUp)
	KeyPageUp       = Key(glfw.KeyPageUp)
	KeyPageDown     = Key(glfw.KeyPageDown)
	KeyHome         = Key(glfw.KeyHome)
	KeyEnd          = Key(glfw.KeyEnd)
	KeyCapsLock     = Key(glfw.KeyCapsLock)
	KeyScrollLock   = Key(glfw.KeyScrollLock)
	KeyNumLock      = Key(glfw.KeyNumLock)
	KeyPrintScreen  = Key(glfw.KeyPrintScreen)
	KeyPause        = Key(glfw.KeyPause)
	KeyF1           = Key(glfw.KeyF1)
	KeyF2           = Key(glfw.KeyF2)
	KeyF3           = Key(glfw.KeyF3)
	KeyF4           = Key(glfw.KeyF4)
	KeyF5           = Key(glfw.KeyF5)
	KeyF6           = Key(glfw.KeyF6)
	KeyF7           = Key(glfw.KeyF7)
	KeyF8           = Key(glfw.KeyF8)
	KeyF9           = Key(glfw.KeyF9)
	KeyF10          = Key(glfw.KeyF10)
	KeyF11          = Key(glfw.KeyF11)
	KeyF12          = Key(glfw.KeyF12)
	KeyF13          = Key(glfw.KeyF13)
	KeyF14          = Key(glfw.KeyF14)
	KeyF15          = Key(glfw.KeyF15)
	KeyF16          = Key(glfw.KeyF16)
	KeyF17          = Key(glfw.KeyF17)
	KeyF18          = Key(glfw.KeyF18)
	KeyF19          = Key(glfw.KeyF19)
	KeyF20          = Key(glfw.KeyF20)
	KeyF21          = Key(glfw.KeyF21)
	KeyF22          = Key(glfw.KeyF22)
	KeyF23          = Key(glfw.KeyF23)
	KeyF24          = Key(glfw.KeyF24)
	KeyF25          = Key(glfw.KeyF25)
	KeyKP0          = Key(glfw.KeyKP0)
	KeyKP1          = Key(glfw.KeyKP1)
	KeyKP2          = Key(glfw.KeyKP2)
	KeyKP3          = Key(glfw.KeyKP3)
	KeyKP4          = Key(glfw.KeyKP4)
	KeyKP5          = Key(glfw.KeyKP5)
	KeyKP6          = Key(glfw.KeyKP6)
	KeyKP7          = Key(glfw.KeyKP7)
	KeyKP8          = Key(glfw.KeyKP8)
	KeyKP9          = Key(glfw.KeyKP9)
	KeyKPDecimal    = Key(glfw.KeyKPDecimal)
	KeyKPDivide     = Key(glfw.KeyKPDivide)
	KeyKPMultiply   = Key(glfw.KeyKPMultiply)
	KeyKPSubtract   = Key(glfw.KeyKPSubtract)
	KeyKPAdd        = Key(glfw.KeyKPAdd)
	KeyKPEnter      = Key(glfw.KeyKPEnter)
	KeyKPEqual      = Key(glfw.KeyKPEqual)
	KeyLeftShift    = Key(glfw.KeyLeftShift)
	KeyLeftControl  = Key(glfw.KeyLeftControl)
	KeyLeftAlt      = Key(glfw.KeyLeftAlt)
	KeyLeftSuper    = Key(glfw.KeyLeftSuper)
	KeyRightShift   = Key(glfw.KeyRightShift)
	KeyRightControl = Key(glfw.KeyRightControl)
	KeyRightAlt     = Key(glfw.KeyRightAlt)
	KeyRightSuper   = Key(glfw.KeyRightSuper)
	KeyMenu         = Key(glfw.KeyMenu)
	KeyLast         = Key(glfw.KeyLast)
)

// ModifierKey corresponds to a modifier key.
type ModifierKey int

// Modifier keys
const (
	ModShift   = ModifierKey(glfw.ModShift)
	ModControl = ModifierKey(glfw.ModControl)
	ModAlt     = ModifierKey(glfw.ModAlt)
	ModSuper   = ModifierKey(glfw.ModSuper)
)

// MouseButton corresponds to a mouse button.
type MouseButton int

// Mouse buttons
const (
	MouseButton1      = MouseButton(glfw.MouseButton1)
	MouseButton2      = MouseButton(glfw.MouseButton2)
	MouseButton3      = MouseButton(glfw.MouseButton3)
	MouseButton4      = MouseButton(glfw.MouseButton4)
	MouseButton5      = MouseButton(glfw.MouseButton5)
	MouseButton6      = MouseButton(glfw.MouseButton6)
	MouseButton7      = MouseButton(glfw.MouseButton7)
	MouseButton8      = MouseButton(glfw.MouseButton8)
	MouseButtonLast   = MouseButton(glfw.MouseButtonLast)
	MouseButtonLeft   = MouseButton(glfw.MouseButtonLeft)
	MouseButtonRight  = MouseButton(glfw.MouseButtonRight)
	MouseButtonMiddle = MouseButton(glfw.MouseButtonMiddle)
)

// StandardCursor corresponds to a standard cursor icon.
type StandardCursor int

// Standard cursors
const (
	ArrowCursor     = StandardCursor(glfw.ArrowCursor)
	IBeamCursor     = StandardCursor(glfw.IBeamCursor)
	CrosshairCursor = StandardCursor(glfw.CrosshairCursor)
	HandCursor      = StandardCursor(glfw.HandCursor)
	HResizeCursor   = StandardCursor(glfw.HResizeCursor)
	VResizeCursor   = StandardCursor(glfw.VResizeCursor)
)

// Action corresponds to a key or button action.
type Action int

const (
	Release = Action(glfw.Release) // The key or button was released.
	Press   = Action(glfw.Press)   // The key or button was pressed.
	Repeat  = Action(glfw.Repeat)  // The key was held down until it repeated.
)

// InputMode corresponds to an input mode.
type InputMode int

// Input modes
const (
	CursorMode             = InputMode(glfw.CursorMode)             // See Cursor mode values
	StickyKeysMode         = InputMode(glfw.StickyKeysMode)         // Value can be either 1 or 0
	StickyMouseButtonsMode = InputMode(glfw.StickyMouseButtonsMode) // Value can be either 1 or 0
)

// Cursor mode values
const (
	CursorNormal   = glfw.CursorNormal
	CursorHidden   = glfw.CursorHidden
	CursorDisabled = glfw.CursorDisabled
)

//
// Window event names using for dispatch and subscribe
//
const (
	OnWindowPos  = "win.OnWindowPos"
	OnWindowSize = "win.OnWindowSize"
	OnKeyUp      = "win.OnKeyUp"
	OnKeyDown    = "win.OnKeyDown"
	OnKeyRepeat  = "win.OnKeyRepeat"
	OnChar       = "win.OnChar"
	OnCursor     = "win.OnCursor"
	OnMouseUp    = "win.OnMouseUp"
	OnMouseDown  = "win.OnMouseDown"
	OnScroll     = "win.OnScroll"
	OnFrame      = "win.OnFrame"
)

// Window position changed event
type PosEvent struct {
	W    IWindow
	Xpos int
	Ypos int
}

// Window size changed
type SizeEvent struct {
	W      IWindow
	Width  int
	Height int
}

// Key pressed in window
type KeyEvent struct {
	W        IWindow
	Keycode  Key
	Scancode int
	Action   Action
	Mods     ModifierKey
}

// Char pressed in window
type CharEvent struct {
	W    IWindow
	Char rune
	Mods ModifierKey
}

// Mouse button event
type MouseEvent struct {
	W      IWindow
	Xpos   float32
	Ypos   float32
	Button MouseButton
	Action Action
	Mods   ModifierKey
}

// Cursor position changed
type CursorEvent struct {
	W    IWindow
	Xpos float32
	Ypos float32
}

// Scroll event
type ScrollEvent struct {
	W       IWindow
	Xoffset float32
	Yoffset float32
}

// New creates and returns a new window of the specified type, width, height and title.
// If full is true, the window will be opened in full screen and the width and height
// parameters will be ignored.
// Currently only "glfw" type is supported.
func New(wtype string, width, height int, title string, full bool) (IWindow, error) {

	if wtype != "glfw" {
		panic("Unsupported window type")
	}
	return newGLFW(width, height, title, full)
}
