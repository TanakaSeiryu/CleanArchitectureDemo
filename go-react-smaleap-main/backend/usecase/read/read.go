package read

//https://qiita.com/mizuko_dev/items/a8a3864e23d82ba2a60d
import (
	"context"

	"github.com/SeiryuTanaka/go-react-smallapp/domain/model"
	domainRepository "github.com/SeiryuTanaka/go-react-smallapp/domain/repository/grade"
)

type GradeReadUsecaseImpl struct {
	GradeRepository domainRepository.GradeRepository
}

func NewGradeReadUseCase(repo domainRepository.GradeRepository) GradeReadUsecase {
	return &GradeReadUsecaseImpl{
		GradeRepository: repo,
	}
}

type GradeReadUsecase interface {
	GetGradesAll(ctx context.Context) ([]model.Grade, error)
}

func (i *GradeReadUsecaseImpl) GetGradesAll(ctx context.Context) ([]model.Grade, error) {
	return i.GradeRepository.GetGradesAll(ctx)
}
