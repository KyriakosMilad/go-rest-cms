test:
	go test ./...
test-view:
	go test ./... -v
push: test
	git push origin master
