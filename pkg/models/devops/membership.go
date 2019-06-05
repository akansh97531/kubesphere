/*
Copyright 2019 The KubeSphere Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package devops

const (
	DevOpsProjectMembershipTableName       = "project_membership"
	DevOpsProjectMembershipUsernameColumn  = "project_membership.username"
	DevOpsProjectMembershipProjectIdColumn = "project_membership.project_id"
	DevOpsProjectMembershipRoleColumn      = "project_membership.role"
)

type DevOpsProjectMembership struct {
	Username  string `json:"username" description:"member's username，username can uniquely identify a user"`
	ProjectId string `json:"project_id" db:"project_id" description:"the devops projects which project membership belongs to"`
	Role      string `json:"role" description:"devops project membership's role type'"`
	Status    string `json:"status" description:"Desperated, status of project membership"`
	GrantBy   string `json:"grand_by,omitempty" description:"Username of the user who assigned the role"`
}

var DevOpsProjectMembershipColumns = GetColumnsFromStruct(&DevOpsProjectMembership{})

func NewDevOpsProjectMemberShip(username, projectId, role, grantBy string) *DevOpsProjectMembership {
	return &DevOpsProjectMembership{
		Username:  username,
		ProjectId: projectId,
		Role:      role,
		Status:    StatusActive,
		GrantBy:   grantBy,
	}
}
