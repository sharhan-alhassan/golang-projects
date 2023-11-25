
## Pointers

```sh
- A pointer points to the memory address of a variable

name := "Sharhan"
ptr = &name                                             # ptr outputs the memory address eg; 0xc000014150

- A deferencer gets the original value from the pointer
data = *ptr

- Thus, a function can take a pointer(ptr) and deference it as seen below

func updateName(x *string) {
    *x = "good boy"
}

- Here, the (x *string) is first dereferencing the pointer first to get its 
original value from its memory location


```

## Receiver Functions

```sh
type bill struct {
	name string
	items map[string]float64
	tip float64
}

- These are methods associated with struct objects

- You right the method and then it takes in the struct as a receiver

- The method is called "format()" which takes a receiver struct called "bill" and 
can be accessed via the "b"
func (b bill) format() string {
    ## 
}

- Meaning, we're limitting the "format()" method only to be called from the "bill" objects or instance
```