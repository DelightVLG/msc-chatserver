.PHONY: generate lint grpcui

generate:
	buf generate

lint:
	buf lint

grpcui:
	grpcui -plaintext localhost:8088
