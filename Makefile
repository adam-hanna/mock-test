all: mock

mock:
	@mockgen -source=mock/interface.go -destination=mock/interface-mock.go -package=mock