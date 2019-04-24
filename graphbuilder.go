// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pnlgui

import (
	// "image/color"
	// "log"
	// "reflect"

	"github.com/goki/gi/gi"
	"github.com/goki/gi/giv"

	// "github.com/goki/gi/units"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

/////////////////////////////////////////////////////////////////////////////
//  GraphBuilder

// GraphBuilder shows a color, using sliders to set values,
type GraphBuilder struct {
	gi.Frame
	// Components   giv.TreeView  `desc:"the component that we view"`
	// NumView ValueView `desc:"inline struct view of the numbers"`
	// TmpSave ValueView `json:"-" xml:"-" desc:"value view that needs to have SaveTmp called on it whenever a change is made to one of the underlying values -- pass this down to any sub-views created from a parent"`
	// ViewSig ki.Signal `json:"-" xml:"-" desc:"signal for valueview -- only one signal sent when a value has been set -- all related value views interconnect with each other to update when others update"`
}

var KiT_GraphBuilder = kit.Types.AddType(&GraphBuilder{}, GraphBuilderProps)

// AddNewGraphBuilder adds a new GraphBuilder to given parent node, with given name.
func AddNewGraphBuilder(parent ki.Ki, name string) *GraphBuilder {
	return parent.AddNewChild(KiT_GraphBuilder, name).(*GraphBuilder)
}

var GraphBuilderProps = ki.Props{
	"background-color": &gi.Prefs.Colors.Background,
	"color":            &gi.Prefs.Colors.Font,
	"max-width":        -1, // needed for everything to stretch
	"max-height":       -1,
}

// Config configures a standard setup of entire view
func (gb *GraphBuilder) Config() {
	gb.Lay = gi.LayoutHoriz
	gb.SetProp("spacing", gi.StdDialogVSpaceUnits)
	config := gb.StdFrameConfig()
	mods, updt := gb.ConfigChildren(config, false)
	if mods {
		gb.ConfigComponents()
	} else {
		updt = gb.UpdateStart()
	}
	gb.UpdateEnd(updt)
}

// SliderLayConfig configures the sliders layout
func (gb *GraphBuilder) ConfigComponents() {
	tv := gb.Child(0).(*giv.TreeView)
	tv.SetRootNode(&AllComponents)
	dc := gb.Child(1).(*DropCanvas)
	dc.Fill = true
}

func (gb *GraphBuilder) StdFrameConfig() kit.TypeAndNameList {
	config := kit.TypeAndNameList{}
	config.Add(giv.KiT_TreeView, "components")
	config.Add(KiT_DropCanvas, "canvas")
	return config
}
