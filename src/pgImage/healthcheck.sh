#!/bin/bash
set -e

pg_isready -U "$ExpDB_USER" -d "$ExpDB_NAME"
