--name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

--name: ListUsers :many
SELECT * FROM users
ORDER BY createdAt;

--name: CreateUser :one
INSERT INTO users (
    email,username, profile_pic, status
) VALUES(
    $1, $2, $3, $4
)
RETURNING *;

--name GetInterest :one
SELECT * FROM interests
WHERE id = $1 LIMIT 1;

--name: ListInterests :many
SELECT * FROM interests
ORDER BY id;

--name: CreateInterest :one
INSERT INTO interests (
    interest_name, interest_img
) VALUES(
    $1, $2
)
RETURNING *;

--name: CreatePlace :one
INSERT INTO places (
    place_name, location,location_name,palce_img,interest_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

--name GetPlaces :many
SELECT * FROM  places
ORDER BY id;