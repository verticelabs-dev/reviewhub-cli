package docker_engine

import (
	"fmt"
	"math/rand"
	"net"
	"reviewhub-cli/orchestrator/core"
	"time"
)

type PortFinder struct {
	start int
	end   int
	taken []int
}

func NewPortFinder(start int, end int) *PortFinder {
	return &PortFinder{start: start, end: end}
}

func (pf *PortFinder) GetUnassignedPort() int {
	logger := core.GetLogger()
	logger.Info().Msg(fmt.Sprintf("Finding unassigned port in range %d - %d", pf.start, pf.end))

	rand.Seed(time.Now().UnixNano())
	for {
		port := rand.Intn(pf.end-pf.start+1) + pf.start
		if !pf.isTaken(port) {
			pf.taken = append(pf.taken, port)

			logger.Info().Msg(fmt.Sprintf("Found unassigned port %d", port))

			return port
		}
	}
}

func (pf *PortFinder) isTaken(port int) bool {
	for _, p := range pf.taken {
		if p == port {

			return true
		}
	}

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	conn.Close()

	return false
}
