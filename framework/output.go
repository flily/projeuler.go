package framework

import (
	"fmt"
	"slices"
	"strings"
)

type Column struct {
	Name  string
	Width int
}

func (c *Column) String() string {
	format := fmt.Sprintf("%%%ds", c.Width)
	return fmt.Sprintf(format, c.Name)
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

func (t *OutputTable) AddHeader(name string, width int) {
	c := Column{
		Name:  name,
		Width: width,
	}

	t.Headers = append(t.Headers, c)
}

func (t *OutputTable) MakeLine(fields ...string) string {
	builder := make([]string, 0, 2+2*len(t.Headers))
	builder = append(builder, "|")

	ext := ""

	for i, h := range t.Headers {
		content := ""
		if i < len(fields) {
			content = fields[i]
		}

		format := fmt.Sprintf(" %%-%ds ", h.Width)
		builder = append(builder, fmt.Sprintf(format, content))
		builder = append(builder, "|")
	}

	if len(ext) > 0 {
		builder = append(builder, " "+ext)
	}

	return strings.Join(builder, "")
}

func (t *OutputTable) Separator() string {
	builder := make([]string, 0, 1+2*len(t.Headers))
	builder = append(builder, "+")

	for _, h := range t.Headers {
		builder = append(builder, strings.Repeat("-", h.Width+2))
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
