package write

//https://qiita.com/mizuko_dev/items/a8a3864e23d82ba2a60d
import (
	domainRepository "github.com/SeiryuTanaka/go-react-smallapp/domain/repository/grade"
)

type GradeWriteUsecaseImpl struct {
	GradeRepository domainRepository.GradeRepository
}
