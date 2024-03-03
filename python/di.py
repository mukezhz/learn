class ApiClient:
    def __init__(self, api_key: str, timeout: int) -> None:
        self.api_key = api_key  # <-- dependency is injected
        self.timeout = timeout  # <-- dependency is injected


class Service:
    def __init__(self, api_client: ApiClient) -> None:
        self.api_client = api_client  # <-- dependency is injected


def main(service: Service) -> None:  # <-- dependency is injected
    print("API KEY AFTER DI:", service.api_client.api_key)


if __name__ == "__main__":
    main(
        service=Service(
            api_client=ApiClient(
                api_key="MY_API_KEY",
                timeout=10,
            ),
        ),
    )
