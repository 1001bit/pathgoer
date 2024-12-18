package pathmodel

import (
	"context"
	"database/sql"
)

type Path struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
	Stats  []Stat `json:"stats"`
	Id     int32  `json:"id"`
}

func (ps *PathStore) CreatePath(ctx context.Context, userId, pathName string) (string, error) {
	row, err := ps.postgresC.QueryRowContext(ctx, `
		INSERT INTO paths (user_id, name) VALUES ($1, $2) RETURNING id
	`, userId, pathName)
	if err != nil {
		return "", err
	}

	id := ""
	err = row.Scan(&id)

	return id, err
}

func (ps *PathStore) UpdatePath(ctx context.Context, pathId string, name string, public bool, askerId string) error {
	result, err := ps.postgresC.ExecContext(ctx, `
		UPDATE paths
		SET name = $1, public = $2
		WHERE id = $3 AND user_id = $4
	`, name, public, pathId, askerId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (ps *PathStore) DeletePath(ctx context.Context, pathId string, askerId string) error {
	result, err := ps.postgresC.ExecContext(ctx, `
		DELETE FROM paths
		WHERE id = $1 AND user_id = $2
	`, pathId, askerId)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (ps *PathStore) GetPathAndOwnerId(ctx context.Context, askerId, pathId string) (Path, string, error) {
	row, err := ps.postgresC.QueryRowContext(ctx, `
		SELECT id, name, public, user_id
		FROM paths
		WHERE id = $1
		AND (public = true OR user_id = $2)
	`, pathId, askerId)
	if err != nil {
		return Path{}, "", err
	}

	userId := ""
	path := Path{
		Stats: []Stat{},
	}

	if err := row.Scan(&path.Id, &path.Name, &path.Public, &userId); err != nil {
		return Path{}, "", err
	}

	path.Stats, err = ps.GetStats(ctx, path.Id)

	return path, userId, err
}

func (ps *PathStore) GetPaths(ctx context.Context, userId, askerId string) ([]*Path, error) {
	rows, err := ps.postgresC.QueryContext(ctx, `
        SELECT id, name, public
        FROM paths
        WHERE user_id = $1 
        AND (public = true OR user_id = $2)
    `, userId, askerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paths []*Path
	pathMap := make(map[int32]*Path)

	for rows.Next() {
		path := &Path{
			Stats: []Stat{},
		}

		if err := rows.Scan(&path.Id, &path.Name, &path.Public); err != nil {
			return nil, err
		}

		paths = append(paths, path)
		pathMap[path.Id] = paths[len(paths)-1]
	}

	err = ps.FetchStatsIntoPaths(ctx, pathMap)
	return paths, err
}
