# GoLang Notes

## Importing Packages
When importing packages, we use the path. But using the package, we use the identifier defined with `package` keyword.

## Initializing Variables
There is no the concept of uninitialized variables or fields in Go. Everything is initialized by default.

```
func main(){
    var foo String;
    fmt.Println(foo) // outputs ""
    var person Person
    fmt.Println(person.FirstName) // outputs ""
    fmt.Println(person.IsMale) // outputs false
}
```