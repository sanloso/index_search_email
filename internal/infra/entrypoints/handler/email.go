package handler

import usecase "emailIndexer/internal/aplication/usecases"

func EmailHandler() error {
	emailUsecase := usecase.NewEmailUseCase()
}
