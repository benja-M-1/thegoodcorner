package models

type Statistic struct {
	Request *FizzbuzzRequest `json:"request"`
	Hits    int              `json:"hits"`
}

func (db *DB) AllStatistics() ([]*Statistic, error) {
	// @Todo There maybe a better way to fetch the data. Here we have to scan all the table AND order it
	// Maybe it is better to update a hit value of the fizzbuzz request instead of writing each request
	// Then we would only have to select the right line.
	// But this would mean to move the most consuming action into the write request which could make
	// the /fizzbuzz end point slower. We don't want that because this endpoint would be more requested than the /statistics one
	q := "SELECT count(id) as hits, int1, int2, str1, str2 FROM fizzbuzz_requests GROUP BY int1, int2, str1, str2 ORDER BY hits DESC"
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*Statistic

	for rows.Next() {
		s := &Statistic{&FizzbuzzRequest{}, 0}
		err := rows.Scan(&s.Hits, &s.Request.Int1, &s.Request.Int2, &s.Request.Str1, &s.Request.Str2)
		if err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	return stats, nil
}
