#!/bin/bash

set -e

# This script is used to build the docker image for the application.
# It is used by the CI/CD pipeline to build the image and push it to the registry.
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE TABLE books (
        id integer,
        title varchar(255)
    );

    INSERT INTO books VALUES (1, 'The Great Gatsby');
EOSQL
