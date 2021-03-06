// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package codecommitiface provides an interface for the AWS CodeCommit.
package codecommitiface

import (
	"github.com/djannot/aws-sdk-go/aws/request"
	"github.com/djannot/aws-sdk-go/service/codecommit"
)

// CodeCommitAPI is the interface type for codecommit.CodeCommit.
type CodeCommitAPI interface {
	BatchGetRepositoriesRequest(*codecommit.BatchGetRepositoriesInput) (*request.Request, *codecommit.BatchGetRepositoriesOutput)

	BatchGetRepositories(*codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error)

	CreateBranchRequest(*codecommit.CreateBranchInput) (*request.Request, *codecommit.CreateBranchOutput)

	CreateBranch(*codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error)

	CreateRepositoryRequest(*codecommit.CreateRepositoryInput) (*request.Request, *codecommit.CreateRepositoryOutput)

	CreateRepository(*codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error)

	DeleteRepositoryRequest(*codecommit.DeleteRepositoryInput) (*request.Request, *codecommit.DeleteRepositoryOutput)

	DeleteRepository(*codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error)

	GetBranchRequest(*codecommit.GetBranchInput) (*request.Request, *codecommit.GetBranchOutput)

	GetBranch(*codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error)

	GetCommitRequest(*codecommit.GetCommitInput) (*request.Request, *codecommit.GetCommitOutput)

	GetCommit(*codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error)

	GetRepositoryRequest(*codecommit.GetRepositoryInput) (*request.Request, *codecommit.GetRepositoryOutput)

	GetRepository(*codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error)

	GetRepositoryTriggersRequest(*codecommit.GetRepositoryTriggersInput) (*request.Request, *codecommit.GetRepositoryTriggersOutput)

	GetRepositoryTriggers(*codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error)

	ListBranchesRequest(*codecommit.ListBranchesInput) (*request.Request, *codecommit.ListBranchesOutput)

	ListBranches(*codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error)

	ListBranchesPages(*codecommit.ListBranchesInput, func(*codecommit.ListBranchesOutput, bool) bool) error

	ListRepositoriesRequest(*codecommit.ListRepositoriesInput) (*request.Request, *codecommit.ListRepositoriesOutput)

	ListRepositories(*codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error)

	ListRepositoriesPages(*codecommit.ListRepositoriesInput, func(*codecommit.ListRepositoriesOutput, bool) bool) error

	PutRepositoryTriggersRequest(*codecommit.PutRepositoryTriggersInput) (*request.Request, *codecommit.PutRepositoryTriggersOutput)

	PutRepositoryTriggers(*codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error)

	TestRepositoryTriggersRequest(*codecommit.TestRepositoryTriggersInput) (*request.Request, *codecommit.TestRepositoryTriggersOutput)

	TestRepositoryTriggers(*codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error)

	UpdateDefaultBranchRequest(*codecommit.UpdateDefaultBranchInput) (*request.Request, *codecommit.UpdateDefaultBranchOutput)

	UpdateDefaultBranch(*codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error)

	UpdateRepositoryDescriptionRequest(*codecommit.UpdateRepositoryDescriptionInput) (*request.Request, *codecommit.UpdateRepositoryDescriptionOutput)

	UpdateRepositoryDescription(*codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error)

	UpdateRepositoryNameRequest(*codecommit.UpdateRepositoryNameInput) (*request.Request, *codecommit.UpdateRepositoryNameOutput)

	UpdateRepositoryName(*codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error)
}

var _ CodeCommitAPI = (*codecommit.CodeCommit)(nil)
