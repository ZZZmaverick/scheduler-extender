package controller

import (
	"log"
	"math/rand"

	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)


const (
	// lucky priority gives a random [0, schedulerapi.MaxPriority] score
	// currently schedulerapi.MaxPriority is 10
	luckyPrioMsg = "pod %v/%v is lucky to get score %v\n"
)

func prioritize(args schedulerapi.ExtenderArgs) *schedulerapi.HostPriorityList {
	pod := args.Pod
	nodes := args.Nodes.Items
	hostPriorityList := make(schedulerapi.HostPriorityList, len(nodes))
	for i, node := range nodes {
		score := rand.Intn(schedulerapi.MaxPriority + 1)
		current_t := time.Now().Format("15:04:05")
                                t_score, _ := strconv.Atoi(current_t[6:])
                                lucky := t_score%2 == 0
                                if lucky {
                                    score = score + t_score%10
                                }
		log.Printf(luckyPrioMsg, pod.Name, pod.Namespace, score)
		hostPriorityList[i] = schedulerapi.HostPriority{
			Host:  node.Name,
			Score: score,
		}
	}
	return &hostPriorityList
}