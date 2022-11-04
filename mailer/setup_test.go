package mailer

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)


var pool *dockertest.Pool
var resource *dockertest.Resource

var mailer = Mail{
	Domain: "localhost",
	Templates: "./testdata/mail",
	Host: "localhost",
	Port: 1026,
	Encryption: "none",
	FromAddress: "me@here.com",
	FromName: "Joe",
	Jobs: make(chan Message, 1),
	Results: make(chan Result, 1),
}

func TestMain(m *testing.M) {
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal("could not connect to docker", err)
	}
	pool = p

	opts := dockertest.RunOptions{
		Repository: "mailhog/mailhog",
		Tag: "latest",
		Env: []string{},
		ExposedPorts: []string{"1025", "8025"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"1025": {
				{HostIP: "0.0.0.0", HostPort: "1026"},
			},
			"8025": {
				{HostIP: "0.0.0.0", HostPort: "8026"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Println(err)
		_ = pool.Purge(resource)
		log.Fatal("Could not start resource")
	}

	time.Sleep(2 * time.Second)

	go mailer.ListenForMail()

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}

	os.Exit(code)
}