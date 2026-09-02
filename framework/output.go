package framework

import (
	"fmt"
	"slices"
	"strings"
)

type Column struct {
	Name  string
	Style Style
}

func (c *Column) String() string {
	return c.Style.Apply(c.Name)
}

type OutputTable struct {
	Headers []Column
}

func NewOutputTable() *OutputTable {
	columns := make([]Column, 0, 16)
	t := &OutputTable{
		Headers: columns,
	}

	return t
}

func NewOutputTableWith(columns []Column) *OutputTable {
	t := &OutputTable{
		Headers: slices.Clone(columns),
	}
	return t
}

func (t *OutputTable) AddHeader(name string, style Style) {
	c := Column{
		Name:  name,
		Style: style,
	}

	t.Headers = append(t.Headers, c)
}

func (t *OutputTable) makeLine(apply func(Column, int) string) string {
	builder := make([]string, 0, 2+2*len(t.Headers))
	builder = append(builder, "|")

	ext := ""

	for i, h := range t.Headers {
		applied := apply(h, i)

		line := " " + applied + " "
		builder = append(builder, line)
		builder = append(builder, "|")
	}

	if len(ext) > 0 {
		builder = append(builder, " "+ext)
	}

	return strings.Join(builder, "")
}

func (t *OutputTable) MakeStyleLine(fields ...DisplayStyle) string {
	return t.makeLine(func(h Column, i int) string {
		if i < len(fields) {
			return h.Style.ApplyWith(fields[i])
		}

		return h.Style.Apply("")
	})
}

func (t *OutputTable) MakeLine(fields ...string) string {
	return t.makeLine(func(h Column, i int) string {
		if i < len(fields) {
			return h.Style.Apply(fields[i])
		}

		return h.Style.Apply("")
	})
}

func (t *OutputTable) Separator() string {
	builder := make([]string, 0, 1+2*len(t.Headers))
	builder = append(builder, "+")

	for _, h := range t.Headers {
		builder = append(builder, strings.Repeat("-", h.Style.Width+2))
		builder = append(builder, "+")
	}

	return strings.Join(builder, "")
}

func (t *OutputTable) PrintSeparator() {
	fmt.Println(t.Separator())
}

func (t *OutputTable) PrintHeader() {
	t.PrintSeparator()

	fields := make([]string, len(t.Headers))
	for i, h := range t.Headers {
		fields[i] = h.Name
	}
	fmt.Println(t.MakeLine(fields...))
	t.PrintSeparator()
}

func (t *OutputTable) PrintLine(fields ...string) {
	fmt.Println(t.MakeLine(fields...))
}

func (t *OutputTable) PrintStyleItems(items ...DisplayStyle) {
	fmt.Println(t.MakeStyleLine(items...))
}
