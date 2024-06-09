package mocks

//go:generate mockgen -source=../internal/repository.go -destination=mock_repository_test.go -package=mocks
//go:generate mockgen -source=../internal/usecase.go -destination=mock_usecase_test.go -package=mocks
