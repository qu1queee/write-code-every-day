.PHONY: test test-archive test-phase0 test-phase1 fmt kata

GO_DIR := go

test:
	cd $(GO_DIR) && go test ./...

test-archive:
	cd $(GO_DIR) && go test ./problems/archive/...

test-phase0:
	cd $(GO_DIR) && go test ./problems/phase0/...

test-phase1:
	cd $(GO_DIR) && go test ./problems/phase1/...

fmt:
	cd $(GO_DIR) && gofmt -w .

kata:
	@test -n "$(PROBLEM)" || (echo "Usage: make kata PROBLEM=problems/phase0/lc0225_stack_using_queues" && exit 1)
	@./scripts/kata.sh "$(PROBLEM)"
