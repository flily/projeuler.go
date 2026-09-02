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

type DisplayStyle struct {
	FgColour Colour
	BgColour Colour
	Booleans TerminalStyleBooleans
	Content  any
}

func NewDisplayStyle(fg Colour, bg Colour, booleans TerminalStyleBooleans, content any) DisplayStyle {
	s := DisplayStyle{
		FgColour: fg,
		BgColour: bg,
		Booleans: booleans,
		Content:  content,
	}

	return s
}

func DefaultDisplayStyle() DisplayStyle {
	return NewDisplayStyle(ColourDefault, ColourDefault, 0, "")
}

func (s DisplayStyle) Reset() DisplayStyle {
	return DefaultDisplayStyle()
}

func (s DisplayStyle) Apply() *color.Color {
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

func (s DisplayStyle) With(content any) DisplayStyle {
	return NewDisplayStyle(s.FgColour, s.BgColour, s.Booleans, content)
}

func (s DisplayStyle) Bold() DisplayStyle {
	return NewDisplayStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleBold, s.Content)
}

func (s DisplayStyle) Italic() DisplayStyle {
	return NewDisplayStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleItalic, s.Content)
}

func (s DisplayStyle) Underline() DisplayStyle {
	return NewDisplayStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleUnderline, s.Content)
}

func (s DisplayStyle) Deleted() DisplayStyle {
	return NewDisplayStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleDeleted, s.Content)
}

func (s DisplayStyle) Colour(colour Colour) DisplayStyle {
	return NewDisplayStyle(colour, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) ToBackgroundColour() DisplayStyle {
	return NewDisplayStyle(ColourDefault, s.FgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) BackgroundColour(colour Colour) DisplayStyle {
	return NewDisplayStyle(s.FgColour, colour, s.Booleans, s.Content)
}

func (s DisplayStyle) Black() DisplayStyle {
	return NewDisplayStyle(ColourBlack, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) Red() DisplayStyle {
	return NewDisplayStyle(ColourRed, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) Green() DisplayStyle {
	return NewDisplayStyle(ColourGreen, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) Yellow() DisplayStyle {
	return NewDisplayStyle(ColourYellow, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) Blue() DisplayStyle {
	return NewDisplayStyle(ColourBlue, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) Magenta() DisplayStyle {
	return NewDisplayStyle(ColourMagenta, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) Cyan() DisplayStyle {
	return NewDisplayStyle(ColourCyan, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) White() DisplayStyle {
	return NewDisplayStyle(ColourWhite, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiBlack() DisplayStyle {
	return NewDisplayStyle(ColourHiBlack, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiRed() DisplayStyle {
	return NewDisplayStyle(ColourHiRed, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiGreen() DisplayStyle {
	return NewDisplayStyle(ColourHiGreen, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiYellow() DisplayStyle {
	return NewDisplayStyle(ColourHiYellow, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiBlue() DisplayStyle {
	return NewDisplayStyle(ColourHiBlue, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiMagenta() DisplayStyle {
	return NewDisplayStyle(ColourHiMagenta, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiCyan() DisplayStyle {
	return NewDisplayStyle(ColourHiCyan, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) HiWhite() DisplayStyle {
	return NewDisplayStyle(ColourHiWhite, s.BgColour, s.Booleans, s.Content)
}

func (s DisplayStyle) BgBlack() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourBlack, s.Booleans, s.Content)
}
func (s DisplayStyle) BgRed() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourRed, s.Booleans, s.Content)
}

func (s DisplayStyle) BgGreen() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourGreen, s.Booleans, s.Content)
}

func (s DisplayStyle) BgYellow() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourYellow, s.Booleans, s.Content)
}

func (s DisplayStyle) BgBlue() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourBlue, s.Booleans, s.Content)
}

func (s DisplayStyle) BgMagenta() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourMagenta, s.Booleans, s.Content)
}

func (s DisplayStyle) BgCyan() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourCyan, s.Booleans, s.Content)
}

func (s DisplayStyle) BgWhite() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourWhite, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiBlack() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiBlack, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiRed() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiRed, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiGreen() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiGreen, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiYellow() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiYellow, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiBlue() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiBlue, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiMagenta() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiMagenta, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiCyan() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiCyan, s.Booleans, s.Content)
}

func (s DisplayStyle) BgHiWhite() DisplayStyle {
	return NewDisplayStyle(s.FgColour, ColourHiWhite, s.Booleans, s.Content)
}

func (s DisplayStyle) Force() DisplayStyle {
	return NewDisplayStyle(s.FgColour, s.BgColour, s.Booleans|TerminalStyleForceColor, s.Content)
}

type Style struct {
	Alignment Alignment
	Width     int
	Precision int
	Padding   string
	Format    FormatType
	Display   DisplayStyle
}

func NewStyleWith(format FormatType, align Alignment, display DisplayStyle, width int, precision int, padding string) Style {
	s := Style{
		Alignment: align,
		Width:     width,
		Precision: precision,
		Padding:   padding,
		Format:    format,
		Display:   display,
	}

	return s
}

// NewStyle creates a new Style based on the given format string.
// [padding] [alignment] [width] [. [precision]] [format]
func NewIntegerStyle(width int) Style {
	s := Style{
		Alignment: AlignRight,
		Width:     width,
		Precision: 0,
		Padding:   " ",
		Format:    FormatTypeDecimal,
		Display:   DefaultDisplayStyle(),
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
		Display:   DefaultDisplayStyle(),
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
		Display:   DefaultDisplayStyle(),
	}

	return s
}

func (s Style) With(display DisplayStyle) Style {
	return NewStyleWith(s.Format, s.Alignment, display, s.Width, s.Precision, s.Padding)
}

func (s Style) WithPadding(padding string) Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display, s.Width, s.Precision, padding)
}

func (s Style) Left() Style {
	return NewStyleWith(s.Format, AlignLeft, s.Display, s.Width, s.Precision, s.Padding)
}

func (s Style) Center() Style {
	return NewStyleWith(s.Format, AlignCenter, s.Display, s.Width, s.Precision, s.Padding)
}

func (s Style) Right() Style {
	return NewStyleWith(s.Format, AlignRight, s.Display, s.Width, s.Precision, s.Padding)
}

func (s Style) Force() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Force(), s.Width, s.Precision, s.Padding)
}

func (s Style) applyColour(value string, display DisplayStyle) string {
	c := display.Apply()
	return c.Sprint(value)
}

func (s Style) Black() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourBlack), s.Width, s.Precision, s.Padding)
}

func (s Style) Red() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourRed), s.Width, s.Precision, s.Padding)
}

func (s Style) Green() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourGreen), s.Width, s.Precision, s.Padding)
}

func (s Style) Blue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourBlue), s.Width, s.Precision, s.Padding)
}

func (s Style) Yellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourYellow), s.Width, s.Precision, s.Padding)
}

func (s Style) Magenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourMagenta), s.Width, s.Precision, s.Padding)
}

func (s Style) Cyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourCyan), s.Width, s.Precision, s.Padding)
}

func (s Style) White() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourWhite), s.Width, s.Precision, s.Padding)
}

func (s Style) HiBlack() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiBlack), s.Width, s.Precision, s.Padding)
}

func (s Style) HiRed() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiRed), s.Width, s.Precision, s.Padding)
}

func (s Style) HiGreen() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiGreen), s.Width, s.Precision, s.Padding)
}

func (s Style) HiBlue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiBlue), s.Width, s.Precision, s.Padding)
}

func (s Style) HiYellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiYellow), s.Width, s.Precision, s.Padding)
}

func (s Style) HiMagenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiMagenta), s.Width, s.Precision, s.Padding)
}

func (s Style) HiCyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiCyan), s.Width, s.Precision, s.Padding)
}

func (s Style) HiWhite() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.Colour(ColourHiWhite), s.Width, s.Precision, s.Padding)
}

func (s Style) BgBlack() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourBlack), s.Width, s.Precision, s.Padding)
}

func (s Style) BgRed() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourRed), s.Width, s.Precision, s.Padding)
}

func (s Style) BgGreen() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourGreen), s.Width, s.Precision, s.Padding)
}

func (s Style) BgYellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourYellow), s.Width, s.Precision, s.Padding)
}

func (s Style) BgBlue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourBlue), s.Width, s.Precision, s.Padding)
}

func (s Style) BgMagenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourMagenta), s.Width, s.Precision, s.Padding)
}

func (s Style) BgCyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourCyan), s.Width, s.Precision, s.Padding)
}

func (s Style) BgWhite() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourWhite), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiBlack() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiBlack), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiRed() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiRed), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiGreen() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiGreen), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiYellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiYellow), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiBlue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiBlue), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiMagenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiMagenta), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiCyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiCyan), s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiWhite() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Display.BackgroundColour(ColourHiWhite), s.Width, s.Precision, s.Padding)
}

func repeatPadding(padding string, length int) string {
	pl := len(padding)
	count := length / pl
	remainder := length % pl
	return strings.Repeat(padding, count) + padding[:remainder]
}

func (s Style) applyAlignment(value string, padding string, display DisplayStyle) string {
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

	return s.applyColour(content, display)
}

func (s Style) applyFloat(value any, display DisplayStyle) string {
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
	return s.applyAlignment(content, s.Padding, display)
}

func (s Style) applyInteger(value any, display DisplayStyle) string {
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
	return s.applyAlignment(content, s.Padding, display)
}

func (s Style) applyString(value string, display DisplayStyle) string {
	return s.applyAlignment(value, s.Padding, display)
}

func (s Style) ApplyWith(display DisplayStyle) string {
	result := ""
	switch v := display.Content.(type) {
	case int, int8, int16, int32, int64:
		result = s.applyInteger(v, display)

	case uint, uint8, uint16, uint32, uint64:
		result = s.applyInteger(v, display)

	case float32, float64:
		result = s.applyFloat(v, display)

	case string:
		result = s.applyString(v, display)

	default:
		content := ""
		if stringer, ok := v.(fmt.Stringer); ok {
			content = stringer.String()
		} else {
			content = fmt.Sprintf("%v", v)
		}

		result = s.applyString(content, display)
	}

	return result
}

func (s Style) Apply(value any) string {
	return s.ApplyWith(s.Display.With(value))
}
