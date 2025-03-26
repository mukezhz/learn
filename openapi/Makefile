gen:
	npx @redocly/cli bundle idl/openapi/users/openapi.yaml --output idl/openapi/users.yaml && mkdir -p server_gen/user_gen && oapi-codegen -package user_gen -generate types,gin-server -o server_gen/user_gen/users.gen.go idl/openapi/users.yaml
	npx @redocly/cli bundle idl/openapi/organizations/openapi.yaml --output idl/openapi/organizations.yaml && mkdir -p server_gen/organization_gen && oapi-codegen -package organization_gen -generate types,gin-server -o server_gen/organization_gen/organizations.gen.go idl/openapi/organizations.yaml
	# npx @redocly/cli bundle idl/openapi/openapi.yaml --output server_gen/openapi.yaml && oapi-codegen -package server_gen -generate types,gin-server -o server_gen/api.gen.go server_gen/openapi.yaml
	# oapi-codegen -config oapi.yaml idl/openapi/openapi.yaml

.PHONY: gen