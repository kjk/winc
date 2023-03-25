/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */

package winc

import (
	"fmt"
	"syscall"

	"github.com/kjk/winc/w32"
)

const (
	FontBold      byte = 0x01
	FontItalic    byte = 0x02
	FontUnderline byte = 0x04
	FontStrikeOut byte = 0x08
)

func init() {
	hDC := w32.GetDC(0)
	defer w32.ReleaseDC(0, hDC)
	ScreenDPIY = w32.GetDeviceCaps(hDC, w32.LOGPIXELSY)

	ncm := w32.GetNonClientMetrics()
	h := int(ncm.LfMessageFont.Height)
	DefaultFont = NewFontPixel("MS Shell Dlg 2", h, 0)
}

type Font struct {
	hfont     w32.HFONT
	family    string
	pointSize int
	style     byte
}

func Unscale(n int, dpi int) int {
	return (n * 72) / dpi
}

func NewFontPixel(family string, size int, style byte) *Font {
	if style > FontBold|FontItalic|FontUnderline|FontStrikeOut {
		panic("Invalid font style")
	}
	pointSize := Unscale(size, ScreenDPIY)
	fmt.Printf("pointSize: %v, size: %v\n", pointSize, size)

	font := Font{
		family:    family,
		pointSize: pointSize,
		style:     style,
	}

	font.hfont = font.createForDPI(ScreenDPIY)
	if font.hfont == 0 {
		panic("CreateFontIndirect failed")
	}

	return &font
}

func NewFont(family string, pointSize int, style byte) *Font {
	if style > FontBold|FontItalic|FontUnderline|FontStrikeOut {
		panic("Invalid font style")
	}

	font := Font{
		family:    family,
		pointSize: pointSize,
		style:     style,
	}

	font.hfont = font.createForDPI(ScreenDPIY)
	if font.hfont == 0 {
		panic("CreateFontIndirect failed")
	}

	return &font
}

func (fnt *Font) createForHeight(height int) w32.HFONT {
	var lf w32.LOGFONT
	lf.Height = int32(-height)
	if fnt.style&FontBold > 0 {
		lf.Weight = w32.FW_BOLD
	} else {
		lf.Weight = w32.FW_NORMAL
	}
	if fnt.style&FontItalic > 0 {
		lf.Italic = 1
	}
	if fnt.style&FontUnderline > 0 {
		lf.Underline = 1
	}
	if fnt.style&FontStrikeOut > 0 {
		lf.StrikeOut = 1
	}
	lf.CharSet = w32.DEFAULT_CHARSET
	lf.OutPrecision = w32.OUT_TT_PRECIS
	lf.ClipPrecision = w32.CLIP_DEFAULT_PRECIS
	lf.Quality = w32.CLEARTYPE_QUALITY
	lf.PitchAndFamily = w32.VARIABLE_PITCH | w32.FF_SWISS

	src, _ := syscall.UTF16FromString(fnt.family)
	dest := lf.FaceName[:]
	copy(dest, src)

	return w32.CreateFontIndirect(&lf)
}

func (fnt *Font) createForDPI(dpi int) w32.HFONT {
	n := w32.MulDiv(fnt.pointSize, dpi, 72)
	return fnt.createForHeight(n)
}

func (fnt *Font) GetHFONT() w32.HFONT {
	return fnt.hfont
}

func (fnt *Font) Bold() bool {
	return fnt.style&FontBold > 0
}

func (fnt *Font) Dispose() {
	if fnt.hfont != 0 {
		w32.DeleteObject(w32.HGDIOBJ(fnt.hfont))
	}
}

func (fnt *Font) Family() string {
	return fnt.family
}

func (fnt *Font) Italic() bool {
	return fnt.style&FontItalic > 0
}

func (fnt *Font) StrikeOut() bool {
	return fnt.style&FontStrikeOut > 0
}

func (fnt *Font) Underline() bool {
	return fnt.style&FontUnderline > 0
}

func (fnt *Font) Style() byte {
	return fnt.style
}
