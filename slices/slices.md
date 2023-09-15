### Some important points about slices in Go

1. a[start:end]
    - Here the range is half open
    - i.e., [start, end)

2. To append one slice to another, make sure to add "..." to the second one:
    - append(s1, s2...)
    - append(s1, 2, 3) allows you to add individual elements to the slice

3. The function signature of `make` is: 
    - `func make([]T, len, cap) []T`
    - So, if you omit the third argument, you are specifying `len` and not `cap`.
