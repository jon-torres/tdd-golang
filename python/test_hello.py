from hello import hello
import pytest


def test_hello_empty_string():
    assert hello("", "") == "Olá, Mundo"


def test_hello_portuguese():
    assert hello("Jon", "") == "Olá, Jon"


def test_hello_french():
    assert hello("Pierre", "French") == "Bonjour, Pierre"


def test_hello_spanish():
    assert hello("Maria", "Spanish") == "Hola, Maria"


def test_hello_unk_language():
    assert hello("Carlos", "Japanese") == "Olá, Carlos"
