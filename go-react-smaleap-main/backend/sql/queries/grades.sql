-- name: GetGradesAll :many
SELECT * FROM grades;

-- name: RegisterGrade :one
INSERT INTO grades (
    name,
    subject_id,
    grade
) VALUES (
    $1, $2, $3
)
RETURNING *;
