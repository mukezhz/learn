from typing import Protocol
from dependency_injector import containers, providers
from dependency_injector.wiring import inject, Provide


class PaymentProcessor(Protocol):
    def process_payment(self, order_id: str, amount: float) -> bool:
        """Process a payment"""
        return False


class PaypalPaymentService(PaymentProcessor):
    def process_payment(self, order_id: str, amount: float) -> bool:
        print(
            f"Processing payment through PayPal for order {order_id} with amount {amount}"
        )
        return True


class StripePaymentService(PaymentProcessor):
    def process_payment(self, order_id: str, amount: float) -> bool:
        print(
            f"Processing payment through Stripe for order {order_id} with amount {amount}"
        )
        return True


class PaymentContainer(containers.DeclarativeContainer):
    print("creating paypal service bean")
    paypal_payment_service = providers.Factory(PaypalPaymentService)
    print("creating stripe service bean")
    stripe_payment_service = providers.Factory(StripePaymentService)


class CheckoutService:
    @inject
    def __init__(
        self,
        payment_processor: PaymentProcessor = Provide[
            PaymentContainer.paypal_payment_service
            # injecting the payment service bean
        ],
    ):
        self.payment_processor = payment_processor

    def process_checkout(self, order_id: str, amount: float):
        return self.payment_processor.process_payment(order_id, amount)


def main():
    checkout_service = CheckoutService()
    checkout_service.process_checkout("order123", 99.99)


if __name__ == "__main__":
    print("creating IOC container...")
    container = PaymentContainer()
    print("wiring the beans...")
    container.wire(modules=[__name__])

    print()
    main()
