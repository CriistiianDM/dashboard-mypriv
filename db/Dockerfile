#PostgreSQL 15.5
FROM postgres:15.5

# Set environment
ENV POSTGRES_DB=mypriv_go
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

# Init database
COPY ./sql/sql.sql /docker-entrypoint-initdb.d/
