.PHONY: all
all: build test

#
# Logging
#

### Colour Definitions
END_COLOR=\x1b[0m
GREEN_COLOR=\x1b[32;01m
RED_COLOR=\x1b[31;01m
YELLOW_COLOR=\x1b[33;01m

### End output
end:
	@echo "$(YELLOW_COLOR)🙏  🙏  🙏$(END_COLOR)"

#
# Project Initialisation
#

### Name of the executable, it's possible to have multiple executables
APP_EXECUTABLE="passport"

### Get a list of all golang packages
ALL_PACKAGES=$(go list ./... | grep -v "vendor")

#
# Recipes for building existing projects
#

### Clean temporary files
clean:
	@echo "$(GREEN_COLOR)Cleaning unwanted files $(END_COLOR)"
	rm -rf passport.yaml
	rm -rf coverage.txt
	rm -rf coverage.tmp
	rm -rf coverage.html
	rm -rf bin/

### Initialisation project for the first time
init:
	@echo "$(GREEN_COLOR)Initialising dep for the first time $(END_COLOR)"
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint

### Update dependencies
update:
	@echo "$(GREEN_COLOR)Running dep ensure $(END_COLOR)"
	dep ensure

### Fix formatting
fmt:
	@echo "$(GREEN_COLOR)Running fmt $(END_COLOR)"
	go fmt ./...

### Run go vet
vet:
	@echo "$(GREEN_COLOR)Running vet $(END_COLOR)"
	go vet ./...

### Check for linting issues
lint:
	@echo "$(GREEN_COLOR)Running lint $(END_COLOR)"
	go list ./... | xargs -L1 golint

### Copy config from template
copy-config:
	@echo "$(GREEN_COLOR)Copying config from sample $(END_COLOR)"
	cp passport.yaml.sample passport.yaml

db-clean:
	@echo "$(GREEN_COLOR)Cleaning the database $(END_COLOR)"
	$(APP_EXECUTABLE) rollback

db-migrate:
	@echo "$(GREEN_COLOR)Migrating the database to the latest state $(END_COLOR)"
	$(APP_EXECUTABLE) migrate

### Manually test all packages
test:
	@echo "$(GREEN_COLOR)Running tests for all packages $(END_COLOR)"
	go test ./... -v -p=5 -race -covermode=atomic -timeout=30s

### Compile a linux and mac binary in the ./bin folder
compile:
	@echo "$(GREEN_COLOR)Compiling linux and mac binaries in ./bin $(END_COLOR)"
	mkdir -p bin/
	go build -o bin/$(APP_EXECUTABLE) ./cmd/$(APP_EXECUTABLE)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(APP_EXECUTABLE)_linux ./cmd/$(APP_EXECUTABLE)

### Calculate test coverage for the whole project (except vendors)
coverage:
	@echo "$(GREEN_COLOR)Calculating test coverage across packages $(END_COLOR)"
	@echo 'mode: atomic' > coverage.txt && echo '' > coverage.tmp && go list ./... | xargs -n1 -I{} sh -c 'go test -p=5 -race -covermode=atomic -coverprofile=coverage.tmp -timeout=30s {} && tail -n +2 coverage.tmp >> coverage.txt'
	go tool cover -html=coverage.txt -o coverage.html
	@rm coverage.tmp
	@echo "$(YELLOW_COLOR)Run open ./coverage.html to view coverage $(END_COLOR)"

### Install all binaries (Repo could have multiple binaries)
install:
	@echo "$(GREEN_COLOR)Installing all binaries $(END_COLOR)"
	go install ./...

### Build the latest source
build: fmt vet lint coverage install end

### Build the latest source for the first time
build_fresh: clean init update fmt vet lint copy-config compile install db-clean db-migrate coverage end

#
# Receipes for docker
#

build_docker:
	@echo "$(GREEN_COLOR)Building a docker image $(END_COLOR)"
	docker build -t build-tanker/passport .

deploy_quay:
	@echo "$(GREEN_COLOR)Pushing the docker image to Quay.io $(END_COLOR)"
	docker login -u="$(QUAY_USERNAME)" -p="$(QUAY_PASSWORD)" quay.io
	docker build -t build-tanker/passport .
	docker tag build-tanker/passport quay.io/build-tanker/passport
	docker push quay.io/build-tanker/passport

#
# Deploy
#

create_deploy:
	kubectl create -f k8s/secrets/passport-secrets.yaml
	kubectl create -f k8s/passport-deploy.yaml

rolling_update:
	kubectl set image deployments/passport passport=quay.io/build-tanker/passport:latest