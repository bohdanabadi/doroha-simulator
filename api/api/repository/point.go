package repository

import (
	"database/sql"
	"errors"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/apperror"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/model"
	"github.com/bohdanabadi/Traffic-Simulation/api/db"
	"github.com/bohdanabadi/Traffic-Simulation/api/observibility"
	"time"
)

func GetRandomStartAndEndPoints() (model.Point, error) {
	var point model.Point

	sqlQuery := `
		WITH random_points AS (
			SELECT point
			FROM road_map_points
			ORDER BY RANDOM()   -- Adjust as necessary
			LIMIT 2
		),
			 starting_point AS (
				 SELECT point
				 FROM random_points
				 LIMIT 1
			 ),
			 ending_point AS (
				 SELECT point
				 FROM random_points
				 OFFSET 1
					 LIMIT 1
			 )
		SELECT
			ST_X(starting_point.point) AS "startingpointx", ST_Y(starting_point.point) AS "startingpointy",
			ST_X(ending_point.point) AS "endingpointx", ST_Y(ending_point.point) AS "endingpointy",
			ST_Distance(starting_point.point::geography, ending_point.point::geography) as "distance"
		FROM starting_point, ending_point
		WHERE
				ST_Distance(starting_point.point::geography, ending_point.point::geography) > 1000
		LIMIT 1;
	`

	//Retry Mechanism the query might return empty result as the 2 points are distances less than 1000
	const maxRetries = 10
	retryCount := 0
	m := observibility.GetMetrics()
	startTime := time.Now()
	for retryCount < maxRetries {
		row := db.DB.Raw(sqlQuery).Row() // (*sql.Row)
		err := row.Scan(&point.StartingPointX, &point.StartingPointY, &point.EndingPointX, &point.EndingPointY, &point.Distance)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				//No result, retry
				retryCount++
				continue
			} else {
				return model.Point{}, err
			}
		} else {
			duration := time.Since(startTime)
			m.DBQueryLatency.Observe(duration.Seconds())
			return point, nil
		}
	}
	//Max Retries reached and no result returned
	return model.Point{}, apperror.NewAppError(503, "Max retries reached without obtaining a result.", sql.ErrNoRows)
}
