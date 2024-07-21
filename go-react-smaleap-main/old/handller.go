package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/SeiryuTanaka/go-react-smallapp/internal/database"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GradeParams struct {
	Name      string `json:"name"`
	SubjectID int    `json:"subject_id"`
	Grade     int    `json:"grade" validate:"required"`
}

// GradeParams のバリデーションを行う
func (r GradeParams) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		// SubjectID:必須フィールド, 1<=SubjectID<=3
		validation.Field(&r.SubjectID, validation.Required, validation.Min(1), validation.Max(3)),
		// Grade:必須フィールド, 0<=SubjectID<=100
		validation.Field(&r.Grade, validation.Min(0), validation.Max(100)),
	)
}

func getGradesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := Queries.GetGradesAll(context.Background())
	if err != nil {
		log.Println("データの取得に失敗", err)
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}

func registerGradeHandler(w http.ResponseWriter, r *http.Request) {
	var params GradeParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err)
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}
	//バリデーションエラーがある場合にHTTPレスポンスを返す
	if err := params.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var processedParam database.RegisterGradeParams
	processedParam.Name = params.Name
	processedParam.SubjectID = int32(params.SubjectID)
	processedParam.Grade = int32(params.Grade)

	data, err := Queries.RegisterGrade(context.Background(), processedParam)
	if err != nil {
		log.Fatal("データの登録に失敗")
	}
	respondWithJSON(w, http.StatusOK, data)

}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("レスポンスの作成に失敗しました")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}
