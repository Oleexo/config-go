.PHONY: lint align

lint:
	golangci-lint run ./...

# Fix struct field alignment (requires betteralign: go install github.com/dkorunic/betteralign/cmd/betteralign@latest)
align:
	@command -v betteralign > /dev/null 2>&1 || (echo "betteralign not found. Install with: go install github.com/dkorunic/betteralign/cmd/betteralign@latest" && exit 1)
	betteralign -apply ./...
