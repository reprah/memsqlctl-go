# memsqlctl-go

This CLI tool implements  memsqlctl's `show-leaves`, `set-license` and `show-license` commands in Go.

## show-leaves
```
$ ./memsqlctl-go show-leaves
+-----------+------+--------------------+-----------+-----------+--------+--------------------+------------------------------+--------+-------------------------+
|   HOST    | PORT | AVAILABILITY GROUP | PAIR HOST | PAIR PORT | STATE  | OPENED CONNECTIONS | AVERAGE ROUNDTRIP LATENCY MS | NODEID | GRACE PERIOD IN SECONDS |
+-----------+------+--------------------+-----------+-----------+--------+--------------------+------------------------------+--------+-------------------------+
| 127.0.0.1 | 3307 | 1                  |           |           | online | 1                  | 0.688                        | 2      |                         |
+-----------+------+--------------------+-----------+-----------+--------+--------------------+------------------------------+--------+-------------------------+
```

## set-license
```
# not a real license

$ ./memsqlctl-go set-license --license <base-64 encoded license>
Set license to <base-64 encoded license>
```

## show-license
```
# not a real license or key

$ ./memsqlctl-go show-license
+-----------------------------+--------------------------------------------------------------------------------------------------------------------------------------------------+
| License                     | ADMzZTc1MGFmYWRjODQ5NTA4NTExZWExMmU1OGYxZmQ3AAAAAAAAAAAEAAAAAAAAAAwwNAIYY27TF54sauDgqbA6hjoiRVZ0R8OVqxxHAhj80+jaNUMLl8MZ/PLYgbQiFk7Ahe3cMZoAAA== |
| License_version             | 4                                                                                                                                                |
| License_capacity            | 4 units                                                                                                                                          |
| Used_instance_license_units | 0                                                                                                                                                |
| License_expiration          | 0                                                                                                                                                |
| License_key                 | 43e750afbdc899508511ea42e5831fd8                                                                                                                 |
| License_type                | free                                                                                                                                             |
+-----------------------------+--------------------------------------------------------------------------------------------------------------------------------------------------+
```

## Installation

Requirements:

- Go
- MemSQL's [cluster-in-a-box](https://hub.docker.com/r/memsql/cluster-in-a-box) Docker image, or a MemSQL cluster listening on port 3306 with default credentials

Steps:

1. `git clone <this repo's git URL>`
2. `cd memsqlctl-go`
3. `go build -o memsqlctl-go` (or `go get`)
