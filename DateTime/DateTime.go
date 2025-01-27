// 日期时间框
// 要美化的话, 就得自绘, 看这个: http://www.xcgui.com/doc-ui/page_draw__month_cal.html
// 我以后有空会翻译过来
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	w := window.New(0, 0, 400, 300, "", 0, xcc.Window_Style_Default)

	dt := widget.NewDateTime(20, 50, 120, 26, w.Handle)
	// 0为日期元素, 1为时间元素.
	dt.SetStyle(0)

	dt2 := widget.NewDateTime(20, 150, 120, 26, w.Handle)
	// 0为日期元素, 1为时间元素.
	dt2.SetStyle(1)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
