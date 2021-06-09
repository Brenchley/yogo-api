CREATE TYPE status AS ENUM (
  'active',
  'inactive'
);
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email text NOT NULL,
    username text NOT NULL,
    profile_pic text,
    status status  NOT NULL,
    createdAt timestamp NOT NULL DEFAULT NOW(),
    updatedAt timestamp  
);

CREATE TABLE interests (
    id SERIAL PRIMARY KEY,
    interest_name text NOT NULL,
    interest_img text NOT NULL,
    updatedAt timestamp  
);

CREATE TABLE places (
    id SERIAL PRIMARY KEY,
    place_name text NOT NULL,
    location text NOT NULL,
    location_name text NOT NULL,
    palce_img text NOT NULL,
    interest_id SERIAL NOT NULL REFERENCES interests(id)
);

CREATE TABLE trips(
    id SERIAL PRIMARY KEY,
    trip_name text NOT NULL,
    cost int NOT NULL,
    start_date timestamp NOT NULL,
    status status  NOT NULL,
    orgernizer SERIAL NOT NULL REFERENCES users(id)
);

CREATE  TABLE trip_members(
    id SERIAL PRIMARY KEY,
    trip_id SERIAL NOT NULL REFERENCES trips(id),
    member SERIAL NOT NULL REFERENCES users(id)
);