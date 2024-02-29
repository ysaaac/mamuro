# Email Indexer

- This is a simple program that gets data from [Enron Email Dataset](https://www.cs.cmu.edu/~enron/) just for learning
  purposes.

## QuickStart:

- All config needed is put at the `docker-compose.yml` file so it's sufficient with running:

```zsh
docker-compose up -d # I prefer the detached mod, but it's totally optional
```

## Objective:

- The objective of this project is to index the dataset into `Zincsearch` to then use it as a powerful tool for
  searching.
- For this project the scope I will consider only the `inbox` and `sent_items` folders for indexes.

## Profiling

- You can see the difference in times into versions in the path `api/indexer/profiling` and also you can run the tests
  with:

```zsh
go test -v ./indexer/profiling
```

Note: remember to replace the version you want to use, just change `v1.IndexData()` to `v1`, `v2`, `vN`.

- Also, you can run the profiling by:

```zsh
export exportPath="./indexer/profiling/v1" && go test -v -cpuprofile "$exportPath/cpu.prof" ./indexer/profiling && go tool pprof -png "$exportPath/cpu.prof" > "$exportPath/cpu.png"

# This second option deletes the img generated in case it exists to replace with the new one
export exportPath="./indexer/profiling/v1" && go test -v -cpuprofile "$exportPath/cpu.prof" ./indexer/profiling && rm -f "$exportPath/cpu.png" && go tool pprof -png "$exportPath/cpu.prof" > "$exportPath/cpu.png"
```

Note: Notice that `exportPath` is used to set where the profiling files will be saved

## Restore Zincsearch Data

- I added some usefull commands to the `docker-compose.yml` file:

### Restore Data

```zsh
# to run the restore cmd you need to have the
# zincsearch-backup.tar.gz file into the backup folder
# or modify the command at the docker-compose
# url to download backup (I uploaded to google drive because is easier but can be at S3 too or any service):
# https://drive.google.com/file/d/1KLCwC9tNOHv1qb_bnCQhr7BLfgW7XoAH/view?usp=sharing
docker-compose run --rm restore
docker-compose up -d zincsearch
```

### Backup Data

```zsh
docker-compose run --rm backup
```

## Structure:

```
mamuro
├── api/                              # backend api -> Indexing and REST API
│   ├── config/                       # env vars and things related to global project configuration
│   │   ├── ...
│   │   └── utils.go
│   │
│   ├── indexer/               
│   │   ├── profiling/                # code versions that will be updated while performance is improved
│   │   │   ├── v1/                   # version of code
│   │   │   │   └── indexing.go       # indexing logic code
│   │   │   │
│   │   │   ├── .../                 
│   │   │   │   └── indexing.go     
│   │   │   │
│   │   │   ├── vN/                 
│   │   │   │   └── indexing.go     
│   │   │   │
│   │   │   └── profiling_test.go     # runs the code with profiling flags to compare benchmarks
│   │   │   
│   │   └── test_files/                # I took some users for the benchmark. Those users data are here.
│   │       ├── user1
│   │       ├── ...
│   │       └── userN
│   │
│   ├── models/                        # all models for the API
│   │
│   ├── src/                           # main logic for the API
│   │   ├── handlers/                  # logic divided for each endpoint
│   │   └── routes.go                  # routes for endpoints
│   │
│   ├── zincsearch/                    # Zincsearch logic to connect to buildin API 
│   │   ├── ...
│   │   └── bulk_v2.go
│   │
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
│
├── client/                            # interface to interact with mails
│   ├── Dockerfile
│   └── .../                           # It follows the standart Vue project structure
│
├── docker-compose.yml
└── README.md
```