// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pnlgui

import (
	// "image/color"
	// "log"
	// "reflect"

	// "github.com/goki/gi/gi"
	"github.com/goki/gi/svg"
	"github.com/goki/gi/oswin/mimedata"
	"github.com/goki/gi/oswin/dnd"
	// "github.com/goki/gi/giv"
	// "github.com/goki/gi/units"
	// "github.com/goki/ki/ki"
	// "github.com/goki/ki/kit"
)

/////////////////////////////////////////////////////////////////////////////
//  DropCanvas

// DropCanvas shows a color, using sliders to set values,
type DropCanvas struct {
	svg.Editor
}

// Drop pops up a menu to determine what specifically to do with dropped items
// satisfies gi.DragNDropper interface and can be overridden by subtypes
func (dc *DropCanvas) Drop(md mimedata.Mimes, mod dnd.DropMods) {
	comp := AddNewComponentView(dc, "new")
	comp.Pos.Y = len(dc.Kids)*100 
	comp.Size.X = 50
	comp.Size.Y = 50 
}



