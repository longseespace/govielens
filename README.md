# govielens

Go client library for accessing the Movielens unpublished API.

[![Build Status](https://travis-ci.org/longseespace/govielens.svg?branch=master)](https://travis-ci.org/longseespace/govielens)

### Usage

```go
import "github.com/longseespace/govielens/movielens"

client := movielens.NewClient(nil)
err := client.Login("your@email.com", "password")
if err != nil {
	my, _ := client.GetMe()
	fmt.Printf("Hi, my email is %s", my.Email)
}
```

### Integration Tests

You can run integration tests from the `tests` directory.