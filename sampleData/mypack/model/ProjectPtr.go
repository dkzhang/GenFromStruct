// Code generated by generator.
package model

import "time"

type ProjectPtr struct {
	ProjectID           *int       `db:"project_id" json:"project_id"`
	ProjectName         *string    `db:"project_name" json:"project_name"`
	ProjectCode         *string    `db:"project_code" json:"project_code"`
	DepartmentCode      *string    `db:"department_code" json:"department_code"`
	Department          *string    `db:"department" json:"department"`
	ChiefID             *int       `db:"chief_id" json:"chief_id"`
	ChiefChineseName    *string    `db:"chief_cn_name" json:"chief_cn_name"`
	ExtraInfo           *string    `db:"extra_info" json:"extra_info"`
	BasicStatus         *int       `db:"basic_status" json:"basic_status"`
	StartDate           *time.Time `db:"start_date" json:"start_date"`
	TotalDaysApply      *int       `db:"total_days_apply" json:"total_days_apply"`
	EndReminderAt       *time.Time `db:"end_reminder_at" json:"end_reminder_at"`
	CpuNodesExpected    *int       `db:"cpu_nodes_expected" json:"cpu_nodes_expected"`
	GpuNodesExpected    *int       `db:"gpu_nodes_expected" json:"gpu_nodes_expected"`
	StorageSizeExpected *int       `db:"storage_size_expected" json:"storage_size_expected"`
	CpuNodesAcquired    *int       `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    *int       `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired *int       `db:"storage_size_acquired" json:"storage_size_acquired"`
	AccountAllocInfo    *string    `db:"account_alloc_info" json:"account_alloc_info"`
	StorageAllocInfo    *string    `db:"storage_alloc_info" json:"storage_alloc_info"`
	CreatedAt           *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt           *time.Time `db:"updated_at" json:"updated_at"`
}

func (c ProjectPtr) toDbMap() map[string]interface{} {
	m := make(map[string]interface{})
	if c.ProjectID != nil {
		m["project_id"] = *c.ProjectID
	}
	if c.ProjectName != nil {
		m["project_name"] = *c.ProjectName
	}
	if c.ProjectCode != nil {
		m["project_code"] = *c.ProjectCode
	}
	if c.DepartmentCode != nil {
		m["department_code"] = *c.DepartmentCode
	}
	if c.Department != nil {
		m["department"] = *c.Department
	}
	if c.ChiefID != nil {
		m["chief_id"] = *c.ChiefID
	}
	if c.ChiefChineseName != nil {
		m["chief_cn_name"] = *c.ChiefChineseName
	}
	if c.ExtraInfo != nil {
		m["extra_info"] = *c.ExtraInfo
	}
	if c.BasicStatus != nil {
		m["basic_status"] = *c.BasicStatus
	}
	if c.StartDate != nil {
		m["start_date"] = *c.StartDate
	}
	if c.TotalDaysApply != nil {
		m["total_days_apply"] = *c.TotalDaysApply
	}
	if c.EndReminderAt != nil {
		m["end_reminder_at"] = *c.EndReminderAt
	}
	if c.CpuNodesExpected != nil {
		m["cpu_nodes_expected"] = *c.CpuNodesExpected
	}
	if c.GpuNodesExpected != nil {
		m["gpu_nodes_expected"] = *c.GpuNodesExpected
	}
	if c.StorageSizeExpected != nil {
		m["storage_size_expected"] = *c.StorageSizeExpected
	}
	if c.CpuNodesAcquired != nil {
		m["cpu_nodes_acquired"] = *c.CpuNodesAcquired
	}
	if c.GpuNodesAcquired != nil {
		m["gpu_nodes_acquired"] = *c.GpuNodesAcquired
	}
	if c.StorageSizeAcquired != nil {
		m["storage_size_acquired"] = *c.StorageSizeAcquired
	}
	if c.AccountAllocInfo != nil {
		m["account_alloc_info"] = *c.AccountAllocInfo
	}
	if c.StorageAllocInfo != nil {
		m["storage_alloc_info"] = *c.StorageAllocInfo
	}
	if c.CreatedAt != nil {
		m["created_at"] = *c.CreatedAt
	}
	if c.UpdatedAt != nil {
		m["updated_at"] = *c.UpdatedAt
	}
	return m
}
