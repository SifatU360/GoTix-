### GO Tickets

- Project Structure
- Database Connection
- Password Hashing
- JWT Authentication
- Request Validation
- Service Layer
- Repository Layer
- Database Relationship
- Concurrency Handling

---
### Tech Stack
- Go
- Echo
- PostgreSQL
- GORM

---
### Folder Structure

```
GoTix/
├── cmd/           // application entry point, config load, database connect, server create, dependancy warren
│   ├── components/
│   │   └── Header.js
│   ├── main.go
│   └── styles.css
├── internal/                // private code
│   └── index.html
├── .gitignore
├── go.mod
├── go.sum
└── Requirements.md
```


>go mod init <project name / github url>

>go mod tidy

---
## Routing
```
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```

## Path Parameters
```
// e.GET("/users/:id", getUser)
func getUser(c *echo.Context) error {
    // user ID from path `user/:id`
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}
```

## Query Parameters

`/show?team=x-men&member=wolverine`
```
// e.GET("/show", show)
func show(c *echo.Context) error {
    // Get team and member from the query string
    team := c.QueryParam("team")
    member := c.QueryParam("member")
    return c.String(http.StatusOK, "team: " + team + ", member: " + member)
}
```

## Form application
`POST` `/save`

| name  | value          |
| ------|:--------------:|
| name  | Sifatullah     |
| email | sifat@gmail.com|

```
 //e.POST("/save", save)
 func save(c *echo.Context) error {
    // Get name and email
    name := c.FormValue("name")
    email := c.FormValue("email")
    return c.String(http.StatusOK, "Name: " + name + ", Email: " + email)
 }
```

Run the following command:
```
$ curl -d "name=Sifatullah" -d "email=sifat@gmail.com" http://localhost:5000/save 
// => Name:Sifatullah, Email: sifat@gmail.com
```

## Handling Request
- Bind `json`, `xml`, `from` or `query` payload into Go struct based on `Content-Type` request hader.
- Render response as `json` or `xml` with status code.

```
type User struct {
    Name string `json:"name" xml:"name" form:"name" query:"name"`
    Email string `json:"email" xml:"email" form:"email" query:"email"`
}

e.POST("/users", func(c *echo.Context) error{
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, u)

    // return c.XML(http.StatusCreated, u)
})
```

>> Request -> Handler -> Service -> Repository