all:
	go build -o GSNIFF main.go

clean:
	rm -rf GSNIFF