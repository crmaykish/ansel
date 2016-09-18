#!/bin/bash
echo "Setting up SQLite database...";

create()
{
    cat create_tables.sql | sqlite3 ansel.db;
}

if [ ! -e ansel.db ]; then
    create;
else
    read -p "Ansel DB file already exists. Delete and recreate? " answer;
    if [ "$answer" = "Y" ] || [ "$answer" == "y" ]; then
        rm ansel.db -f;
        create;
    fi
fi

echo "Done!";