CREATE TABLE fizzbuzz_requests (
    id serial PRIMARY KEY,
    int1 BIGINT NOT NULL,
    int2 BIGINT NOT NULL,
    str1 VARCHAR (355) NOT NULL,
    str2 VARCHAR (355) NOT NULL
);

CREATE INDEX request ON fizzbuzz_requests (int1, int2, str1, str2);
