# go-mod-sample
Experimenting with ```go modules``` and [sqlboiler](https://github.com/volatiletech/sqlboiler).  

## Requires
- go 1.11 (go modules)
- [sqlboiler](https://github.com/volatiletech/sqlboiler)

## Build & Run
Set database credentials in ```sqlboiler.toml```

Execute ```init.sql```

Run ```go mod download``` or ```go mod vendor```

Build project ```go build```

Execute binary.

## Modification

Steps below should work provided there were no changes to [sqlboiler](https://github.com/volatiletech/sqlboiler).

```
go get -u -t github.com/volatiletech/sqlboiler

# Also install the driver of your choice, there exists pqsl, mysql, mssql
# These are separate binaries.
go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
```

Alter tables in database.

Run ```go generate ./...``` to generate models. For more info head over to  [sqlboiler](https://github.com/volatiletech/sqlboiler).

```//go:generate sqlboiler --wipe --tag db psql```  
- wipe: clears old models
- tag: adds an additional tag called ```db```
- psql: uses postgres driver and configuration