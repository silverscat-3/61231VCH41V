#!/bin/bash
while read line
do
    ./61231VCH41V.out learning "${line}" /var/db/61231VCH41V.sqlite3
done < ./output.txt
