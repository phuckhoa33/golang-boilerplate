// Package base extends its model abstraction structures to include full audit
// capabilities. This file introduces a struct that models can embed to gain
// not only creation and update timestamps but also a deletion timestamp,
// supporting soft delete functionality.
package base

import "time"

// FullAuditModelAbstract extends AuditModelAbstract to include a deletion timestamp.
// This struct is intended to be embedded in models that require full audit capabilities,
// including tracking the soft deletion of records.
type FullAuditModelAbstract struct {
	// AuditModelAbstract is embedded to inherit the CreatedAt and UpdatedAt fields,
	// providing automatic tracking of creation and update timestamps.
	*AuditModelAbstract
	// DeleteAt stores the timestamp of when the model instance was marked as deleted.
	// It supports soft delete functionality, where records are not actually removed
	// from the database but are instead marked as deleted.
	DeleteAt time.Time `json:"deleteAt" gorm:"type:varchar(200);"`
}
