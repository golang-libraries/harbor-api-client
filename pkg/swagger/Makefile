BIN_DIR=/tmp/.local/bin
SWAGGER_CLI=${BIN_DIR}/swagger
SWAGGER_VERSION=0.29.0

HARBOR_VERSION=v2.5.0
HARBOR_API_VERSION=v2.0

install_swagger:
ifeq (,$(wildcard ${SWAGGER_CLI}))
		mkdir -p ${BIN_DIR} && \
 		curl -L -o ${SWAGGER_CLI}  https://github.com/go-swagger/go-swagger/releases/download/v${SWAGGER_VERSION}/swagger_darwin_arm64
		chmod +x ${SWAGGER_CLI}
		${SWAGGER_CLI} version
else
		@printf "swagger cli, version: %s is already installed in: %s\n"  ${SWAGGER_VERSION}  ${SWAGGER_CLI}
endif

client: install_swagger
	rm -rf ./$(HARBOR_API_VERSION)
	mkdir -p $(HARBOR_API_VERSION)
	curl -o ./$(HARBOR_API_VERSION)/swagger.yaml https://raw.githubusercontent.com/goharbor/harbor/${HARBOR_VERSION}/api/v2.0/swagger.yaml
	echo $(HARBOR_VERSION) > ./$(HARBOR_API_VERSION)/HARBOR_VERSION.txt
	${SWAGGER_CLI}  generate client -f ./$(HARBOR_API_VERSION)/swagger.yaml -A harbor  --target=./$(HARBOR_API_VERSION)/