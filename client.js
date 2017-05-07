const grpc = require("grpc");
const argv = require('minimist')(process.argv.slice(2));

const addressbook = grpc.load('./proto/addressbook.proto').addressbook

const secrets = grpc.credentials.createInsecure()
const client = new addressbook.AddressBook('localhost:12345', secrets);

function list() {
  const call = client.listPerson(new addressbook.NoArgs());
  call.on('data', function(res) { console.log(res); });
  //call.on('end', function() { console.log("end"); });
  //call.on('status', function(status) { console.log("status:", status); });
}

function add(argv) {
  const p = {
    name: argv[0],
    age: argv[1],
  };
  client.addPerson(p, function() {});
}

switch (argv._[0]) {
  case "list": list(); break;
  case "add": add(argv._.slice(1)); break;
}
