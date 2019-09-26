# GO-CACHE

key-value 缓存

local cache + redis + db

## 解决问题
1. LRU缓存淘汰（or local Memcached） 
2. 缓存击穿
    - Redis 加互斥锁，保护下游DB
3. 缓存穿透 
    - BloomFilter 检测不存在的key
4. 缓存雪崩 
    1. local cache + redis + db
    2. redis + redis + db