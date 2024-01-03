FROM postgis/postgis:15-3.3-alpine

# Optional: Copy initialization scripts
COPY ./init.sql /docker-entrypoint-initdb.d/
COPY ./road_map_points.sql /docker-entrypoint-initdb.d/

# Optional: Set environment variables
ENV POSTGRES_USER=trafficUser21
ENV POSTGRES_DB=traffic_db


# Password is Morelanaj
