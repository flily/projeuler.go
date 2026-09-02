package framework

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type (
	Alignment  int
	FormatType int
	Colour     int
)

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight

	FormatTypeDefault FormatType = iota
	FormatTypeDecimal
	FormatTypeOctal
	FormatTypeHexadecimalUpper
	FormatTypeHexadecimalLower
	FormatTypeFloat
	FormatTypeString

	ColourDefault Colour = iota
	ColourBlack
	ColourRed
	ColourGreen
	ColourYellow
	ColourBlue
	ColourMagenta
	ColourCyan
	ColourWhite
	ColourHiBlack
	ColourHiRed
	ColourHiGreen
	ColourHiYellow
	ColourHiBlue
	ColourHiMagenta
	ColourHiCyan
	ColourHiWhite
)

var (
	fgColourMap = map[Colour]color.Attribute{
		ColourBlack:     color.FgBlack,
		ColourRed:       color.FgRed,
		ColourGreen:     color.FgGreen,
		ColourYellow:    color.FgYellow,
		ColourBlue:      color.FgBlue,
		ColourMagenta:   color.FgMagenta,
		ColourCyan:      color.FgCyan,
		ColourWhite:     color.FgWhite,
		ColourHiBlack:   color.FgHiBlack,
		ColourHiRed:     color.FgHiRed,
		ColourHiGreen:   color.FgHiGreen,
		ColourHiYellow:  color.FgHiYellow,
		ColourHiBlue:    color.FgHiBlue,
		ColourHiMagenta: color.FgHiMagenta,
		ColourHiCyan:    color.FgHiCyan,
		ColourHiWhite:   color.FgHiWhite,
	}

	bgColourMap = map[Colour]color.Attribute{
		ColourBlack:     color.BgBlack,
		ColourRed:       color.BgRed,
		ColourGreen:     color.BgGreen,
		ColourYellow:    color.BgYellow,
		ColourBlue:      color.BgBlue,
		ColourMagenta:   color.BgMagenta,
		ColourCyan:      color.BgCyan,
		ColourWhite:     color.BgWhite,
		ColourHiBlack:   color.BgHiBlack,
		ColourHiRed:     color.BgHiRed,
		ColourHiGreen:   color.BgHiGreen,
		ColourHiYellow:  color.BgHiYellow,
		ColourHiBlue:    color.BgHiBlue,
		ColourHiMagenta: color.BgHiMagenta,
		ColourHiCyan:    color.BgHiCyan,
		ColourHiWhite:   color.BgHiWhite,
	}
)

type TerminalStyleBooleans uint32

const (
	TerminalStyleBold       TerminalStyleBooleans = 0x0001
	TerminalStyleItalic     TerminalStyleBooleans = 0x0002
	TerminalStyleUnderline  TerminalStyleBooleans = 0x0004
	TerminalStyleDeleted    TerminalStyleBooleans = 0x0008
	TerminalStyleForceColor TerminalStyleBooleans = 0x0010
)

type TerminalStyle struct {
	FgColour Colour
	BgColour Colour
	Booleans TerminalStyleBooleans
}

func NewTerminalStyle(fg Colour, bg Colour, booleans TerminalStyleBooleans) TerminalStyle {
	s := TerminalStyle{
		FgColour: fg,
		BgColour: bg,
		Booleans: booleans,
	}

	return s
}

func DefaultTerminalStyle() TerminalStyle {
	return NewTerminalStyle(ColourDefault, ColourDefault, 0)
}

func (s TerminalStyle) Reset() TerminalStyle {
	return DefaultTerminalStyle()
}

func (s TerminalStyle) Apply() *color.Color {
	attrs := make([]color.Attribute, 0, 6)

	if fgAttr, found := fgColourMap[s.FgColour]; found {
		attrs = append(attrs, fgAttr)
	}

	if bgAttr, found := bgColourMap[s.BgColour]; found {
		attrs = append(attrs, bgAttr)
	}

	c := color.New(attrs...)

	if s.Booleans&TerminalStyleBold != 0 {
		c.Add(color.Bold)
	}

	if s.Booleans&TerminalStyleItalic != 0 {
		c.Add(color.Italic)
	}

	if s.Booleans&TerminalStyleUnderline != 0 {
		c.Add(color.Underline)
	}

	if s.Booleans&TerminalStyleDeleted != 0 {
		c.Add(color.CrossedOut)
	}

	if s.Booleans&TerminalStyleForceColor != 0 {
		c.EnableColor()
	}

	return c
}

func (s TerminalStyle) Bold() TerminalStyle {
	return NewTerminalStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleBold)
}

func (s TerminalStyle) Italic() TerminalStyle {
	return NewTerminalStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleItalic)
}

func (s TerminalStyle) Underline() TerminalStyle {
	return NewTerminalStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleUnderline)
}

func (s TerminalStyle) Deleted() TerminalStyle {
	return NewTerminalStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleDeleted)
}

func (s TerminalStyle) Colour(colour Colour) TerminalStyle {
	return NewTerminalStyle(colour, s.BgColour, s.Booleans)
}

func (s TerminalStyle) BackgroundColour(colour Colour) TerminalStyle {
	return NewTerminalStyle(s.FgColour, colour, s.Booleans)
}

func (s TerminalStyle) Force() TerminalStyle {
	return NewTerminalStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleForceColor)
}

type Style struct {
	Alignment Alignment
	Width     int
	Precision int
	Padding   string
	Format    FormatType
	Terminal  TerminalStyle
}

func NewStyleWith(format FormatType, align Alignment, termStyle TerminalStyle, width int, precision int, padding string) Style {
	s := Style{
		Alignment: align,
		Width:     width,
		Precision: precision,
		Padding:   padding,
		Format:    format,
		Terminal:  termStyle,
	}

	return s
}

// NewStyle creates a new Style based on the given format string.
// [padding] [alignment] [width] [. [precision]] [format]
func NewIntegerStyle(width int, precision int) Style {
	s := Style{
		Alignment: AlignRight,
		Width:     width,
		Precision: precision,
		Padding:   " ",
		Format:    FormatTypeDecimal,
		Terminal:  TerminalStyle{},
	}

	return s
}

func NewFloatStyle(width int, precision int) Style {
	s := Style{
		Alignment: AlignRight,
		Width:     width,
		Precision: precision,
		Padding:   " ",
		Format:    FormatTypeFloat,
		Terminal:  TerminalStyle{},
	}

	return s
}

func NewGenericStyle(width int) Style {
	s := Style{
		Alignment: AlignLeft,
		Width:     width,
		Precision: 0,
		Padding:   " ",
		Format:    FormatTypeDefault,
		Terminal:  TerminalStyle{},
	}

	return s
}

func (s Style) WithPadding(padding string) Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal, s.Width, s.Precision, padding)
}

func (s Style) Left() Style {
	return NewStyleWith(s.Format, AlignLeft, s.Terminal, s.Width, s.Precision, s.Padding)
}

func (s Style) Center() Style {
	return NewStyleWith(s.Format, AlignCenter, s.Terminal, s.Width, s.Precision, s.Padding)
}

func (s Style) Right() Style {
	return NewStyleWith(s.Format, AlignRight, s.Terminal, s.Width, s.Precision, s.Padding)
}

func (s Style) Reset() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal.Reset(), s.Width, s.Precision, s.Padding)
}

func (s Style) Bold() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal.Bold(), s.Width, s.Precision, s.Padding)
}

func (s Style) Underline() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal.Underline(), s.Width, s.Precision, s.Padding)
}

func (s Style) Italic() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal.Italic(), s.Width, s.Precision, s.Padding)
}

func (s Style) Colour(colour Colour) Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal.Colour(colour), s.Width, s.Precision, s.Padding)
}

func (s Style) Background(colour Colour) Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal.BackgroundColour(colour), s.Width, s.Precision, s.Padding)
}

func (s Style) Black() Style {
	return s.Colour(ColourBlack)
}

func (s Style) Red() Style {
	return s.Colour(ColourRed)
}

func (s Style) Green() Style {
	return s.Colour(ColourGreen)
}

func (s Style) Blue() Style {
	return s.Colour(ColourBlue)
}

func (s Style) Yellow() Style {
	return s.Colour(ColourYellow)
}

func (s Style) Magenta() Style {
	return s.Colour(ColourMagenta)
}

func (s Style) Cyan() Style {
	return s.Colour(ColourCyan)
}

func (s Style) White() Style {
	return s.Colour(ColourWhite)
}

func (s Style) HiBlack() Style {
	return s.Colour(ColourHiBlack)
}

func (s Style) HiRed() Style {
	return s.Colour(ColourHiRed)
}

func (s Style) HiGreen() Style {
	return s.Colour(ColourHiGreen)
}

func (s Style) HiBlue() Style {
	return s.Colour(ColourHiBlue)
}

func (s Style) HiYellow() Style {
	return s.Colour(ColourHiYellow)
}

func (s Style) HiMagenta() Style {
	return s.Colour(ColourHiMagenta)
}

func (s Style) HiCyan() Style {
	return s.Colour(ColourHiCyan)
}

func (s Style) HiWhite() Style {
	return s.Colour(ColourHiWhite)
}

func (s Style) Force() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Terminal.Force(), s.Width, s.Precision, s.Padding)
}

func (s Style) applyColour(value string) string {
	c := s.Terminal.Apply()

	return c.Sprint(value)
}

func repeatPadding(padding string, length int) string {
	pl := len(padding)
	count := length / pl
	remainder := length % pl
	return strings.Repeat(padding, count) + padding[:remainder]
}

func (s Style) applyAlignment(value string, padding string) string {
	remain := s.Width - len(value)
	content := value
	if remain > 0 && len(padding) > 0 {
		if s.Alignment == AlignCenter {
			left := remain / 2
			right := remain - left
			content = repeatPadding(padding, left) + content + repeatPadding(padding, right)

		} else {
			fill := repeatPadding(padding, remain)
			if s.Alignment == AlignRight {
				content = fill + content
			} else {
				content = content + fill
			}
		}
	}

	return s.applyColour(content)
}

func (s Style) applyFloat(value any) string {
	parts := make([]string, 0, 5)
	parts = append(parts, "%")

	if s.Precision >= 0 {
		parts = append(parts, fmt.Sprintf(".%d", s.Precision))
	}

	switch s.Format {
	case FormatTypeDefault, FormatTypeFloat:
		parts = append(parts, "f")

	default:
		err := fmt.Errorf("unsupported format type '%v' for %T", s.Format, value)
		panic(err)
	}

	format := strings.Join(parts, "")
	content := fmt.Sprintf(format, value)
	return s.applyAlignment(content, s.Padding)
}

func (s Style) applyInteger(value any) string {
	parts := make([]string, 0, 5)
	parts = append(parts, "%")

	switch s.Format {
	case FormatTypeDefault, FormatTypeDecimal:
		parts = append(parts, "d")

	case FormatTypeOctal:
		parts = append(parts, "o")

	case FormatTypeHexadecimalLower:
		parts = append(parts, "x")

	case FormatTypeHexadecimalUpper:
		parts = append(parts, "X")

	default:
		err := fmt.Errorf("unsupported format type '%v' for %T", s.Format, value)
		panic(err)
	}

	format := strings.Join(parts, "")
	content := fmt.Sprintf(format, value)
	return s.applyAlignment(content, s.Padding)
}

func (s Style) applyString(value string) string {
	return s.applyAlignment(value, s.Padding)
}

func (s Style) Apply(value any) string {
	result := ""
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		result = s.applyInteger(v)

	case uint, uint8, uint16, uint32, uint64:
		result = s.applyInteger(v)

	case float32, float64:
		result = s.applyFloat(v)

	case string:
		result = s.applyString(v)

	default:
		content := ""
		if stringer, ok := v.(fmt.Stringer); ok {
			content = stringer.String()
		} else {
			content = fmt.Sprintf("%v", v)
		}

		result = s.applyString(content)
	}

	return result
}
