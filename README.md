# Email Indexer

- This is a simple program that gets data from [Enron Email Dataset](https://www.cs.cmu.edu/~enron/) just for learning
  purposes.

## Objective:

- The objective of this project is to index the dataset into `Zincsearch` to then use it as a powerful tool for
  searching.
- For this project the scope I will consider only the `inbox` and `sent_items` folders for indexes.

## Structure:

```
mamuro-backend
├── config/                # env vars and things related to global project configuration
│   ├── ...
│   └── utils.go
│
├── profiling/              # code versions that will be updated while performance is improved
│   └── v1/                 # version of code
│       └── indexing.go     # logic code
│
├── test_file/              # I took some users for the benchmark. Those users data are here.
│   ├── user1
│   ├── ...
│   └── userN
│
├── zincsearch/             # Zincsearch logic to connect to buildin API 
│   ├── ...
│   └── bulk_v2.go
│
├── models/                 # Models of the documents that will be indexed
│   ├── ...
│   └── modelN.go
│
├── go.mod
├── main.go
└── README.md
```