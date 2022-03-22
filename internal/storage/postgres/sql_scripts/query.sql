-- name: AddPerson :execresult
INSERT INTO persons (
	person_name, phone, addressID
) VALUES (
	?, ?, ?
);