.PHONY: all brain expose swagger clean

brain:
	go run server/brain/brain.go

expose:
	go run server/expose/rest-app.go