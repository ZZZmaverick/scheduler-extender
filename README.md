#Extender 

./controller/priorities.go
打分函数。方法上，先生成0到schedulerapi.MaxPriority之间的随机数（包括两端），然后判断当前时间的秒数是否为偶数，是则在先前随机数的基础上加秒数的个位作为最终分数，否则采用最初的随机数作为最终分数。 

./controller/pridicates.go
过滤函数。先由filter函数调用podFitsOnNode函数，判断是否有某个pod适合某个node，而podFitsOnNode函数通过调用LuckyPredicate函数，进而判断pod是否合适node。此处的LuckyPredicate函数采用了时间加随机数的方法。使用当前时间的秒数加上一个整型随机数0或1，判断加和是否为偶数，是则为lucky，否则不是。 
