package handler

import (
	"fmt"
	"team01/internal/server/config"
	"time"
)

type Cluster struct {
	NodesList []Node `json:"nodes_list"`
}

func NewCluster() *Cluster {
	return &Cluster{
		NodesList: make([]Node, 0),
	}
}

func (c *Cluster) Monitor(cfg *config.ServerConfig) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if cfg.CurrentPort == cfg.LeaderPort {
			c.CheckFollowers()
		} else {
			DoHeartbeat(cfg, c)
		}
	}
}

func (c *Cluster) PrintNodesList() {
	fmt.Println("Nodes List:")
	for _, node := range c.NodesList {
		fmt.Printf("%d - %s\n", node.Port, node.Role)
	}
}

func (c *Cluster) AppendNode(node Node) {
	c.NodesList = append(c.NodesList, node)

	if node.Role != "Leader" {
		fmt.Printf("the node on port %d has been registered\n", node.Port)
	}

	c.PrintNodesList()
}

func (c *Cluster) CheckFollowers() {
	for i, node := range c.NodesList {
		if node.Role == "Leader" {
			continue
		}

		if time.Since(node.LastActive) > HeartbeatTimeout {
			fmt.Printf("Node on port %d is dead\n", node.Port)
			c.NodesList = append(c.NodesList[:i], c.NodesList[i+1:]...)
			c.PrintNodesList()
		}
	}
}

func (c *Cluster) isExistNode(node Node) bool {
	if len(c.NodesList) == 0 {
		return true
	}

	for _, n := range c.NodesList {
		if n.Port == node.Port {
			return false
		}
	}

	return true
}

func (c *Cluster) updateLastActive(node Node) {
	for i, n := range c.NodesList {
		if n.Port == node.Port {
			c.NodesList[i].LastActive = node.LastActive
			break
		}
	}
}

func (c *Cluster) IsEqual(other *Cluster) bool {
	if len(c.NodesList) != len(other.NodesList) {
		return false
	}

	for i, node := range c.NodesList {
		if node.Port != other.NodesList[i].Port || node.Role != other.NodesList[i].Role {
			return false
		}
	}

	return true
}

func (c *Cluster) isEmpty() bool {
	if len(c.NodesList) == 0 {
		return true
	}

	return false
}

func (c *Cluster) makeNewLeader() int {
	fmt.Println("Leader is dead")

	c.NodesList = c.NodesList[1:]

	c.NodesList[0].Role = "Leader"
	leaderPort := c.NodesList[0].Port

	fmt.Printf("New leader is node on port %d\n", leaderPort)

	return leaderPort
}

func (c *Cluster) update(other *Cluster) {
	c.NodesList = other.NodesList
}
