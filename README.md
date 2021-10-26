# golang-basics
A tutorial to explore golang concepts

### Some details on golang
- Golang is a compiled language!
- It can run files directly, without the need of an intermediate like JVM.
- Can produce different executables for different OS.

### Advantages
- Already in lot of prod clusetrs!
- Can be used to make System Apps to Web Apps

## Types
- Type is case-sensitive (almost always)
- Var type should be knoiwn in advance
- Everything has a type

### Basic Types in Golang
- String
- Bool
- Integer (uint8, uint64, int8, int64, uintptr)
- Float (float32, float64)
- Complex

### Advanced types in Golang
- Array
- Slices
- Maps
- Structs
- Pointers

### And many more...
- Functions
- Channels
- etc.

### To see output
- ```go run main.go```

### To build the code
- Simply run ```go build```
- To build for some non-host OS, ```GOOS="os_name" go build```

### Memory Management
- No need to do it manually; golang does it for us!
- Garbage Collection (GC) happens automatically
- ```new()```: Allocate memory, not init; zeroed storage; get mem address
- ```make()```: Allocate memory and init; non-zeroed storage; get mem address