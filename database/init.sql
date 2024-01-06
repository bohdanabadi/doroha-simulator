CREATE EXTENSION IF NOT EXISTS postgis;

CREATE SCHEMA IF NOT EXISTS traffic_db;
--
-- -- Set the search path to include your schema and public
-- SET search_path TO traffic_db, public;

CREATE TABLE IF NOT EXISTS journeys (
                                        id SERIAL PRIMARY KEY,
                                        starting_point_x FLOAT NOT NULL,
                                        starting_point_y FLOAT NOT NULL,
                                        ending_point_x FLOAT NOT NULL,
                                        ending_point_y FLOAT NOT NULL,
                                        distance FLOAT NOT NULL,
                                        date_create TIMESTAMP NOT NULL,
                                        attempts SMALLINT,
                                        status VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS road_map_points (
                                               point GEOMETRY NOT NULL
);
