def hello(name, language=""):
    if not name:
        name = "Mundo"

    return greeting_prefix(language) + name


def greeting_prefix(language):
    prefixes = {
        "French": "Bonjour, ",
        "Spanish": "Hola, ",
    }

    return prefixes.get(language, "Olá, ")


if __name__ == "__main__":
    print(hello("Carlos", "Japanese"))
