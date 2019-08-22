

A variable of interface type stores a pair: the concrete value assigned to the variable(reflect.Value), and that value's type descriptor(reflect.Type). From  the reflect.Value it's easy to get to the reflect.Type

Another is that both `Type` and `Value` have a `Kind` method that returns a constant indicating what sort of item is stored: `Uint`, `Float64`, `Slice`, and so on. The kind of a reflection object describes the underlying type, not the static type

###  three law of reflection

* Reflection goes frm interface value to reflection object

* Reflection goes from reflection object to interface value
* To modify a reflection object, the value must be settable

### references

[laws-of-reflection](https://blog.golang.org/laws-of-reflection)

[Go Data Structures: Interfaces](https://research.swtch.com/interfaces)