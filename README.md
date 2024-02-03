# ag-csv-table-generator
General Assembly CSV Table generator

Used to generate a CSV table with header: `Nr. Crt., Proprietate, Nume Proprietar, Semnătură` based on our current owners database (in CSV format).

The scripts checks which owners are members in the Association and groups their properties by their name.

## Build
Run `/build.sh` script. The executable will be generated in binary/ folder.

## Run
./binary/gen-ag-table <input csv>
