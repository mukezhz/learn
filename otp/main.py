import pyotp
import time

# This should be stored differently in the database for the user
shared_secret = 'JBSWY3DPEHPK3PXP'

def verify_totp(submitted_otp):
    """
    Verifies the submitted OTP against the current time-based OTP.

    :param submitted_otp: The OTP submitted by the user.
    :return: True if the OTP is valid, False otherwise.
    """
    totp = pyotp.TOTP(shared_secret)

    return totp.verify(submitted_otp)

totp = pyotp.TOTP(shared_secret, interval=5)
print("Current OTP:", totp.now())

# Simulate user submitting an OTP
user_submitted_otp = input("Enter OTP: ")

# Verify the OTP
is_valid = verify_totp(user_submitted_otp)

if is_valid:
    print("OTP is valid.")
else:
    print("OTP is invalid or expired.")

