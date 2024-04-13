package list

type node struct {
    val  any
    next *node
}

type List struct {
    head *node
    tail *node
    size int
}
