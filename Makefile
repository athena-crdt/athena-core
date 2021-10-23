.PHONY: codegen
codegen:
	# Protoc Version: libprotoc 3.17.3
	# If your protoc unnecessarily modifies the proto generated files
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/**.proto
