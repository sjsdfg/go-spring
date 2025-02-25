module github.com/go-spring/starter-go-mongo

go 1.14

require (
	github.com/go-spring/spring-base v1.1.0-rc2
	github.com/go-spring/spring-core v1.1.0-rc2
	go.mongodb.org/mongo-driver v1.7.3
)

replace github.com/go-spring/spring-core => ../../spring/spring-core
