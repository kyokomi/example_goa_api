test:
	go test -cover $$(go list ./... | grep -vE '(/vendor/|/mock|/gen|/_design)')

lint:
	go list ./... | grep -vE '(/vendor/|/mock|/gen|/_design)' | xargs -L1 golint

vet:
	go vet $$(go list ./... | grep -vE '(/vendor/|/mock|/gen|/_design)')

