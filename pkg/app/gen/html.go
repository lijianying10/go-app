//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type tag struct {
	Name          string
	Type          tagType
	Doc           string
	Attrs         []attr
	EventHandlers []eventHandler
}

type tagType int

const (
	parent tagType = iota
	privateParent
	selfClosing
)

var tags = []tag{
	{
		// A:
		Name: "A",
		Doc:  "defines a hyperlink.",
		Attrs: withGlobalAttrs(attrsByNames(
			"download",
			"href",
			"hreflang",
			"media",
			"ping",
			"rel",
			"target",
			"type",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Abbr",
		Doc:           "defines an abbreviation or an acronym.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Address",
		Doc:           "defines contact information for the author/owner of a document.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Area",
		Type: selfClosing,
		Doc:  "defines an area inside an image-map.",
		Attrs: withGlobalAttrs(attrsByNames(
			"alt",
			"coords",
			"download",
			"href",
			"hreflang",
			"media",
			"rel",
			"shape",
			"target",
			"type",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Article",
		Doc:           "defines an article.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Aside",
		Doc:           "defines content aside from the page content.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Audio",
		Doc:  "defines sound content.",
		Attrs: withGlobalAttrs(attrsByNames(
			"autoplay",
			"controls",
			"crossorigin",
			"loop",
			"muted",
			"preload",
			"src",
		)...),
		EventHandlers: withMediaEventHandlers(withGlobalEventHandlers()...),
	},

	// B:
	{
		Name:          "B",
		Doc:           "defines bold text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Base",
		Type: selfClosing,
		Doc:  "specifies the base URL/target for all relative URLs in a document.",
		Attrs: withGlobalAttrs(attrsByNames(
			"href",
			"target",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Bdi",
		Doc:           "isolates a part of text that might be formatted in a different direction from other text outside it.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Bdo",
		Doc:           "overrides the current text direction.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Blockquote",
		Doc:  "defines a section that is quoted from another source.",
		Attrs: withGlobalAttrs(attrsByNames(
			"cite",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:  "Body",
		Type:  privateParent,
		Doc:   "defines the document's body.",
		Attrs: withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(eventHandlersByName(
			"onafterprint",
			"onbeforeprint",
			"onbeforeunload",
			"onerror",
			"onhashchange",
			"onload",
			"onmessage",
			"onoffline",
			"ononline",
			"onpagehide",
			"onpageshow",
			"onpopstate",
			"onresize",
			"onstorage",
			"onunload",
		)...),
	},
	{
		Name:          "Br",
		Type:          selfClosing,
		Doc:           "defines a single line break.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Button",
		Doc:  "defines a clickable button.",
		Attrs: withGlobalAttrs(attrsByNames(
			"autofocus",
			"disabled",
			"form",
			"formaction",
			"formenctype",
			"formmethod",
			"formnovalidate",
			"formtarget",
			"name",
			"type",
			"value",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// C:
	{
		Name: "Canvas",
		Doc:  "is used to draw graphics on the fly.",
		Attrs: withGlobalAttrs(attrsByNames(
			"height",
			"width",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Caption",
		Doc:           "defines a table caption.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Cite",
		Doc:           "defines the title of a work.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Code",
		Doc:           "defines a piece of computer code.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Col",
		Type: selfClosing,
		Doc:  "specifies column properties for each column within a colgroup element.",
		Attrs: withGlobalAttrs(attrsByNames(
			"span",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "ColGroup",
		Doc:  "specifies a group of one or more columns in a table for formatting.",
		Attrs: withGlobalAttrs(attrsByNames(
			"span",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// D:
	{
		Name: "Data",
		Doc:  "links the given content with a machine-readable translation.",
		Attrs: withGlobalAttrs(attrsByNames(
			"value",
		)...),
	},
	{
		Name:          "DataList",
		Doc:           "specifies a list of pre-defined options for input controls.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Dd",
		Doc:           "defines a description/value of a term in a description list.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Del",
		Doc:  "defines text that has been deleted from a document.",
		Attrs: withGlobalAttrs(attrsByNames(
			"cite",
			"datetime",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Details",
		Doc:  "defines additional details that the user can view or hide.",
		Attrs: withGlobalAttrs(attrsByNames(
			"open",
		)...),
		EventHandlers: withGlobalEventHandlers(eventHandlersByName(
			"ontoggle",
		)...),
	},
	{
		Name:          "Dfn",
		Doc:           "represents the defining instance of a term.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Dialog",
		Doc:  "defines a dialog box or window.",
		Attrs: withGlobalAttrs(attrsByNames(
			"open",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Div",
		Doc:           "defines a section in a document.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Dl",
		Doc:           "defines a description list.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Dt",
		Doc:           "defines a term/name in a description list.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},

	// E:
	{
		Name: "Elem",
		Doc:  "represents an customizable HTML element.",
		Attrs: withGlobalAttrs(attrsByNames(
			"xmlns",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "ElemSelfClosing",
		Type: selfClosing,
		Doc:  "represents a self closing custom HTML element.",
		Attrs: withGlobalAttrs(attrsByNames(
			"xmlns",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Em",
		Doc:           "defines emphasized text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Embed",
		Type: selfClosing,
		Doc:  "defines a container for an external (non-HTML) application.",
		Attrs: withGlobalAttrs(attrsByNames(
			"height",
			"src",
			"type",
			"width",
		)...),
		EventHandlers: withMediaEventHandlers(withGlobalEventHandlers()...),
	},

	// F:
	{
		Name: "FieldSet",
		Doc:  "groups related elements in a form.",
		Attrs: withGlobalAttrs(attrsByNames(
			"disabled",
			"form",
			"name",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "FigCaption",
		Doc:           "defines a caption for a figure element.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Figure",
		Doc:           "specifies self-contained content.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Footer",
		Doc:           "defines a footer for a document or section.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Form",
		Doc:  "defines an HTML form for user input.",
		Attrs: withGlobalAttrs(attrsByNames(
			"accept-charset",
			"action",
			"autocomplete",
			"enctype",
			"method",
			"name",
			"novalidate",
			"target",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// H:
	{
		Name:          "H1",
		Doc:           "defines HTML heading.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "H2",
		Doc:           "defines HTML heading.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "H3",
		Doc:           "defines HTML heading.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "H4",
		Doc:           "defines HTML heading.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "H5",
		Doc:           "defines HTML heading.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "H6",
		Doc:           "defines HTML heading.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:  "Head",
		Doc:   "defines information about the document.",
		Attrs: withGlobalAttrs(attrsByNames()...),
	},
	{
		Name:          "Header",
		Doc:           "defines a header for a document or section.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Hr",
		Type:          selfClosing,
		Doc:           "defines a thematic change in the content.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:  "Html",
		Type:  privateParent,
		Doc:   "defines the root of an HTML document.",
		Attrs: withGlobalAttrs(),
	},

	// I:
	{
		Name:          "I",
		Doc:           "defines a part of text in an alternate voice or mood.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "IFrame",
		Doc:  "defines an inline frame.",
		Attrs: withGlobalAttrs(attrsByNames(
			"allow",
			"allowfullscreen",
			"allowpaymentrequest",
			"height",
			"name",
			"referrerpolicy",
			"sandbox",
			"src",
			"srcdoc",
			"width",
			"loading",
		)...),
		EventHandlers: withGlobalEventHandlers(eventHandlersByName(
			"onload",
		)...,
		),
	},
	{
		Name: "Img",
		Type: selfClosing,
		Doc:  "defines an image.",
		Attrs: withGlobalAttrs(attrsByNames(
			"alt",
			"crossorigin",
			"height",
			"ismap",
			"sizes",
			"src",
			"srcset",
			"usemap",
			"width",
		)...),
		EventHandlers: withMediaEventHandlers(withGlobalEventHandlers(
			eventHandlersByName(
				"onload",
			)...,
		)...),
	},
	{
		Name: "Input",
		Type: selfClosing,
		Doc:  "defines an input control.",
		Attrs: withGlobalAttrs(attrsByNames(
			"accept",
			"alt",
			"autocomplete",
			"autofocus",
			"capture",
			"checked",
			"dirname",
			"disabled",
			"form",
			"formaction",
			"formenctype",
			"formmethod",
			"formnovalidate",
			"formtarget",
			"height",
			"list",
			"max",
			"maxlength",
			"min",
			"multiple",
			"name",
			"pattern",
			"placeholder",
			"readonly",
			"required",
			"size",
			"src",
			"step",
			"type",
			"value",
			"width",
		)...),
		EventHandlers: withGlobalEventHandlers(eventHandlersByName(
			"onload",
		)...,
		),
	},
	{
		Name:          "Ins",
		Doc:           "defines a text that has been inserted into a document.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},

	// K:
	{
		Name:          "Kbd",
		Doc:           "defines keyboard input.",
		Attrs:         withGlobalAttrs(attrsByNames()...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// L:
	{
		Name: "Label",
		Doc:  "defines a label for an input element.",
		Attrs: withGlobalAttrs(attrsByNames(
			"for",
			"form",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Legend",
		Doc:           "defines a caption for a fieldset element.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Li",
		Doc:  "defines a list item.",
		Attrs: withGlobalAttrs(attrsByNames(
			"value",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Link",
		Type: selfClosing,
		Doc:  "defines the relationship between a document and an external resource (most used to link to style sheets).",
		Attrs: withGlobalAttrs(attrsByNames(
			"crossorigin",
			"href",
			"hreflang",
			"media",
			"rel",
			"sizes",
			"type",
		)...),
		EventHandlers: withGlobalEventHandlers(eventHandlersByName(
			"onload",
		)...),
	},

	// M:
	{
		Name:          "Main",
		Doc:           "specifies the main content of a document.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Map",
		Doc:  "defines a client-side image-map.",
		Attrs: withGlobalAttrs(attrsByNames(
			"name",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Mark",
		Doc:           "defines marked/highlighted text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Meta",
		Type: selfClosing,
		Doc:  ".",
		Attrs: withGlobalAttrs(attrsByNames(
			"charset",
			"content",
			"http-equiv",
			"name",
			"property",
		)...),
	},
	{
		Name: "Meter",
		Doc:  "defines a scalar measurement within a known range (a gauge).",
		Attrs: withGlobalAttrs(attrsByNames(
			"form",
			"high",
			"low",
			"max",
			"min",
			"optimum",
			"value",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// N:
	{
		Name:          "Nav",
		Doc:           "defines navigation links.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:  "NoScript",
		Doc:   "defines an alternate content for users that do not support client-side scripts.",
		Attrs: withGlobalAttrs(attrsByNames()...),
	},

	// O:
	{
		Name: "Object",
		Doc:  "defines an embedded object.",
		Attrs: withGlobalAttrs(attrsByNames(
			"data",
			"form",
			"height",
			"name",
			"type",
			"usemap",
			"width",
		)...),
		EventHandlers: withMediaEventHandlers(withGlobalEventHandlers()...),
	},
	{
		Name: "Ol",
		Doc:  "defines an ordered list.",
		Attrs: withGlobalAttrs(attrsByNames(
			"reversed",
			"start",
			"type",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "OptGroup",
		Doc:  "defines a group of related options in a drop-down list.",
		Attrs: withGlobalAttrs(attrsByNames(
			"disabled",
			"label",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Option",
		Doc:  "defines an option in a drop-down list.",
		Attrs: withGlobalAttrs(attrsByNames(
			"disabled",
			"label",
			"selected",
			"value",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Output",
		Doc:  ".",
		Attrs: withGlobalAttrs(attrsByNames(
			"for",
			"form",
			"name",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// P:
	{
		Name:          "P",
		Doc:           "defines a paragraph.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Param",
		Type: selfClosing,
		Doc:  "defines a parameter for an object.",
		Attrs: withGlobalAttrs(attrsByNames(
			"name",
			"value",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Picture",
		Doc:           "defines a container for multiple image resources.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Pre",
		Doc:           "defines preformatted text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Progress",
		Doc:  "represents the progress of a task.",
		Attrs: withGlobalAttrs(attrsByNames(
			"max",
			"value",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// Q:
	{
		Name: "Q",
		Doc:  "defines a short quotation.",
		Attrs: withGlobalAttrs(attrsByNames(
			"cite",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},

	// R:
	{
		Name:          "Rp",
		Doc:           "defines what to show in browsers that do not support ruby annotations.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Rt",
		Doc:           "defines an explanation/pronunciation of characters (for East Asian typography).",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Ruby",
		Doc:           "defines a ruby annotation (for East Asian typography).",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},

	// S:
	{
		Name:          "S",
		Doc:           "Defines text that is no longer correct.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Samp",
		Doc:           "defines sample output from a computer program.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Script",
		Doc:  "defines a client-side script.",
		Attrs: withGlobalAttrs(attrsByNames(
			"async",
			"charset",
			"crossorigin",
			"defer",
			"src",
			"type",
		)...),
		EventHandlers: eventHandlersByName("onload"),
	},
	{
		Name:          "Section",
		Doc:           "defines a section in a document.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Select",
		Doc:  "defines a drop-down list.",
		Attrs: withGlobalAttrs(attrsByNames(
			"autofocus",
			"disabled",
			"form",
			"multiple",
			"name",
			"required",
			"size",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Small",
		Doc:           "defines smaller text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Source",
		Type: selfClosing,
		Doc:  ".",
		Attrs: withGlobalAttrs(attrsByNames(
			"src",
			"srcset",
			"media",
			"sizes",
			"type",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Span",
		Doc:           "defines a section in a document.",
		Attrs:         withGlobalAttrs(attrsByNames()...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Strong",
		Doc:           "defines important text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Style",
		Doc:  "defines style information for a document.",
		Attrs: withGlobalAttrs(attrsByNames(
			"media",
			"type",
		)...),
		EventHandlers: withGlobalEventHandlers(eventHandlersByName(
			"onload",
		)...),
	},
	{
		Name:          "Sub",
		Doc:           "defines subscripted text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Summary",
		Doc:           "defines a visible heading for a details element.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Sup",
		Doc:           "defines superscripted text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},

	// T:
	{
		Name:          "Table",
		Doc:           "defines a table.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "TBody",
		Doc:           "groups the body content in a table.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Td",
		Doc:  "defines a cell in a table.",
		Attrs: withGlobalAttrs(attrsByNames(
			"colspan",
			"headers",
			"rowspan",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:  "Template",
		Doc:   "defines a template.",
		Attrs: withGlobalAttrs(attrsByNames()...),
	},
	{
		Name: "Textarea",
		Doc:  "defines a multiline input control (text area).",
		Attrs: withGlobalAttrs(attrsByNames(
			"autofocus",
			"cols",
			"dirname",
			"disabled",
			"form",
			"maxlength",
			"name",
			"placeholder",
			"readonly",
			"required",
			"rows",
			"wrap",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "TFoot",
		Doc:           "groups the footer content in a table.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Th",
		Doc:  "defines a header cell in a table.",
		Attrs: withGlobalAttrs(attrsByNames(
			"abbr",
			"colspan",
			"headers",
			"rowspan",
			"scope",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "THead",
		Doc:           "groups the header content in a table",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Time",
		Doc:  "defines a date/time.",
		Attrs: withGlobalAttrs(attrsByNames(
			"datetime",
		)...),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:  "Title",
		Doc:   "defines a title for the document.",
		Attrs: withGlobalAttrs(attrsByNames()...),
	},
	{
		Name:          "Tr",
		Doc:           "defines a row in a table.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},

	// U:
	{
		Name:          "U",
		Doc:           "defines text that should be stylistically different from normal text.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name:          "Ul",
		Doc:           "defines an unordered list.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},

	// V:
	{
		Name:          "Var",
		Doc:           "defines a variable.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},
	{
		Name: "Video",
		Doc:  "defines a video or movie.",
		Attrs: withGlobalAttrs(attrsByNames(
			"autoplay",
			"controls",
			"crossorigin",
			"height",
			"loop",
			"muted",
			"poster",
			"preload",
			"src",
			"width",
		)...),
		EventHandlers: withMediaEventHandlers(withGlobalEventHandlers()...),
	},
	{
		Name:          "Wbr",
		Doc:           "defines a possible line-break.",
		Attrs:         withGlobalAttrs(),
		EventHandlers: withGlobalEventHandlers(),
	},

	//SVG label start

	{
		Name: "SVGvkern",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The <vkern> SVG element allows to fine-tweak the vertical distance between two glyphs in top-to-bottom fonts. This process is known as kerning.`,
		Attrs: attrsSVGByNames(),
	},

	{
		Name: "SVGclipPath",
		Doc: `The <clipPath> SVG element defines a clipping path, to be used by the clip-path property.
A clipping path restricts the region to which paint can be applied. Conceptually, parts of the drawing that lie outside of the region bounded by the clipping path are not drawn.`,
		Attrs: attrsSVGByNames(
			"clip-path",
			"clipPathUnits",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-rule",
			"color",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
	},

	{
		Name: "SVGdesc",
		Doc: `The <desc> element provides an accessible, long-text description of any SVG container element or graphics element.
Text in a <desc> element is not rendered as part of the graphic. If the element can be described by visible text, it is possible to reference that text with the aria-describedby attribute. If aria-describedby is used, it will take precedence over <desc>.
The hidden text of a <desc> element can also be concatenated with the visible text of other elements using multiple IDs in an aria-describedby value. In that case, the <desc> element must provide an ID for reference.`,
		Attrs: attrsSVGByNames(
			"Core",
			"id",
			"Styling",
			"class",
			"style",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...),
	},

	{
		Name: "SVGline",
		Doc:  `The <line> element is an SVG basic shape used to create a line connecting two points.`,
		Attrs: attrsSVGByNames(
			"x1",
			"x2",
			"y1",
			"y2",
			"pathLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...),
	},

	{
		Name: "SVGrect",
		Doc:  `The <rect> element is a basic SVG shape that draws rectangles, defined by their position, width, and height. The rectangles may have their corners rounded.`,
		Attrs: attrsSVGByNames(
			"x",
			"y",
			"width",
			"height",
			"rx",
			"ry",
			"pathLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...),
	},

	{
		Name: "SVGellipse",
		Doc: `The <ellipse> element is an SVG basic shape, used to create ellipses based on a center coordinate, and both their x and y radius.

Note: Ellipses are unable to specify the exact orientation of the ellipse (if, for example, you wanted to draw an ellipse tilted at a 45 degree angle), but it can be rotated by using the transform attribute.
`,
		Attrs: attrsSVGByNames(
			"transform",
			"cx",
			"cy",
			"rx",
			"ry",
			"pathLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...),
	},

	{
		Name: "SVGfeSpecularLighting",
		Doc: `The <feSpecularLighting> SVG filter primitive lights a source graphic using the alpha channel as a bump map. The resulting image is an RGBA image based on the light color. The lighting calculation follows the standard specular component of the Phong lighting model. The resulting image depends on the light color, light position and surface geometry of the input bump map. The result of the lighting calculation is added. The filter primitive assumes that the viewer is at infinity in the z direction.
This filter primitive produces an image which contains the specular reflection part of the lighting calculation. Such a map is intended to be combined with a texture using the add term of the arithmetic <feComposite> method. Multiple light sources can be simulated by adding several of these light maps before applying it to the texture image.`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"in",
			"surfaceScale",
			"specularConstant",
			"specularExponent",
			"kernelUnitLength",
		),
	},

	{
		Name: "SVGfeSpotLight",
		Doc: `
The <feSpotLight> SVG filter primitive defines a light source that can be used to create a spotlight effect.
It is used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>.
`,
		Attrs: attrsSVGByNames(
			"x",
			"y",
			"z",
			"pointsAtX",
			"pointsAtY",
			"pointsAtZ",
			"specularExponent",
			"limitingConeAngle",
		),
	},

	{
		Name: "Svg",
		Doc: `The svg element is a container that defines a new coordinate system and viewport. It is used as the outermost element of SVG documents, but it can also be used to embed an SVG fragment inside an SVG or HTML document.

Note: The xmlns attribute is only required on the outermost svg element of SVG documents. It is unnecessary for inner svg elements or inside HTML documents.
`,
		Attrs: attrsSVGByNames(
			"viewBox",
			"baseProfile",
			"contentScriptType",
			"contentStyleType",
			"height",
			"preserveAspectRatio",
			"version",
			"width",
			"x",
			"y",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGDocumentEventHandler(withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...)...)...),
	},

	{
		Name: "SVGfeTurbulence",
		Doc:  `The <feTurbulence> SVG filter primitive creates an image using the Perlin turbulence function. It allows the synthesis of artificial textures like clouds or marble. The resulting image will fill the entire filter primitive subregion.`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"baseFrequency",
			"numOctaves",
			"seed",
			"stitchTiles",
			"type",
		),
	},

	{
		Name: "SVGaltGlyphDef",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The <altGlyphDef> SVG element defines a substitution representation for glyphs.`,
		Attrs: attrsSVGByNames(),
	},

	{
		Name: "SVGdefs",
		Doc: `The <defs> element is used to store graphical objects that will be used at a later time. Objects created inside a <defs> element are not rendered directly. To display them you have to reference them (with a <use> element for example).
Graphical objects can be referenced from anywhere, however, defining these objects inside of a <defs> element promotes understandability of the SVG content and is beneficial to the overall accessibility of the document.`,
		Attrs: attrsSVGByNames(
			"Core",
			"id",
			"lang",
			"Styling",
			"class",
			"style",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...)...),
	},

	{
		Name: "SVGfeMerge",
		Doc:  `The <feMerge> SVG element allows filter effects to be applied concurrently instead of sequentially. This is achieved by other filters storing their output via the result attribute and then accessing it in a <feMergeNode> child.`,
		Attrs: attrsSVGByNames(
			"result",
			"class",
			"style",
		),
	},

	{
		Name: "SVGfeMorphology",
		Doc:  `The <feMorphology> SVG filter primitive is used to erode or dilate the input image. Its usefulness lies especially in fattening or thinning effects.`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"in",
			"operator",
			"radius",
		),
	},

	{
		Name:  "SVGmetadata",
		Doc:   `The <metadata> SVG element adds metadata to SVG content. Metadata is structured information about data. The contents of <metadata> should be elements from other XML namespaces such as RDF, FOAF, etc.`,
		Attrs: attrsSVGByNames(),
	},

	{
		Name: "SVGpath",
		Doc:  `The <path> SVG element is the generic element to define a shape. All the basic shapes can be created with a path element.`,
		Attrs: attrsSVGByNames(
			"d",
			"pathLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...),
	},

	{
		Name: "SVGpolygon",
		Doc: `The <polygon> element defines a closed shape consisting of a set of connected straight line segments. The last point is connected to the first point.
For open shapes, see the <polyline> element.`,
		Attrs: attrsSVGByNames(
			"points",
			"pathLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...),
	},

	{
		Name: "SVGtspan",
		Doc:  `The SVG <tspan> element defines a subtext within a <text> element or another <tspan> element. It allows for adjustment of the style and/or position of that subtext as needed.`,
		Attrs: attrsSVGByNames(
			"x",
			"y",
			"dx",
			"dy",
			"rotate",
			"lengthAdjust",
			"textLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"dominant-baseline",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"text-anchor",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName()...)...),
	},

	{
		Name: "SVGanimate",
		Doc:  `The SVG <animate> element provides a way to animate an attribute of an element over time.`,
		Attrs: attrsSVGByNames(
			"begin",
			"dur",
			"end",
			"min",
			"max",
			"restart",
			"repeatCount",
			"repeatDur",
			"fill",
			"calcMode",
			"values",
			"keyTimes",
			"keySplines",
			"From",
			"To",
			"by",
			"attributeName",
			"additive",
			"accumulate",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(withSVGAnimationEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...)...),
	},

	{
		Name: "SVGaltGlyphItem",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The <altGlyphItem> element provides a set of candidates for glyph substitution by the <altGlyph> element.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGg",
		Doc: `The <g> SVG element is a container used to group other SVG elements.
Transformations applied to the <g> element are performed on its child elements, and its attributes are inherited by its children. It can also group multiple elements to be referenced later with the <use> element.`,
		Attrs: attrsSVGByNames(
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name: "SVGglyphRef",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The glyphRef element provides a single possible glyph to the referencing <altGlyph> substitution.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGmarker",
		Doc: `The <marker> element defines the graphic that is to be used for drawing arrowheads or polymarkers on a given <path>, <line>, <polyline> or <polygon> element.
Markers are attached to shapes using the marker-start, marker-mid, and marker-end properties.`,
		Attrs: attrsSVGByNames(
			"marker-start",
			"marker-mid",
			"marker-end",
			"markerHeight",
			"markerUnits",
			"markerWidth",
			"orient",
			"preserveAspectRatio",
			"refX",
			"refY",
			"viewBox",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGaltGlyph",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The <altGlyph> SVG element allows sophisticated selection of the glyphs used to render its child character data.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeMergeNode",
		Doc:  `The feMergeNode takes the result of another filter to be processed by its parent <feMerge>.`,
		Attrs: attrsSVGByNames(
			"Core",
			"in",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfePointLight",
		Doc:  `The <fePointLight> filter primitive defines a light source which allows to create a point light effect. It that can be used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>.`,
		Attrs: attrsSVGByNames(
			"x",
			"y",
			"z",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGtitle",
		Doc: `The <title> element provides an accessible, short-text description of any SVG container element or graphics element.
Text in a <title> element is not rendered as part of the graphic, but browsers usually display it as a tooltip. If an element can be described by visible text, it is recommended to reference that text with an aria-labelledby attribute rather than using the <title> element.

Note: For backward compatibility with SVG 1.1, <title> elements should be the first child element of their parent.
`,
		Attrs: attrsSVGByNames(
			"Core",
			"id",
			"Styling",
			"class",
			"style",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name: "SVGuse",
		Doc:  `The <use> element takes nodes from within the SVG document, and duplicates them somewhere else.`,
		Attrs: attrsSVGByNames(
			"x",
			"y",
			"width",
			"height",
			"href",
			"viewBox",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name:          "SVGmpath",
		Doc:           `The <mpath> sub-element for the <animateMotion> element provides the ability to reference an external <path> element as the definition of a motion path.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGset",
		Doc: `The SVG <set> element provides a simple means of just setting the value of an attribute for a specified duration.
It supports all attribute types, including those that cannot reasonably be interpolated, such as string and boolean values. For attributes that can be reasonably be interpolated, the <animate> is usually preferred.

Note: The <set> element is non-additive. The additive and accumulate attributes are not allowed, and will be ignored if specified.
`,
		Attrs: attrsSVGByNames(
			"additive",
			"accumulate",
			"To",
			"begin",
			"dur",
			"end",
			"min",
			"max",
			"restart",
			"repeatCount",
			"repeatDur",
			"fill",
			"attributeName",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(withSVGAnimationEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...)...),
	},

	{
		Name: "SVGa",
		Doc: `The <a> SVG element creates a hyperlink to other web pages, files, locations in the same page, email addresses, or any other URL. It is very similar to HTML's <a> element.
SVG's <a> element is a container, which means you can create a link around text (like in HTML) but also around any shape.`,
		Attrs: attrsSVGByNames(
			"href",
			"target",
			"Core",
			"id",
			"lang",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...)...),
	},

	{
		Name: "SVGglyph",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
A <glyph> defines a single glyph in an SVG font.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeDistantLight",
		Doc:  `The <feDistantLight> filter primitive defines a distant light source that can be used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>.`,
		Attrs: attrsSVGByNames(
			"azimuth",
			"elevation",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name:          "SVGfeFuncA",
		Doc:           `The <feFuncA> SVG filter primitive defines the transfer function for the alpha component of the input graphic of its parent <feComponentTransfer> element.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGpattern",
		Doc: `The <pattern> element defines a graphics object which can be redrawn at repeated x- and y-coordinate intervals ("tiled") to cover an area.
The <pattern> is referenced by the fill and/or stroke attributes on other graphics elements to fill or stroke those elements with the referenced pattern.`,
		Attrs: attrsSVGByNames(
			"fill",
			"stroke",
			"height",
			"href",
			"patternContentUnits",
			"patternTransform",
			"patternUnits",
			"preserveAspectRatio",
			"viewBox",
			"width",
			"x",
			"y",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name:          "SVGfeFuncB",
		Doc:           `The <feFuncB> SVG filter primitive defines the transfer function for the blue component of the input graphic of its parent <feComponentTransfer> element.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeTile",
		Doc:  `The <feTile> SVG filter primitive allows to fill a target rectangle with a repeated, tiled pattern of an input image. The effect is similar to the one of a <pattern>.`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"in",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGforeignObject",
		Doc:  `The <foreignObject> SVG element includes elements from a different XML namespace. In the context of a browser, it is most likely (X)HTML.`,
		Attrs: attrsSVGByNames(
			"height",
			"width",
			"x",
			"y",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGDocumentEventHandler(withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...)...)...),
	},

	{
		Name: "SVGanimateTransform",
		Doc:  `The animateTransform element animates a transformation attribute on its target element, thereby allowing animations to control translation, scaling, rotation, and/or skewing.`,
		Attrs: attrsSVGByNames(
			"by",
			"From",
			"To",
			"type",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeBlend",
		Doc:  `The <feBlend> SVG filter primitive composes two objects together ruled by a certain blending mode. This is similar to what is known from image editing software when blending two layers. The mode is defined by the mode attribute.`,
		Attrs: attrsSVGByNames(
			"mode",
			"class",
			"style",
			"in",
			"in2",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeComposite",
		Doc: `The <feComposite> SVG filter primitive performs the combination of two input images pixel-wise in image space using one of the Porter-Duff compositing operations: over, in, atop, out, xor, lighter, or arithmetic.
The table below shows each of these operations using an image of the MDN logo composited with a red circle:



  Operation
  Description


  
	over 
	  
	
  
  
	The source graphic defined by the in attribute
	(the MDN logo) is placed over the destination graphic defined by the
	in2 attribute (the circle).
	
	  This is the default operation, which will be used if no
	  operation or an unsupported operation is specified.
	
  


  
	in 
	  
	
  
  
	The parts of the source graphic defined by the in attribute
	that overlap the destination graphic defined in the
	in2 attribute, replace the destination graphic.
  


  
	out 
	  
	
  
  
	The parts of the source graphic defined by the in attribute
	that fall outside the destination graphic defined in the
	in2 attribute, are displayed.
  


  
	atop 
	  
	
  
  
	The parts of the source graphic defined in the
	in attribute, which overlap the destination graphic defined
	in the in2 attribute, replace the destination graphic. The
	parts of the destination graphic that do not overlap with the source
	graphic stay untouched.
  


  
	xor 
	  
	
  
  
	The non-overlapping regions of the source graphic defined in the
	in attribute and the destination graphic defined in the
	in2 attribute are combined.
  


  
	lighter 
	  
	
  
  
	The sum of the source graphic defined in the in attribute
	and the destination graphic defined in the in2 attribute is
	displayed.
  


  
	
	  arithmetic
	  
	  
	
  
  
	
	  The arithmetic operation is useful for combining the
	  output from the <feDiffuseLighting> and
	  <feSpecularLighting> filters with texture
	  data. If the arithmetic operation is chosen, each result
	  pixel is computed using the following formula:
	
	result = k1*i1*i2 + k2*i1 + k3*i2 + k4
	where:
	
	  
		i1 and i2 indicate the corresponding pixel
		channel values of the input image, which map to
		in and in2 respectively
	  
	  
		k1, k2,
		k3, and k4 indicate the
		values of the attributes with the same name.
	  
	
  


`,
		Attrs: attrsSVGByNames(
			"in",
			"in2",
			"k1",
			"k2",
			"k3",
			"k4",
			"class",
			"style",
			"operator",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeFlood",
		Doc:  `The <feFlood> SVG filter primitive fills the filter subregion with the color and opacity defined by flood-color and flood-opacity.`,
		Attrs: attrsSVGByNames(
			"flood-color",
			"flood-opacity",
			"class",
			"style",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGmask",
		Doc:  `The <mask> element defines an alpha mask for compositing the current object into the background. A mask is used/referenced using the mask property.`,
		Attrs: attrsSVGByNames(
			"mask",
			"height",
			"maskContentUnits",
			"maskUnits",
			"x",
			"y",
			"width",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"opacity",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGradialGradient",
		Doc: `The <radialGradient> element lets authors define radial gradients that can be applied to fill or stroke of graphical elements.

Note: Don't be confused with CSS radial-gradient() as CSS gradients can only apply to HTML elements where SVG gradient can only apply to SVG elements.
`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGtextPath",
		Doc:  `To render text along the shape of a <path>, enclose the text in a <textPath> element that has an href attribute with a reference to the <path> element.`,
		Attrs: attrsSVGByNames(
			"href",
			"lengthAdjust",
			"method",
			"path",
			"side",
			"spacing",
			"startOffset",
			"textLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name: "SVGdiscard",
		Doc: `The <discard> SVG element allows authors to specify the time at which particular elements are to be discarded, thereby reducing the resources required by an SVG user agent. This is particularly useful to help SVG viewers conserve memory while displaying long-running documents.
The <discard> element may occur wherever the <animate> element may.`,
		Attrs: attrsSVGByNames(
			"begin",
			"href",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGhkern",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The <hkern> SVG element allows to fine-tweak the horizontal distance between two glyphs. This process is known as kerning.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGsymbol",
		Doc: `The <symbol> element is used to define graphical template objects which can be instantiated by a <use> element.
The use of <symbol> elements for graphics that are used multiple times in the same document adds structure and semantics. Documents that are rich in structure may be rendered graphically, as speech, or as Braille, and thus promote accessibility.`,
		Attrs: attrsSVGByNames(
			"height",
			"preserveAspectRatio",
			"refX",
			"refY",
			"viewBox",
			"width",
			"x",
			"y",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...)...),
	},

	{
		Name: "SVGtref",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The textual content for a <text> SVG element can be either character data directly embedded within the <text> element or the character data content of a referenced element, where the referencing is specified with a <tref> element.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeDisplacementMap",
		Doc: `The <feDisplacementMap> SVG filter primitive uses the pixel values from the image from in2 to spatially displace the image from in.
The formula for the transformation looks like this:
P'(x,y) ← P(x + scale * (XC(x,y) - 0.5), y + scale * (YC(x,y) - 0.5))
where P(x,y) is the input image, in, and P'(x,y) is the destination. XC(x,y) and YC(x,y) are the component values of the channel designated by xChannelSelector and yChannelSelector.`,
		Attrs: attrsSVGByNames(
			"in2",
			"in",
			"xChannelSelector",
			"yChannelSelector",
			"class",
			"style",
			"scale",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfilter",
		Doc:  `The <filter> SVG element defines a custom filter effect by grouping atomic filter primitives. It is never rendered itself, but must be used by the filter attribute on SVG elements, or the filter CSS property for SVG/HTML elements.`,
		Attrs: attrsSVGByNames(
			"filter",
			"class",
			"style",
			"x",
			"y",
			"width",
			"height",
			"filterRes",
			"filterUnits",
			"primitiveUnits",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGpolyline",
		Doc:  `The <polyline> SVG element is an SVG basic shape that creates straight lines connecting several points. Typically a polyline is used to create open shapes as the last point doesn't have to be connected to the first point. For closed shapes see the <polygon> element.`,
		Attrs: attrsSVGByNames(
			"points",
			"pathLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name: "SVGfeGaussianBlur",
		Doc:  `The <feGaussianBlur> SVG filter primitive blurs the input image by the amount specified in stdDeviation, which defines the bell-curve.`,
		Attrs: attrsSVGByNames(
			"stdDeviation",
			"class",
			"style",
			"in",
			"edgeMode",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfont",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.
The <font> SVG element defines a font to be used for text layout.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGstyle",
		Doc: `The SVG <style> element allows style sheets to be embedded directly within SVG content.

Note: SVG's style element has the same attributes as the corresponding element in HTML (see HTML's <style> element).
`,
		Attrs: attrsSVGByNames(
			"type",
			"media",
			"title",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name: "SVGanimateMotion",
		Doc: `The SVG <animateMotion> element provides a way to define how an element moves along a motion path.

Note: To reuse an existing path, it will be necessary to use an <mpath> element inside the <animateMotion> element instead of the path attribute.
`,
		Attrs: attrsSVGByNames(
			"path",
			"keyPoints",
			"keyTimes",
			"d",
			"rotate",
			"calcMode",
			"begin",
			"dur",
			"end",
			"min",
			"max",
			"restart",
			"repeatCount",
			"repeatDur",
			"fill",
			"values",
			"keySplines",
			"From",
			"To",
			"by",
			"attributeName",
			"additive",
			"accumulate",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(withSVGAnimationEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...)...),
	},

	{
		Name: "SVGfeColorMatrix",
		Doc: `The <feColorMatrix> SVG filter element changes colors based on a transformation matrix. Every pixel's color value [R,G,B,A] is matrix multiplied by a 5 by 5 color matrix to create new color [R',G',B',A'].

Note: The prime symbol ' is used in mathematics indicate the result of a transformation.

| R' |     | r1 r2 r3 r4 r5 |   | R |
| G' |     | g1 g2 g3 g4 g5 |   | G |
| B' |  =  | b1 b2 b3 b4 b5 | * | B |
| A' |     | a1 a2 a3 a4 a5 |   | A |
| 1  |     | 0  0  0  0  1  |   | 1 |

In simplified terms, below is how each color channel in the new pixel is calculated. The last row is ignored because its values are constant.
R' = r1*R + r2*G + r3*B + r4*A + r5
G' = g1*R + g2*G + g3*B + g4*A + g5
B' = b1*R + b2*G + b3*B + b4*A + b5
A' = a1*R + a2*G + a3*B + a4*A + a5

Take the amount of red in the new pixel, or R':
It is the sum of:

r1 times the old pixel's red R,
r2 times the old pixel's green G,
r3 times of the old pixel's blue B,
r4 times the old pixel's alpha A,
plus a shift r5.

These specified amounts can be any real number, though the final R' will be clamped between 0 and 1. The same goes for G', B', and A'.
R'      =      r1 * R      +        r2 * G      +       r3 * B      +       r4 * A       +       r5
New red = [ r1 * old red ] + [ r2 * old green ] + [ r3 * old Blue ] + [ r4 * old Alpha ] + [ shift of r5 ]

If, say, we want to make a completely black image redder, we can make the r5 a positive real number x, boosting the redness on every pixel of the new image by x.
An identity matrix looks like this:
 R G B A W
R' | 1 0 0 0 0 |
G' | 0 1 0 0 0 |
B' | 0 0 1 0 0 |
A' | 0 0 0 1 0 |

In it, every new value is exactly 1 times its old value, with nothing else added. It is recommended to start manipulating the matrix from here.`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"in",
			"type",
			"values",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name:          "SVGlinearGradient",
		Doc:           `The <linearGradient> element lets authors define linear gradients to apply to other SVG elements.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGscript",
		Doc: `The SVG script element allows to add scripts to an SVG document.

Note: While SVG's script element is equivalent to the HTML <script> element, it has some discrepancies, like it uses the href attribute instead of src and it doesn't support ECMAScript modules so far (See browser compatibility below for details)
`,
		Attrs: attrsSVGByNames(
			"href",
			"type",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name:          "SVGfeFuncR",
		Doc:           `The <feFuncR> SVG filter primitive defines the transfer function for the red component of the input graphic of its parent <feComponentTransfer> element.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeImage",
		Doc:  `The <feImage> SVG filter primitive fetches image data from an external source and provides the pixel data as output (meaning if the external source is an SVG image, it is rasterized.)`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"preserveAspectRatio",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeOffset",
		Doc:  `The <feOffset> SVG filter primitive allows to offset the input image. The input image as a whole is offset by the values specified in the dx and dy attributes.`,
		Attrs: attrsSVGByNames(
			"dx",
			"dy",
			"class",
			"style",
			"in",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeDiffuseLighting",
		Doc: `The <feDiffuseLighting> SVG filter primitive lights an image using the alpha channel as a bump map. The resulting image, which is an RGBA opaque image, depends on the light color, light position and surface geometry of the input bump map.
The light map produced by this filter primitive can be combined with a texture image using the multiply term of the arithmetic operator of the <feComposite> filter primitive. Multiple light sources can be simulated by adding several of these light maps together before applying it to the texture image.`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"in",
			"surfaceScale",
			"diffuseConstant",
			"kernelUnitLength",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGview",
		Doc:  `A view is a defined way to view the image, like a zoom level or a detail view.`,
		Attrs: attrsSVGByNames(
			"viewBox",
			"preserveAspectRatio",
			"zoomAndPan",
			"viewTarget",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGcircle",
		Doc:  `The <circle> SVG element is an SVG basic shape, used to draw circles based on a center point and a radius.`,
		Attrs: attrsSVGByNames(
			"cx",
			"cy",
			"r",
			"pathLength",
			"Core",
			"id",
			"tabindex",
			"Styling",
			"class",
			"style",
			"Conditional_Processing",
			"requiredExtensions",
			"systemLanguage",
			"Presentation",
			"clip-path",
			"clip-rule",
			"color",
			"color-interpolation",
			"color-rendering",
			"cursor",
			"display",
			"fill",
			"fill-opacity",
			"fill-rule",
			"filter",
			"mask",
			"opacity",
			"pointer-events",
			"shape-rendering",
			"stroke",
			"stroke-dasharray",
			"stroke-dashoffset",
			"stroke-linecap",
			"stroke-linejoin",
			"stroke-miterlimit",
			"stroke-opacity",
			"stroke-width",
			"transform",
			"vector-effect",
			"visibility",
		),
		EventHandlers: withSVGGraphicalEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name: "SVGcursor",
		Doc: `Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.

Note: The CSS cursor property should be used instead of this element.

The <cursor> SVG element can be used to define a platform-independent custom cursor. A recommended approach for defining a platform-independent custom cursor is to create a PNG image and define a cursor element that references the PNG image and identifies the exact position within the image which is the pointer position (i.e., the hot spot).
The PNG format is recommended because it supports the ability to define a transparency mask via an alpha channel. If a different image format is used, this format should support the definition of a transparency mask (two options: provide an explicit alpha channel or use a particular pixel color to indicate transparency). If the transparency mask can be determined, the mask defines the shape of the cursor; otherwise, the cursor is an opaque rectangle. Typically, the other pixel information (e.g., the R, G and B channels) defines the colors for those parts of the cursor which are not masked out. Note that cursors usually contain at least two colors so that the cursor can be visible over most backgrounds.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeComponentTransfer",
		Doc: `The <feComponentTransfer> SVG filter primitive performs color-component-wise remapping of data for each pixel. It allows operations like brightness adjustment, contrast adjustment, color balance or thresholding.
The calculations are performed on non-premultiplied color values. The colors are modified by changing each channel (R, G, B, and A) to the result of what the children <feFuncR>, <feFuncB>, <feFuncG>, and <feFuncA> return. If more than one of the same element is provided, the last one specified is used, and if no element is supplied to modify one of the channels, the effect is the same is if an identity transformation had been given for that channel.`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"in",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGfeConvolveMatrix",
		Doc: `The <feConvolveMatrix> SVG filter primitive applies a matrix convolution filter effect. A convolution combines pixels in the input image with neighboring pixels to produce a resulting image. A wide variety of imaging operations can be achieved through convolutions, including blurring, edge detection, sharpening, embossing and beveling.
A matrix convolution is based on an n-by-m matrix (the convolution kernel) which describes how a given pixel value in the input image is combined with its neighboring pixel values to produce a resulting pixel value. Each result pixel is determined by applying the kernel matrix to the corresponding source pixel and its neighboring pixels. The basic convolution formula which is applied to each color value for a given pixel is:

COLORX,Y = (
SUM I=0 to [orderY-1] {
SUM J=0 to [orderX-1] {
SOURCE X-targetX+J, Y-targetY+I * kernelMatrixorderX-J-1, orderY-I-1
}
}
) / divisor + bias * ALPHAX,Y

where "orderX" and "orderY" represent the X and Y values for the 'order' attribute, "targetX" represents the value of the 'targetX' attribute, "targetY" represents the value of the 'targetY' attribute, "kernelMatrix" represents the value of the 'kernelMatrix' attribute, "divisor" represents the value of the 'divisor' attribute, and "bias" represents the value of the 'bias' attribute.
Note in the above formulas that the values in the kernel matrix are applied such that the kernel matrix is rotated 180 degrees relative to the source and destination images in order to match convolution theory as described in many computer graphics textbooks.
To illustrate, suppose you have an input image which is 5 pixels by 5 pixels, whose color values for one of the color channels are as follows:
0    20  40 235 235
100 120 140 235 235
200 220 240 235 235
225 225 255 255 255
225 225 255 255 255

and you define a 3-by-3 convolution kernel as follows:
1 2 3
4 5 6
7 8 9

Let's focus on the color value at the second row and second column of the image (source pixel value is 120). Assuming the simplest case (where the input image's pixel grid aligns perfectly with the kernel's pixel grid) and assuming default values for attributes 'divisor', 'targetX' and 'targetY', then resulting color value will be:
(9*0   + 8*20  + 7*40 +
6*100 + 5*120 + 4*140 +
3*200 + 2*220 + 1*240) / (9+8+7+6+5+4+3+2+1)
`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"in",
			"order",
			"kernelMatrix",
			"divisor",
			"bias",
			"targetX",
			"targetY",
			"edgeMode",
			"kernelUnitLength",
			"preserveAlpha",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGimage",
		Doc: `The <image> SVG element includes images inside SVG documents. It can display raster image files or other SVG files.
The only image formats SVG software must support are JPEG, PNG, and other SVG files. Animated GIF behavior is undefined.
SVG files displayed with <image> are treated as an image: external resources aren't loaded, :visited styles aren't applied, and they cannot be interactive. To include dynamic SVG elements, try <use> with an external URL. To include SVG files and run scripts inside them, try <object> inside of <foreignObject>.

Note: The HTML spec defines <image> as a synonym for <img> while parsing HTML. This specific element and its behavior only apply inside SVG documents or inline SVGs.
`,
		Attrs: attrsSVGByNames(
			"class",
			"style",
			"transform",
			"x",
			"y",
			"width",
			"height",
			"href",
			"preserveAspectRatio",
			"crossorigin",
		),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	{
		Name: "SVGstop",
		Doc:  `The SVG <stop> element defines a color and its position to use on a gradient. This element is always a child of a <linearGradient> or <radialGradient> element.`,
		Attrs: attrsSVGByNames(
			"offset",
			"stop-color",
			"stop-opacity",
			"Core",
			"id",
			"Styling",
			"class",
			"style",
			"Presentation",
			"color",
			"display",
			"visibility",
		),
		EventHandlers: withSVGDocumentElementEventHandler(withSVGGlobalEventHandler(svgEventHandlersByName("onbegin", "onend", "onrepeat")...)...),
	},

	{
		Name:          "SVGfeFuncG",
		Doc:           `The <feFuncG> SVG filter primitive defines the transfer function for the green component of the input graphic of its parent <feComponentTransfer> element.`,
		Attrs:         attrsSVGByNames(),
		EventHandlers: svgEventHandlersByName("onbegin", "onend", "onrepeat"),
	},

	//SVG label end
}

type attr struct {
	Name         string
	NameOverride string
	Type         string
	Key          bool
	Doc          string
}

var attrs = map[string]attr{
	// A:
	"abbr": {
		Name: "Abbr",
		Type: "string",
		Doc:  "specifies an abbreviated version of the content in a header cell.",
	},
	"accept": {
		Name: "Accept",
		Type: "string",
		Doc:  "specifies the types of files that the server accepts (only for file type).",
	},
	"allow": {
		Name: "Allow",
		Type: "string",
		Doc:  "specifies a feature policy. Can be called multiple times to set multiple policies.",
	},
	"allowfullscreen": {
		Name: "AllowFullscreen",
		Type: "bool|force",
		Doc:  "reports whether an iframe can activate fullscreen mode.",
	},
	"allowpaymentrequest": {
		Name: "AllowPaymentRequest",
		Type: "bool|force",
		Doc:  "reports whether an iframe should be allowed to invoke the Payment Request API",
	},
	"aria-*": {
		Name: "Aria",
		Type: "aria|value",
		Doc:  "stores accessible rich internet applications (ARIA) data.",
	},
	"attribute": {
		Name: "Attr",
		Type: "attr|value",
		Doc:  "sets the named attribute with the given value.",
	},
	"accept-charset": {
		Name:         "AcceptCharset",
		NameOverride: "accept-charset",
		Type:         "string",
		Doc:          "specifies the character encodings that are to be used for the form submission.",
	},
	"accesskey": {
		Name: "AccessKey",
		Type: "string",
		Doc:  "specifies a shortcut key to activate/focus an element.",
	},
	"action": {
		Name: "Action",
		Type: "string",
		Doc:  "specifies where to send the form-data when a form is submitted.",
	},
	"alt": {
		Name: "Alt",
		Type: "string",
		Doc:  "specifies an alternate text when the original element fails to display.",
	},
	"async": {
		Name: "Async",
		Type: "bool",
		Doc:  "specifies that the script is executed asynchronously (only for external scripts).",
	},
	"autocomplete": {
		Name: "AutoComplete",
		Type: "on/off",
		Doc:  "specifies whether the element should have autocomplete enabled.",
	},
	"autofocus": {
		Name: "AutoFocus",
		Type: "bool",
		Doc:  "specifies that the element should automatically get focus when the page loads.",
	},
	"autoplay": {
		Name: "AutoPlay",
		Type: "bool",
		Doc:  "specifies that the audio/video will start playing as soon as it is ready.",
	},

	// C:
	"capture": {
		Name: "Capture",
		Type: "string",
		Doc:  "specifies the capture input method in file upload controls",
	},
	"charset": {
		Name: "Charset",
		Type: "string",
		Doc:  "specifies the character encoding.",
	},
	"checked": {
		Name: "Checked",
		Type: "bool",
		Doc:  "specifies that an input element should be pre-selected when the page loads (for checkbox or radio types).",
	},
	"cite": {
		Name: "Cite",
		Type: "url",
		Doc:  "specifies a URL which explains the quote/deleted/inserted text.",
	},
	"class": {
		Name: "Class",
		Type: "string|class",
		Doc:  "specifies one or more classnames for an element (refers to a class in a style sheet).",
	},
	"cols": {
		Name: "Cols",
		Type: "int",
		Doc:  "specifies the visible width of a text area.",
	},
	"colspan": {
		Name: "ColSpan",
		Type: "int",
		Doc:  "specifies the number of columns a table cell should span.",
	},
	"content": {
		Name: "Content",
		Type: "string",
		Doc:  "gives the value associated with the http-equiv or name attribute.",
	},
	"contenteditable": {
		Name: "ContentEditable",
		Type: "bool",
		Doc:  "specifies whether the content of an element is editable or not.",
	},
	"controls": {
		Name: "Controls",
		Type: "bool",
		Doc:  "specifies that audio/video controls should be displayed (such as a play/pause button etc).",
	},
	"coords": {
		Name: "Coords",
		Type: "string",
		Doc:  "specifies the coordinates of the area.",
	},
	"crossorigin": {
		Name: "CrossOrigin",
		Type: "string",
		Doc:  "sets the mode of the request to an HTTP CORS Request.",
	},

	// D:
	"data": {
		Name: "Data",
		Type: "url",
		Doc:  "specifies the URL of the resource to be used by the object.",
	},
	"data-*": {
		Name: "DataSet",
		Type: "data|value",
		Doc:  "stores custom data private to the page or application.",
	},
	"datetime": {
		Name: "DateTime",
		Type: "string",
		Doc:  "specifies the date and time.",
	},
	"default": {
		Name: "Default",
		Type: "bool",
		Doc:  "specifies that the track is to be enabled if the user's preferences do not indicate that another track would be more appropriate.",
	},
	"defer": {
		Name: "Defer",
		Type: "bool",
		Doc:  "specifies that the script is executed when the page has finished parsing (only for external scripts).",
	},
	"dir": {
		Name: "Dir",
		Type: "string",
		Doc:  "specifies the text direction for the content in an element.",
	},
	"dirname": {
		Name: "DirName",
		Type: "string",
		Doc:  "specifies that the text direction will be submitted.",
	},
	"disabled": {
		Name: "Disabled",
		Type: "bool",
		Doc:  "specifies that the specified element/group of elements should be disabled.",
	},
	"download": {
		Name: "Download",
		Type: "string",
		Doc:  "specifies that the target will be downloaded when a user clicks on the hyperlink.",
	},
	"draggable": {
		Name: "Draggable",
		Type: "bool",
		Doc:  "specifies whether an element is draggable or not.",
	},

	// E:
	"enctype": {
		Name: "EncType",
		Type: "string",
		Doc:  "specifies how the form-data should be encoded when submitting it to the server (only for post method).",
	},

	// F:
	"for": {
		Name: "For",
		Type: "string",
		Doc:  "specifies which form element(s) a label/calculation is bound to.",
	},
	"form": {
		Name: "Form",
		Type: "string",
		Doc:  "specifies the name of the form the element belongs to.",
	},
	"formaction": {
		Name: "FormAction",
		Type: "string",
		Doc:  "specifies where to send the form-data when a form is submitted. Only for submit type.",
	},
	"formenctype": {
		Name: "FormEncType",
		Type: "string",
		Doc:  "specifies how form-data should be encoded before sending it to a server. Only for submit type.",
	},
	"formmethod": {
		Name: "FormMethod",
		Type: "string",
		Doc:  "specifies how to send the form-data (which HTTP method to use). Only for submit type.",
	},
	"formnovalidate": {
		Name: "FormNoValidate",
		Type: "bool",
		Doc:  "specifies that the form-data should not be validated on submission. Only for submit type.",
	},
	"formtarget": {
		Name: "FormTarget",
		Type: "string",
		Doc:  "specifies where to display the response after submitting the form. Only for submit type.",
	},

	// H:
	"headers": {
		Name: "Headers",
		Type: "string",
		Doc:  "specifies one or more headers cells a cell is related to.",
	},
	"height": {
		Name: "Height",
		Type: "int",
		Doc:  "specifies the height of the element (in pixels).",
	},
	"hidden": {
		Name: "Hidden",
		Type: "bool",
		Doc:  "specifies that an element is not yet, or is no longer relevant.",
	},
	"high": {
		Name: "High",
		Type: "float64",
		Doc:  "specifies the range that is considered to be a high value.",
	},
	"href": {
		Name: "Href",
		Type: "url",
		Doc:  "specifies the URL of the page the link goes to.",
	},
	"hreflang": {
		Name: "HrefLang",
		Type: "string",
		Doc:  "specifies the language of the linked document.",
	},
	"http-equiv": {
		Name:         "HTTPEquiv",
		NameOverride: "http-equiv",
		Type:         "string",
		Doc:          "provides an HTTP header for the information/value of the content attribute.",
	},

	// I:
	"id": {
		Name: "ID",
		Type: "string",
		Doc:  "specifies a unique id for an element.",
	},
	"ismap": {
		Name: "IsMap",
		Type: "bool",
		Doc:  "specifies an image as a server-side image-map.",
	},

	// K:
	"kind": {
		Name: "Kind",
		Type: "string",
		Doc:  "specifies the kind of text track.",
	},

	// L:
	"label": {
		Name: "Label",
		Type: "string",
		Doc:  "specifies a shorter label for the option.",
	},
	"lang": {
		Name: "Lang",
		Type: "string",
		Doc:  "specifies the language of the element's content.",
	},
	"list": {
		Name: "List",
		Type: "string",
		Doc:  "refers to a datalist element that contains pre-defined options for an input element.",
	},
	"loading": {
		Name: "Loading",
		Type: "string",
		Doc:  "indicates how the browser should load the iframe (eager|lazy).",
	},
	"loop": {
		Name: "Loop",
		Type: "bool",
		Doc:  "specifies that the audio/video will start over again, every time it is finished.",
	},
	"low": {
		Name: "Low",
		Type: "float64",
		Doc:  "specifies the range that is considered to be a low value.",
	},

	// M:
	"max": {
		Name: "Max",
		Type: "any",
		Doc:  "Specifies the maximum value.",
	},
	"maxlength": {
		Name: "MaxLength",
		Type: "int",
		Doc:  "specifies the maximum number of characters allowed in an element.",
	},
	"media": {
		Name: "Media",
		Type: "string",
		Doc:  "specifies what media/device the linked document is optimized for.",
	},
	"method": {
		Name: "Method",
		Type: "string",
		Doc:  "specifies the HTTP method to use when sending form-data.",
	},
	"min": {
		Name: "Min",
		Type: "any",
		Doc:  "specifies a minimum value.",
	},
	"multiple": {
		Name: "Multiple",
		Type: "bool",
		Doc:  "specifies that a user can enter more than one value.",
	},
	"muted": {
		Name: "Muted",
		Type: "bool",
		Doc:  "specifies that the audio output of the video should be muted.",
	},

	// N:
	"name": {
		Name: "Name",
		Type: "string",
		Doc:  "specifies the name of the element.",
	},
	"novalidate": {
		Name: "NoValidate",
		Type: "bool",
		Doc:  "specifies that the form should not be validated when submitted.",
	},

	// O:
	"open": {
		Name: "Open",
		Type: "bool",
		Doc:  "specifies that the details should be visible (open) to the user.",
	},
	"optimum": {
		Name: "Optimum",
		Type: "float64",
		Doc:  "specifies what value is the optimal value for the gauge.",
	},

	// P:
	"pattern": {
		Name: "Pattern",
		Type: "string",
		Doc:  "specifies a regular expression that an input element's value is checked against.",
	},
	"ping": {
		Name: "Ping",
		Type: "string",
		Doc:  "specifies a list of URLs to be notified if the user follows the hyperlink.",
	},
	"placeholder": {
		Name: "Placeholder",
		Type: "string",
		Doc:  "specifies a short hint that describes the expected value of the element.",
	},
	"poster": {
		Name: "Poster",
		Type: "string",
		Doc:  "specifies an image to be shown while the video is downloading, or until the user hits the play button.",
	},
	"preload": {
		Name: "Preload",
		Type: "string",
		Doc:  "specifies if and how the author thinks the audio/video should be loaded when the page loads.",
	},
	"property": {
		Name: "Property",
		Type: "string",
		Doc:  "specifies the property name.",
	},

	// R:
	"readonly": {
		Name: "ReadOnly",
		Type: "bool",
		Doc:  "specifies that the element is read-only.",
	},
	"referrerpolicy": {
		Name: "ReferrerPolicy",
		Type: "string",
		Doc:  "specifies how much/which referrer information that will be sent when processing the iframe attributes",
	},
	"rel": {
		Name: "Rel",
		Type: "string",
		Doc:  "specifies the relationship between the current document and the linked document.",
	},
	"required": {
		Name: "Required",
		Type: "bool",
		Doc:  "specifies that the element must be filled out before submitting the form.",
	},
	"reversed": {
		Name: "Reversed",
		Type: "bool",
		Doc:  "specifies that the list order should be descending (9,8,7...).",
	},
	"role": {
		Name: "Role",
		Type: "string",
		Doc:  "specifies to parsing software the exact function of an element (and its children).",
	},
	"rows": {
		Name: "Rows",
		Type: "int",
		Doc:  "specifies the visible number of lines in a text area.",
	},
	"rowspan": {
		Name: "Rowspan",
		Type: "int",
		Doc:  "specifies the number of rows a table cell should span.",
	},

	// S:
	"sandbox": {
		Name: "Sandbox",
		Type: "any",
		Doc:  "enables an extra set of restrictions for the content in an iframe.",
	},
	"scope": {
		Name: "Scope",
		Type: "string",
		Doc:  "specifies whether a header cell is a header for a column, row, or group of columns or rows.",
	},
	"selected": {
		Name: "Selected",
		Type: "bool",
		Doc:  "specifies that an option should be pre-selected when the page loads.",
	},
	"shape": {
		Name: "Shape",
		Type: "string",
		Doc:  "specifies the shape of the area.",
	},
	"size": {
		Name: "Size",
		Type: "int",
		Doc:  "specifies the width.",
	},
	"sizes": {
		Name: "Sizes",
		Type: "string",
		Doc:  "specifies the size of the linked resource.",
	},
	"span": {
		Name: "Span",
		Type: "int",
		Doc:  "specifies the number of columns to span.",
	},
	"spellcheck": {
		Name: "Spellcheck",
		Type: "bool|force",
		Doc:  "specifies whether the element is to have its spelling and grammar checked or not.",
	},
	"src": {
		Name: "Src",
		Type: "url",
		Doc:  "specifies the URL of the media file.",
	},
	"srcdoc": {
		Name: "SrcDoc",
		Type: "string",
		Doc:  "specifies the HTML content of the page to show in the iframe.",
	},
	"srclang": {
		Name: "SrcLang",
		Type: "string",
		Doc:  `specifies the language of the track text data (required if kind = "subtitles").`,
	},
	"srcset": {
		Name: "SrcSet",
		Type: "url",
		Doc:  "specifies the URL of the image to use in different situations.",
	},
	"start": {
		Name: "Start",
		Type: "int",
		Doc:  "specifies the start value of the ordered list.",
	},
	"step": {
		Name: "Step",
		Type: "float64",
		Doc:  "specifies the legal number intervals for an input field.",
	},
	"style": {
		Name: "Style",
		Type: "style",
		Doc:  "specifies a CSS style for an element. Can be called multiple times to set multiple css styles.",
	},
	"styles": {
		Name: "Styles",
		Type: "style|map",
		Doc:  "specifies CSS styles for an element. Can be called multiple times to set multiple css styles.",
	},

	// T:
	"tabindex": {
		Name: "TabIndex",
		Type: "int",
		Doc:  "specifies the tabbing order of an element.",
	},
	"target": {
		Name: "Target",
		Type: "string",
		Doc:  "specifies the target for where to open the linked document or where to submit the form.",
	},
	"title": {
		Name: "Title",
		Type: "string",
		Doc:  "specifies extra information about an element.",
	},
	"type": {
		Name: "Type",
		Type: "string",
		Doc:  "specifies the type of element.",
	},

	// U:
	"usemap": {
		Name: "UseMap",
		Type: "string",
		Doc:  "specifies an image as a client-side image-map.",
	},

	// V:
	"value": {
		Name: "Value",
		Type: "any",
		Doc:  "specifies the value of the element.",
	},

	// W:
	"width": {
		Name: "Width",
		Type: "int",
		Doc:  "specifies the width of the element.",
	},
	"wrap": {
		Name: "Wrap",
		Type: "string",
		Doc:  "specifies how the text in a text area is to be wrapped when submitted in a form.",
	},
	"xmlns": {
		Name: "XMLNS",
		Type: "xmlns",
		Doc:  "specifies the xml namespace of the element.",
	},
}

var svgattrs = map[string]attr{
	"filterRes": {
		Name: "Filterres",
		Type: "attr|value",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"pointsAtY": {
		Name: "Pointsaty",
		Type: "string",
		Doc:  "The pointsAtY attribute represents the y location in the coordinate system established by attribute primitiveUnits on the <filter> element of the point at which the light source is pointing.",
	},

	"dy": {
		Name: "Dy",
		Type: "string",
		Doc:  "The dy attribute indicates a shift along the y-axis on the position of an element or its content.",
	},

	"marker-start": {
		Name:         "MarkerStart",
		NameOverride: "marker-start",
		Type:         "string",
		Doc:          "The marker-start attribute defines the arrowhead or polymarker that will be drawn at the first vertex of the given shape.",
	},

	"font-stretch": {
		Name:         "FontStretch",
		NameOverride: "font-stretch",
		Type:         "attr|value",
		Doc:          "The font-stretch attribute indicates the desired amount of condensing or expansion in the glyphs used to render the text.",
	},

	"contentScriptType": {
		Name: "Contentscripttype",
		Type: "attr|value",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"d": {
		Name: "D",
		Type: "string",
		Doc:  "The d attribute defines a path to be drawn.",
	},

	"k1": {
		Name: "K1",
		Type: "string",
		Doc:  "The k1 attribute defines one of the values to be used within the arithmetic operation of the <feComposite> filter primitive.",
	},

	"Styling": {
		Name: "Styling",
		Type: "attr|value",
		Doc:  "The SVG styling attributes are all the attributes that can be specified on any SVG element to apply CSS styling effects.",
	},

	"x": {
		Name: "X",
		Type: "string",
		Doc:  "The x attribute defines an x-axis coordinate in the user coordinate system.",
	},

	"cx": {
		Name: "Cx",
		Type: "string",
		Doc:  "The cx attribute define the x-axis coordinate of a center point.",
	},

	"surfaceScale": {
		Name: "Surfacescale",
		Type: "string",
		Doc:  "The surfaceScale attribute represents the height of the surface for a light filter primitive.",
	},

	"side": {
		Name: "Side",
		Type: "string",
		Doc:  "Experimental: This is an experimental technologyCheck the Browser compatibility table carefully before using this in production.",
	},

	"values": {
		Name: "Values",
		Type: "attr|value",
		Doc:  "The values attribute has different meanings, depending upon the context where it's used, either it defines a sequence of values used over the course of an animation, or it's a list of numbers for a color matrix, which is interpreted differently depending on the type of color change to be performed.",
	},

	"target": {
		Name: "Target",
		Type: "string",
		Doc:  "The target attribute should be used when there are multiple possible targets for the ending resource, such as when the parent document is embedded within an HTML or XHTML document, or is viewed with a tabbed browser. This attribute specifies the name of the browsing context (e.g., a browser tab or an (X)HTML iframe or object element) into which a document is to be opened when the link is activated:",
	},

	"clipPathUnits": {
		Name: "Clippathunits",
		Type: "string",
		Doc:  "The clipPathUnits attribute indicates which coordinate system to use for the contents of the <clipPath> element.",
	},

	"id": {
		Name: "Id",
		Type: "string",
		Doc:  "The id attribute assigns a unique name to an element.",
	},

	"x2": {
		Name: "X2",
		Type: "string",
		Doc:  "The x2 attribute is used to specify the second x-coordinate for drawing an SVG element that requires more than one coordinate. Elements that only need one coordinate use the x attribute instead.",
	},

	"specularConstant": {
		Name: "Specularconstant",
		Type: "string",
		Doc:  "The specularConstant attribute controls the ratio of reflection of the specular lighting. It represents the ks value in the Phong lighting model. The bigger the value the stronger the reflection.",
	},

	"accumulate": {
		Name: "Accumulate",
		Type: "attr|value",
		Doc:  "The accumulate attribute controls whether or not an animation is cumulative.",
	},

	"edgeMode": {
		Name: "Edgemode",
		Type: "attr|value",
		Doc:  "The edgeMode attribute determines how to extend the input image as necessary with color values so that the matrix operations can be applied when the kernel is positioned at or near the edge of the input image.",
	},

	"pointer-events": {
		Name:         "PointerEvents",
		NameOverride: "pointer-events",
		Type:         "string",
		Doc:          "The pointer-events attribute is a presentation attribute that allows defining whether or when an element may be the target of a mouse event.",
	},

	"rx": {
		Name: "Rx",
		Type: "string",
		Doc:  "The rx attribute defines a radius on the x-axis.",
	},

	"pointsAtX": {
		Name: "Pointsatx",
		Type: "string",
		Doc:  "The pointsAtX attribute represents the x location in the coordinate system established by attribute primitiveUnits on the <filter> element of the point at which the light source is pointing.",
	},

	"lengthAdjust": {
		Name: "Lengthadjust",
		Type: "string",
		Doc:  "The lengthAdjust attribute controls how the text is stretched into the length defined by the textLength attribute.",
	},

	"kernelMatrix": {
		Name: "Kernelmatrix",
		Type: "string",
		Doc:  "The kernelMatrix attribute defines the list of numbers that make up the kernel matrix for the <feConvolveMatrix> element.",
	},

	"opacity": {
		Name: "Opacity",
		Type: "string",
		Doc:  "The opacity attribute specifies the transparency of an object or of a group of objects, that is, the degree to which the background behind the element is overlaid.",
	},

	"calcMode": {
		Name: "Calcmode",
		Type: "attr|value",
		Doc:  "The calcMode attribute specifies the interpolation mode for the animation.",
	},

	"marker-mid": {
		Name:         "MarkerMid",
		NameOverride: "marker-mid",
		Type:         "string",
		Doc:          "The marker-mid attribute defines the arrowhead or polymarker that will be drawn at all interior vertices of the given shape.",
	},

	"k2": {
		Name: "K2",
		Type: "string",
		Doc:  "The k2 attribute defines one of the values to be used within the arithmetic operation of the <feComposite> filter primitive.",
	},

	"pointsAtZ": {
		Name: "Pointsatz",
		Type: "string",
		Doc:  "The pointsAtZ attribute represents the y location in the coordinate system established by attribute primitiveUnits on the <filter> element of the point at which the light source is pointing, assuming that, in the initial local coordinate system, the positive z-axis comes out towards the person viewing the content and assuming that one unit along the z-axis equals one unit in x and y.",
	},

	"result": {
		Name: "Result",
		Type: "string",
		Doc:  "The result attribute defines the assigned name for this filter primitive. If supplied, then graphics that result from processing this filter primitive can be referenced by an in attribute on a subsequent filter primitive within the same <filter> element. If no value is provided, the output will only be available for re-use as the implicit input into the next filter primitive if that filter primitive provides no value for its in attribute.",
	},

	"text-anchor": {
		Name:         "TextAnchor",
		NameOverride: "text-anchor",
		Type:         "string",
		Doc:          "The text-anchor attribute is used to align (start-, middle- or end-alignment) a string of pre-formatted text or auto-wrapped text where the wrapping area is determined from the inline-size property relative to a given point.",
	},

	"maskUnits": {
		Name: "Maskunits",
		Type: "string",
		Doc:  "The maskUnits attribute indicates which coordinate system to use for the geometry properties of the <mask> element.",
	},

	"path": {
		Name: "Path",
		Type: "string",
		Doc:  "The path attribute has two different meanings, either it defines a text path along which the characters of a text are rendered, or a motion path along which a referenced element is animated.",
	},

	"stroke-width": {
		Name:         "StrokeWidth",
		NameOverride: "stroke-width",
		Type:         "string",
		Doc:          "The stroke-width attribute is a presentation attribute defining the width of the stroke to be applied to the shape.",
	},

	"seed": {
		Name: "Seed",
		Type: "string",
		Doc:  "The seed attribute represents the starting number for the pseudo random number generator of the <feTurbulence> filter primitive.",
	},

	"operator": {
		Name: "Operator",
		Type: "string",
		Doc:  "The operator attribute has two meanings based on the context it's used in. Either it defines the compositing or morphing operation to be performed.",
	},

	"dominant-baseline": {
		Name:         "DominantBaseline",
		NameOverride: "dominant-baseline",
		Type:         "string",
		Doc:          "The dominant-baseline attribute specifies the dominant baseline, which is the baseline used to align the box's text and inline-level contents. It also indicates the default alignment baseline of any boxes participating in baseline alignment in the box's alignment context.",
	},

	"crossorigin": {
		Name: "Crossorigin",
		Type: "string",
		Doc:  "The crossorigin attribute, valid on the <image> element, provides support for CORS, defining how the element handles crossorigin requests, thereby enabling the configuration of the CORS requests for the element's fetched data. It is a CORS settings attribute.",
	},

	"cy": {
		Name: "Cy",
		Type: "string",
		Doc:  "The cy attribute define the y-axis coordinate of a center point.",
	},

	"viewBox": {
		Name: "Viewbox",
		Type: "string",
		Doc:  "The viewBox attribute defines the position and dimension, in user space, of an SVG viewport.",
	},

	"radius": {
		Name: "Radius",
		Type: "attr|value",
		Doc:  "The radius attribute represents the radius (or radii) for the operation on a given <feMorphology> filter primitive.",
	},

	"repeatCount": {
		Name: "Repeatcount",
		Type: "string",
		Doc:  "The repeatCount attribute indicates the number of times an animation will take place.",
	},

	"order": {
		Name: "Order",
		Type: "string",
		Doc:  "The order attribute indicates the size of the matrix to be used by a <feConvolveMatrix> element.",
	},

	"repeatDur": {
		Name: "Repeatdur",
		Type: "string",
		Doc:  "The repeatDur attribute specifies the total duration for repeating an animation.",
	},

	"elevation": {
		Name: "Elevation",
		Type: "string",
		Doc:  "The elevation attribute specifies the direction angle for the light source from the XY plane towards the Z-axis, in degrees. Note that the positive Z-axis points towards the viewer of the content.",
	},

	"flood-color": {
		Name:         "FloodColor",
		NameOverride: "flood-color",
		Type:         "string",
		Doc:          "The flood-color attribute indicates what color to use to flood the current filter primitive subregion.",
	},

	"xChannelSelector": {
		Name: "Xchannelselector",
		Type: "string",
		Doc:  "The xChannelSelector attribute indicates which color channel from in2 to use to displace the pixels in in along the x-axis.",
	},

	"transform": {
		Name: "Transform",
		Type: "string",
		Doc:  "The transform attribute defines a list of transform definitions that are applied to an element and the element's children.",
	},

	"keyTimes": {
		Name: "Keytimes",
		Type: "string",
		Doc:  "The keyTimes attribute represents a list of time values used to control the pacing of the animation.",
	},

	"href": {
		Name: "Href",
		Type: "string",
		Doc:  "The href attribute defines a link to a resource as a reference URL. The exact meaning of that link depends on the context of each element using it.",
	},

	"patternTransform": {
		Name: "Patterntransform",
		Type: "string",
		Doc:  "The patternTransform attribute defines a list of transform definitions that are applied to a pattern tile.",
	},

	"scale": {
		Name: "Scale",
		Type: "string",
		Doc:  "The scale attribute defines the displacement scale factor to be used on a <feDisplacementMap> filter primitive. The amount is expressed in the coordinate system established by the primitiveUnits attribute on the <filter> element.",
	},

	"keyPoints": {
		Name: "Keypoints",
		Type: "string",
		Doc:  "The keyPoints attribute indicates the simple duration of an animation.",
	},

	"kernelUnitLength": {
		Name: "Kernelunitlength",
		Type: "attr|value",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"baseFrequency": {
		Name: "Basefrequency",
		Type: "string",
		Doc:  "The baseFrequency attribute represents the base frequency parameter for the noise function of the <feTurbulence> filter primitive.",
	},

	"in#workaround_for_backgroundimage": {
		Name: "In#Workaround_for_backgroundimage",
		Type: "attr|value",
		Doc:  "The in attribute identifies input for the given filter primitive.",
	},

	"style": {
		Name: "Style",
		Type: "string",
		Doc:  "The style attribute allows to style an element using CSS declarations. It functions identically to the style attribute in HTML.",
	},

	"color": {
		Name: "Color",
		Type: "string",
		Doc:  "The color attribute is used to provide a potential indirect value, currentcolor, for the fill, stroke, stop-color, flood-color, and lighting-color attributes.",
	},

	"in": {
		Name: "In",
		Type: "string",
		Doc:  "The in attribute identifies input for the given filter primitive.",
	},

	"preserveAspectRatio": {
		Name: "Preserveaspectratio",
		Type: "string",
		Doc:  "The preserveAspectRatio attribute indicates how an element with a viewBox providing a given aspect ratio must fit into a viewport with a different aspect ratio.",
	},

	"Core": {
		Name: "Core",
		Type: "attr|value",
		Doc:  "The SVG core attributes are all the common attributes that can be specified on any SVG element.",
	},

	"maskContentUnits": {
		Name: "Maskcontentunits",
		Type: "string",
		Doc:  "The maskContentUnits attribute indicates which coordinate system to use for the contents of the <mask> element.",
	},

	"width": {
		Name: "Width",
		Type: "string",
		Doc:  "The width attribute defines the horizontal length of an element in the user coordinate system.",
	},

	"patternContentUnits": {
		Name: "Patterncontentunits",
		Type: "string",
		Doc:  "The patternContentUnits attribute indicates which coordinate system to use for the contents of the <pattern> element.",
	},

	"method": {
		Name: "Method",
		Type: "attr|value",
		Doc:  "Experimental: This is an experimental technologyCheck the Browser compatibility table carefully before using this in production.",
	},

	"spacing": {
		Name: "Spacing",
		Type: "attr|value",
		Doc:  "The spacing attribute indicates how the user agent should determine the spacing between typographic characters that are to be rendered along a path.",
	},

	"primitiveUnits": {
		Name: "Primitiveunits",
		Type: "attr|value",
		Doc:  "The primitiveUnits attribute specifies the coordinate system for the various length values within the filter primitives and for the attributes that define the filter primitive subregion.",
	},

	"fill-rule": {
		Name:         "FillRule",
		NameOverride: "fill-rule",
		Type:         "string",
		Doc:          "The fill-rule attribute is a presentation attribute defining the algorithm to use to determine the inside part of a shape.",
	},

	"specularExponent": {
		Name: "Specularexponent",
		Type: "string",
		Doc:  "The specularExponent attribute controls the focus for the light source. The bigger the value the brighter the light.",
	},

	"markerUnits": {
		Name: "Markerunits",
		Type: "attr|value",
		Doc:  "The markerUnits attribute defines the coordinate system for the markerWidth and markerHeight attributes and the contents of the <marker>.",
	},

	"refY": {
		Name: "Refy",
		Type: "attr|value",
		Doc:  "The refY attribute defines the y coordinate of an element's reference point.",
	},

	"version": {
		Name: "Version",
		Type: "string",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"font-size-adjust": {
		Name:         "FontSizeAdjust",
		NameOverride: "font-size-adjust",
		Type:         "string",
		Doc:          "The font-size-adjust attribute allows authors to specify an aspect value for an element that will preserve the x-height of the first choice font in a substitute font.",
	},

	"systemLanguage": {
		Name: "Systemlanguage",
		Type: "string",
		Doc:  "The systemLanguage attribute represents a list of supported language tags. This list is matched against the language defined in the user preferences.",
	},

	"stroke-linejoin": {
		Name:         "StrokeLinejoin",
		NameOverride: "stroke-linejoin",
		Type:         "string",
		Doc:          "The stroke-linejoin attribute is a presentation attribute defining the shape to be used at the corners of paths when they are stroked.",
	},

	"cursor": {
		Name: "Cursor",
		Type: "attr|value",
		Doc:  "SVG Attribute reference home",
	},

	"viewTarget": {
		Name: "Viewtarget",
		Type: "attr|value",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"r": {
		Name: "R",
		Type: "string",
		Doc:  "The r attribute defines the radius of a circle.",
	},

	"stroke-dasharray": {
		Name:         "StrokeDasharray",
		NameOverride: "stroke-dasharray",
		Type:         "string",
		Doc:          "The stroke-dasharray attribute is a presentation attribute defining the pattern of dashes and gaps used to paint the outline of the shape;",
	},

	"Events#global_event_attributes": {
		Name: "Events#Global_event_attributes",
		Type: "attr|value",
		Doc:  "Event attributes always have their name starting with \"on\" followed by the name of the event for which they are intended. They specifies some script to run when the event of the given type is dispatched to the element on which the attributes are specified.",
	},

	"To": {
		Name: "To",
		Type: "attr|value",
		Doc:  "The to attribute indicates the final value of the attribute that will be modified during the animation.",
	},

	"markerHeight": {
		Name: "Markerheight",
		Type: "attr|value",
		Doc:  "The markerHeight attribute represents the height of the viewport into which the <marker> is to be fitted when it is rendered according to the viewBox and preserveAspectRatio attributes.",
	},

	"stdDeviation": {
		Name: "Stddeviation",
		Type: "string",
		Doc:  "The stdDeviation attribute defines the standard deviation for the blur operation.",
	},

	"rotate": {
		Name: "Rotate",
		Type: "string",
		Doc:  "The rotate attribute specifies how the animated element rotates as it travels along a path specified in an <animateMotion> element.",
	},

	"Events#animation_event_attributes": {
		Name: "Events#Animation_event_attributes",
		Type: "attr|value",
		Doc:  "Event attributes always have their name starting with \"on\" followed by the name of the event for which they are intended. They specifies some script to run when the event of the given type is dispatched to the element on which the attributes are specified.",
	},

	"font-variant": {
		Name:         "FontVariant",
		NameOverride: "font-variant",
		Type:         "string",
		Doc:          "The font-variant attribute indicates whether the text is to be rendered using variations of the font's glyphs.",
	},

	"zoomAndPan": {
		Name: "Zoomandpan",
		Type: "string",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"bias": {
		Name: "Bias",
		Type: "attr|value",
		Doc:  "The bias attribute shifts the range of the filter. After applying the kernelMatrix of the <feConvolveMatrix> element to the input image to yield a number and applied the divisor attribute, the bias attribute is added to each component. This allows representation of values that would otherwise be clamped to 0 or 1.",
	},

	"contentStyleType": {
		Name: "Contentstyletype",
		Type: "attr|value",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"dx": {
		Name: "Dx",
		Type: "string",
		Doc:  "The dx attribute indicates a shift along the x-axis on the position of an element or its content.",
	},

	"min": {
		Name: "Min",
		Type: "string",
		Doc:  "The min attribute specifies the minimum value of the active animation duration.",
	},

	"orient": {
		Name: "Orient",
		Type: "string",
		Doc:  "The orient attribute indicates how a marker is rotated when it is placed at its position on the shape.",
	},

	"mode": {
		Name: "Mode",
		Type: "string",
		Doc:  "The mode attribute defines the blending mode on the <feBlend> filter primitive.",
	},

	"k3": {
		Name: "K3",
		Type: "string",
		Doc:  "The k3 attribute defines one of the values to be used within the arithmetic operation of the <feComposite> filter primitive.",
	},

	"preserveAlpha": {
		Name: "Preservealpha",
		Type: "string",
		Doc:  "the preserveAlpha attribute indicates how a <feConvolveMatrix> element handles alpha transparency.",
	},

	"patternUnits": {
		Name: "Patternunits",
		Type: "string",
		Doc:  "The patternUnits attribute indicates which coordinate system to use for the geometry properties of the <pattern> element.",
	},

	"font-style": {
		Name:         "FontStyle",
		NameOverride: "font-style",
		Type:         "string",
		Doc:          "The font-style attribute specifies whether the text is to be rendered using a normal, italic, or oblique face.",
	},

	"stop-color": {
		Name:         "StopColor",
		NameOverride: "stop-color",
		Type:         "attr|value",
		Doc:          "The stop-color attribute indicates what color to use at a gradient stop.",
	},

	"z": {
		Name: "Z",
		Type: "string",
		Doc:  "The z attribute defines the location along the z-axis for a light source in the coordinate system established by the primitiveUnits attribute on the <filter> element, assuming that, in the initial coordinate system, the positive z-axis comes out towards the person viewing the content and assuming that one unit along the z-axis equals one unit in x and y.",
	},

	"baseProfile": {
		Name: "Baseprofile",
		Type: "string",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"points": {
		Name: "Points",
		Type: "string",
		Doc:  "The points attribute defines a list of points. Each point is defined by a pair of number representing a X and a Y coordinate in the user coordinate system. If the attribute contains an odd number of coordinates, the last one will be ignored.",
	},

	"attributeName": {
		Name: "Attributename",
		Type: "string",
		Doc:  "The attributeName attribute indicates the name of the CSS property or attribute of the target element that is going to be changed during an animation.",
	},

	"Events#document_event_attributes": {
		Name: "Events#Document_event_attributes",
		Type: "attr|value",
		Doc:  "Event attributes always have their name starting with \"on\" followed by the name of the event for which they are intended. They specifies some script to run when the event of the given type is dispatched to the element on which the attributes are specified.",
	},

	"begin": {
		Name: "Begin",
		Type: "string",
		Doc:  "The begin attribute defines when an animation should begin or when an element should be discarded.",
	},

	"targetX": {
		Name: "Targetx",
		Type: "attr|value",
		Doc:  "The targetX attribute determines the positioning in horizontal direction of the convolution matrix relative to a given target pixel in the input image. The leftmost column of the matrix is column number zero. The value must be such that: 0 <= targetX < order X.",
	},

	"Events#document_element_event_attributes": {
		Name: "Events#Document_element_event_attributes",
		Type: "attr|value",
		Doc:  "Event attributes always have their name starting with \"on\" followed by the name of the event for which they are intended. They specifies some script to run when the event of the given type is dispatched to the element on which the attributes are specified.",
	},

	"x1": {
		Name: "X1",
		Type: "string",
		Doc:  "The x1 attribute is used to specify the first x-coordinate for drawing an SVG element that requires more than one coordinate. Elements that only need one coordinate use the x attribute instead.",
	},

	"Events#graphical_event_attributes": {
		Name: "Events#Graphical_event_attributes",
		Type: "attr|value",
		Doc:  "Event attributes always have their name starting with \"on\" followed by the name of the event for which they are intended. They specifies some script to run when the event of the given type is dispatched to the element on which the attributes are specified.",
	},

	"color-interpolation": {
		Name:         "ColorInterpolation",
		NameOverride: "color-interpolation",
		Type:         "attr|value",
		Doc:          "The color-interpolation attribute specifies the color space for gradient interpolations, color animations, and alpha compositing.",
	},

	"refX": {
		Name: "Refx",
		Type: "attr|value",
		Doc:  "The refX attribute defines the x coordinate of an element's reference point.",
	},

	"k4": {
		Name: "K4",
		Type: "string",
		Doc:  "The k4 attribute defines one of the values to be used within the arithmetic operation of the <feComposite> filter primitive.",
	},

	"Presentation": {
		Name: "Presentation",
		Type: "attr|value",
		Doc:  "SVG presentation attributes are CSS properties that can be used as attributes on SVG elements.",
	},

	"stroke-opacity": {
		Name:         "StrokeOpacity",
		NameOverride: "stroke-opacity",
		Type:         "string",
		Doc:          "The stroke-opacity attribute is a presentation attribute defining the opacity of the paint server (color, gradient, pattern, etc.) applied to the stroke of a shape.",
	},

	"limitingConeAngle": {
		Name: "Limitingconeangle",
		Type: "string",
		Doc:  "The limitingConeAngle attribute represents the angle in degrees between the spot light axis (i.e. the axis between the light source and the point to which it is pointing at) and the spot light cone. So it defines a limiting cone which restricts the region where the light is projected. No light is projected outside the cone.",
	},

	"additive": {
		Name: "Additive",
		Type: "attr|value",
		Doc:  "The additive attribute controls whether or not an animation is additive.",
	},

	"stop-opacity": {
		Name:         "StopOpacity",
		NameOverride: "stop-opacity",
		Type:         "attr|value",
		Doc:          "The stop-opacity attribute defines the opacity of a given color gradient stop.",
	},

	"requiredFeatures": {
		Name: "Requiredfeatures",
		Type: "string",
		Doc:  "Deprecated: This feature is no longer recommended. Though some browsers might still support it, it may have already been removed from the relevant web standards, may be in the process of being dropped, or may only be kept for compatibility purposes. Avoid using it, and update existing code if possible; see the compatibility table at the bottom of this page to guide your decision. Be aware that this feature may cease to work at any time.",
	},

	"clip-path": {
		Name:         "ClipPath",
		NameOverride: "clip-path",
		Type:         "string",
		Doc:          "The clip-path presentation attribute defines or associates a clipping path with the element it is related to.",
	},

	"y2": {
		Name: "Y2",
		Type: "string",
		Doc:  "The y2 attribute is used to specify the second y-coordinate for drawing an SVG element that requires more than one coordinate. Elements that only need one coordinate use the y attribute instead.",
	},

	"flood-opacity": {
		Name:         "FloodOpacity",
		NameOverride: "flood-opacity",
		Type:         "string",
		Doc:          "The flood-opacity attribute indicates the opacity value to use across the current filter primitive subregion.",
	},

	"type": {
		Name: "Type",
		Type: "attr|value",
		Doc:  "The type attribute is a generic attribute and it has different meaning based on the context in which it's used.",
	},

	"keySplines": {
		Name: "Keysplines",
		Type: "string",
		Doc:  "The keySplines attribute defines a set of Bézier curve control points associated with the keyTimes list, defining a cubic Bézier function that controls interval pacing.",
	},

	"class": {
		Name: "Class",
		Type: "string",
		Doc:  "« SVG Attribute reference home",
	},

	"stroke-dashoffset": {
		Name:         "StrokeDashoffset",
		NameOverride: "stroke-dashoffset",
		Type:         "string",
		Doc:          "The stroke-dashoffset attribute is a presentation attribute defining an offset on the rendering of the associated dash array.",
	},

	"y1": {
		Name: "Y1",
		Type: "string",
		Doc:  "The y1 attribute is used to specify the first y-coordinate for drawing an SVG element that requires more than one coordinate. Elements that only need one coordinate use the y attribute instead.",
	},

	"pathLength": {
		Name: "Pathlength",
		Type: "string",
		Doc:  "The pathLength attribute lets authors specify a total length for the path, in user units. This value is then used to calibrate the browser's distance calculations with those of the author, by scaling all distance computations using the ratio pathLength / (computed value of path length).",
	},

	"diffuseConstant": {
		Name: "Diffuseconstant",
		Type: "string",
		Doc:  "The diffuseConstant attribute represents the kd value in the Phong lighting model. In SVG, this can be any non-negative number.",
	},

	"stroke-miterlimit": {
		Name:         "StrokeMiterlimit",
		NameOverride: "stroke-miterlimit",
		Type:         "string",
		Doc:          "The stroke-miterlimit attribute is a presentation attribute defining a limit on the ratio of the miter length to the stroke-width used to draw a miter join. When the limit is exceeded, the join is converted from a miter to a bevel.",
	},

	"tabindex": {
		Name: "Tabindex",
		Type: "string",
		Doc:  "The tabindex attribute allows you to control whether an element is focusable and to define the relative order of the element for the purposes of sequential focus navigation.",
	},

	"end": {
		Name: "End",
		Type: "string",
		Doc:  "The end attribute defines an end value for the animation that can constrain the active duration.",
	},

	"startOffset": {
		Name: "Startoffset",
		Type: "string",
		Doc:  "The startOffset attribute defines an offset from the start of the path for the initial current text position along the path after converting the path to the <textPath> element's coordinate system.",
	},

	"font-family": {
		Name:         "FontFamily",
		NameOverride: "font-family",
		Type:         "string",
		Doc:          "The font-family attribute indicates which font family will be used to render the text, specified as a prioritized list of font family names and/or generic family names.",
	},

	"targetY": {
		Name: "Targety",
		Type: "attr|value",
		Doc:  "The targetY attribute determines the positioning in vertical direction of the convolution matrix relative to a given target pixel in the input image. The topmost row of the matrix is row number zero. The value must be such that: 0 <= targetY < order Y.",
	},

	"Conditional_Processing": {
		Name: "Conditional_processing",
		Type: "attr|value",
		Doc:  "The SVG conditional processing attributes are all the attributes that can be specified on some SVG elements to control whether or not the element on which it appears should be rendered.",
	},

	"display": {
		Name: "Display",
		Type: "string",
		Doc:  "The display attribute lets you control the rendering of graphical or container elements.",
	},

	"shape-rendering": {
		Name:         "ShapeRendering",
		NameOverride: "shape-rendering",
		Type:         "string",
		Doc:          "The shape-rendering attribute provides hints to the renderer about what tradeoffs to make when rendering shapes like paths, circles, or rectangles.",
	},

	"lang": {
		Name: "Lang",
		Type: "string",
		Doc:  "The lang attribute specifies the primary language used in contents and attributes containing text content of particular elements.",
	},

	"by": {
		Name: "By",
		Type: "string",
		Doc:  "The by attribute specifies a relative offset value for an attribute that will be modified during an animation.",
	},

	"filterUnits": {
		Name: "Filterunits",
		Type: "attr|value",
		Doc:  "The filterUnits attribute defines the coordinate system for the attributes x, y, width and height.",
	},

	"divisor": {
		Name: "Divisor",
		Type: "string",
		Doc:  "The divisor attribute specifies the value by which the resulting number of applying the kernelMatrix of a <feConvolveMatrix> element to the input image color value is divided to yield the destination color value.",
	},

	"max": {
		Name: "Max",
		Type: "string",
		Doc:  "The max attribute specifies the maximum value of the active animation duration.",
	},

	"yChannelSelector": {
		Name: "Ychannelselector",
		Type: "string",
		Doc:  "The yChannelSelector attribute indicates which color channel from in2 to use to displace the pixels in in along the y-axis.",
	},

	"stroke-linecap": {
		Name:         "StrokeLinecap",
		NameOverride: "stroke-linecap",
		Type:         "string",
		Doc:          "The stroke-linecap attribute is a presentation attribute defining the shape to be used at the end of open subpaths when they are stroked.",
	},

	"vector-effect": {
		Name:         "VectorEffect",
		NameOverride: "vector-effect",
		Type:         "string",
		Doc:          "The vector-effect property specifies the vector effect to use when drawing an object. Vector effects are applied before any of the other compositing operations, i.e. filters, masks and clips.",
	},

	"height": {
		Name: "Height",
		Type: "string",
		Doc:  "The height attribute defines the vertical length of an element in the user coordinate system.",
	},

	"stitchTiles": {
		Name: "Stitchtiles",
		Type: "string",
		Doc:  "The stitchTiles attribute defines how the Perlin Noise tiles behave at the border.",
	},

	"media": {
		Name: "Media",
		Type: "string",
		Doc:  "The media attribute specifies a media query that must be matched for a style sheet to apply.",
	},

	"clip-rule": {
		Name:         "ClipRule",
		NameOverride: "clip-rule",
		Type:         "string",
		Doc:          "« SVG Attribute reference home",
	},

	"ry": {
		Name: "Ry",
		Type: "string",
		Doc:  "The ry attribute defines a radius on the y-axis.",
	},

	"textLength": {
		Name: "Textlength",
		Type: "string",
		Doc:  "The textLength attribute, available on SVG <text> and <tspan> elements, lets you specify the width of the space into which the text will draw. The user agent will ensure that the text does not extend farther than that distance, using the method or methods specified by the lengthAdjust attribute. By default, only the spacing between characters is adjusted, but the glyph size can also be adjusted if you change lengthAdjust.",
	},

	"dur": {
		Name: "Dur",
		Type: "string",
		Doc:  "The dur attribute indicates the simple duration of an animation.",
	},

	"y": {
		Name: "Y",
		Type: "string",
		Doc:  "The y attribute defines a y-axis coordinate in the user coordinate system.",
	},

	"numOctaves": {
		Name: "Numoctaves",
		Type: "string",
		Doc:  "The numOctaves attribute defines the number of octaves for the noise function of the <feTurbulence> primitive.",
	},

	"restart": {
		Name: "Restart",
		Type: "string",
		Doc:  "The restart attribute specifies whether or not an animation can restart.",
	},

	"fill": {
		Name: "Fill",
		Type: "string",
		Doc:  "The fill attribute has two different meanings. For shapes and text it's a presentation attribute that defines the color (or any SVG paint servers like gradients or patterns) used to paint the element; for animation it defines the final state of the animation.",
	},

	"fill-opacity": {
		Name:         "FillOpacity",
		NameOverride: "fill-opacity",
		Type:         "string",
		Doc:          "The fill-opacity attribute is a presentation attribute defining the opacity of the paint server (color, gradient, pattern, etc.) applied to a shape.",
	},

	"filter": {
		Name: "Filter",
		Type: "string",
		Doc:  "The filter attribute specifies the filter effects defined by the <filter> element that shall be applied to its element.",
	},

	"stroke": {
		Name: "Stroke",
		Type: "string",
		Doc:  "The stroke attribute is a presentation attribute defining the color (or any SVG paint servers like gradients or patterns) used to paint the outline of the shape;",
	},

	"marker-end": {
		Name:         "MarkerEnd",
		NameOverride: "marker-end",
		Type:         "string",
		Doc:          "The marker-end attribute defines the arrowhead or polymarker that will be drawn at the final vertex of the given shape.",
	},

	"markerWidth": {
		Name: "Markerwidth",
		Type: "attr|value",
		Doc:  "The markerWidth attribute represents the width of the viewport into which the <marker> is to be fitted when it is rendered according to the viewBox and preserveAspectRatio attributes.",
	},

	"azimuth": {
		Name: "Azimuth",
		Type: "string",
		Doc:  "The azimuth attribute specifies the direction angle for the light source on the XY plane (clockwise), in degrees from the x axis.",
	},

	"font-size": {
		Name:         "FontSize",
		NameOverride: "font-size",
		Type:         "string",
		Doc:          "The font-size attribute refers to the size of the font from baseline to baseline when multiple lines of text are set solid in a multiline layout environment.",
	},

	"mask": {
		Name: "Mask",
		Type: "string",
		Doc:  "The mask attribute is a presentation attribute mainly used to bind a given <mask> element with the element the attribute belongs to.",
	},

	"visibility": {
		Name: "Visibility",
		Type: "string",
		Doc:  "The visibility attribute lets you control the visibility of graphical elements. With a value of hidden or collapse the current graphics element is invisible.",
	},

	"From": {
		Name: "From",
		Type: "attr|value",
		Doc:  "The from attribute indicates the initial value of the attribute that will be modified during the animation.",
	},

	"font-weight": {
		Name:         "FontWeight",
		NameOverride: "font-weight",
		Type:         "string",
		Doc:          "The font-weight attribute refers to the boldness or lightness of the glyphs used to render the text, relative to other fonts in the same font family.",
	},

	"in2": {
		Name: "In2",
		Type: "attr|value",
		Doc:  "The in2 attribute identifies the second input for the given filter primitive. It works exactly like the in attribute.",
	},

	"color-rendering": {
		Name:         "ColorRendering",
		NameOverride: "color-rendering",
		Type:         "attr|value",
		Doc:          "",
	},

	"offset": {
		Name: "Offset",
		Type: "attr|value",
		Doc:  "",
	},

	"title": {
		Name: "Title",
		Type: "attr|value",
		Doc:  "",
	},

	"requiredExtensions": {
		Name: "Requiredextensions",
		Type: "attr|value",
		Doc:  "",
	},
}

func attrsSVGByNames(names ...string) []attr {
	res := make([]attr, 0, len(names))
	for _, n := range names {
		attr, ok := svgattrs[n]
		if !ok {
			panic("unkowmn attr: " + n)
		}
		res = append(res, attr)
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.Compare(res[i].Name, res[j].Name) <= 0
	})

	return res
}

func attrsByNames(names ...string) []attr {
	res := make([]attr, 0, len(names))
	for _, n := range names {
		attr, ok := attrs[n]
		if !ok {
			panic("unkowmn attr: " + n)
		}
		res = append(res, attr)
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.Compare(res[i].Name, res[j].Name) <= 0
	})

	return res
}

func withGlobalAttrs(attrs ...attr) []attr {
	attrs = append(attrs, attrsByNames(
		"accesskey",
		"aria-*",
		"class",
		"contenteditable",
		"data-*",
		"dir",
		"draggable",
		"hidden",
		"id",
		"lang",
		"role",
		"spellcheck",
		"style",
		"styles",
		"tabindex",
		"title",
		"attribute",
	)...)

	sort.Slice(attrs, func(i, j int) bool {
		return strings.Compare(attrs[i].Name, attrs[j].Name) <= 0
	})

	return attrs
}

type eventHandler struct {
	Name string
	Doc  string
}

var svgEventHandlers = map[string]eventHandler{"onbegin": {
	Name: "OnBegin",
},
	"onend": {
		Name: "OnEnd",
	},
	"onrepeat": {
		Name: "OnRepeat",
	},
	"oncopy": {
		Name: "OnCopy",
	},
	"oncut": {
		Name: "OnCut",
	},
	"onpaste": {
		Name: "OnPaste",
	},
	"onabort": {
		Name: "OnAbort",
	},
	" onerror": {
		Name: "OoNerror",
	},
	" onresize": {
		Name: "OoNresize",
	},
	" onscroll": {
		Name: "OoNscroll",
	},
	" onunload": {
		Name: "OoNunload",
	},
	"onactivate": {
		Name: "OnActivate",
	},
	"onfocusin": {
		Name: "OnFocusin",
	},
	"onfocusout": {
		Name: "OnFocusout",
	},
	"oncancel": {
		Name: "OnCancel",
	},
	"oncanplay": {
		Name: "OnCanplay",
	},
	"oncanplaythrough": {
		Name: "OnCanplaythrough",
	},
	"onchange": {
		Name: "OnChange",
	},
	"onclick": {
		Name: "OnClick",
	},
	"onclose": {
		Name: "OnClose",
	},
	"oncuechange": {
		Name: "OnCuechange",
	},
	"ondblclick": {
		Name: "OnDblclick",
	},
	"ondrag": {
		Name: "OnDrag",
	},
	"ondragend": {
		Name: "OnDragend",
	},
	"ondragenter": {
		Name: "OnDragenter",
	},
	"ondragleave": {
		Name: "OnDragleave",
	},
	"ondragover": {
		Name: "OnDragover",
	},
	"ondragstart": {
		Name: "OnDragstart",
	},
	"ondrop": {
		Name: "OnDrop",
	},
	"ondurationchange": {
		Name: "OnDurationchange",
	},
	"onemptied": {
		Name: "OnEmptied",
	},
	"onended": {
		Name: "OnEnded",
	},
	"onerror": {
		Name: "OnError",
	},
	"onfocus": {
		Name: "OnFocus",
	},
	"oninput": {
		Name: "OnInput",
	},
	"oninvalid": {
		Name: "OnInvalid",
	},
	"onkeydown": {
		Name: "OnKeyDown",
	},
	"onkeypress": {
		Name: "OnKeyPress",
	},
	"onkeyup": {
		Name: "OnKeyUp",
	},
	"onload": {
		Name: "OnLoad",
	},
	"onloadeddata": {
		Name: "OnLoadeddata",
	},
	"onloadedmetadata": {
		Name: "OnLoadedmetadata",
	},
	"onloadstart": {
		Name: "OnLoadstart",
	},
	"onmousedown": {
		Name: "OnMousedown",
	},
	"onmouseenter": {
		Name: "OnMouseenter",
	},
	"onmouseleave": {
		Name: "OnMouseleave",
	},
	"onmousemove": {
		Name: "OnMousemove",
	},
	"onmouseout": {
		Name: "OnMouseout",
	},
	"onmouseover": {
		Name: "OnMouseover",
	},
	"onmouseup": {
		Name: "OnMouseup",
	},
	"onmousewheel": {
		Name: "OnMousewheel",
	},
	"onpause": {
		Name: "OnPause",
	},
	"onplay": {
		Name: "OnPlay",
	},
	"onplaying": {
		Name: "OnPlaying",
	},
	"onprogress": {
		Name: "OnProgress",
	},
	"onratechange": {
		Name: "OnRatechange",
	},
	"onreset": {
		Name: "OnReset",
	},
	"onresize": {
		Name: "OnResize",
	},
	"onscroll": {
		Name: "OnScroll",
	},
	"onseeked": {
		Name: "OnSeeked",
	},
	"onseeking": {
		Name: "OnSeeking",
	},
	"onselect": {
		Name: "OnSelect",
	},
	"onshow": {
		Name: "OnShow",
	},
	"onstalled": {
		Name: "OnStalled",
	},
	"onsubmit": {
		Name: "OnSubmit",
	},
	"onsuspend": {
		Name: "OnSuspend",
	},
	"ontimeupdate": {
		Name: "OnTimeupdate",
	},
	"ontoggle": {
		Name: "OnToggle",
	},
	"onvolumechange": {
		Name: "OnVolumechange",
	},
	"onwaiting": {
		Name: "OnWaiting",
	},
}

var eventHandlers = map[string]eventHandler{
	// Window events:
	"onafterprint": {
		Name: "OnAfterPrint",
		Doc:  "runs the given handler after the document is printed.",
	},
	"onbeforeprint": {
		Name: "OnBeforePrint",
		Doc:  "calls the given handler before the document is printed.",
	},
	"onbeforeunload": {
		Name: "OnBeforeUnload",
		Doc:  "calls the given handler when the document is about to be unloaded.",
	},
	"onerror": {
		Name: "OnError",
		Doc:  "calls the given handler when an error occurs.",
	},
	"onhashchange": {
		Name: "OnHashChange",
		Doc:  "calls the given handler when there has been changes to the anchor part of the a URL.",
	},
	"onload": {
		Name: "OnLoad",
		Doc:  "calls the given handler after the element is finished loading.",
	},
	"onmessage": {
		Name: "OnMessage",
		Doc:  "calls then given handler when a message is triggered.",
	},
	"onoffline": {
		Name: "OnOffline",
		Doc:  "calls the given handler when the browser starts to work offline.",
	},
	"ononline": {
		Name: "OnOnline",
		Doc:  "calls the given handler when the browser starts to work online.",
	},
	"onpagehide": {
		Name: "OnPageHide",
		Doc:  "calls the given handler when a user navigates away from a page.",
	},
	"onpageshow": {
		Name: "OnPageShow",
		Doc:  "calls the given handler when a user navigates to a page.",
	},
	"onpopstate": {
		Name: "OnPopState",
		Doc:  "calls the given handler when the window's history changes.",
	},
	"onresize": {
		Name: "OnResize",
		Doc:  "calls the given handler when the browser window is resized.",
	},
	"onstorage": {
		Name: "OnStorage",
		Doc:  "calls the given handler when a Web Storage area is updated.",
	},
	"onunload": {
		Name: "OnUnload",
		Doc:  "calls the given handler once a page has unloaded (or the browser window has been closed).",
	},

	// Form events:
	"onblur": {
		Name: "OnBlur",
		Doc:  "calls the given handler when the element loses focus.",
	},
	"onchange": {
		Name: "OnChange",
		Doc:  "calls the given handler when the value of the element is changed.",
	},
	"oncontextmenu": {
		Name: "OnContextMenu",
		Doc:  "calls the given handler when a context menu is triggered.",
	},
	"onfocus": {
		Name: "OnFocus",
		Doc:  "calls the given handler when the element gets focus.",
	},
	"oninput": {
		Name: "OnInput",
		Doc:  "calls the given handler when an element gets user input.",
	},
	"oninvalid": {
		Name: "OnInvalid",
		Doc:  "calls the given handler when an element is invalid.",
	},
	"onreset": {
		Name: "OnReset",
		Doc:  "calls the given handler when the Reset button in a form is clicked.",
	},
	"onsearch": {
		Name: "OnSearch",
		Doc:  `calls the given handler when the user writes something in a search field.`,
	},
	"onselect": {
		Name: "OnSelect",
		Doc:  "calls the given handler after some text has been selected in an element.",
	},
	"onsubmit": {
		Name: "OnSubmit",
		Doc:  "calls the given handler when a form is submitted.",
	},

	// Keyboard events:
	"onkeydown": {
		Name: "OnKeyDown",
		Doc:  "calls the given handler when a user is pressing a key.",
	},
	"onkeypress": {
		Name: "OnKeyPress",
		Doc:  "calls the given handler when a user presses a key.",
	},
	"onkeyup": {
		Name: "OnKeyUp",
		Doc:  "calls the given handler when a user releases a key.",
	},

	// Mouse events:
	"onclick": {
		Name: "OnClick",
		Doc:  "calls the given handler when there is a mouse click on the element.",
	},
	"ondblclick": {
		Name: "OnDblClick",
		Doc:  "calls the given handler when there is a mouse double-click on the element.",
	},
	"onmousedown": {
		Name: "OnMouseDown",
		Doc:  "calls the given handler when a mouse button is pressed down on an element.",
	},
	"onmouseenter": {
		Name: "OnMouseEnter",
		Doc:  "calls the given handler when a mouse button is initially moved so that its hotspot is within the element at which the event was fired.",
	},
	"onmouseleave": {
		Name: "OnMouseLeave",
		Doc:  "calls the given handler when the mouse pointer is fired when the pointer has exited the element and all of its descendants.",
	},
	"onmousemove": {
		Name: "OnMouseMove",
		Doc:  "calls the given handler when the mouse pointer is moving while it is over an element.",
	},
	"onmouseout": {
		Name: "OnMouseOut",
		Doc:  "calls the given handler when the mouse pointer moves out of an element.",
	},
	"onmouseover": {
		Name: "OnMouseOver",
		Doc:  "calls the given handler when the mouse pointer moves over an element.",
	},
	"onmouseup": {
		Name: "OnMouseUp",
		Doc:  "calls the given handler when a mouse button is released over an element.",
	},
	"onwheel": {
		Name: "OnWheel",
		Doc:  "calls the given handler when the mouse wheel rolls up or down over an element.",
	},

	// Drag events:
	"ondrag": {
		Name: "OnDrag",
		Doc:  "calls the given handler when an element is dragged.",
	},
	"ondragend": {
		Name: "OnDragEnd",
		Doc:  "calls the given handler at the end of a drag operation.",
	},
	"ondragenter": {
		Name: "OnDragEnter",
		Doc:  "calls the given handler when an element has been dragged to a valid drop target.",
	},
	"ondragleave": {
		Name: "OnDragLeave",
		Doc:  "calls the given handler when an element leaves a valid drop target.",
	},
	"ondragover": {
		Name: "OnDragOver",
		Doc:  "calls the given handler when an element is being dragged over a valid drop target.",
	},
	"ondragstart": {
		Name: "OnDragStart",
		Doc:  "calls the given handler at the start of a drag operation.",
	},
	"ondrop": {
		Name: "OnDrop",
		Doc:  "calls the given handler when dragged element is being dropped.",
	},
	"onscroll": {
		Name: "OnScroll",
		Doc:  "calls the given handler when an element's scrollbar is being scrolled.",
	},

	// Clipboard event:
	"oncopy": {
		Name: "OnCopy",
		Doc:  "calls the given handler when the user copies the content of an element.",
	},
	"oncut": {
		Name: "OnCut",
		Doc:  "calls the given handler when the user cuts the content of an element.",
	},
	"onpaste": {
		Name: "OnPaste",
		Doc:  "calls the given handler when the user pastes some content in an element.",
	},

	// Media events:
	"onabort": {
		Name: "OnAbort",
		Doc:  "calls the given handler on abort.",
	},
	"oncanplay": {
		Name: "OnCanPlay",
		Doc:  "calls the given handler when a file is ready to start playing (when it has buffered enough to begin).",
	},
	"oncanplaythrough": {
		Name: "OnCanPlayThrough",
		Doc:  "calls the given handler when a file can be played all the way to the end without pausing for buffering.",
	},
	"oncuechange": {
		Name: "OnCueChange",
		Doc:  "calls the given handler when the cue changes in a track element.",
	},
	"ondurationchange": {
		Name: "OnDurationChange",
		Doc:  "calls the given handler when the length of the media changes.",
	},
	"onemptied": {
		Name: "OnEmptied",
		Doc:  "calls the given handler when something bad happens and the file is suddenly unavailable (like unexpectedly disconnects).",
	},
	"onended": {
		Name: "OnEnded",
		Doc:  "calls the given handler when the media has reach the end.",
	},
	"onloadeddata": {
		Name: "OnLoadedData",
		Doc:  "calls the given handler when media data is loaded.",
	},
	"onloadedmetadata": {
		Name: "OnLoadedMetaData",
		Doc:  "calls the given handler when meta data (like dimensions and duration) are loaded.",
	},
	"onloadstart": {
		Name: "OnLoadStart",
		Doc:  "calls the given handler just as the file begins to load before anything is actually loaded.",
	},
	"onpause": {
		Name: "OnPause",
		Doc:  "calls the given handler when the media is paused either by the user or programmatically.",
	},
	"onplay": {
		Name: "OnPlay",
		Doc:  "calls the given handler when the media is ready to start playing.",
	},
	"onplaying": {
		Name: "OnPlaying",
		Doc:  "calls the given handler when the media actually has started playing.",
	},
	"onprogress": {
		Name: "OnProgress",
		Doc:  "calls the given handler when the browser is in the process of getting the media data.",
	},
	"onratechange": {
		Name: "OnRateChange",
		Doc:  "calls the given handler each time the playback rate changes (like when a user switches to a slow motion or fast forward mode).",
	},
	"onseeked": {
		Name: "OnSeeked",
		Doc:  "calls the given handler when the seeking attribute is set to false indicating that seeking has ended.",
	},
	"onseeking": {
		Name: "OnSeeking",
		Doc:  "calls the given handler when the seeking attribute is set to true indicating that seeking is active.",
	},
	"onstalled": {
		Name: "OnStalled",
		Doc:  "calls the given handler when the browser is unable to fetch the media data for whatever reason.",
	},
	"onsuspend": {
		Name: "OnSuspend",
		Doc:  "calls the given handler when fetching the media data is stopped before it is completely loaded for whatever reason.",
	},
	"ontimeupdate": {
		Name: "OnTimeUpdate",
		Doc:  "calls the given handler when the playing position has changed (like when the user fast forwards to a different point in the media).",
	},
	"onvolumechange": {
		Name: "OnVolumeChange",
		Doc:  `calls the given handler each time the volume is changed which (includes setting the volume to "mute").`,
	},
	"onwaiting": {
		Name: "OnWaiting",
		Doc:  "calls the given handler when the media has paused but is expected to resume (like when the media pauses to buffer more data).",
	},

	// Miscs events:
	"ontoggle": {
		Name: "OnToggle",
		Doc:  "calls the given handler when the user opens or closes the details element.",
	},
}

func svgEventHandlersByName(names ...string) []eventHandler {
	res := make([]eventHandler, 0, len(names))
	for _, n := range names {
		h, ok := svgEventHandlers[n]
		if !ok {
			panic("unknown event handler: " + n)
		}
		res = append(res, h)
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.Compare(res[i].Name, res[j].Name) <= 0
	})

	return res
}

func eventHandlersByName(names ...string) []eventHandler {
	res := make([]eventHandler, 0, len(names))
	for _, n := range names {
		h, ok := eventHandlers[n]
		if !ok {
			panic("unknown event handler: " + n)
		}
		res = append(res, h)
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.Compare(res[i].Name, res[j].Name) <= 0
	})

	return res
}

func withSVGAnimationEventHandler(handlers ...eventHandler) []eventHandler {
	begin := false
	end := false
	repeat := false
	for _, h := range handlers {
		if h.Name == "OnBegin" {
			begin = true
		}
		if h.Name == "OnEnd" {
			end = true
		}
		if h.Name == "OnRepeat" {
			repeat = true
		}
	}
	if !begin {
		handlers = append(handlers, svgEventHandlersByName(
			"onbegin",
		)...)
	}

	if !end {
		handlers = append(handlers, svgEventHandlersByName(
			"onend",
		)...)
	}
	if !repeat {
		handlers = append(handlers, svgEventHandlersByName(
			"onrepeat",
		)...)
	}

	sort.Slice(handlers, func(i, j int) bool {
		return strings.Compare(handlers[i].Name, handlers[j].Name) <= 0
	})

	return handlers
}

func withSVGDocumentElementEventHandler(handlers ...eventHandler) []eventHandler {
	handlers = append(handlers, svgEventHandlersByName(
		"oncopy", "oncut", "onpaste",
	)...)

	sort.Slice(handlers, func(i, j int) bool {
		return strings.Compare(handlers[i].Name, handlers[j].Name) <= 0
	})

	return handlers
}

func withSVGDocumentEventHandler(handlers ...eventHandler) []eventHandler {
	handlers = append(handlers, svgEventHandlersByName(
		"onabort", " onerror", " onresize", " onscroll", " onunload",
	)...)

	sort.Slice(handlers, func(i, j int) bool {
		return strings.Compare(handlers[i].Name, handlers[j].Name) <= 0
	})

	return handlers
}

func withSVGGraphicalEventHandler(handlers ...eventHandler) []eventHandler {
	handlers = append(handlers, svgEventHandlersByName(
		"onactivate", "onfocusin", "onfocusout",
	)...)

	sort.Slice(handlers, func(i, j int) bool {
		return strings.Compare(handlers[i].Name, handlers[j].Name) <= 0
	})

	return handlers
}

func withSVGGlobalEventHandler(handlers ...eventHandler) []eventHandler {
	handlers = append(handlers, svgEventHandlersByName(
		"oncancel",
		"oncanplay",
		"oncanplaythrough",
		"onchange",
		"onclick",
		"onclose",
		"oncuechange",
		"ondblclick",
		"ondrag",
		"ondragend",
		"ondragenter",
		"ondragleave",
		"ondragover",
		"ondragstart",
		"ondrop",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onfocus",
		"oninput",
		"oninvalid",
		"onkeydown",
		"onkeypress",
		"onkeyup",
		"onload",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onmousedown",
		"onmouseenter",
		"onmouseleave",
		"onmousemove",
		"onmouseout",
		"onmouseover",
		"onmouseup",
		"onmousewheel",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onreset",
		"onresize",
		"onscroll",
		"onseeked",
		"onseeking",
		"onselect",
		"onshow",
		"onstalled",
		"onsubmit",
		"onsuspend",
		"ontimeupdate",
		"ontoggle",
		"onvolumechange",
		"onwaiting",
	)...)

	sort.Slice(handlers, func(i, j int) bool {
		return strings.Compare(handlers[i].Name, handlers[j].Name) <= 0
	})

	return handlers
}

func withGlobalEventHandlers(handlers ...eventHandler) []eventHandler {
	handlers = append(handlers, eventHandlersByName(
		"onblur",
		"onchange",
		"oncontextmenu",
		"onfocus",
		"oninput",
		"oninvalid",
		"onreset",
		"onsearch",
		"onselect",
		"onsubmit",

		"onkeydown",
		"onkeypress",
		"onkeyup",

		"onclick",
		"ondblclick",
		"onmousedown",
		"onmouseenter",
		"onmouseleave",
		"onmousemove",
		"onmouseout",
		"onmouseover",
		"onmouseup",
		"onwheel",

		"ondrag",
		"ondragend",
		"ondragenter",
		"ondragleave",
		"ondragover",
		"ondragstart",
		"ondrop",
		"onscroll",

		"oncopy",
		"oncut",
		"onpaste",
	)...)

	sort.Slice(handlers, func(i, j int) bool {
		return strings.Compare(handlers[i].Name, handlers[j].Name) <= 0
	})

	return handlers
}

func withMediaEventHandlers(handlers ...eventHandler) []eventHandler {
	handlers = append(handlers, eventHandlersByName(
		"onabort",
		"oncanplay",
		"oncanplaythrough",
		"oncuechange",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onseeked",
		"onseeking",
		"onstalled",
		"onsuspend",
		"ontimeupdate",
		"onvolumechange",
		"onwaiting",
	)...)

	sort.Slice(handlers, func(i, j int) bool {
		return strings.Compare(handlers[i].Name, handlers[j].Name) <= 0
	})

	return handlers
}

func main() {
	generateHTMLGo()
	generateHTMLTestGo()
}

func generateHTMLGo() {
	f, err := os.Create("html_gen.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package app")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// Code generated by go generate; DO NOT EDIT.")
	fmt.Fprintln(f, `
import (
	"fmt"
	"strings"
)
		`)

	for _, t := range tags {
		writeInterface(f, t)

		switch t.Name {
		case "Elem", "ElemSelfClosing":
			fmt.Fprintf(f, `
			/* %s returns an HTML element that %s */
			func %s(tag string) HTML%s {
				e := &html%s{
					htmlElement: htmlElement{
						tag: tag,
						isSelfClosing: %v,
					},
				}

				return e
			}
			`,
				t.Name,
				t.Doc,
				t.Name,
				t.Name,
				t.Name,
				t.Type == selfClosing,
			)

		default:
			fmt.Fprintf(f, `
			/* %s returns an HTML element that %s */
			func %s() HTML%s {
				e := &html%s{
					htmlElement: htmlElement{
						tag: "%s",
						isSelfClosing: %v,
					},
				}

				return e
			}
			`,
				t.Name,
				t.Doc,
				t.Name,
				t.Name,
				t.Name,
				strings.ToLower(t.Name),
				t.Type == selfClosing,
			)
		}

		fmt.Fprintln(f)
		fmt.Fprintln(f)
		writeStruct(f, t)
		fmt.Fprintln(f)
		fmt.Fprintln(f)
	}

}

func writeInterface(w io.Writer, t tag) {
	fmt.Fprintf(w, `
		// HTML%s is the interface that describes a "%s" HTML element.
		type HTML%s interface {
			UI
		`,
		t.Name,
		strings.ToLower(t.Name),
		t.Name,
	)

	switch t.Type {
	case parent:
		fmt.Fprintf(w, `
			// Body set the content of the element.
			Body(elems ...UI) HTML%s 
		`, t.Name)

		fmt.Fprintf(w, `
			// Text sets the content of the element with a text node containing the stringified given value.
			Text(v any) HTML%s
		`, t.Name)

	case privateParent:
		fmt.Fprintf(w, `
			privateBody(elems ...UI) HTML%s 
		`, t.Name)
	}

	for _, a := range t.Attrs {
		fmt.Fprintln(w)
		fmt.Fprintln(w)

		fmt.Fprintf(w, "/* %s %s\n*/\n", a.Name, a.Doc)
		writeAttrFunction(w, a, t, true)
	}

	fmt.Fprintln(w)

	fmt.Fprintf(w, `
		// On registers the given event handler to the specified event.
		On(event string, h EventHandler, scope ...any) HTML%s 
	`, t.Name)

	for _, e := range t.EventHandlers {
		fmt.Fprintln(w)
		fmt.Fprintln(w)

		fmt.Fprintf(w, "/* %s %s\n*/\n", e.Name, e.Doc)
		writeEventFunction(w, e, t, true)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "}")
}

func writeStruct(w io.Writer, t tag) {
	fmt.Fprintf(w, `type html%s struct {
			htmlElement
		}`, t.Name)

	switch t.Type {
	case parent:
		fmt.Fprintf(w, `
			func (e *html%s) Body(v ...UI) HTML%s {
				e.setChildren(v...)
				return e
			}
			`,
			t.Name,
			t.Name,
		)

		if t.Name == "Textarea" {
			fmt.Fprintf(w, `
			func (e *html%s) Text(v any) HTML%s {
				e.setAttr("value", v)
				return e
			}
			`,
				t.Name,
				t.Name,
			)
		} else {
			fmt.Fprintf(w, `
			func (e *html%s) Text(v any) HTML%s {
				return e.Body(Text(v))
			}
			`,
				t.Name,
				t.Name,
			)
		}

	case privateParent:
		fmt.Fprintf(w, `
			func (e *html%s) privateBody(v ...UI) HTML%s {
				e.setChildren(v...)
				return e
			}
			`,
			t.Name,
			t.Name,
		)
	}

	for _, a := range t.Attrs {
		fmt.Fprintln(w)
		fmt.Fprintln(w)

		writeAttrFunction(w, a, t, false)
	}

	fmt.Fprintln(w)

	fmt.Fprintf(w, `
		func (e *html%s) On(event string, h EventHandler, scope ...any)  HTML%s {
			e.setEventHandler(event, h, scope...)
			return e
		}
		`,
		t.Name,
		t.Name,
	)

	for _, e := range t.EventHandlers {
		fmt.Fprintln(w)
		fmt.Fprintln(w)

		writeEventFunction(w, e, t, false)
	}
}

func writeAttrFunction(w io.Writer, a attr, t tag, isInterface bool) {
	if !isInterface {
		fmt.Fprintf(w, "func (e *html%s)", t.Name)
	}

	var attrName string
	if a.NameOverride != "" {
		attrName = strings.ToLower(a.NameOverride)
	} else {
		attrName = strings.ToLower(a.Name)
	}

	switch a.Type {
	case "data|value":
		fmt.Fprintf(w, `%s(k string, v any) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				e.setAttr("data-"+k, fmt.Sprintf("%s", v))
				return e
			}`, "%v")
		}

	case "attr|value":
		fmt.Fprintf(w, `%s(n string, v any) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				e.setAttr(n, v)
				return e
			}`)
		}

	case "aria|value":
		fmt.Fprintf(w, `%s(k string, v any) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				e.setAttr("aria-"+k, fmt.Sprintf("%s", v))
				return e
			}`, "%v")
		}

	case "style":
		fmt.Fprintf(w, `%s(k, v string) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				e.setAttr("style", k+":"+v)
				return e
			}`)
		}

	case "style|map":
		fmt.Fprintf(w, `%s(s map[string]string) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				for k, v := range s {
					e.Style(k, v)
				}
				return e
			}`)
		}

	case "on/off":
		fmt.Fprintf(w, `%s(v bool) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				s := "off"
				if (v) {
					s = "on"
				}
	
				e.setAttr("%s", s)
				return e
			}`, attrName)
		}

	case "bool|force":
		fmt.Fprintf(w, `%s(v bool) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				s := "false"
				if (v) {
					s = "true"
				}
	
				e.setAttr("%s", s)
				return e
			}`, attrName)
		}

	case "url":
		fmt.Fprintf(w, `%s(v string) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				e.setAttr("%s", v)
				return e
			}`, attrName)
		}

	case "string|class":
		fmt.Fprintf(w, `%s(v ...string) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				e.setAttr("%s", strings.Join(v, " "))
				return e
			}`, attrName)
		}

	case "xmlns":
		fmt.Fprintf(w, `%s(v string) HTML%s`, a.Name, t.Name)
		if !isInterface {
			fmt.Fprintln(w, `{
				e.xmlns = v
				return e
			}`)
		}

	default:
		fmt.Fprintf(w, `%s(v %s) HTML%s`, a.Name, a.Type, t.Name)
		if !isInterface {
			fmt.Fprintf(w, `{
				e.setAttr("%s", v)
				return e
			}`, attrName)
		}
	}
}

func writeEventFunction(w io.Writer, e eventHandler, t tag, isInterface bool) {
	if !isInterface {
		fmt.Fprintf(w, `func (e *html%s)`, t.Name)
	}

	fmt.Fprintf(w, `%s (h EventHandler, scope ...any) HTML%s`, e.Name, t.Name)
	if !isInterface {
		fmt.Fprintf(w, `{
			e.setEventHandler("%s", h, scope...)
			return e
		}`, strings.TrimPrefix(strings.ToLower(e.Name), "on"))
	}
}

func generateHTMLTestGo() {
	f, err := os.Create("html_gen_test.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package app")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// Code generated by go generate; DO NOT EDIT.")
	fmt.Fprintln(f, `
import (
	"testing"
)
		`)

	for _, t := range tags {
		fmt.Fprintln(f)
		fmt.Fprintf(f, `func Test%s(t *testing.T) {`, t.Name)
		fmt.Fprintln(f)

		switch t.Name {
		case "Elem", "ElemSelfClosing":
			fmt.Fprintf(f, `elem := %s("div")`, t.Name)

		default:
			fmt.Fprintf(f, `elem := %s()`, t.Name)
		}

		fmt.Fprintln(f)

		for _, a := range t.Attrs {
			fmt.Fprintf(f, `elem.%s(`, a.Name)

			switch a.Type {
			case "data|value", "aria|value", "attr|value":
				fmt.Fprintln(f, `"foo", "bar")`)

			case "style":
				fmt.Fprintln(f, `"color", "deepskyblue")`)

			case "style|map":
				fmt.Fprintln(f, `map[string]string{"color": "pink"})`)

			case "bool", "bool|force", "on/off":
				fmt.Fprintln(f, `true)`)
				fmt.Fprintf(f, `elem.%s(false)`, a.Name)
				fmt.Fprintln(f)

			case "int":
				fmt.Fprintln(f, `42)`)

			case "string":
				fmt.Fprintln(f, `"foo")`)

			case "url":
				fmt.Fprintln(f, `"http://foo.com")`)

			case "string|class":
				fmt.Fprintln(f, `"foo bar")`)

			case "xmlns":
				fmt.Fprintln(f, `"http://www.w3.org/2000/svg")`)

			default:
				fmt.Fprintln(f, `42)`)
			}
		}

		if len(t.EventHandlers) != 0 {
			fmt.Fprint(f, `
				h := func(ctx Context, e Event) {}
			`)
			fmt.Fprintf(f, `elem.On("click", h)`)
			fmt.Fprintln(f)
		}

		for _, e := range t.EventHandlers {
			fmt.Fprintf(f, `elem.%s(h)`, e.Name)
			fmt.Fprintln(f)
		}

		switch t.Type {
		case parent:
			fmt.Fprintln(f, `elem.Text("hello")`)

		case privateParent:
			fmt.Fprintln(f, `elem.privateBody(Text("hello"))`)
		}

		fmt.Fprintln(f, "}")
	}
}
