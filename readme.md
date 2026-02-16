my-ls/
│
├── main.go
│
├── internal/
│   ├── app/
│   │   └── run.go
│   │
│   ├── flags/
│   │   └── parser.go
│   │
│   ├── filesystem/
│   │   ├── reader.go
│   │   ├── stat.go
│   │   └── recursive.go
│   │
│   ├── sorter/
│   │   └── sort.go
│   │
│   ├── formatter/
│   │   ├── long.go
│   │   └── basic.go
│   │
│   └── utils/
│       └── helpers.go
│
└── go.mod

////////////////////////////////////mars

main → app.Run()
          ↓
      parse flags
          ↓
      for each path:
          ↓
      read entries
          ↓
      filter hidden
          ↓
      sort
          ↓
      print (basic or long)
          ↓
      if recursive → walk

    ////////////////////////////////////note!
    starting with flag -R makes my ls use func recall it selfe recuresively;
    so each time print > read > sort > print > for each dir: recall again..
    /////////////////////////////////////////
    program is traversal engine + sorting + formatting pipeline