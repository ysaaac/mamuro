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
│   │   │   └── start_profiling.go     # runs the code with profiling flags to compare benchmarks
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