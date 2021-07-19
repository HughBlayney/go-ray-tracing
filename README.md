# go-ray-tracing
A simple ray tracer, written while I learn Go

## Stuff I've Learnt

A collection of stuff I've picked up along the way. Will eventually be organsed in order from dumb -> interesting.

- Go doesn't have the functionality to run code at struct construction / initialisation (init in Python). It seems that the best way to get around this is by using a dedicated constructor function (e.g. MakeRay). Optionally, the class itself could be not exported, to enforce the usage of this function, but that then creates issues if you want to refer to that struct in e.g. func definitions.

- The Go `:= range` syntax is amazing, with the functionality of Python's `enumerate` in a more concise way.