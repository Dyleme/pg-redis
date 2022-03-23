-- name: AddPerson :execresult
INSERT INTO persons (
	person_name, phone, addressID
) VALUES (
	?, ?, ?
);

-- name: PersonList :many
SELECT * FROM persons
ORDER BY person_name;

-- name: PersonByID :one
SELECT * FROM persons as p 
LEFT JOIN addresses as a ON p.addressID = a.id
WHERE p.id = ?;

-- name: DeletePerson :exec
DELETE FROM persons 
WHERE id = ?;