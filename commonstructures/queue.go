package commonstructures

//probably only structure ill use


func enqueue(queue []string,file string) []string{
	x:=append(queue,file)
	return x
}

func dequeue(queue[] string) ([]string){
	return queue[1:]
}