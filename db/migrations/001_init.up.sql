CREATE TABLE houses (
    id VARCHAR PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    head_of_house TEXT
);

CREATE TABLE students (
    id VARCHAR PRIMARY KEY UNIQUE,
    first_name TEXT NOT NULL UNIQUE,
    last_name TEXT NOT NULL UNIQUE,
    age INT,
    house_id VARCHAR REFERENCES houses(id)
);
