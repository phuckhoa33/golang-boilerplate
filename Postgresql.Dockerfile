# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set environment variables for the PostgreSQL database
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD 332003
ENV POSTGRES_DB postgres

# Copy the SQL script to initialize the database
COPY migration/postgresql/relationship-tables.sql /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port
EXPOSE 5432