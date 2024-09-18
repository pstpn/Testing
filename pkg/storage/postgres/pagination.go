package postgres

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

type SortDirection int8

const (
	ASC SortDirection = iota
	DESC
)

func (s SortDirection) String() string {
	switch s {
	case DESC:
		return "DESC"
	default:
		return "ASC"
	}
}

func SortDirectionFromString(dir string) SortDirection {
	switch dir {
	case "ASC":
		return ASC
	default:
		return DESC
	}
}

type SortOptions struct {
	Direction SortDirection
	Columns   []string
}

func (s SortOptions) Format() string {
	return fmt.Sprintf("%s %s", strings.Join(s.Columns, ","), s.Direction.String())
}

type FilterOptions struct {
	Pattern string
	Column  string
}

type Pagination struct {
	PageNumber int64
	PageSize   int64
	Filter     FilterOptions
	Sort       SortOptions
}

func (p *Pagination) ToSQL(s squirrel.SelectBuilder) squirrel.SelectBuilder {
	if p.Sort.Columns[0] != "" {
		s = s.OrderBy(p.Sort.Format())
	}

	if p.Filter.Column != "" {
		s = s.Where(squirrel.ILike{p.Filter.Column + "::text": fmt.Sprintf("%%%s%%", p.Filter.Pattern)})
	}
	if p.PageSize > 0 {
		s = s.Limit(uint64(p.PageSize))
	}
	if p.PageNumber >= 0 && p.PageSize > 0 {
		s = s.Offset(uint64(p.PageNumber * p.PageSize))
	}

	return s
}
