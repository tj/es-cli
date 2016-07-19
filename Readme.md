
# es(1)

Elasticsearch stats CLI, at least the start of one.

## Installation

```
$ go get github.com/tj/es
```

## Setup

First assign __ES_ADDR__:

```
export ES_ADDR=http://localhost:5555
```

## Node Stats

```
$ es nodes

Ronan the Accuser

     Documents: 44,577,077
        Memory: 1.2 GB free (27%) – 3.0 GB used (72%)
      JVM Heap: 2.1 GB committed – 1.7 GB used
        JVM GC: 51 MB young – 4.3 MB survivor – 1.6 GB old
    Field Data: 1.1 GB (0 evictions)
  Filter Cache: 195 MB (58677 evictions)
   Query Cache: 0 B (0 evictions)
      ID Cache: 0 B
      Segments: 51 MB (873)

Plug

     Documents: 45,870,106
        Memory: 1.1 GB free (26%) – 3.0 GB used (73%)
      JVM Heap: 2.1 GB committed – 1.8 GB used
        JVM GC: 114 MB young – 4.0 MB survivor – 1.7 GB old
    Field Data: 1.2 GB (0 evictions)
  Filter Cache: 196 MB (59867 evictions)
   Query Cache: 0 B (0 evictions)
      ID Cache: 0 B
      Segments: 51 MB (843)

Dementia

     Documents: 44,171,226
        Memory: 1.1 GB free (27%) – 3.0 GB used (72%)
      JVM Heap: 2.1 GB committed – 1.8 GB used
        JVM GC: 122 MB young – 1.7 MB survivor – 1.6 GB old
    Field Data: 1.1 GB (0 evictions)
  Filter Cache: 193 MB (56776 evictions)
   Query Cache: 0 B (0 evictions)
      ID Cache: 0 B
      Segments: 50 MB (860)
```

## Badges

![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)
[![](http://apex.sh/images/badge.svg)](https://apex.sh/ping/)

---

> [tjholowaychuk.com](http://tjholowaychuk.com) &nbsp;&middot;&nbsp;
> GitHub [@tj](https://github.com/tj) &nbsp;&middot;&nbsp;
> Twitter [@tjholowaychuk](https://twitter.com/tjholowaychuk)
