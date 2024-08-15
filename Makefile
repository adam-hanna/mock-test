all: mock

mock:
	@mockgen -source=foo/interface.go -destination=interface-mock.go -package=main