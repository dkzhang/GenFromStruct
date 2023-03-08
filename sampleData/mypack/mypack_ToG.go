//go:generate dktool arg1 "multiword arg"

package mypack

import (
	"time"
)

type Project_ToG struct {
	// Static 8
	ID               int    `pgType:"SERIAL" pgA:"PRIMARY KEY" db:"project_id" json:"project_id"`
	ProjectName      string `pgType:"varchar(256)" db:"project_name" json:"project_name"`
	ProjectCode      string `pgType:"varchar(32)" db:"project_code" json:"project_code"`
	DepartmentCode   string `db:"department_code" json:"department_code"`
	Department       string `db:"department" json:"department"`
	ChiefID          int    `db:"chief_id" json:"chief_id"`
	ChiefChineseName string `db:"chief_cn_name" json:"chief_cn_name"`
	ExtraInfo        string `db:"extra_info" json:"extra_info"`

	// Status 4
	BasicStatus int `db:"basic_status" json:"basic_status"`

	// Apply Info 6
	StartDate           time.Time `db:"start_date" json:"start_date"`
	TotalDaysApply      int       `db:"total_days_apply" json:"total_days_apply"`
	EndReminderAt       time.Time
	CpuNodesExpected    int `db:"cpu_nodes_expected" json:"cpu_nodes_expected"`
	GpuNodesExpected    int `db:"gpu_nodes_expected" json:"gpu_nodes_expected"`
	StorageSizeExpected int `db:"storage_size_expected" json:"storage_size_expected"`

	// Alloc 3+2
	CpuNodesAcquired    int `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int
	StorageSizeAcquired int    `db:"storage_size_acquired" json:"storage_size_acquired"`
	AccountAllocInfo    string `db:"account_alloc_info" json:"account_alloc_info"`
	StorageAllocInfo    string

	CreatedAt time.Time
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type User_ToG struct {
	ID          int    `pgType:"SERIAL" pgA:"PRIMARY KEY" db:"id" json:"id"`
	ChineseName string `db:"cn_name" json:"cn_name"`
	Name        string // user's name
}
