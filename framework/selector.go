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

	return true
}

func (i *SelectorInfo) MatchSolution(index int, s SolutionEntry) bool {
	if i.Int != nil {
		if *i.Int+1 == index {
			return true
		}
	}

	if strings.Contains(s.Name, i.String) {
		return true
	}

	return false
}

func (i *SelectorInfo) MatchSolutions(p *Problem) []bool {
	result := make([]bool, len(p.Methods))

	for index, s := range p.Methods {
		if i.MatchSolution(index, s) {
			result[index] = true
		}
	}

	return result
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
	Problem   SelectorInfo
	Solutions []*SelectorInfo
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
		pid = NewAbsentSelectorInfo()
	}

	selector := &Selector{
		Problem:   *pid,
		Solutions: make([]*SelectorInfo, 0),
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
				selector.Solutions = append(selector.Solutions, info)
				itemStart = next + 1
			}
		}

		if itemStart < len(content) {
			// flush
			s := string(content[itemStart:])
			info := NewSelectorInfo(s)
			selector.Solutions = append(selector.Solutions, info)
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

	if len(s.Solutions) != len(o.Solutions) {
		return false
	}

	for i := range s.Solutions {
		if !s.Solutions[i].Equals(o.Solutions[i]) {
			return false
		}
	}

	return true
}

func mergeBooleanMap(a []bool, b []bool) []bool {
	if len(a) != len(b) {
		panic("length mismatch")
	}

	result := make([]bool, len(a))
	for i := range a {
		result[i] = a[i] || b[i]
	}

	return result
}

func (s *Selector) MatchedSolutions(p *Problem) []string {
	result := make([]string, 0, len(p.Methods))

	if s == nil || len(s.Solutions) <= 0 {
		for _, method := range p.Methods {
			result = append(result, method.Name)
		}

	} else {
		matched := make([]bool, len(p.Methods))
		for _, solution := range s.Solutions {
			part := solution.MatchSolutions(p)
			matched = mergeBooleanMap(matched, part)
		}

		for i := range p.Methods {
			if matched[i] {
				result = append(result, p.Methods[i].Name)
			}
		}
	}

	return result
}

func (s *Selector) Match(p *Problem) bool {
	if !s.Problem.MatchProblem(p) {
		return false
	}

	matched := s.MatchedSolutions(p)
	if len(p.Methods) > 0 {
		return len(matched) > 0
	}

	return true
}

type SelectorCollection []Selector

func ParseSelectorCollection(selectors []string) (SelectorCollection, error) {
	result := make(SelectorCollection, 0, len(selectors))
	for _, s := range selectors {
		selector, err := ParseSelector(s)
		if err != nil {
			return nil, err
		}
		result = append(result, *selector)
	}

	return result, nil
}

func (c SelectorCollection) Match(p *Problem) (*Selector, bool) {
	for _, selector := range c {
		if selector.Match(p) {
			return &selector, true
		}
	}

	return nil, len(c) == 0
}
