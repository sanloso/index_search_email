package usecase

type EmailUseCase struct {
	zincsearchAdapter ZincSearchAdapter
}

func NewEmailUseCase(zinc zincsearchAdapter) *EmailUseCase {
	return &EmailUseCase{
		zincsearchAdapter: zsa,
	}
}
