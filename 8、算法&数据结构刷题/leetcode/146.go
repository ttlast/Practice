ype Node struct {
    Key int
    Val int
    Pre *Node
    Next *Node
}

type LRUCache struct {
    Mm map[int]*Node
    Size int
    Cap int
    Headp *Node
    Tailp *Node
}


func Constructor(capacity int) LRUCache {
    var lruCache LRUCache
    lruCache.Cap = capacity
    lruCache.Mm = make(map[int]*Node)
    return lruCache
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.Mm[key]; ok {
        //fmt.Println(node, this.Headp, this.Tailp)
        this.updateNodeToHead(node)
        return node.Val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.Mm[key]; ok {
        node.Val = value
        this.updateNodeToHead(node)
    } else {
        this.Size ++
        newNode := &Node{key, value, nil, this.Headp}
        if this.Headp == nil {
            this.Tailp = newNode
        } else {
            this.Headp.Pre = newNode
        }
        this.Mm[key] = newNode
        this.Headp = newNode

        if this.Size > this.Cap {
            this.Size --
            delete(this.Mm, this.Tailp.Key)
            this.Tailp = this.Tailp.Pre
            this.Tailp.Next = nil
        }
    }
}

func (this *LRUCache) updateNodeToHead(node *Node) {
    if node == this.Headp {
            return
        }

        if node == this.Tailp {
            this.Tailp = this.Tailp.Pre
        }
        if node.Pre != nil {
            node.Pre.Next = node.Next
        }
        if node.Next != nil {
            node.Next.Pre = node.Pre
        }

        node.Pre = nil
        node.Next = this.Headp
        this.Headp.Pre = node
        this.Headp = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */