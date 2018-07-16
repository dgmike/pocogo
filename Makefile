run:
	ps aux | grep go-build | grep -v grep | awk '{print $$2}' | xargs -rn1 kill; \
	go run main.go &
