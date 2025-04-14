package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ManagerLayout(
	app *tview.Application,
	// pages *tview.Pages,
	// pageName string,
	// switchFun func(pages *tview.Pages, pageName string),
) *tview.Flex {
	bucketLayout := BucketNameLayout()
	btnLayout := ButtonsLayout(app)
	consoleLayout := ConsoleLayout()
	browserLayout := BrowserLayout(consoleLayout)

	// 整體佈局
	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(bucketLayout, 1, 0, false).
		AddItem(browserLayout, 0, 5, false).
		AddItem(btnLayout, 1, 0, true).
		AddItem(consoleLayout, 5, 0, false)

	layout.SetBorder(true)

	return layout
}

func BrowserLayout(console *tview.TextView) *tview.Flex {
	prefixTreeView := PrefixTreeLayout()
	fileListView := FileListLayout()

	prefixTreeView.SetSelectedFunc(func(node *tview.TreeNode) {
		ref := node.GetReference()
		if ref != nil {
			prefix := ref.(string)
			fileListView.Clear().
				AddItem(fmt.Sprintf("%s_File_1", prefix), "", 0, nil).
				AddItem(fmt.Sprintf("%s_File_2", prefix), "", 0, nil).
				AddItem(fmt.Sprintf("%s_File_3", prefix), "", 0, nil)
			console.SetText(fmt.Sprintf("Selected Prefix: %s", prefix))
		}
	})

	flex := tview.NewFlex().
		AddItem(prefixTreeView, 0, 1, true).
		AddItem(fileListView, 0, 2, false)

	return flex
}

func PrefixTreeLayout() *tview.TreeView {
	rootNode := tview.NewTreeNode("Prefixes").SetColor(tcell.ColorGreen)
	for i := 1; i <= 5; i++ {
		prefix := fmt.Sprintf("Prefix_%d", i)
		child := tview.NewTreeNode(prefix).SetReference(prefix)
		rootNode.AddChild(child)
	}

	tree := tview.NewTreeView().SetRoot(rootNode).SetCurrentNode(rootNode)
	tree.SetBorder(true).SetTitle("Prefixes")

	return tree
}

func FileListLayout() *tview.List {
	list := tview.NewList().
		AddItem("File_1", "", 0, nil).
		AddItem("File_2", "", 0, nil).
		AddItem("File_3", "", 0, nil)

	list.SetBorder(true).SetTitle("Files")

	return list
}

func BucketNameLayout() *tview.TextView {
	return tview.NewTextView().SetText("Bucket: ")
}

func ConsoleLayout() *tview.TextView {
	console := tview.NewTextView().
		SetText("console...").
		SetDynamicColors(true).
		SetScrollable(true)

	console.SetTitle(" Console ").SetBorder(true)

	return console
}

func ButtonsLayout(app *tview.Application) *tview.Flex {
	// 下方按鈕區
	inputField := tview.NewInputField().
		SetLabel("Upload Path: ").
		SetFieldWidth(40)
	uploadBtn := tview.NewButton("Upload")
	downloadBtn := tview.NewButton("Download")
	deleteBtn := tview.NewButton("Delete")
	exitBtn := tview.NewButton("Exit").SetSelectedFunc(func() {
		app.Stop()
	})

	layout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(inputField, 55, 0, true).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(uploadBtn, 10, 0, false).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(downloadBtn, 10, 0, false).
		AddItem(tview.NewBox(), 2, 0, false).
		AddItem(deleteBtn, 10, 0, false).
		AddItem(tview.NewBox(), 2, 0, false).
		AddItem(exitBtn, 10, 0, false)

	return layout
}
