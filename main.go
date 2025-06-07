package main

import (
	"fmt"
	"log"
	"strings"

	"go.i3wm.org/i3/v4"
)

func main() {
	for sub := i3.Subscribe(i3.WindowEventType); sub.Next(); {
		event := sub.Event()
		if winEvent, ok := event.(*i3.WindowEvent); ok {
			switch winEvent.Change {
			case "move", "new", "title", "close":
				err := Rename()
				if err != nil {
					log.Printf("failed to rename workspaces: %s", err)
				}
			}
		}
	}
}

func Rename() error {
	wsNum, err := getWorkspaceNumberMap()
	if err != nil {
		return err
	}

	tree, err := i3.GetTree()
	if err != nil {
		return err
	}

	workspaces := findWorkspace(tree.Root)
	commands := make([]string, 0, len(workspaces))
	for _, ws := range workspaces {
		if ws.Workspace.Name == "__i3_scratch" {
			continue
		}
		num := int64(0)
		if n, ok := wsNum[i3.WorkspaceID(ws.Workspace.ID)]; ok {
			num = n
		} else {
			return fmt.Errorf("workspace %s not found in workspace map", ws.Workspace.Name)
		}

		windowNames := make([]string, 0, len(ws.Windows))
		for _, w := range ws.Windows {
			windowNames = append(windowNames, w.WindowProperties.Class)
		}
		newName := fmt.Sprintf("%d: %s", num, strings.Join(windowNames, "|"))

		commands = append(commands, buildRenameCommand(ws.Workspace.Name, newName))
	}

	if len(commands) > 0 {
		fullCmd := strings.Join(commands, "; ")
		_, err := i3.RunCommand(fullCmd)
		if err != nil {
			return fmt.Errorf("failed to run i3 command: %w", err)
		}
	}

	return nil
}

const (
	NodeTypeRoot      = "root"
	NodeTypeOutput    = "output"
	NodeTypeWorkspace = "workspace"
	NodeTypeCon       = "con"
	NodeTypeDockArea  = "dockarea"
)

type workspaceInfo struct {
	Workspace *i3.Node
	Windows   []*i3.Node
}

func findWorkspace(n *i3.Node) []workspaceInfo {
	var result []workspaceInfo
	if n.Type != NodeTypeWorkspace {
		for _, child := range n.Nodes {
			result = append(result, findWorkspace(child)...)
		}
		return result
	}

	// This is a workspace node, collect all windows in it.
	windows := []*i3.Node{}
	for _, child := range n.Nodes {
		windows = append(windows, visitWindow(child)...)
	}
	for _, child := range n.FloatingNodes {
		windows = append(windows, visitWindow(child)...)
	}
	result = append(result, workspaceInfo{
		Workspace: n,
		Windows:   windows,
	})
	return result
}

func visitWindow(n *i3.Node) []*i3.Node {
	if isLeafNode(n) {
		return []*i3.Node{n}
	}

	leaves := []*i3.Node{}
	for _, child := range n.Nodes {
		leaves = append(leaves, visitWindow(child)...)
	}
	for _, child := range n.FloatingNodes {
		leaves = append(leaves, visitWindow(child)...)
	}
	return leaves
}

func isLeafNode(n *i3.Node) bool {
	return len(n.Nodes) == 0 && len(n.FloatingNodes) == 0 && n.Type == NodeTypeCon
}

func getWorkspaceNumberMap() (map[i3.WorkspaceID]int64, error) {
	workspaces, err := i3.GetWorkspaces()
	if err != nil {
		return nil, err
	}
	wsNum := map[i3.WorkspaceID]int64{}
	for _, w := range workspaces {
		wsNum[w.ID] = w.Num
	}
	return wsNum, nil
}

func buildRenameCommand(oldName, newName string) string {
	oldNameEscaped := strings.ReplaceAll(oldName, "\"", "\\\"")
	newNameEscaped := strings.ReplaceAll(newName, "\"", "\\\"")
	return fmt.Sprintf("rename workspace \"%s\" to \"%s\"", oldNameEscaped, newNameEscaped)
}
