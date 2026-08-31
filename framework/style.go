package framework

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Alignment int
type FormatType int

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
)

type Style struct {
	Alignment Alignment
	Width     int
	Precision int
	Padding   string
	Format    FormatType
	Attribute color.Attribute
}

func NewStyleWith(format FormatType, align Alignment, attr color.Attribute, width int, precision int, padding string) Style {
	s := Style{
		Alignment: align,
		Width:     width,
		Precision: precision,
		Padding:   padding,
		Format:    format,
		Attribute: attr,
	}

	return s
}

// NewStyle creates a new Style based on the given format string.
// [padding] [alignment] [width] [. [precision]] [format]
func NewIntegerStyle(width int, precision int) Style {
	s := Style{
		Alignment: AlignLeft,
		Width:     width,
		Precision: precision,
		Padding:   " ",
		Format:    FormatTypeDecimal,
		Attribute: color.Reset,
	}

	return s
}

func (s Style) Reset() Style {
	return NewStyleWith(s.Format, s.Alignment, color.Reset, s.Width, s.Precision, s.Padding)
}

func (s Style) Bold() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.Bold, s.Width, s.Precision, s.Padding)
}

func (s Style) Underline() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.Underline, s.Width, s.Precision, s.Padding)
}

func (s Style) Italic() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.Italic, s.Width, s.Precision, s.Padding)
}

func (s Style) Black() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgBlack, s.Width, s.Precision, s.Padding)
}

func (s Style) Red() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgRed, s.Width, s.Precision, s.Padding)
}

func (s Style) Green() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgGreen, s.Width, s.Precision, s.Padding)
}

func (s Style) Blue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgBlue, s.Width, s.Precision, s.Padding)
}

func (s Style) Yellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgYellow, s.Width, s.Precision, s.Padding)
}

func (s Style) Magenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgMagenta, s.Width, s.Precision, s.Padding)
}

func (s Style) Cyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgCyan, s.Width, s.Precision, s.Padding)
}

func (s Style) White() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgWhite, s.Width, s.Precision, s.Padding)
}

func (s Style) HiBlack() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiBlack, s.Width, s.Precision, s.Padding)
}

func (s Style) HiRed() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiRed, s.Width, s.Precision, s.Padding)
}

func (s Style) HiGreen() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiGreen, s.Width, s.Precision, s.Padding)
}

func (s Style) HiBlue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiBlue, s.Width, s.Precision, s.Padding)
}

func (s Style) HiYellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiYellow, s.Width, s.Precision, s.Padding)
}

func (s Style) HiMagenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiMagenta, s.Width, s.Precision, s.Padding)
}

func (s Style) HiCyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiCyan, s.Width, s.Precision, s.Padding)
}

func (s Style) HiWhite() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.FgHiWhite, s.Width, s.Precision, s.Padding)
}

func (s Style) BgBlack() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgBlack, s.Width, s.Precision, s.Padding)
}

func (s Style) BgRed() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgRed, s.Width, s.Precision, s.Padding)
}

func (s Style) BgGreen() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgGreen, s.Width, s.Precision, s.Padding)
}

func (s Style) BgBlue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgBlue, s.Width, s.Precision, s.Padding)
}

func (s Style) BgYellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgYellow, s.Width, s.Precision, s.Padding)
}

func (s Style) BgMagenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgMagenta, s.Width, s.Precision, s.Padding)
}

func (s Style) BgCyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgCyan, s.Width, s.Precision, s.Padding)
}

func (s Style) BgWhite() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgWhite, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiBlack() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiBlack, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiRed() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiRed, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiGreen() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiGreen, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiBlue() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiBlue, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiYellow() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiYellow, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiMagenta() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiMagenta, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiCyan() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiCyan, s.Width, s.Precision, s.Padding)
}

func (s Style) BgHiWhite() Style {
	return NewStyleWith(s.Format, s.Alignment, s.Attribute|color.BgHiWhite, s.Width, s.Precision, s.Padding)
}

func (s Style) applyColour(value string) string {
	return color.Set(s.Attribute).Sprint(value)
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

func (s Style) applyInteger(value any) string {
	parts := make([]string, 0, 5)
	parts = append(parts, "%")

	if len(s.Padding) > 0 {
		parts = append(parts, s.Padding)
	}

	if s.Width > 0 {
		parts = append(parts, fmt.Sprintf("%d", s.Width))
		if s.Precision > 0 {
			parts = append(parts, fmt.Sprintf(".%d", s.Precision))
		}
	}

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
	return s.applyAlignment(content, " ")
}

func (s Style) applyString(value string) string {
	parts := make([]string, 0, 5)
	parts = append(parts, "%")

	if s.Width > 0 {
		parts = append(parts, fmt.Sprintf("%d", s.Width))
	}

	parts = append(parts, "s")

	format := strings.Join(parts, "")
	content := fmt.Sprintf(format, value)
	return s.applyAlignment(content, s.Padding)
}

func (s Style) Apply(value any) string {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return s.applyInteger(value)

	case uint, uint8, uint16, uint32, uint64:
		return s.applyInteger(value)

	default:
		content := fmt.Sprintf("%v", value)
		return s.applyString(content)
	}
}
