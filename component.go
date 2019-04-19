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
	// "github.com/goki/gi/giv"
	// "github.com/goki/gi/units"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

/////////////////////////////////////////////////////////////////////////////
//  Component

// Component shows a color, using sliders to set values,
type Component struct {
	ki.Node
	Params map[string]string
}

var KiT_Component = kit.Types.AddType(&Component{}, nil)

// AddNewComponent adds a new Component to given parent node, with given name.
func AddNewComponent(parent ki.Ki, name string) *Component {
	return parent.AddNewChild(KiT_Component, name).(*Component)
}

var AllComponents Component

func init() {
	AllComponents.InitName(&AllComponents, "AllComponents")
	AddNewComponent(&AllComponents, "TransferMechanism")
	AddNewComponent(&AllComponents, "ProcessingMechanism")
	AddNewComponent(&AllComponents, "IntegratorMechanism")
}

/////////////////////////////////////////////////////////////////////////////
//  ComponentView

// ComponentView shows a color, using sliders to set values,
type ComponentView struct {
	svg.Rect
}

var KiT_ComponentView = kit.Types.AddType(&ComponentView{}, nil)

// AddNewComponentView adds a new ComponentView to given parent node, with given name.
func AddNewComponentView(parent ki.Ki, name string) *ComponentView {
	return parent.AddNewChild(KiT_ComponentView, name).(*ComponentView)
}
