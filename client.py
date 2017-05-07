import click
import grpc
import addressbook_pb2 as ab
import person_pb2


def new_stub():
    channel = grpc.insecure_channel('localhost:12345')
    return ab.AddressBookStub(channel)


@click.group()
@click.pass_context
def cli(ctx):
    pass


@cli.command()
@click.pass_context
@click.option('--name', help=f'Name', required=True)
@click.option('--age', help=f'Age', required=True, type=int)
def add(ctx, name, age):
    stub = new_stub()
    stub.AddPerson(person_pb2.Person(name=name, age=age))


@cli.command()
@click.pass_context
def list(ctx):
    stub = new_stub()
    for p in stub.ListPerson(ab.NoArgs()):
        print(p)


if __name__ == "__main__":
    cli(obj={})
