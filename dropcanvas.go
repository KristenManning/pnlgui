// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pnlgui

import (
	// "image/color"
	// "log"
	// "reflect"

	"fmt"

	"github.com/goki/gi/gi"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/dnd"
	"github.com/goki/gi/oswin/mimedata"
	"github.com/goki/gi/svg"

	// "github.com/goki/gi/giv"
	// "github.com/goki/gi/units"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

/////////////////////////////////////////////////////////////////////////////
//  DropCanvas

// DropCanvas shows a color, using sliders to set values,
type DropCanvas struct {
	svg.Editor
}

var KiT_DropCanvas = kit.Types.AddType(&DropCanvas{}, DropCanvasProps)

var DropCanvasProps = ki.Props{
	"background-color": "blue",
	"color":            &gi.Prefs.Colors.Font,
	"min-width":        "20em",
	"min-height":       "20em",
	"max-width":        -1,
	"max-height":       -1,
}

// Drop pops up a menu to determine what specifically to do with dropped items
// satisfies gi.DragNDropper interface and can be overridden by subtypes
func (dc *DropCanvas) Drop(md mimedata.Mimes, mod dnd.DropMods) {
	fmt.Printf("in dropcanvas drop\n")
	comp := AddNewComponentView(dc, "new")
	updt := dc.UpdateStart()
	comp.Pos.Y = float32(len(dc.Kids) * 100)
	// comp.Pos.Y = 50
	comp.Size.X = 50
	comp.Size.Y = 50
	dc.DragNDropFinalize(mod)
	dc.UpdateEnd(updt)
}

// DragNDropTarget handles a drag-n-drop onto this node
func (dc *DropCanvas) DragNDropTarget(de *dnd.Event) {
	de.Target = dc.This()
	if de.Mod == dnd.DropLink {
		de.Mod = dnd.DropCopy // link not supported -- revert to copy
	}
	de.SetProcessed()
	if dpr, ok := dc.This().(gi.DragNDropper); ok {
		dpr.Drop(de.Data, de.Mod)
	} else {
		dc.Drop(de.Data, de.Mod)
	}
}

// DragNDropFinalize is called to finalize actions on the Source node prior to
// performing target actions -- mod must indicate actual action taken by the
// target, including ignore
func (dc *DropCanvas) DragNDropFinalize(mod dnd.DropMods) {
	if dc.Viewport == nil {
		return
	}
	dc.Viewport.Win.FinalizeDragNDrop(mod)
}

// DropCanvasEvents: this is key for processing DND events
func (dc *DropCanvas) DropCanvasEvents() {
	dc.ConnectEvent(oswin.DNDEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		if recv == nil {
			return
		}
		de := d.(*dnd.Event)
		dcc := recv.Embed(KiT_DropCanvas).(*DropCanvas)
		switch de.Action {
		case dnd.Start:
		case dnd.DropOnTarget:
			dcc.DragNDropTarget(de)
		case dnd.DropFmSource:
		}
	})
	dc.ConnectEvent(oswin.DNDFocusEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		if recv == nil {
			return
		}
		de := d.(*dnd.FocusEvent)
		dcc := recv.Embed(KiT_DropCanvas).(*DropCanvas)
		switch de.Action {
		case dnd.Enter:
			dcc.Viewport.Win.DNDSetCursor(de.Mod)
		case dnd.Exit:
			dcc.Viewport.Win.DNDNotCursor()
		case dnd.Hover:
		}
	})
}

func (dc *DropCanvas) ConnectEvents2D() {
	dc.EditorEvents()
	dc.DropCanvasEvents()
}
