
<img src="/public/ghostly.jpg"  width="300" height="300">

# Ghostly
-----------------------------------------------------------------------------
## Ghostly is a simple, lightweight, and fast full-stack framework for Golang

## Functionality: 
> Object Relation Mapper (ORM) that is database agnostic

> A fully functional Database Migration system

> A fully featured user authentication system that can be installed with a single command, which includes:

> A password reset system

> Session based authentication (for web based applications)

> Token based authentication (for APIs and systems built with front ends like React and Vue)

> A fully featured templating system (using both Go templates and Jet templates)

> A complete caching system that supports Redis and Badger

> Easy session management, with cookie, database (MySQL and Postgres), Redis stores

> Simple response types for HTML, XML, JSON, and file downloads

> Form validation

> JSON validation

> A complete mailing system which supports SMTP servers, and third party APIs including MailGun, SparkPost, and SendGrid

> A command line application which allows for easy generation of emails, handlers, database models

> the command line application will allow us to create a ready-to-go web application by tying a single command: ghostly new <myproject>

## Notice
There is coverage and CI for both Linux, Mac and Windows environments, but I make no guarantees about the bin version working on Windows.
Must be Go version 1.17 or higher

## Installation

As a library

```shell
go get github.com/dominic-wassef/ghostly@latest
```

or if you want to use it as a bin command I will list the exact steps below:


Step 1. 
Make a workfolder on your Desktop and cd into it
```shell
mkdir Ghostly-App
```
```shell
cd Ghostly-App
```

Step 2. 
Clone the repository 
```shell
git clone git@github.com:Dominic-Wassef/ghostly.git
```

Step 3. 
cd into directory and build the binary with the Makefile at root level of the ghostly project
```shell
cd ghostly
```
```shell
make build
```

Step 4. 
cd into the dist directory of the ghostly application and copy it to your Desktop
```shell
cd dist
```
```shell
cp ./ghostly ~/Desktop
```

## Usage

Once above steps have been followed, you can show all ghostly command by going to your Desktop and run:
```shell
./ghostly
```

Making a new project:
```shell
./ghostly new $("PROJECT-NAME")
```

Then cd into your newly made Go project:
```shell
cd $("PROJECT-NAME")
```

Run the project by using the makefile in your new project directory
```shell
make start
```

Here are the types for the Ghostly Framework

```go
type Ghostly struct {
	AppName       string
	Debug         bool
	Version       string
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	RootPath      string
	Routes        *chi.Mux
	Render        *render.Render
	Session       *scs.SessionManager
	DB            Database
	JetViews      *jet.Set
	config        config
	EncryptionKey string
	Cache         cache.Cache
	Scheduler     *cron.Cron
	Mail          mailer.Mail
	Server        Server
}
```

Below types are for Server and Config:

```go
type Server struct {
	ServerName string
	Port       string
	Secure     bool
	URL        string
}

type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
	database    databaseConfig
	redis       redisConfig
}
```

For full documentation please refer to the package on:
[Ghostly Documentation](https://pkg.go.dev/github.com/dominic-wassef/ghostly@v1.3.0)

## Who?

The full library [ghostly](https://github.com/dominic-wassef/ghostly) was written by [Dominic-Wassef](https://github.com/Dominic-Wassef)