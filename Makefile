export STRIPE_SECRET=sk_test_dowTjfu0YyTszRI47qYV0sTa
export STRIPE_KEY=pk_test_6oqsofc2sazbHrgCQcfAeiE4
export GOSTRIPE_PORT=4002
export API_PORT=4001
export DSN=root@tcp(localhost:3306)/widgets?parseTime=true&tls=false

## build: builds all binaries
.PHONY: build
build: clean build_front build_back
	@printf "All binaries built!\n"

## clean: cleans all binaries and runs go clean
.PHONY: clean
clean:
	@echo "Cleaning..."
	@- rm -f dist/*
	@go clean
	@echo "Cleaned!"

## build_front: builds the front end
.PHONY: build_front
build_front:
	@echo "Building front end..."
	@go build -o dist/gostripe ./cmd/web
	@echo "Front end built!"

## build_back: builds the back end
.PHONY: build_back
build_back:
	@echo "Building back end..."
	@go build -o dist/gostripe_api ./cmd/api
	@echo "Back end built!"

## start: starts front and back end
.PHONY: start
start: start_front start_back

## start_front: starts the front end
.PHONY: start_front
start_front: build_front
	@echo "Starting the front end..."
	##@env STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET} ./dist/gostripe -port=${GOSTRIPE_PORT} -dsn="${DSN}" &
	@env STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET} ./dist/gostripe -port=${GOSTRIPE_PORT}
	@echo "Front end running!"

## start_back: starts the back end
.PHONY: start_back
start_back: build_back
	@echo "Starting the back end..."
	##@env STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET} ./dist/gostripe_api -port=${API_PORT}  -dsn="${DSN}" &
	@env STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET} ./dist/gostripe_api -port=${API_PORT}
	@echo "Back end running!"

## stop: stops the front and back end
.PHONY: stop
stop: stop_front stop_back
	@echo "All applications stopped"

## stop_front: stops the front end
.PHONY: stop_front
stop_front:
	@echo "Stopping the front end..."
	@-pkill -SIGTERM -f "gostripe -port=${GOSTRIPE_PORT}"
	@echo "Stopped front end"

## stop_back: stops the back end
.PHONY: stop_back
stop_back:
	@echo "Stopping the back end..."
	@-pkill -SIGTERM -f "gostripe_api -port=${API_PORT}"
	@echo "Stopped back end"


