package handler

import "context"

type segmentCreateService interface {
	Create(ctx context.Context, name string) (int64, error) // используем контекст, в ответе обязательно предусматриваем ошибку
}

type SegmentCreateHandler struct {
	segmentCreateService segmentCreateService
}

func New(segmentCreateService segmentCreateService) *SegmentCreateHandler {
	return &SegmentCreateHandler{segmentCreateService: segmentCreateService}
}

// в хэндлер должен прилететь контекст, и далее по вызовам он будет продолжать передаваться
func (s *SegmentCreateHandler) Handle(ctx context.Context, request string, response int64) error {
	// валидация параметра ???
	//segmentID := s.segmentCreateService.Create(ctx, request)

	return nil
}
