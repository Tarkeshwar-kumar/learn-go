package main

func main(){
	var arr [10]int
	for i := 0; i<10; i+=1 {
		arr[i] = i
	}

	sum := 0

	for i := 0; i<10; i+=1 {
		sum += arr[i]
	}
	print(sum)
}