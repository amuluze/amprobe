// Package schema
// Date       : 2024/9/4 15:14
// Author     : Amu
// Description:
package schema

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Remark   string `json:"remark"`
	IsAdmin  int    `json:"isAdmin"`
	Status   int    `json:"status"`
	Roles    []Role `json:"roles"`
}

type Role struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Remark    string     `json:"remark"`
	Status    int        `json:"status"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Status int    `json:"status"`
}

type UserQueryArgs struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}

type UserQueryReply struct {
	Data  []User `json:"data"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
}

type UserCreateArgs struct {
	Username string   `json:"username" validate:"required"`
	Remark   *string  `json:"remark" validate:"required"`
	Status   int      `json:"status" validate:"required"`
	IsAdmin  int      `json:"isAdmin" validate:"required"`
	RoleIDs  []string `json:"role_ids" validate:"required"`
}

type UserUpdateArgs struct {
	Username *string  `json:"username"`
	Remark   *string  `json:"remark"`
	Status   *int     `json:"status"`
	IsAdmin  *int     `json:"isAdmin"`
	RoleIDs  []string `json:"role_ids"`
}

type UserDeleteArgs struct {
	ID  *string  `json:"id"`
	IDs []string `json:"ids"`
}

type RoleQueryArgs struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}

type RoleQueryReply struct {
	Data  []Role `json:"data"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
}

type RoleCreateArgs struct {
	Name        string   `json:"name" validate:"required"`
	Remark      string   `json:"remark"`
	Status      int      `json:"status" validate:"required"`
	ResourceIDs []string `json:"resource_ids"`
}

type RoleUpdateArgs struct {
	Name        *string  `json:"name"`
	Remark      *string  `json:"remark"`
	Status      *int     `json:"status"`
	ResourceIDs []string `json:"resource_ids"`
}

type RoleDeleteArgs struct {
	ID  *string  `json:"id"`
	IDs []string `json:"ids"`
}

type ResourceQueryArgs struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}

type ResourceQueryReply struct {
	Data  []Resource `json:"data"`
	Total int        `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}
