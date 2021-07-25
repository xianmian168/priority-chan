我们知道golang的select是随机执行case语句的，那现在有一种场景：有一个方法持续不间断地从high和low中分别接收任务high和任务low，
如何确保当high和low同时达到就绪状态时，优先执行任务high，在没有任务high的时候再去执行任务low呢？请修改下面的代码以满足需求
```go
func handle(low, high <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case highJob := <-high:
			log.Println("highJob: ", highJob)
		case lowJob := <-low:
			log.Println("lowJob: ", lowJob)
		}
	}
}
```
