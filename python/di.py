from dependency_injector import containers, providers
from dependency_injector.wiring import Provide, inject, required


class ApiClient:
    def __init__(self, api_key: str, timeout: int) -> None:
        self.api_key = api_key  # <-- dependency is injected
        self.timeout = timeout  # <-- dependency is injected


class Service:
    def __init__(self, api_client: ApiClient) -> None:
        self.api_client = api_client  # <-- dependency is injected


# def main(service: Service) -> None:  # <-- dependency is injected
#     print("API KEY AFTER DI:", service.api_client.api_key)


# if __name__ == "__main__":
#     main(
#         service=Service(
#             api_client=ApiClient(
#                 api_key="MY_API_KEY",
#                 timeout=10,
#             ),
#         ),
#     )


class Container(containers.DeclarativeContainer):
    config = providers.Configuration()

    api_client = providers.Singleton(
        ApiClient,
        api_key=config.api_key,
        timeout=config.timeout,
    )

    service = providers.Factory(
        Service,
        api_client=api_client,
    )


@inject
def main(service: Service = Provide[Container.service]) -> None:
    print("API KEY AFTER::", service.api_client.api_key)


if __name__ == "__main__":
    container = Container()
    container.config.api_key.from_env(name="API_KEY", default="TEST", required=True)
    # container.config.api_key.from_env("API_KEY", required=True)
    container.config.timeout.from_env(name="TIMEOUT", default="TEST", required=True)
    container.wire(modules=[__name__])

    main()  # <-- dependency is injected automatically
