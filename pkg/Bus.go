package pkg

import (
	"fmt"
)

type Bus struct {
}

func (r *Bus) Route(startPoint int, endPoint int) {
	avgSpeed := 15
	total := endPoint - startPoint
	totalTime := total * 50
	fmt.Printf("Bus Road A:[%d] to B:[%d] Average speed :[%d] Total :[%d] Total time :[%d]\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
