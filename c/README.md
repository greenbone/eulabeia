# eulabeia-c-library

This library is intended to be used as an relatively easy to use connection to eulabeia.

It provides the actual library implementation within [libeulabeia](./libeulabeia) as well as [examples](./example).

## Build

To build eulabeia-c-library and it's examples we use cmake. 

To setup a local environment you can run:

```
cd ./libeulabeia
cmake -Bbuild -DCMAKE_EXPORT_PACKAGE_REGISTRY=1 -DCMAKE_EXPORT_COMPILE_COMMANDS=1
```
this will create a build directory, create compile_commands.json (if you're using clangd), register libeulabeia as package within `~/.cmake/packages` so that other projects (like `examples`) can use `find_package(Eulabeia REQUIRED)` without having to install it system wide.

Afterwards you can run:
```
cmake --build build
cmake --build build -- test
```

To build the library and run it's tests.

To build the example you have to build the library before as described above; afterwards you can simply run:

```
cd ./example
cmake -Bbuild
cmake --build build
```

## Run the examples

To run the examples you need to have:
- redis
- director
- sensor
- mqtt broker
running; to simplify that there is a `Makefile` within [eulabeia](https://github.com/greenbone/eulabeia) so that you can run:

```
make build-container # optional
make start-container
```

within the root dir of eulabeia.

This will open a local port `9138` to the mqtt broker and start a director, sensor and redis as required.

After that you can simply tun: `./example/build/start_scan`
