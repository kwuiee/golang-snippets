package main

import "fmt"
import "github.com/pelletier/go-toml"
import "github.com/pelletier/go-toml/query"

type Postgres struct {
	User     string
	Password string
}

type Config struct {
	Postgres Postgres
}

func main() {
	config, _ := toml.Load(`
	[postgres]
	user = "pelletier"
	password = "mypassword"`)
	// retrieve data directly
	user := config.Get("postgres.user").(string)
	fmt.Println("user=", user)

	// or using an intermediate object
	postgresConfig := config.Get("postgres").(*toml.Tree)
	password := postgresConfig.Get("password").(string)
	fmt.Println("password=", password)

	doc := []byte(`
[Postgres]
User = "pelletier"
Password = "mypassword"`)

	config2 := Config{}
	toml.Unmarshal(doc, &config2)
	fmt.Println("user=", config2.Postgres.User)

	// use a query to gather elements without walking the tree
	q, _ := query.Compile("$..[user,password]")
	results := q.Execute(config)
	for ii, item := range results.Values() {
		fmt.Printf("Query result %d: %v\n", ii, item)
	}
}
