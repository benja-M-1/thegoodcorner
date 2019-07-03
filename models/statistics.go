package models

type Statistic struct {
	Request *Request `json:"request"`
	Hits    int      `json:"hits"`
}

func (db *DB) AllStatistics() ([]*Statistic, error) {
	q := "SELECT count(id) as Hits, int1, int2, str1, str2 FROM requests GROUP BY int1, int2, str1, str2 ORDER BY Hits DESC"
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*Statistic

	for rows.Next() {
		s := &Statistic{&Request{}, 0}
		err := rows.Scan(&s.Hits, &s.Request.Int1, &s.Request.Int2, &s.Request.Str1, &s.Request.Str2)
		if err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	return stats, nil
}
