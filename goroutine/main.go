package main


import "fmt"
import "sync"


func worker(name string , wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(name, "started working")
	fmt.Println(name, "finished working")
}

func main(){
	var wg sync.WaitGroup
	workers:=[]string{"worker1", "worker2", "worker3", "worker4", "worker5", "worker6", "worker7", "worker8", "worker9", "worker10"}
	for i:=0;i<len(workers);i++{
		wg.Add(1)
		go worker(workers[i], &wg)
		fmt.Println("started---->>>>>", workers[i], "from main")
	}
	wg.Wait()
	fmt.Println("all workers finished")
}