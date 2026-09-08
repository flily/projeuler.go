package framework

import (
	"strconv"
	"strings"
)

type SelectorInfo struct {
	String string
	Int    *int
}

func newSelectorInfo(s string, n *int) *SelectorInfo {
	i := &SelectorInfo{
		String: s,
		Int:    n,
	}

	return i
}

func NewAbsentSelectorInfo() *SelectorInfo {
	i := &SelectorInfo{
		String: "",
		Int:    nil,
	}

	return i
}

func NewSelectorInfo(s string) *SelectorInfo {
	valNum := (*int)(nil)
	num, err := strconv.Atoi(s)
	if err == nil {
		valNum = &num
	}

	return newSelectorInfo(s, valNum)
}

func (i *SelectorInfo) Equals(o *SelectorInfo) bool {
	if i.Int != nil && o.Int != nil {
		if *i.Int != *o.Int {
			return false
		}
	}

	if i.Int == nil || o.Int == nil {
		return false
	}

	if i.String != o.String {
		return false
	}

	return true
}

func (i *SelectorInfo) IsAbsent() bool {
	return i.String == ""
}

func (i *SelectorInfo) MatchProblem(p *Problem) bool {
	if i.Int != nil {
		if p.Id == int(*i.Int) {
			return true
		}
	}

	if i.String != "" {
		return strings.Contains(p.Title, i.String)
	}

	return false
}

type SelectorParseError struct {
	Raw     string
	Message string
	Start   int
	End     int
}

func NewSelectorParseError(raw string) *SelectorParseError {
	e := &SelectorParseError{
		Raw:     raw,
		Message: "",
		Start:   0,
		End:     0,
	}

	return e
}

func (e *SelectorParseError) On(start, end int) *SelectorParseError {
	e.Start = start
	e.End = end
	return e
}

func (e *SelectorParseError) With(message string) *SelectorParseError {
	e.Message = message
	return e
}

func (e *SelectorParseError) Error() string {
	lines := make([]string, 0, 3)
	indent := strings.Repeat(" ", e.Start)
	indicator := "^" + strings.Repeat("^", e.End-e.Start)

	lines = append(lines, e.Raw)
	lines = append(lines, indent+indicator)
	lines = append(lines, indent+e.Message)
	return strings.Join(lines, "\n")
}

type Selector struct {
	Problem  SelectorInfo
	Solution []*SelectorInfo
}

func readSelectorInfo(content []rune, start int) (*SelectorInfo, int) {
	i := start
	for ; i < len(content); i++ {
		if content[i] == '.' || content[i] == ',' || content[i] == '{' || content[i] == '}' {
			break
		}
	}

	s := string(content[start:i])
	info := NewSelectorInfo(s)
	return info, i
}

func ParseSelector(s string) (*Selector, error) {
	errBase := NewSelectorParseError(s) // error base

	content := []rune(s)
	pid, next := readSelectorInfo(content, 0)
	if next <= 0 {
		err := errBase.On(0, 0).
			With("no problem selector found")
		return nil, err
	}

	selector := &Selector{
		Problem:  *pid,
		Solution: make([]*SelectorInfo, 0),
	}

	if next < len(content) {
		if content[next] != '.' {
			err := errBase.On(next, next+1).
				With("expect '.' after problem selector")
			return nil, err
		}

		next += 1
		itemStart := next
		stop := false
		started, closed := false, false
		for ; !stop && next < len(content); next += 1 {
			switch content[next] {
			case '{':
				if started {
					err := errBase.On(next, next+1).
						With("nested '{' is not allowed")
					return nil, err
				}
				started = true
				itemStart += 1

			case '}':
				if !started {
					err := errBase.On(next, next+1).
						With("unexpected '}'")
					return nil, err
				}
				closed = true
				stop = true
				fallthrough

			case ',':
				s := string(content[itemStart:next])
				info := NewSelectorInfo(s)
				selector.Solution = append(selector.Solution, info)
				itemStart = next + 1
			}
		}

		if itemStart < len(content) {
			// flush
			s := string(content[itemStart:])
			info := NewSelectorInfo(s)
			selector.Solution = append(selector.Solution, info)
		}

		if started != closed {
			err := errBase.On(next, next+1).
				With("braces not closed")
			return nil, err
		}
	}

	return selector, nil
}

func (s *Selector) Equal(o *Selector) bool {
	if !s.Problem.Equals(&o.Problem) {
		return false
	}

	if len(s.Solution) != len(o.Solution) {
		return false
	}

	for i := range s.Solution {
		if !s.Solution[i].Equals(o.Solution[i]) {
			return false
		}
	}

	return true
}
