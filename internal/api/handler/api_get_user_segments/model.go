package api_get_user_segments

import "segment-manager/internal/api/model"

type Response struct {
	Segments []Segment          `json:"segments"`
	Error    *model.ResultError `json:"error,omitempty"`
}

type Segment struct {
	SegmentID int64  `json:"segment_id"`
	Slug      string `json:"slug"`
}
