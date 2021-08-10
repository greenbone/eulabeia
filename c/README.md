# Build

```
cd ./libeulabeia
cmake -Bbuild -DCMAKE_EXPORT_PACKAGE_REGISTRY=1 -DCMAKE_EXPORT_COMPILE_COMMANDS=1
cmake --build build
cmake --build build -- test
```

```
cd ./example
cmake -Bbuild -DCMAKE_EXPORT_PACKAGE_REGISTRY=1
cmake --build build
```
