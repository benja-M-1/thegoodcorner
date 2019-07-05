CREATE TABLE fizzbuzz_requests (
    id serial PRIMARY KEY,
    int1 BIGINT NOT NULL,
    int2 BIGINT NOT NULL,
    str1 VARCHAR (355) NOT NULL,
    str2 VARCHAR (355) NOT NULL
);

CREATE INDEX request ON fizzbuzz_requests (int1, int2, str1, str2);

/* Populate the table with many values */
DO $$
BEGIN
    FOR inserts IN 1..10 LOOP
        INSERT INTO fizzbuzz_requests(int1, int2, str1, str2) SELECT generate_series(1,10) AS int1, generate_series(1, 10) as int2, 'fizz', 'buzz';
    END LOOP;
END; $$
/************************************************************/
/* I used the following lines to debug the SQL performances */
/************************************************************/

/*
EXPLAIN INSERT INTO fizzbuzz_requests(int1, int2, str1, str2) VALUES (2, 4, 'buzz', 'fizz');
EXPLAIN SELECT count(id) as Hits, int1, int2, str1, str2 FROM fizzbuzz_requests GROUP BY int1, int2, str1, str2 ORDER BY Hits DESC;
*/