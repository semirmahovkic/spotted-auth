set type=%~1

go test -race ./... -tags=%type% -v