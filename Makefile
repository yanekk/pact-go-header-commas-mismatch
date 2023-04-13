test:
	docker build -t pact-go-authorization-header-error .
	docker run -it pact-go-authorization-header-error
#	docker rm -it pact-go-authorization-header-error
