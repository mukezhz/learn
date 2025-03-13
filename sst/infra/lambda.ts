/// <reference path="../.sst/platform/config.d.ts" />

const greetHandler = new sst.aws.Function("Greeter", {
    handler: "src/functions/greeter",
    runtime: 'go',
})

const gw = new sst.aws.ApiGatewayV2("Api")
gw.route("GET /{subpath+}", greetHandler.arn)
gw.route("POST /{subpath+}", greetHandler.arn)

export {}