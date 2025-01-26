# Careerise Backend


### File structure

```
.
├── Makefile 
├── README.md
├── go.mod
├── handlers   <--(http route handlers)
├── internal 
│   ├── controller <--(business logics)
│   ├── middleware <--(middleware)
│   └── model
├── main.go
├── migrations <--(db migrations)
├── mysql <--(docker volume for mysql)
├── temp
└── views
```

#### Current Features (before january 21)

- [ ] db schema
- [ ] login system (cookie based)
- [ ] filtering of mentor
- [ ] booking of a mentor

#### Conventions
- 1. Everyone is a student until he applied to be a mentor
- 2. only name, email and gender needed to create profile
- 3. rest of the profile details can be set after
- 4.    
