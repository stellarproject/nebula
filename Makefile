PACKAGES=$(shell go list ./... | grep -v /vendor/)

all:
	protobuild --quiet ${PACKAGES}
	go install . ./...
