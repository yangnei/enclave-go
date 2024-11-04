package api

import (
	"net/url"
	"strconv"
)

type TimeRange struct {
	StartMs int64
	EndMs   int64
}

func (t *TimeRange) GetUrlValues() url.Values {
	values := url.Values{}
	if t.StartMs != 0 {
		values.Set("startTime", strconv.FormatInt(t.StartMs, 10))
	}
	if t.EndMs != 0 {
		values.Set("endTime", strconv.FormatInt(t.EndMs, 10))
	}
	return values
}

type Paging struct {
	Limit  int
	Cursor string
}

func (p *Paging) GetUrlValues() url.Values {
	values := url.Values{}
	if p.Limit != 0 {
		values.Set("limit", strconv.Itoa(p.Limit))
	}
	if p.Cursor != "" {
		values.Set("cursor", p.Cursor)
	}
	return values
}

type PagingAndTimeRange struct {
	Paging
	TimeRange
}

func (p *PagingAndTimeRange) GetUrlValues() url.Values {
	values := p.GetUrlValues()
	timeValues := p.GetUrlValues()
	for k, v := range timeValues {
		for _, value := range v {
			values.Add(k, value)
		}
	}
	return values
}

// PageInfo represents pagination information.
type PageInfo struct {
	NextCursor string `json:"nextCursor,omitempty"`
	PrevCursor string `json:"prevCursor,omitempty"`
}

// Response represents the generic API response.
type Response[T any] struct {
	Success   bool   `json:"success"`
	Result    T      `json:"result,omitempty"`
	Error     string `json:"error,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
}

// PaginatedResponse represents the generic API paged response.
type PaginatedResponse[T any] struct {
	Response[T]
	PageInfo PageInfo `json:"pageInfo"`
}

// Error represents an error returned by the API.
type Error struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}
