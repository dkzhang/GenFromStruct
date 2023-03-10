package mypack

import "time"

type Info struct {
	ProjectID           int       `db:"project_id" json:"project_id"`
	ProjectName         string    `db:"project_name" json:"project_name"`
	ProjectCode         string    `db:"project_code" json:"project_code"`
	DepartmentCode      string    `db:"department_code" json:"department_code"`
	Department          string    `db:"department" json:"department"`
	ChiefID             int       `db:"chief_id" json:"chief_id"`
	ChiefChineseName    string    `db:"chief_cn_name" json:"chief_cn_name"`
	ExtraInfo           string    `db:"extra_info" json:"extra_info"`
	BasicStatus         int       `db:"basic_status" json:"basic_status"`
	StartDate           time.Time `db:"start_date" json:"start_date"`
	TotalDaysApply      int       `db:"total_days_apply" json:"total_days_apply"`
	EndReminderAt       time.Time `db:"end_reminder_at" json:"end_reminder_at"`
	CpuNodesExpected    int       `db:"cpu_nodes_expected" json:"cpu_nodes_expected"`
	GpuNodesExpected    int       `db:"gpu_nodes_expected" json:"gpu_nodes_expected"`
	StorageSizeExpected int       `db:"storage_size_expected" json:"storage_size_expected"`
	CpuNodesAcquired    int       `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int       `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired int       `db:"storage_size_acquired" json:"storage_size_acquired"`
	AccountAllocInfo    string    `db:"account_alloc_info" json:"account_alloc_info"`
	StorageAllocInfo    string    `db:"storage_alloc_info" json:"storage_alloc_info"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at"`
}
type User struct {
	ID          int    `db:"id" json:"id"`
	ChineseName string `db:"cn_name" json:"cn_name"`
	Name        string `db:"name" json:"name"` // user's name

}

var SchemaInfo = ` 
		CREATE TABLE info ( 
			project_id SERIAL PRIMARY KEY, 
			project_name varchar(256), 
			project_code varchar(32), 
			department_code text, 
			department text, 
			chief_id integer, 
			chief_cn_name text, 
			extra_info text, 
			basic_status integer, 
			start_date TIMESTAMP WITH TIME ZONE, 
			total_days_apply integer, 
			end_reminder_at TIMESTAMP WITH TIME ZONE, 
			cpu_nodes_expected integer, 
			gpu_nodes_expected integer, 
			storage_size_expected integer, 
			cpu_nodes_acquired integer, 
			gpu_nodes_acquired integer, 
			storage_size_acquired integer, 
			account_alloc_info text, 
			storage_alloc_info text, 
			created_at TIMESTAMP WITH TIME ZONE, 
			updated_at TIMESTAMP WITH TIME ZONE, 
); 
`
var SchemaUser = ` 
		CREATE TABLE user ( 
			id SERIAL PRIMARY KEY, 
			cn_name text, 
			name text, 
); 
`
