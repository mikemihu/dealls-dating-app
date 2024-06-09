package mocks

//go:generate mockgen -source=../internal/repository.go -destination=mock_repository.go -package=mocks
//go:generate mockgen -source=../internal/usecase.go -destination=mock_usecase.go -package=mocks
