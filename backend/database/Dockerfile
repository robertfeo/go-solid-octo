# Use the official PostgreSQL image
FROM postgres:latest

# Set the necessary environment variables for initialization
ENV POSTGRES_DB=library
ENV POSTGRES_USER=library_db
ENV POSTGRES_PASSWORD=password

# Optional: Initialize the database with a script
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432

# This image comes with an ENTRYPOINT that runs the PostgreSQL server
