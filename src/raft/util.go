package raft

import "log"

// Debugging
const Debug = 0

func DPrintf(format string, a ...interface{}) {
	if Debug > 0 {
		log.Printf(format, a...)
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
