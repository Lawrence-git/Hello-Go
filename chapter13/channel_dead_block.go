package main

//dead  lock
func channelDeadBlock() {
	strChan := make(chan string)
	strChan <- "Main"
}

func chanBlock(strChan chan string) {
	strChan <- "Func chanBlock"
}
