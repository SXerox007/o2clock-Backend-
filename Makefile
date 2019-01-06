.PHONY: all brain expose clean

brain:
	go run server/brain/brain.go

expose:
	go run server/expose/rest-app.go
