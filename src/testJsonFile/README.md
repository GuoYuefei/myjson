### 测试结果记录
> 1.1-xx.json,代表第一次改动下的第一次测试 测试json的内容来自xx.json
> 2.3-xx.json,代表第三次改动下的第三次测试 测试json的内容来自xx.json

##### 一体机电脑 i7-4770S CPU@3.10GHz  3.10GHz 内存16G

|  测试id | 测试次数 | ns/op | B/op |	allcs/op	|
| ---- | ----|----|-----|-------|
|目标|5000000|3165|200|6|
|1.1-xx.json|30000|41168|19297|1067|
|1.2-xx.json|50000|38955|19338|1067|
|1.3-xx.json|50000|38682|19338|1067|
|2.1-xx.json|50000|29384|13082|687|
|2.2-xx.json|50000|28982|13082|687|
|2.3-xx.json|50000|30688|13082|687|
|3.1-xx.json|50000|25950|12250|636|
|3.2-xx.json|50000|26712|12250|636|
|3.3-xx.json|50000|26224|12249|636|
|3.4-xx.json|50000|25150|12249|636|
|4.1-xx.json|100000|23627|11975|619|
|4.2-xx.json|50000|29172|11939|619|
|4.3-xx.json|50000|24751|11939|619|
|4.4-xx.json|50000|24779|11939|619|
|4.5-xx.json|50000|24662|11939|619|
|5.1-xx.json|100000|23730|17415|557|
|5.2-xx.json|100000|23230|17415|577|
|5.3-xx.json|50000|22007|17379|577|
|5.4-xx.json|100000|22801|17415|557|
|5.5-xx.json|50000|22842|17379|577|
|6.1-xx.json|200000|9485|7370|228|
|6.2-xx.json|200000|9328|7370|228|
|6.3-xx.json|200000|9865|7370|228|
|6.4-xx.json|100000|10046|7359|228|
|6.5-xx.json|200000|9919|7370|228|
|第二天office.1|500000|3298|200|6|
|第二天office.2|300000|3420|200|6|
|第二天office.3|300000|3765|200|6|
|第二天office.4|500000|3204|200|6|
|第二天office.5|500000|3720|200|6|
|第二头off.Aver| * * * * |3481.4|200|6|
|8.1-xx.json|300000|4279|1814|57|
|8.2-xx.json|500000|3685|1900|57|
|8.3-xx.json|500000|3611|1900|57|
|8.4-xx.json|500000|3900|1900|57|
|8.5-xx.json|500000|3636|1900|57|
|8-xx.json.aver|* * * *|3822.2|约1900|57|
|9.1-xx.json|500000|3682|1719|52|
|9.2-xx.json|500000|3522|1719|52|
|9.3-xx.json|500000|3655|1719|52|
|9.4-xx.json|500000|3483|1719|52|
|9.5-xx.json|500000|3591|1719|52|
|9-xx.json.aver|500000|3586.6|1719|52|
||||||
||||||
||||||
||||||
||||||

##### 个人笔记本 A 锐龙R7 2700u 内存8G双通道

|  测试id | 测试次数 | ns/op | B/op |	allcs/op	|
| ---- | ----|----|-----|-------|
|目标|200000|9611|200|6|
|7.1-xx.json|100000|19067|4623|228|
|7.2-xx.json|100000|12878|4623|228|
|7.3-xx.json|100000|13152|4623|228|
|7.4-xx.json|100000|13008|4623|228|



##### 一体机 配置同上

| 测试id       | 测试次数 | ns/op | B/op | allcs/op |
| ------------ | -------- | ----- | ---- | -------- |
| 目标         | 200000   | 7758  | 304  | 20       |
| 9.1-big.json | 200000   | 9152  | 3825 | 117      |
| 9.2-big.json | 200000   | 9259  | 3825 | 117      |
| 9.3-big.json | 200000   | 9215  | 3825 | 117      |
| 9.4-big.json | 200000   | 9726  | 3825 | 117      |
| 9.5-big.json | 200000   | 9470  | 3825 | 117      |
|              |          |       |      |          |
|              |          |       |      |          |
|              |          |       |      |          |







```json
"key12":[
  {
    "key13":134567
  },
  {
    "key14":14567
  }
]
```

