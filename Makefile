
.PHONY: docker
docker: tmpl.linux.amd64
	docker build -t pastjean/tmpl .

tmpl.linux.amd64:
	GOOS=linux go build -o tmpl.linux.amd64 .

tmpl.darwin.amd64:
	GOOS=darwin go build -o tmpl.darwin.amd64 .