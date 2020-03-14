func merge(A []int, m int, B []int, n int)  {
    alen := m + n

    ai := m-1
    bi := n-1
    for i:=alen-1; i>=0; i-- {
        if ai < 0 {
            A[i] = B[bi]
            bi --
            continue
        }
        if bi < 0 {
            A[i] = A[ai]
            ai --
            continue
        }

        if A[ai] > B[bi] {
            A[i] = A[ai]
            ai --
            continue
        }
        A[i] = B[bi]
        bi --
    }
}