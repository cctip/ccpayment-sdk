from setuptools import setup, find_packages

setup(
    name="ccpayment-sdk",
    version="1.0.0",
    description="Official Python SDK for CCPayment API v2",
    author="CCPayment",
    author_email="support@ccpayment.com",
    packages=find_packages(),
    install_requires=[
        "requests>=2.28.0",
    ],
    python_requires=">=3.8",
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
    ],
)
