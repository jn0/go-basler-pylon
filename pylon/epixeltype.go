// Copyright (C) jno, 2019; License: MIT

// EPixelType is an `enum` (i.e. `int`).
// It is supposed to have 32 bits like 0xabCD1234
// where 0xab000000 is "pixel class" (mono/color/custom),
// 0x00CD0000 is "pixel width" (bits per pixel), and
// 0x00001234 is "pixel subtype" (just serial number).
// See $PYLON_HOME/include/pylon/PixelType.h for details
package pylon

import "fmt"

type EPixelType int

func (pt *EPixelType) String() string {
	if !pt.IsValid() {
		return fmt.Sprintf("<EPixelType(%08x)Invalid>", *pt)
	}
	if pt.IsUndefined() {
		return "<EPixelType:UNDEFINED>"
	}
	return fmt.Sprintf("<EPixelType:%s(%d)%04x>",
			   pt.ClassName(), pt.Bits(), pt.Type())
}

func (pt *EPixelType) ClassName() string {
	var class string = ""
	if pt.IsCustom() {
		class = "CUSTOM"
	}
	if pt.IsColor() {
		if class != "" {
			class += ","
		}
		class += "COLOR"
	}
	if pt.IsMono() {
		if class != "" {
			class += ","
		}
		class += "MONO"
	}
	return class
}

func (pt *EPixelType) IsValid() bool {
	class := int((*pt >> 24) & 0x0ff)
	return pt.IsUndefined() || class & ^(0x01 | 0x02 | 0x80) == 0
}

func (pt *EPixelType) IsMono() bool {
	class := int((*pt >> 24) & 0x0ff)
	return class & 0x01 == 0x01
}

func (pt *EPixelType) IsColor() bool {
	class := int((*pt >> 24) & 0x0ff)
	return class & 0x02 == 0x02
}

func (pt *EPixelType) IsCustom() bool {
	class := int((*pt >> 24) & 0x0ff)
	return class & 0x80 == 0x80
}

func (pt *EPixelType) Bits() int {
	return int((*pt >> 16) & 0x0ff)
}

func (pt *EPixelType) Type() int {
	return int(*pt & 0x0ffff)
}

func (pt *EPixelType) IsUndefined() bool {
	return *pt == -1
}

