package grade

import (
	"context"

	"github.com/SeiryuTanaka/go-react-smallapp/domain/model"
)

type GradeRepository interface {
	GetGradesAll(ctx context.Context) ([]model.Grade, error)
	RegisterGrade(ctx context.Context, arg model.Grade) (model.Grade, error)
}
