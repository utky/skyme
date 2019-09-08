# ---------------------------------------------------------
# Variables
# ---------------------------------------------------------

# Version info for binaries
GIT_REVISION := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

# Compiler flags
VPREFIX := github.com/utky/skyme/version
GO_LDFLAGS   := -s -w -X $(VPREFIX).Branch=$(GIT_BRANCH) -X $(VPREFIX).Revision=$(GIT_REVISION)
GO_FLAGS     := -ldflags "-extldflags \"-static\" $(GO_LDFLAGS)" -tags netgo


# ---------------------------------------------------------
# skyme
# ---------------------------------------------------------
.PHONY: skyme
APP_SKYME := cmd/skyme/skyme
skyme: $(APP_SKYME)

$(APP_SKYME): cmd/skyme/main.go pkg/**/*.go
	CGO_ENABLED=0 go build $(GO_FLAGS) -o $@ ./$(@D)

# ---------------------------------------------------------
# Gobal
# ---------------------------------------------------------
.PHONY: clean test
test: skyme
	./run_test.sh ./$(APP_SKYME)

APPS := $(APP_SKYME)
clean:
	rm -f $(APPS)
