package treesitter

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

// Returns true if point p is inside range r
func PointInRange(p ts.Point, r ts.Range) bool {
	if r.StartPoint.Row == r.EndPoint.Row {
		return p.Row == r.StartPoint.Row && p.Column >= r.StartPoint.Column && p.Column <= r.EndPoint.Column
	}
	return p.Row >= r.StartPoint.Row && p.Row <= r.EndPoint.Row
}

// Returns true if range a and b overlaps.
func RangeIntersect(a ts.Range, b ts.Range) bool {
	return (a.StartByte >= b.StartByte && a.EndByte <= b.EndByte) ||
		(b.StartByte >= a.StartByte && b.EndByte <= a.EndByte)
}
