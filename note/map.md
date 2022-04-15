# map
Golang的map使用哈希表作为底层实现，一个哈希表里可以有多个哈希表节点，也即bucket，而每个bucket就保存了map中的一个或一组键值对。

# 数据结构
```go
type hmap struct {
    count     int // 当前保存的元素个数
    ...
    B         uint8
    ...
    buckets    unsafe.Pointer // bucket数组指针，数组的大小为2^B
    ...
}
```
bucket
```
type bmap struct {
    tophash [8]uint8 //存储哈希值的高8位
    data    byte[1]  //key value数据:key/key/key/.../value/value/value...
    overflow *bmap   //溢出bucket的地址
}
```

#  哈希冲突
当有两个或以上数量的键被哈希到了同一个bucket时，我们称这些键发生了冲突。Go使用链地址法来解决键冲突。 由于每个bucket可以存放8个键值对，所以同一个bucket存放超过8个键值对时就会再创建一个键值对，用类似链表的方式将bucket连接起来。

#  负载因子
`loadFactor := count / (2^B)`
count 就是 map 的元素个数，2^B 表示 bucket 数量。

哈希表需要将负载因子控制在合适的大小，超过其阀值需要进行rehash，也即键值对重新组织：

哈希因子过小，说明空间利用率低
哈希因子过大，说明冲突严重，存取效率低

# 渐进式扩容
条件
* 负载因子>6.5
* 溢出桶数量超过 2的B次方（B<=15）
> overflow 的 bucket 数量过多：当 B 小于 15，也就是 bucket 总数 2^B 小于 2^15 时，如果 overflow 的 bucket 数量超过 2^B；当 B >= 15，也就是 bucket 总数 2^B 大于等于 2^15，如果 overflow 的 bucket 数量超过 2^15。

### 增量扩容
扩容容量为2倍，然后在访问map的时候搬迁，每次搬迁2个键值对。搬迁过程中，新的写入到新buckets中。

### 等量扩容
极端场景下，比如不断地增删，而键值对正好集中在一小部分的bucket，这样会造成overflow的bucket数量增多，但负载因子又不高，从而无法执行增量搬迁的情况，此时进行一次等量扩容，即buckets数量不变，经过重新组织后overflow的bucket数量会减少，即节省了空间又会提高访问效率。
> 主要是减少溢出桶数量

# 查找过程
1. 对 key 进行 hash
2. 取后 B 位，和 buckets 数量进行取模，确定落到哪个 bucket 上
3. 取前 8 位，在 tophash 上查找，找到则在该 bucket 中的key比较，有则用偏移量计算地址，取出
4. 否则继续到溢出桶找
5. 如果处于搬迁过程，先去 odl buckets 找
6. 找不到返回对应类型的零值

# 插入过程
1. 对 key 进行 hash
2. 取后 B 位，和 buckets 数量进行取模，确定落到哪个 bucket 
3. 在该 bucket 上寻找，或在溢出桶寻找
4. 存在则更新，否则插入
5. 如果 bucket 满了，创建溢出桶进行存储