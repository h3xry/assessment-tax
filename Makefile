run:
	@echo "Running main.go..."
	go run main.go

test-all:
	@echo "Running tests..."
	go test ./...

coverage:
	@echo "Running tests coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

scanner:
	@echo "Running scanner with sonarqube..."
	go test -v ./... -coverprofile=coverage.out;
	sonar-scanner \
	-Dsonar.go.coverage.reportPaths=coverage.out \
	-Dsonar.language=go \
	-Dsonar.sources=. \
	-Dsonar.tests=. \
	-Dsonar.test.inclusions=**/*_test.go \
	-Dsonar.host.url=${SONAR_HOST_URL} \
	-Dsonar.projectKey=${SONAR_PROJECT_KEY} \
	-Dsonar.login=${SONAR_LOGIN};