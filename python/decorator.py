


def decorator(from_):
    def wrap(fn):
        def inner_fn(*args):
            print(f"Argument: {args} from {from_}")
            fn(args[0])
        return inner_fn
    return wrap


@decorator("from B")
def hello(name):
    print(f"Hello {name}")


if __name__ == '__main__':
    hello("ABC")
