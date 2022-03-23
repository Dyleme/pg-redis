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
