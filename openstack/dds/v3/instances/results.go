package instances

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type Instance struct {
	Id               string         `json:"id"`
	Name             string         `json:"name"`
	DataStore        DataStore      `json:"datastore"`
	Status           string         `json:"status"`
	Region           string         `json:"region"`
	AvailabilityZone string         `json:"availability_zone"`
	VpcId            string         `json:"vpc_id"`
	SubnetId         string         `json:"subnet_id"`
	SecurityGroupId  string         `json:"security_group_id"`
	DiskEncryptionId string         `json:"disk_encryption_id"`
	Mode             string         `json:"mode"`
	Flavor           []Flavor       `json:"flavor"`
	BackupStrategy   BackupStrategy `json:"backup_strategy"`
}

func (r CreateResult) Extract() (*Instance, error) {
	var response Instance
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteInstanceResult struct {
	commonResult
}

type DeleteInstanceResponse struct {
	JobId string `json:"job_id"`
}

func (r DeleteInstanceResult) Extract() (*DeleteInstanceResponse, error) {
	var response DeleteInstanceResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListInstanceResult struct {
	commonResult
}

type ListInstanceResponse struct {
	Instances  []InstanceResponse `json:"instances"`
	TotalCount int                `json:"total_count"`
}

type InstanceResponse struct {
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	Status            string         `json:"status"`
	Port              int            `json:"port"`
	Mode              string         `json:"mode"`
	Region            string         `json:"region"`
	DataStore         DataStore      `json:"datastore"`
	Engine            string         `json:"engine"`
	Created           string         `json:"created"`
	Updated           string         `json:"updated"`
	DbUserName        string         `json:"db_user_name"`
	Ssl               int            `json:"ssl"`
	VpcId             string         `json:"vpc_id"`
	SubnetId          string         `json:"subnet_id"`
	SecurityGroupId   string         `json:"security_group_id"`
	BackupStrategy    BackupStrategy `json:"backup_strategy"`
	MaintenanceWindow string         `json:"maintenance_window"`
	Groups            []Group        `json:"groups"`
	DiskEncryptionId  string         `json:"disk_encryption_id"`
	TimeZone          string         `json:"time_zone"`
}

type Group struct {
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Volume Volume `json:"volume"`
	Nodes  Nodes  `json:"nodes"`
}

type Volume struct {
	Size string `json:"size"`
	Used string `json:"used"`
}

type Nodes struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	Role             string `json:"role"`
	PrivateIP        string `json:"private_ip"`
	PublicIP         string `json:"public_ip"`
	SpecCode         string `json:"spec_code"`
	AvailabilityZone string `json:"availability_zone"`
}

type InstancePage struct {
	pagination.SinglePageBase
}

func (r InstancePage) IsEmpty() (bool, error) {
	data, err := ExtractInstances(r)
	if err != nil {
		return false, err
	}
	return len(data.Instances) == 0, err
}

func ExtractInstances(r pagination.Page) (ListInstanceResponse, error) {
	var s ListInstanceResponse
	err := (r.(InstancePage)).ExtractInto(&s)
	return s, err
}