test:
	go test ./...
test-view:
	go test ./... -v
push: test
	git push origin master
migrate:
	go run . -migrate
drop:
	go run . -drop
fresh:
	go run . -fresh