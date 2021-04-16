configure: ## Download dependencies
	go mod download

lint: configure ## Lint the repository with golang-ci lint
	golangci-lint run --max-same-issues 0 --max-issues-per-linter 0 $(if $(CI),--out-format code-climate > gl-code-quality-report.json 2>golangci-stderr-output)

test: configure ## Run all tests
	go test -v

bench: configure ## Run all benchmarks
	go test -bench=.