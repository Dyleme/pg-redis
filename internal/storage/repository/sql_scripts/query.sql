-- name: AddAddress :execresult
INSERT INTO addresses (
	country, city, street, house, apartments
) VALUES (
	?, ?, ?, ?, ?
);

-- name: AddressList :many
SELECT * from addresses
ORDER BY country, city, street, house, apartments DESC;

-- name: AddressById :one
SELECT * from addresses
WHERE id = ?;

-- name: AddressID :one
SELECT id from addresses
WHERE country = ? and city = ? and street = ? and house = ? and apartments <=> ?;


-- name: AddPerson :execresult
INSERT INTO persons (
	person_name, phone, addressID
) VALUES (
	?, ?, ?
);

-- name: PersonList :many
SELECT * FROM persons as p 
LEFT JOIN addresses as a ON p.addressID = a.id;

-- name: PersonByID :one
SELECT * FROM persons as p 
LEFT JOIN addresses as a ON p.addressID = a.id
WHERE p.id = ?;

-- name: DeletePerson :exec
DELETE FROM persons 
WHERE id = ?;