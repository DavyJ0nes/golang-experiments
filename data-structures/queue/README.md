# Queue

> Simple implementation of a queue, written in a TDD way.

Influenced heaviliy by Jon Calhoun's video: [link](https://youtu.be/fjmsOQFKxho)

## Usage

See Goplayground example [here](https://play.golang.org/p/0UTD5rjend6)

#### Run Tests:

```shell
go test -v
```

#### Create, Add Item and Remove Item from Queue

```go
q := &Queue{
	slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
}

// Get Length of Queue
qLen := q.Len()

// Add Item to Back of Queue
q.Enqueue(10)

// Remove Item from Front of Queue
val, err := q.Dequeue()
if err != nil {
  panic(err)
}

fmt.Println(val) // Will be 1
```
