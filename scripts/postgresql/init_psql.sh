#!/bin/bash

psql -h db -U postgres -a -f scripts/postgresql/create_db.sql
