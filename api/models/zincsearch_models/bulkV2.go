package zincsearch_models

import "api/models"

type BulkV2 struct {
	IndexName string          `json:"index"`
	Records   []*models.Email `json:"records"`
}
