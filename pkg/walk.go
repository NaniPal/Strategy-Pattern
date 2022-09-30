package pkg

import (
	"fmt"
)

type Walk struct {
}

func (r *Walk) Route(startPoint int, endPoint int) {
	avgSpeed := 8
	total := endPoint - startPoint
	totalTime := total * 50
	fmt.Printf("Walk Road A:[%d] to B:[%d] Average speed :[%d] Total :[%d] Total time :[%d]\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
