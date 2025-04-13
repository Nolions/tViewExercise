package exercise

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

func NewTreeView(path string) *tview.TreeView {
	//建立 TreeView 根節點
	root := tview.NewTreeNode(path).
		SetColor(tcell.ColorRed)

	// 用 root 當作 TreeView 的根
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// 替root加上子節點
	addSubNodes := setSubNodes()
	addSubNodes(root, path)

	return tree
}

func setSubNodes() func(target *tview.TreeNode, path string) {
	return func(target *tview.TreeNode, path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}
}
