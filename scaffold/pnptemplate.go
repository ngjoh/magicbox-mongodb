package scaffold

import "encoding/xml"

type Field struct {
	Text                             string `xml:",chardata"`
	DisplayName                      string `xml:"DisplayName,attr"`
	Format                           string `xml:"Format,attr"`
	IsModern                         string `xml:"IsModern,attr"`
	Name                             string `xml:"Name,attr"`
	Title                            string `xml:"Title,attr"`
	Type                             string `xml:"Type,attr"`
	ID                               string `xml:"ID,attr"`
	SourceID                         string `xml:"SourceID,attr"`
	StaticName                       string `xml:"StaticName,attr"`
	ColName                          string `xml:"ColName,attr"`
	RowOrdinal                       string `xml:"RowOrdinal,attr"`
	AppendOnly                       string `xml:"AppendOnly,attr"`
	IsolateStyles                    string `xml:"IsolateStyles,attr"`
	RichText                         string `xml:"RichText,attr"`
	RichTextMode                     string `xml:"RichTextMode,attr"`
	Required                         string `xml:"Required,attr"`
	EnforceUniqueValues              string `xml:"EnforceUniqueValues,attr"`
	List                             string `xml:"List,attr"`
	ShowField                        string `xml:"ShowField,attr"`
	Mult                             string `xml:"Mult,attr"`
	Sortable                         string `xml:"Sortable,attr"`
	UnlimitedLengthInDocumentLibrary string `xml:"UnlimitedLengthInDocumentLibrary,attr"`
	DisplaceOnUpgrade                string `xml:"DisplaceOnUpgrade,attr"`
	ShowInFileDlg                    string `xml:"ShowInFileDlg,attr"`
	ReadOnly                         string `xml:"ReadOnly,attr"`
	FromBaseType                     string `xml:"FromBaseType,attr"`
	Hidden                           string `xml:"Hidden,attr"`
	CanToggleHidden                  string `xml:"CanToggleHidden,attr"`
	MaxLength                        string `xml:"MaxLength,attr"`
	Version                          string `xml:"Version,attr"`
	Indexed                          string `xml:"Indexed,attr"`
	FriendlyDisplayFormat            string `xml:"FriendlyDisplayFormat,attr"`
	Description                      string `xml:"Description,attr"`
	NumLines                         string `xml:"NumLines,attr"`
	FillInChoice                     string `xml:"FillInChoice,attr"`
	RelationshipDeleteBehavior       string `xml:"RelationshipDeleteBehavior,attr"`
	Decimals                         string `xml:"Decimals,attr"`
	Percentage                       string `xml:"Percentage,attr"`
	UserSelectionScope               string `xml:"UserSelectionScope,attr"`
	UserSelectionMode                string `xml:"UserSelectionMode,attr"`
	CustomFormatter                  string `xml:"CustomFormatter,attr"`
	UserDisplayOptions               string `xml:"UserDisplayOptions,attr"`
	Group                            string `xml:"Group,attr"`
	CommaSeparator                   string `xml:"CommaSeparator,attr"`
	CustomUnitOnRight                string `xml:"CustomUnitOnRight,attr"`
	Unit                             string `xml:"Unit,attr"`
	ClientSideComponentId            string `xml:"ClientSideComponentId,attr"`
	Sealed                           string `xml:"Sealed,attr"`
	AllowDeletion                    string `xml:"AllowDeletion,attr"`
	ShowInNewForm                    string `xml:"ShowInNewForm,attr"`
	ShowInEditForm                   string `xml:"ShowInEditForm,attr"`
	ShowInDisplayForm                string `xml:"ShowInDisplayForm,attr"`
	ShowInViewForms                  string `xml:"ShowInViewForms,attr"`
	ShowInListSettings               string `xml:"ShowInListSettings,attr"`
	ShowInVersionHistory             string `xml:"ShowInVersionHistory,attr"`
	FieldRef                         string `xml:"FieldRef,attr"`
	Viewable                         string `xml:"Viewable,attr"`
	JSON                             string `xml:"Json,attr"`
	WebId                            string `xml:"WebId,attr"`
	Default                          string `xml:"Default"`
	CHOICES                          struct {
		Text   string   `xml:",chardata"`
		CHOICE []string `xml:"CHOICE"`
	} `xml:"CHOICES"`
	Customization struct {
		Text            string `xml:",chardata"`
		ArrayOfProperty struct {
			Text     string `xml:",chardata"`
			Property []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"Name"`
				Value struct {
					Text string `xml:",chardata"`
					Q1   string `xml:"q1,attr"`
					Type string `xml:"type,attr"`
					P4   string `xml:"p4,attr"`
					Q2   string `xml:"q2,attr"`
					Q3   string `xml:"q3,attr"`
					Q4   string `xml:"q4,attr"`
					Q5   string `xml:"q5,attr"`
					Q6   string `xml:"q6,attr"`
					Q7   string `xml:"q7,attr"`
					Q8   string `xml:"q8,attr"`
					Q9   string `xml:"q9,attr"`
					Q10  string `xml:"q10,attr"`
					Q11  string `xml:"q11,attr"`
					Q12  string `xml:"q12,attr"`
					Q13  string `xml:"q13,attr"`
				} `xml:"Value"`
			} `xml:"Property"`
		} `xml:"ArrayOfProperty"`
	} `xml:"Customization"`
}

type Fields struct {
	Text  string  `xml:",chardata"`
	Field []Field `xml:"Field"`
}
type ListInstance struct {
	Text                   string `xml:",chardata"`
	Title                  string `xml:"Title,attr"`
	Description            string `xml:"Description,attr"`
	DocumentTemplate       string `xml:"DocumentTemplate,attr"`
	TemplateType           string `xml:"TemplateType,attr"`
	URL                    string `xml:"Url,attr"`
	EnableVersioning       string `xml:"EnableVersioning,attr"`
	MinorVersionLimit      string `xml:"MinorVersionLimit,attr"`
	MaxVersionLimit        string `xml:"MaxVersionLimit,attr"`
	DraftVersionVisibility string `xml:"DraftVersionVisibility,attr"`
	TemplateFeatureID      string `xml:"TemplateFeatureID,attr"`
	EnableFolderCreation   string `xml:"EnableFolderCreation,attr"`
	DefaultDisplayFormUrl  string `xml:"DefaultDisplayFormUrl,attr"`
	DefaultEditFormUrl     string `xml:"DefaultEditFormUrl,attr"`
	DefaultNewFormUrl      string `xml:"DefaultNewFormUrl,attr"`
	ImageUrl               string `xml:"ImageUrl,attr"`
	IrmExpire              string `xml:"IrmExpire,attr"`
	IrmReject              string `xml:"IrmReject,attr"`
	IsApplicationList      string `xml:"IsApplicationList,attr"`
	ValidationFormula      string `xml:"ValidationFormula,attr"`
	ValidationMessage      string `xml:"ValidationMessage,attr"`
	ContentTypesEnabled    string `xml:"ContentTypesEnabled,attr"`
	OnQuickLaunch          string `xml:"OnQuickLaunch,attr"`
	EnableAttachments      string `xml:"EnableAttachments,attr"`
	ReadSecurity           string `xml:"ReadSecurity,attr"`
	WriteSecurity          string `xml:"WriteSecurity,attr"`
	EnableMinorVersions    string `xml:"EnableMinorVersions,attr"`
	ContentTypeBindings    struct {
		Text               string `xml:",chardata"`
		ContentTypeBinding []struct {
			Text          string `xml:",chardata"`
			ContentTypeID string `xml:"ContentTypeID,attr"`
			Default       string `xml:"Default,attr"`
		} `xml:"ContentTypeBinding"`
	} `xml:"ContentTypeBindings"`
	Views struct {
		Text string `xml:",chardata"`
		View []struct {
			Text              string `xml:",chardata"`
			Name              string `xml:"Name,attr"`
			DefaultView       string `xml:"DefaultView,attr"`
			MobileView        string `xml:"MobileView,attr"`
			MobileDefaultView string `xml:"MobileDefaultView,attr"`
			Type              string `xml:"Type,attr"`
			DisplayName       string `xml:"DisplayName,attr"`
			URL               string `xml:"Url,attr"`
			Level             string `xml:"Level,attr"`
			BaseViewID        string `xml:"BaseViewID,attr"`
			ContentTypeID     string `xml:"ContentTypeID,attr"`
			ImageUrl          string `xml:"ImageUrl,attr"`
			RecurrenceRowset  string `xml:"RecurrenceRowset,attr"`
			ToolbarTemplate   string `xml:"ToolbarTemplate,attr"`
			Query             struct {
				Text    string `xml:",chardata"`
				OrderBy struct {
					Text     string `xml:",chardata"`
					FieldRef struct {
						Text      string `xml:",chardata"`
						Name      string `xml:"Name,attr"`
						Ascending string `xml:"Ascending,attr"`
					} `xml:"FieldRef"`
				} `xml:"OrderBy"`
				Where struct {
					Text              string `xml:",chardata"`
					DateRangesOverlap struct {
						Text     string `xml:",chardata"`
						FieldRef []struct {
							Text string `xml:",chardata"`
							Name string `xml:"Name,attr"`
						} `xml:"FieldRef"`
						Value struct {
							Text  string `xml:",chardata"`
							Type  string `xml:"Type,attr"`
							Month string `xml:"Month"`
							Now   string `xml:"Now"`
						} `xml:"Value"`
					} `xml:"DateRangesOverlap"`
					Neq struct {
						Text     string `xml:",chardata"`
						FieldRef struct {
							Text string `xml:",chardata"`
							Name string `xml:"Name,attr"`
						} `xml:"FieldRef"`
						Value struct {
							Text string `xml:",chardata"`
							Type string `xml:"Type,attr"`
						} `xml:"Value"`
					} `xml:"Neq"`
					Eq struct {
						Text     string `xml:",chardata"`
						FieldRef struct {
							Text string `xml:",chardata"`
							Name string `xml:"Name,attr"`
						} `xml:"FieldRef"`
						Value struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"Type,attr"`
							UserID string `xml:"UserID"`
						} `xml:"Value"`
					} `xml:"Eq"`
					Contains struct {
						Text     string `xml:",chardata"`
						FieldRef struct {
							Text string `xml:",chardata"`
							Name string `xml:"Name,attr"`
						} `xml:"FieldRef"`
						Value struct {
							Text string `xml:",chardata"`
							Type string `xml:"Type,attr"`
						} `xml:"Value"`
					} `xml:"Contains"`
				} `xml:"Where"`
				GroupBy struct {
					Text     string `xml:",chardata"`
					Collapse string `xml:"Collapse,attr"`
					FieldRef struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name,attr"`
					} `xml:"FieldRef"`
				} `xml:"GroupBy"`
			} `xml:"Query"`
			ViewFields struct {
				Text     string `xml:",chardata"`
				FieldRef []struct {
					Text string `xml:",chardata"`
					Name string `xml:"Name,attr"`
				} `xml:"FieldRef"`
			} `xml:"ViewFields"`
			RowLimit struct {
				Text  string `xml:",chardata"`
				Paged string `xml:"Paged,attr"`
			} `xml:"RowLimit"`
			JSLink          string `xml:"JSLink"`
			CustomFormatter string `xml:"CustomFormatter"`
			ViewData        struct {
				Text     string `xml:",chardata"`
				FieldRef []struct {
					Text string `xml:",chardata"`
					Name string `xml:"Name,attr"`
					Type string `xml:"Type,attr"`
				} `xml:"FieldRef"`
			} `xml:"ViewData"`
			Aggregations struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value,attr"`
			} `xml:"Aggregations"`
			CalendarViewStyles struct {
				Text              string `xml:",chardata"`
				CalendarViewStyle []struct {
					Text     string `xml:",chardata"`
					Title    string `xml:"Title,attr"`
					Type     string `xml:"Type,attr"`
					Template string `xml:"Template,attr"`
					Sequence string `xml:"Sequence,attr"`
					Default  string `xml:"Default,attr"`
				} `xml:"CalendarViewStyle"`
			} `xml:"CalendarViewStyles"`
		} `xml:"View"`
	} `xml:"Views"`
	Fields    Fields `xml:"Fields"`
	FieldRefs struct {
		Text     string `xml:",chardata"`
		FieldRef []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"ID,attr"`
			Name        string `xml:"Name,attr"`
			Required    string `xml:"Required,attr"`
			DisplayName string `xml:"DisplayName,attr"`
		} `xml:"FieldRef"`
	} `xml:"FieldRefs"`
	Security struct {
		Text                 string `xml:",chardata"`
		BreakRoleInheritance struct {
			Text                string `xml:",chardata"`
			CopyRoleAssignments string `xml:"CopyRoleAssignments,attr"`
			ClearSubscopes      string `xml:"ClearSubscopes,attr"`
			RoleAssignment      []struct {
				Text           string `xml:",chardata"`
				Principal      string `xml:"Principal,attr"`
				RoleDefinition string `xml:"RoleDefinition,attr"`
			} `xml:"RoleAssignment"`
		} `xml:"BreakRoleInheritance"`
	} `xml:"Security"`
	PropertyBagEntries struct {
		Text             string `xml:",chardata"`
		PropertyBagEntry struct {
			Text      string `xml:",chardata"`
			Key       string `xml:"Key,attr"`
			Value     string `xml:"Value,attr"`
			Overwrite string `xml:"Overwrite,attr"`
		} `xml:"PropertyBagEntry"`
	} `xml:"PropertyBagEntries"`
	FieldDefaults struct {
		Text         string `xml:",chardata"`
		FieldDefault struct {
			Text      string `xml:",chardata"`
			FieldName string `xml:"FieldName,attr"`
		} `xml:"FieldDefault"`
	} `xml:"FieldDefaults"`
}

type Provisioning struct {
	XMLName     xml.Name `xml:"Provisioning"`
	Text        string   `xml:",chardata"`
	Pnp         string   `xml:"pnp,attr"`
	Preferences struct {
		Text      string `xml:",chardata"`
		Generator string `xml:"Generator,attr"`
	} `xml:"Preferences"`
	Templates struct {
		Text                 string `xml:",chardata"`
		ID                   string `xml:"ID,attr"`
		ProvisioningTemplate struct {
			Text             string `xml:",chardata"`
			ID               string `xml:"ID,attr"`
			Version          string `xml:"Version,attr"`
			BaseSiteTemplate string `xml:"BaseSiteTemplate,attr"`
			Scope            string `xml:"Scope,attr"`
			WebSettings      struct {
				Text                        string `xml:",chardata"`
				RequestAccessEmail          string `xml:"RequestAccessEmail,attr"`
				NoCrawl                     string `xml:"NoCrawl,attr"`
				WelcomePage                 string `xml:"WelcomePage,attr"`
				SiteLogo                    string `xml:"SiteLogo,attr"`
				AlternateCSS                string `xml:"AlternateCSS,attr"`
				MasterPageUrl               string `xml:"MasterPageUrl,attr"`
				CustomMasterPageUrl         string `xml:"CustomMasterPageUrl,attr"`
				HubSiteUrl                  string `xml:"HubSiteUrl,attr"`
				CommentsOnSitePagesDisabled string `xml:"CommentsOnSitePagesDisabled,attr"`
				QuickLaunchEnabled          string `xml:"QuickLaunchEnabled,attr"`
				MembersCanShare             string `xml:"MembersCanShare,attr"`
				HorizontalQuickLaunch       string `xml:"HorizontalQuickLaunch,attr"`
				SearchScope                 string `xml:"SearchScope,attr"`
				SearchBoxInNavBar           string `xml:"SearchBoxInNavBar,attr"`
			} `xml:"WebSettings"`
			SiteSettings struct {
				Text                                   string `xml:",chardata"`
				AllowDesigner                          string `xml:"AllowDesigner,attr"`
				AllowCreateDeclarativeWorkflow         string `xml:"AllowCreateDeclarativeWorkflow,attr"`
				AllowSaveDeclarativeWorkflowAsTemplate string `xml:"AllowSaveDeclarativeWorkflowAsTemplate,attr"`
				AllowSavePublishDeclarativeWorkflow    string `xml:"AllowSavePublishDeclarativeWorkflow,attr"`
				SearchBoxInNavBar                      string `xml:"SearchBoxInNavBar,attr"`
				SearchCenterUrl                        string `xml:"SearchCenterUrl,attr"`
			} `xml:"SiteSettings"`
			RegionalSettings struct {
				Text                  string `xml:",chardata"`
				AdjustHijriDays       string `xml:"AdjustHijriDays,attr"`
				AlternateCalendarType string `xml:"AlternateCalendarType,attr"`
				CalendarType          string `xml:"CalendarType,attr"`
				Collation             string `xml:"Collation,attr"`
				FirstDayOfWeek        string `xml:"FirstDayOfWeek,attr"`
				FirstWeekOfYear       string `xml:"FirstWeekOfYear,attr"`
				LocaleId              string `xml:"LocaleId,attr"`
				ShowWeeks             string `xml:"ShowWeeks,attr"`
				Time24                string `xml:"Time24,attr"`
				TimeZone              string `xml:"TimeZone,attr"`
				WorkDayEndHour        string `xml:"WorkDayEndHour,attr"`
				WorkDays              string `xml:"WorkDays,attr"`
				WorkDayStartHour      string `xml:"WorkDayStartHour,attr"`
			} `xml:"RegionalSettings"`
			PropertyBagEntries struct {
				Text             string `xml:",chardata"`
				PropertyBagEntry []struct {
					Text      string `xml:",chardata"`
					Key       string `xml:"Key,attr"`
					Value     string `xml:"Value,attr"`
					Overwrite string `xml:"Overwrite,attr"`
					Indexed   string `xml:"Indexed,attr"`
				} `xml:"PropertyBagEntry"`
			} `xml:"PropertyBagEntries"`
			Security struct {
				Text                     string `xml:",chardata"`
				AssociatedOwnerGroup     string `xml:"AssociatedOwnerGroup,attr"`
				AssociatedMemberGroup    string `xml:"AssociatedMemberGroup,attr"`
				AssociatedVisitorGroup   string `xml:"AssociatedVisitorGroup,attr"`
				AdditionalAdministrators struct {
					Text string `xml:",chardata"`
					User struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name,attr"`
					} `xml:"User"`
				} `xml:"AdditionalAdministrators"`
				AdditionalOwners struct {
					Text string `xml:",chardata"`
					User []struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name,attr"`
					} `xml:"User"`
				} `xml:"AdditionalOwners"`
				AdditionalMembers struct {
					Text string `xml:",chardata"`
					User struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name,attr"`
					} `xml:"User"`
				} `xml:"AdditionalMembers"`
				AdditionalVisitors struct {
					Text string `xml:",chardata"`
					User []struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name,attr"`
					} `xml:"User"`
				} `xml:"AdditionalVisitors"`
				Permissions string `xml:"Permissions"`
			} `xml:"Security"`
			Navigation struct {
				Text                          string `xml:",chardata"`
				AddNewPagesToNavigation       string `xml:"AddNewPagesToNavigation,attr"`
				CreateFriendlyUrlsForNewPages string `xml:"CreateFriendlyUrlsForNewPages,attr"`
				GlobalNavigation              struct {
					Text                 string `xml:",chardata"`
					NavigationType       string `xml:"NavigationType,attr"`
					StructuralNavigation struct {
						Text                string `xml:",chardata"`
						RemoveExistingNodes string `xml:"RemoveExistingNodes,attr"`
					} `xml:"StructuralNavigation"`
				} `xml:"GlobalNavigation"`
				CurrentNavigation struct {
					Text                 string `xml:",chardata"`
					NavigationType       string `xml:"NavigationType,attr"`
					StructuralNavigation struct {
						Text                string `xml:",chardata"`
						RemoveExistingNodes string `xml:"RemoveExistingNodes,attr"`
						NavigationNode      []struct {
							Text           string `xml:",chardata"`
							Title          string `xml:"Title,attr"`
							URL            string `xml:"Url,attr"`
							IsExternal     string `xml:"IsExternal,attr"`
							NavigationNode []struct {
								Text           string `xml:",chardata"`
								Title          string `xml:"Title,attr"`
								URL            string `xml:"Url,attr"`
								IsExternal     string `xml:"IsExternal,attr"`
								NavigationNode []struct {
									Text  string `xml:",chardata"`
									Title string `xml:"Title,attr"`
									URL   string `xml:"Url,attr"`
								} `xml:"NavigationNode"`
							} `xml:"NavigationNode"`
						} `xml:"NavigationNode"`
					} `xml:"StructuralNavigation"`
				} `xml:"CurrentNavigation"`
			} `xml:"Navigation"`
			Lists struct {
				Text         string         `xml:",chardata"`
				ListInstance []ListInstance `xml:"ListInstance"`
			} `xml:"Lists"`
			Features struct {
				Text         string `xml:",chardata"`
				SiteFeatures struct {
					Text    string `xml:",chardata"`
					Feature struct {
						Text string `xml:",chardata"`
						ID   string `xml:"ID,attr"`
					} `xml:"Feature"`
				} `xml:"SiteFeatures"`
				WebFeatures struct {
					Text    string `xml:",chardata"`
					Feature []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"ID,attr"`
					} `xml:"Feature"`
				} `xml:"WebFeatures"`
			} `xml:"Features"`
			CustomActions struct {
				Text             string `xml:",chardata"`
				WebCustomActions struct {
					Text         string `xml:",chardata"`
					CustomAction struct {
						Text                          string `xml:",chardata"`
						Name                          string `xml:"Name,attr"`
						Location                      string `xml:"Location,attr"`
						Title                         string `xml:"Title,attr"`
						Sequence                      string `xml:"Sequence,attr"`
						Rights                        string `xml:"Rights,attr"`
						RegistrationType              string `xml:"RegistrationType,attr"`
						ClientSideComponentId         string `xml:"ClientSideComponentId,attr"`
						ClientSideComponentProperties string `xml:"ClientSideComponentProperties,attr"`
					} `xml:"CustomAction"`
				} `xml:"WebCustomActions"`
			} `xml:"CustomActions"`
			ComposedLook struct {
				Text           string `xml:",chardata"`
				Name           string `xml:"Name,attr"`
				ColorFile      string `xml:"ColorFile,attr"`
				FontFile       string `xml:"FontFile,attr"`
				BackgroundFile string `xml:"BackgroundFile,attr"`
				Version        string `xml:"Version,attr"`
			} `xml:"ComposedLook"`
			ApplicationLifecycleManagement struct {
				Text string `xml:",chardata"`
				Apps struct {
					Text string `xml:",chardata"`
					App  []struct {
						Text   string `xml:",chardata"`
						AppId  string `xml:"AppId,attr"`
						Action string `xml:"Action,attr"`
					} `xml:"App"`
				} `xml:"Apps"`
			} `xml:"ApplicationLifecycleManagement"`
			ClientSidePages struct {
				Text           string `xml:",chardata"`
				ClientSidePage struct {
					Text                 string `xml:",chardata"`
					PromoteAsNewsArticle string `xml:"PromoteAsNewsArticle,attr"`
					PromoteAsTemplate    string `xml:"PromoteAsTemplate,attr"`
					Overwrite            string `xml:"Overwrite,attr"`
					Layout               string `xml:"Layout,attr"`
					EnableComments       string `xml:"EnableComments,attr"`
					Title                string `xml:"Title,attr"`
					ThumbnailUrl         string `xml:"ThumbnailUrl,attr"`
					PageName             string `xml:"PageName,attr"`
					Header               struct {
						Text                   string `xml:",chardata"`
						Type                   string `xml:"Type,attr"`
						LayoutType             string `xml:"LayoutType,attr"`
						ShowTopicHeader        string `xml:"ShowTopicHeader,attr"`
						ShowPublishDate        string `xml:"ShowPublishDate,attr"`
						ShowBackgroundGradient string `xml:"ShowBackgroundGradient,attr"`
						TopicHeader            string `xml:"TopicHeader,attr"`
						AlternativeText        string `xml:"AlternativeText,attr"`
						Authors                string `xml:"Authors,attr"`
						AuthorByLineId         string `xml:"AuthorByLineId,attr"`
					} `xml:"Header"`
					Sections struct {
						Text    string `xml:",chardata"`
						Section struct {
							Text     string `xml:",chardata"`
							Order    string `xml:"Order,attr"`
							Type     string `xml:"Type,attr"`
							Controls struct {
								Text          string `xml:",chardata"`
								CanvasControl struct {
									Text                    string `xml:",chardata"`
									WebPartType             string `xml:"WebPartType,attr"`
									ControlId               string `xml:"ControlId,attr"`
									Order                   string `xml:"Order,attr"`
									Column                  string `xml:"Column,attr"`
									CanvasControlProperties struct {
										Text                  string `xml:",chardata"`
										CanvasControlProperty struct {
											Text  string `xml:",chardata"`
											Key   string `xml:"Key,attr"`
											Value string `xml:"Value,attr"`
										} `xml:"CanvasControlProperty"`
									} `xml:"CanvasControlProperties"`
								} `xml:"CanvasControl"`
							} `xml:"Controls"`
						} `xml:"Section"`
					} `xml:"Sections"`
				} `xml:"ClientSidePage"`
			} `xml:"ClientSidePages"`
			Header struct {
				Text          string `xml:",chardata"`
				Layout        string `xml:"Layout,attr"`
				MenuStyle     string `xml:"MenuStyle,attr"`
				ShowSiteTitle string `xml:"ShowSiteTitle,attr"`
			} `xml:"Header"`
			Footer struct {
				Text                string `xml:",chardata"`
				Enabled             string `xml:"Enabled,attr"`
				RemoveExistingNodes string `xml:"RemoveExistingNodes,attr"`
				BackgroundEmphasis  string `xml:"BackgroundEmphasis,attr"`
			} `xml:"Footer"`
		} `xml:"ProvisioningTemplate"`
	} `xml:"Templates"`
}
