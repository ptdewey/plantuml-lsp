package language

type LangDefs struct {
	Defs []string

	// CompletionItemKind:
	// Text = 1; Method = 2; Function = 3; Constructor = 4; Field = 5; Variable
	// = 6; Class = 7; Interface = 8;
	// Module = 9; Property = 10; Unit = 11; Value = 12; Enum = 13;
	// Keyword = 14; Snippet = 15; Color = 16; File = 17;
	// Reference = 18; Folder = 19; EnumMember = 20; Constant = 21;
	// Struct = 22; Event = 23; Operator = 24; TypeParameter = 25;
	Kind int
}

// PlantUML language definitions acquired from 'plantuml -language'
var (
	Types = LangDefs{
		Defs: []string{
			"abstract", "action", "actor", "agent", "annotation", "archimate", "artifact",
			"boundary", "card", "class", "cloud", "collections", "component", "control",
			"database", "diamond", "entity", "enum", "exception", "file", "folder",
			"frame", "hexagon", "interface", "json", "label", "map", "metaclass", "node",
			"object", "package", "participant", "person", "process", "protocol", "queue",
			"rectangle", "relationship", "stack", "state", "storage", "struct", "usecase",
		},
		Kind: 7,
	}

	Keywords = LangDefs{
		Defs: []string{
			"@endboard", "@endbpm", "@endchen", "@endchronology", "@endcreole",
			"@endcute", "@enddef", "@endditaa", "@enddot", "@endebnf", "@endfiles",
			"@endflow", "@endgantt", "@endgit", "@endhcl", "@endjcckit", "@endjson",
			"@endlatex", "@endmath", "@endmindmap", "@endnwdiag", "@endproject",
			"@endregex", "@endsalt", "@endtree", "@enduml", "@endwbs", "@endwire",
			"@endyaml", "@startboard", "@startbpm", "@startchen", "@startchronology",
			"@startcreole", "@startcute", "@startdef", "@startditaa", "@startdot",
			"@startebnf", "@startfiles", "@startflow", "@startgantt", "@startgit",
			"@starthcl", "@startjcckit", "@startjson", "@startlatex", "@startmath",
			"@startmindmap", "@startnwdiag", "@startproject", "@startregex", "@startsalt",
			"@starttree", "@startuml", "@startwbs", "@startwire", "@startyaml", "across",
			"activate", "again", "allow_mixing", "allowmixing", "also", "alt", "as",
			"attribute", "attributes", "autonumber", "bold", "bottom", "box", "break",
			"caption", "center", "circle", "circled", "circles", "color", "create",
			"critical", "dashed", "deactivate", "description", "destroy", "detach",
			"dotted", "down", "else", "elseif", "empty", "end", "endcaption", "endfooter",
			"endheader", "endif", "endlegend", "endtitle", "endwhile", "false", "field",
			"fields", "footbox", "footer", "fork", "group", "header", "hide", "hnote",
			"if", "is", "italic", "kill", "left", "left to right direction", "legend",
			"link", "loop", "mainframe", "member", "members", "method", "methods",
			"namespace", "newpage", "normal", "note", "of", "on", "opt", "order", "over",
			"package", "page", "par", "partition", "plain", "private", "protected",
			"public", "ref", "repeat", "return", "right", "rnote", "rotate", "show",
			"skin", "skinparam", "split", "sprite", "start", "stereotype", "stereotypes",
			"stop", "style", "then", "title", "together", "top",
			"top to bottom direction", "true", "up", "while",
		},
		Kind: 14,
	}

	Preprocessors = LangDefs{
		Defs: []string{
			"!assert", "!define", "!definelong", "!dump_memory", "!else",
			"!enddefinelong", "!endfunction", "!endif", "!endprocedure", "!endsub",
			"!exit", "!function", "!if", "!ifdef", "!ifndef", "!import", "!include",
			"!local", "!log", "!pragma", "!procedure", "!return", "!startsub",
			"!theme", "!undef", "!unquoted",
		},
		Kind: 14,
	}

	Arrows = LangDefs{
		Defs: []string{
			"--", "..", "-->", "<--", "--*", "*--", "--o", "o--", "<|--", "--|>",
			"..|>", "<|..", "*-->", "<--*", "o-->", "<--o", ".", "->", "<-", "-*",
			"*-", "-o", "o-", "<|-", "-|>", ".|>", "<|.", "*->", "<-*", "o->",
			"<-o", "-[hidden]-->", "-[hidden]->", "--down->", "--up->", "--left->",
			"--right->", "..down..>", "..up..>", "..left..>", "..right..>",
			"--d->", "--u->", "--l->", "--r->", "..d..>", "..u..>", "..l..>",
			"..r..>",
		},
		Kind: 24,
	}

	Skinparameters = LangDefs{
		Defs: []string{
			"ActivityBackgroundColor", "ActivityBorderColor",
			"ActivityBorderThickness", "ActivityDiamondFontColor",
			"ActivityDiamondFontName", "ActivityDiamondFontSize",
			"ActivityDiamondFontStyle", "ActivityFontColor", "ActivityFontName",
			"ActivityFontSize", "ActivityFontStyle", "ActorBackgroundColor",
			"ActorBorderColor", "ActorFontColor", "ActorFontName", "ActorFontSize",
			"ActorFontStyle", "ActorStereotypeFontColor", "ActorStereotypeFontName",
			"ActorStereotypeFontSize", "ActorStereotypeFontStyle",
			"AgentBorderThickness", "AgentFontColor", "AgentFontName",
			"AgentFontSize", "AgentFontStyle", "AgentStereotypeFontColor",
			"AgentStereotypeFontName", "AgentStereotypeFontSize",
			"AgentStereotypeFontStyle", "ArchimateBorderThickness",
			"ArchimateFontColor", "ArchimateFontName", "ArchimateFontSize",
			"ArchimateFontStyle", "ArchimateStereotypeFontColor",
			"ArchimateStereotypeFontName", "ArchimateStereotypeFontSize",
			"ArchimateStereotypeFontStyle", "ArrowFontColor", "ArrowFontName",
			"ArrowFontSize", "ArrowFontStyle", "ArrowHeadColor",
			"ArrowLollipopColor", "ArrowMessageAlignment", "ArrowThickness",
			"ArtifactFontColor", "ArtifactFontName", "ArtifactFontSize",
			"ArtifactFontStyle", "ArtifactStereotypeFontColor",
			"ArtifactStereotypeFontName", "ArtifactStereotypeFontSize",
			"ArtifactStereotypeFontStyle", "BackgroundColor",
			"BiddableBackgroundColor", "BiddableBorderColor", "BoundaryFontColor",
			"BoundaryFontName", "BoundaryFontSize", "BoundaryFontStyle",
			"BoundaryStereotypeFontColor", "BoundaryStereotypeFontName",
			"BoundaryStereotypeFontSize", "BoundaryStereotypeFontStyle",
			"BoxPadding", "CaptionFontColor", "CaptionFontName", "CaptionFontSize",
			"CaptionFontStyle", "CardBorderThickness", "CardFontColor",
			"CardFontName", "CardFontSize", "CardFontStyle",
			"CardStereotypeFontColor", "CardStereotypeFontName",
			"CardStereotypeFontSize", "CardStereotypeFontStyle",
			"CircledCharacterFontColor", "CircledCharacterFontName",
			"CircledCharacterFontSize", "CircledCharacterFontStyle",
			"CircledCharacterRadius", "ClassAttributeFontColor",
			"ClassAttributeFontName", "ClassAttributeFontSize",
			"ClassAttributeFontStyle", "ClassAttributeIconSize",
			"ClassBackgroundColor", "ClassBorderColor", "ClassBorderThickness",
			"ClassFontColor", "ClassFontName", "ClassFontSize", "ClassFontStyle",
			"ClassStereotypeFontColor", "ClassStereotypeFontName",
			"ClassStereotypeFontSize", "ClassStereotypeFontStyle", "CloudFontColor",
			"CloudFontName", "CloudFontSize", "CloudFontStyle",
			"CloudStereotypeFontColor", "CloudStereotypeFontName",
			"CloudStereotypeFontSize", "CloudStereotypeFontStyle",
			"ColorArrowSeparationSpace", "ComponentBorderThickness",
			"ComponentFontColor", "ComponentFontName", "ComponentFontSize",
			"ComponentFontStyle", "ComponentStereotypeFontColor",
			"ComponentStereotypeFontName", "ComponentStereotypeFontSize",
			"ComponentStereotypeFontStyle", "ComponentStyle", "ConditionEndStyle",
			"ConditionStyle", "ControlFontColor", "ControlFontName",
			"ControlFontSize", "ControlFontStyle", "ControlStereotypeFontColor",
			"ControlStereotypeFontName", "ControlStereotypeFontSize",
			"ControlStereotypeFontStyle", "DatabaseFontColor", "DatabaseFontName",
			"DatabaseFontSize", "DatabaseFontStyle", "DatabaseStereotypeFontColor",
			"DatabaseStereotypeFontName", "DatabaseStereotypeFontSize",
			"DatabaseStereotypeFontStyle", "DefaultFontColor", "DefaultFontName",
			"DefaultFontSize", "DefaultFontStyle", "DefaultMonospacedFontName",
			"DefaultTextAlignment", "DesignedBackgroundColor",
			"DesignedBorderColor", "DesignedDomainBorderThickness",
			"DesignedDomainFontColor", "DesignedDomainFontName",
			"DesignedDomainFontSize", "DesignedDomainFontStyle",
			"DesignedDomainStereotypeFontColor", "DesignedDomainStereotypeFontName",
			"DesignedDomainStereotypeFontSize", "DesignedDomainStereotypeFontStyle",
			"DiagramBorderColor", "DiagramBorderThickness", "DomainBackgroundColor",
			"DomainBorderColor", "DomainBorderThickness", "DomainFontColor",
			"DomainFontName", "DomainFontSize", "DomainFontStyle",
			"DomainStereotypeFontColor", "DomainStereotypeFontName",
			"DomainStereotypeFontSize", "DomainStereotypeFontStyle", "Dpi",
			"EntityFontColor", "EntityFontName", "EntityFontSize",
			"EntityFontStyle", "EntityStereotypeFontColor",
			"EntityStereotypeFontName", "EntityStereotypeFontSize",
			"EntityStereotypeFontStyle", "FileFontColor", "FileFontName",
			"FileFontSize", "FileFontStyle", "FileStereotypeFontColor",
			"FileStereotypeFontName", "FileStereotypeFontSize",
			"FileStereotypeFontStyle", "FixCircleLabelOverlapping",
			"FolderFontColor", "FolderFontName", "FolderFontSize",
			"FolderFontStyle", "FolderStereotypeFontColor",
			"FolderStereotypeFontName", "FolderStereotypeFontSize",
			"FolderStereotypeFontStyle", "FooterFontColor", "FooterFontName",
			"FooterFontSize", "FooterFontStyle", "FrameFontColor", "FrameFontName",
			"FrameFontSize", "FrameFontStyle", "FrameStereotypeFontColor",
			"FrameStereotypeFontName", "FrameStereotypeFontSize",
			"FrameStereotypeFontStyle", "GenericDisplay", "Guillemet",
			"Handwritten", "HeaderFontColor", "HeaderFontName", "HeaderFontSize",
			"HeaderFontStyle", "HexagonBorderThickness", "HexagonFontColor",
			"HexagonFontName", "HexagonFontSize", "HexagonFontStyle",
			"HexagonStereotypeFontColor", "HexagonStereotypeFontName",
			"HexagonStereotypeFontSize", "HexagonStereotypeFontStyle",
			"HyperlinkColor", "HyperlinkUnderline", "IconIEMandatoryColor",
			"IconPackageBackgroundColor", "IconPackageColor",
			"IconPrivateBackgroundColor", "IconPrivateColor",
			"IconProtectedBackgroundColor", "IconProtectedColor",
			"IconPublicBackgroundColor", "IconPublicColor", "InterfaceFontColor",
			"InterfaceFontName", "InterfaceFontSize", "InterfaceFontStyle",
			"InterfaceStereotypeFontColor", "InterfaceStereotypeFontName",
			"InterfaceStereotypeFontSize", "InterfaceStereotypeFontStyle",
			"LabelFontColor", "LabelFontName", "LabelFontSize", "LabelFontStyle",
			"LabelStereotypeFontColor", "LabelStereotypeFontName",
			"LabelStereotypeFontSize", "LabelStereotypeFontStyle",
			"LegendBorderThickness", "LegendFontColor", "LegendFontName",
			"LegendFontSize", "LegendFontStyle", "LexicalBackgroundColor",
			"LexicalBorderColor", "LifelineStrategy", "Linetype",
			"MachineBackgroundColor", "MachineBorderColor",
			"MachineBorderThickness", "MachineFontColor", "MachineFontName",
			"MachineFontSize", "MachineFontStyle", "MachineStereotypeFontColor",
			"MachineStereotypeFontName", "MachineStereotypeFontSize",
			"MachineStereotypeFontStyle", "MaxAsciiMessageLength", "MaxMessageSize",
			"MinClassWidth", "Monochrome", "NodeFontColor", "NodeFontName",
			"NodeFontSize", "NodeFontStyle", "NodeStereotypeFontColor",
			"NodeStereotypeFontName", "NodeStereotypeFontSize",
			"NodeStereotypeFontStyle", "Nodesep", "NoteBackgroundColor",
			"NoteBorderColor", "NoteBorderThickness", "NoteFontColor",
			"NoteFontName", "NoteFontSize", "NoteFontStyle", "NoteShadowing",
			"NoteTextAlignment", "ObjectAttributeFontColor",
			"ObjectAttributeFontName", "ObjectAttributeFontSize",
			"ObjectAttributeFontStyle", "ObjectBorderThickness", "ObjectFontColor",
			"ObjectFontName", "ObjectFontSize", "ObjectFontStyle",
			"ObjectStereotypeFontColor", "ObjectStereotypeFontName",
			"ObjectStereotypeFontSize", "ObjectStereotypeFontStyle",
			"PackageBorderThickness", "PackageFontColor", "PackageFontName",
			"PackageFontSize", "PackageFontStyle", "PackageStereotypeFontColor",
			"PackageStereotypeFontName", "PackageStereotypeFontSize",
			"PackageStereotypeFontStyle", "PackageStyle", "PackageTitleAlignment",
			"Padding", "PageBorderColor", "PageExternalColor", "PageMargin",
			"ParticipantFontColor", "ParticipantFontName", "ParticipantFontSize",
			"ParticipantFontStyle", "ParticipantPadding",
			"ParticipantStereotypeFontColor", "ParticipantStereotypeFontName",
			"ParticipantStereotypeFontSize", "ParticipantStereotypeFontStyle",
			"PartitionBorderThickness", "PartitionFontColor", "PartitionFontName",
			"PartitionFontSize", "PartitionFontStyle", "PathHoverColor",
			"PersonBorderThickness", "PersonFontColor", "PersonFontName",
			"PersonFontSize", "PersonFontStyle", "PersonStereotypeFontColor",
			"PersonStereotypeFontName", "PersonStereotypeFontSize",
			"PersonStereotypeFontStyle", "QueueBorderThickness", "QueueFontColor",
			"QueueFontName", "QueueFontSize", "QueueFontStyle",
			"QueueStereotypeFontColor", "QueueStereotypeFontName",
			"QueueStereotypeFontSize", "QueueStereotypeFontStyle", "Ranksep",
			"RectangleBorderThickness", "RectangleFontColor", "RectangleFontName",
			"RectangleFontSize", "RectangleFontStyle",
			"RectangleStereotypeFontColor", "RectangleStereotypeFontName",
			"RectangleStereotypeFontSize", "RectangleStereotypeFontStyle",
			"RequirementBackgroundColor", "RequirementBorderColor",
			"RequirementBorderThickness", "RequirementFontColor",
			"RequirementFontName", "RequirementFontSize", "RequirementFontStyle",
			"RequirementStereotypeFontColor", "RequirementStereotypeFontName",
			"RequirementStereotypeFontSize", "RequirementStereotypeFontStyle",
			"ResponseMessageBelowArrow", "RoundCorner", "SameClassWidth",
			"SequenceActorBorderThickness", "SequenceArrowThickness",
			"SequenceBoxBorderColor", "SequenceBoxFontColor", "SequenceBoxFontName",
			"SequenceBoxFontSize", "SequenceBoxFontStyle", "SequenceDelayFontColor",
			"SequenceDelayFontName", "SequenceDelayFontSize",
			"SequenceDelayFontStyle", "SequenceDividerBorderThickness",
			"SequenceDividerFontColor", "SequenceDividerFontName",
			"SequenceDividerFontSize", "SequenceDividerFontStyle",
			"SequenceGroupBodyBackgroundColor", "SequenceGroupBorderThickness",
			"SequenceGroupFontColor", "SequenceGroupFontName",
			"SequenceGroupFontSize", "SequenceGroupFontStyle",
			"SequenceGroupHeaderFontColor", "SequenceGroupHeaderFontName",
			"SequenceGroupHeaderFontSize", "SequenceGroupHeaderFontStyle",
			"SequenceLifeLineBorderColor", "SequenceLifeLineBorderThickness",
			"SequenceMessageAlignment", "SequenceMessageTextAlignment",
			"SequenceNewpageSeparatorColor", "SequenceParticipant",
			"SequenceParticipantBorderThickness", "SequenceReferenceAlignment",
			"SequenceReferenceBackgroundColor", "SequenceReferenceBorderThickness",
			"SequenceReferenceFontColor", "SequenceReferenceFontName",
			"SequenceReferenceFontSize", "SequenceReferenceFontStyle",
			"SequenceReferenceHeaderBackgroundColor", "SequenceStereotypeFontColor",
			"SequenceStereotypeFontName", "SequenceStereotypeFontSize",
			"SequenceStereotypeFontStyle", "Shadowing", "StackFontColor",
			"StackFontName", "StackFontSize", "StackFontStyle",
			"StackStereotypeFontColor", "StackStereotypeFontName",
			"StackStereotypeFontSize", "StackStereotypeFontStyle",
			"StateAttributeFontColor", "StateAttributeFontName",
			"StateAttributeFontSize", "StateAttributeFontStyle", "StateBorderColor",
			"StateFontColor", "StateFontName", "StateFontSize", "StateFontStyle",
			"StateMessageAlignment", "StereotypePosition", "StorageFontColor",
			"StorageFontName", "StorageFontSize", "StorageFontStyle",
			"StorageStereotypeFontColor", "StorageStereotypeFontName",
			"StorageStereotypeFontSize", "StorageStereotypeFontStyle", "Style",
			"SvglinkTarget", "SwimlaneBorderThickness", "SwimlaneTitleFontColor",
			"SwimlaneTitleFontName", "SwimlaneTitleFontSize",
			"SwimlaneTitleFontStyle", "SwimlaneWidth", "SwimlaneWrapTitleWidth",
			"TabSize", "TimingFontColor", "TimingFontName", "TimingFontSize",
			"TimingFontStyle", "TitleBorderRoundCorner", "TitleBorderThickness",
			"TitleFontColor", "TitleFontName", "TitleFontSize", "TitleFontStyle",
			"UsecaseBorderThickness", "UsecaseFontColor", "UsecaseFontName",
			"UsecaseFontSize", "UsecaseFontStyle", "UsecaseStereotypeFontColor",
			"UsecaseStereotypeFontName", "UsecaseStereotypeFontSize",
			"UsecaseStereotypeFontStyle", "WrapWidth",
		},
		Kind: 10,
	}

	Colors = LangDefs{
		Defs: []string{
			"APPLICATION", "AliceBlue", "AntiqueWhite", "Aqua", "Aquamarine",
			"Azure", "BUSINESS", "Beige", "Bisque", "Black", "BlanchedAlmond",
			"Blue", "BlueViolet", "Brown", "BurlyWood", "CadetBlue", "Chartreuse",
			"Chocolate", "Coral", "CornflowerBlue", "Cornsilk", "Crimson", "Cyan",
			"DarkBlue", "DarkCyan", "DarkGoldenRod", "DarkGray", "DarkGreen",
			"DarkGrey", "DarkKhaki", "DarkMagenta", "DarkOliveGreen", "DarkOrchid",
			"DarkRed", "DarkSalmon", "DarkSeaGreen", "DarkSlateBlue",
			"DarkSlateGray", "DarkSlateGrey", "DarkTurquoise", "DarkViolet",
			"Darkorange", "DeepPink", "DeepSkyBlue", "DimGray", "DimGrey",
			"DodgerBlue", "FireBrick", "FloralWhite", "ForestGreen", "Fuchsia",
			"Gainsboro", "GhostWhite", "Gold", "GoldenRod", "Gray", "Green",
			"GreenYellow", "Grey", "HoneyDew", "HotPink", "IMPLEMENTATION",
			"IndianRed", "Indigo", "Ivory", "Khaki", "Lavender", "LavenderBlush",
			"LawnGreen", "LemonChiffon", "LightBlue", "LightCoral", "LightCyan",
			"LightGoldenRodYellow", "LightGray", "LightGreen", "LightGrey",
			"LightPink", "LightSalmon", "LightSeaGreen", "LightSkyBlue",
			"LightSlateGray", "LightSlateGrey", "LightSteelBlue", "LightYellow",
			"Lime", "LimeGreen", "Linen", "MOTIVATION", "Magenta", "Maroon",
			"MediumAquaMarine", "MediumBlue", "MediumOrchid", "MediumPurple",
			"MediumSeaGreen", "MediumSlateBlue", "MediumSpringGreen",
			"MediumTurquoise", "MediumVioletRed", "MidnightBlue", "MintCream",
			"MistyRose", "Moccasin", "NavajoWhite", "Navy", "OldLace", "Olive",
			"OliveDrab", "Orange", "OrangeRed", "Orchid", "PHYSICAL",
			"PaleGoldenRod", "PaleGreen", "PaleTurquoise", "PaleVioletRed",
			"PapayaWhip", "PeachPuff", "Peru", "Pink", "Plum", "PowderBlue",
			"Purple", "Red", "RosyBrown", "RoyalBlue", "STRATEGY", "SaddleBrown",
			"Salmon", "SandyBrown", "SeaGreen", "SeaShell", "Sienna", "Silver",
			"SkyBlue", "SlateBlue", "SlateGray", "SlateGrey", "Snow", "SpringGreen",
			"SteelBlue", "TECHNOLOGY", "Tan", "Teal", "Thistle", "Tomato",
			"Turquoise", "Violet", "Wheat", "White", "WhiteSmoke", "Yellow",
			"YellowGreen",
		},
		Kind: 16,
	}
)
