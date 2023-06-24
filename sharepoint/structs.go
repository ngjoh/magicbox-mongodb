package sharepoint

import "encoding/xml"

// Provisioning was generated 2023-06-22 12:04:09 by https://xml-to-go.github.io/ in Ukraine.
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
					User []struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name,attr"`
					} `xml:"User"`
				} `xml:"AdditionalMembers"`
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
							NavigationNode []struct {
								Text           string `xml:",chardata"`
								Title          string `xml:"Title,attr"`
								URL            string `xml:"Url,attr"`
								IsExternal     string `xml:"IsExternal,attr"`
								NavigationNode []struct {
									Text       string `xml:",chardata"`
									Title      string `xml:"Title,attr"`
									URL        string `xml:"Url,attr"`
									IsExternal string `xml:"IsExternal,attr"`
								} `xml:"NavigationNode"`
							} `xml:"NavigationNode"`
						} `xml:"NavigationNode"`
					} `xml:"StructuralNavigation"`
				} `xml:"CurrentNavigation"`
			} `xml:"Navigation"`
			SiteFields struct {
				Text  string `xml:",chardata"`
				Field []struct {
					Text               string `xml:",chardata"`
					ID                 string `xml:"ID,attr"`
					Name               string `xml:"Name,attr"`
					SourceID           string `xml:"SourceID,attr"`
					Required           string `xml:"Required,attr"`
					Hidden             string `xml:"Hidden,attr"`
					ShowInFileDlg      string `xml:"ShowInFileDlg,attr"`
					ShowInNewForm      string `xml:"ShowInNewForm,attr"`
					ShowInEditForm     string `xml:"ShowInEditForm,attr"`
					DisplaceOnUpgrade  string `xml:"DisplaceOnUpgrade,attr"`
					StaticName         string `xml:"StaticName,attr"`
					Type               string `xml:"Type,attr"`
					DisplayName        string `xml:"DisplayName,attr"`
					Sortable           string `xml:"Sortable,attr"`
					Mult               string `xml:"Mult,attr"`
					Group              string `xml:"Group,attr"`
					Sealed             string `xml:"Sealed,attr"`
					AllowDeletion      string `xml:"AllowDeletion,attr"`
					ReadOnly           string `xml:"ReadOnly,attr"`
					ShowInDisplayForm  string `xml:"ShowInDisplayForm,attr"`
					ShowInListSettings string `xml:"ShowInListSettings,attr"`
					Viewable           string `xml:"Viewable,attr"`
				} `xml:"Field"`
			} `xml:"SiteFields"`
			Lists struct {
				Text         string `xml:",chardata"`
				ListInstance []struct {
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
					EnableAttachments      string `xml:"EnableAttachments,attr"`
					DefaultDisplayFormUrl  string `xml:"DefaultDisplayFormUrl,attr"`
					DefaultEditFormUrl     string `xml:"DefaultEditFormUrl,attr"`
					DefaultNewFormUrl      string `xml:"DefaultNewFormUrl,attr"`
					ImageUrl               string `xml:"ImageUrl,attr"`
					IrmExpire              string `xml:"IrmExpire,attr"`
					IrmReject              string `xml:"IrmReject,attr"`
					IsApplicationList      string `xml:"IsApplicationList,attr"`
					ValidationFormula      string `xml:"ValidationFormula,attr"`
					ValidationMessage      string `xml:"ValidationMessage,attr"`
					EnableFolderCreation   string `xml:"EnableFolderCreation,attr"`
					ContentTypesEnabled    string `xml:"ContentTypesEnabled,attr"`
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
							TabularView       string `xml:"TabularView,attr"`
							RecurrenceRowset  string `xml:"RecurrenceRowset,attr"`
							MobileUrl         string `xml:"MobileUrl,attr"`
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
								GroupBy struct {
									Text       string `xml:",chardata"`
									Collapse   string `xml:"Collapse,attr"`
									GroupLimit string `xml:"GroupLimit,attr"`
									FieldRef   struct {
										Text string `xml:",chardata"`
										Name string `xml:"Name,attr"`
									} `xml:"FieldRef"`
								} `xml:"GroupBy"`
								Where struct {
									Text string `xml:",chardata"`
									Eq   struct {
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
								} `xml:"Where"`
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
							ViewType2       string `xml:"ViewType2"`
							ViewData        struct {
								Text     string `xml:",chardata"`
								FieldRef []struct {
									Text string `xml:",chardata"`
									Name string `xml:"Name,attr"`
									Type string `xml:"Type,attr"`
								} `xml:"FieldRef"`
							} `xml:"ViewData"`
						} `xml:"View"`
					} `xml:"Views"`
					FieldRefs struct {
						Text     string `xml:",chardata"`
						FieldRef []struct {
							Text        string `xml:",chardata"`
							ID          string `xml:"ID,attr"`
							Name        string `xml:"Name,attr"`
							DisplayName string `xml:"DisplayName,attr"`
							Required    string `xml:"Required,attr"`
						} `xml:"FieldRef"`
					} `xml:"FieldRefs"`
					Fields struct {
						Text  string `xml:",chardata"`
						Field []struct {
							Text                  string `xml:",chardata"`
							ID                    string `xml:"ID,attr"`
							DisplayName           string `xml:"DisplayName,attr"`
							Description           string `xml:"Description,attr"`
							Name                  string `xml:"Name,attr"`
							Format                string `xml:"Format,attr"`
							IsolateStyles         string `xml:"IsolateStyles,attr"`
							RichText              string `xml:"RichText,attr"`
							AppendOnly            string `xml:"AppendOnly,attr"`
							RichTextMode          string `xml:"RichTextMode,attr"`
							Title                 string `xml:"Title,attr"`
							Type                  string `xml:"Type,attr"`
							StaticName            string `xml:"StaticName,attr"`
							SourceID              string `xml:"SourceID,attr"`
							ColName               string `xml:"ColName,attr"`
							RowOrdinal            string `xml:"RowOrdinal,attr"`
							ClientSideComponentId string `xml:"ClientSideComponentId,attr"`
							FillInChoice          string `xml:"FillInChoice,attr"`
							CustomFormatter       string `xml:"CustomFormatter,attr"`
							FriendlyDisplayFormat string `xml:"FriendlyDisplayFormat,attr"`
							List                  string `xml:"List,attr"`
							UserDisplayOptions    string `xml:"UserDisplayOptions,attr"`
							UserSelectionMode     string `xml:"UserSelectionMode,attr"`
							UserSelectionScope    string `xml:"UserSelectionScope,attr"`
							ColName2              string `xml:"ColName2,attr"`
							RowOrdinal2           string `xml:"RowOrdinal2,attr"`
							Sealed                string `xml:"Sealed,attr"`
							Group                 string `xml:"Group,attr"`
							AllowDeletion         string `xml:"AllowDeletion,attr"`
							ShowInNewForm         string `xml:"ShowInNewForm,attr"`
							ShowInEditForm        string `xml:"ShowInEditForm,attr"`
							ShowInDisplayForm     string `xml:"ShowInDisplayForm,attr"`
							ShowInViewForms       string `xml:"ShowInViewForms,attr"`
							ShowInListSettings    string `xml:"ShowInListSettings,attr"`
							ShowInVersionHistory  string `xml:"ShowInVersionHistory,attr"`
							Required              string `xml:"Required,attr"`
							Hidden                string `xml:"Hidden,attr"`
							CanToggleHidden       string `xml:"CanToggleHidden,attr"`
							Indexed               string `xml:"Indexed,attr"`
							ReadOnly              string `xml:"ReadOnly,attr"`
							Viewable              string `xml:"Viewable,attr"`
							JSON                  string `xml:"Json,attr"`
							WebId                 string `xml:"WebId,attr"`
							ShowField             string `xml:"ShowField,attr"`
							Version               string `xml:"Version,attr"`
							Mult                  string `xml:"Mult,attr"`
							Sortable              string `xml:"Sortable,attr"`
							CHOICES               struct {
								Text   string   `xml:",chardata"`
								CHOICE []string `xml:"CHOICE"`
							} `xml:"CHOICES"`
							Default       string `xml:"Default"`
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
						} `xml:"Field"`
					} `xml:"Fields"`
					FieldDefaults struct {
						Text         string `xml:",chardata"`
						FieldDefault struct {
							Text      string `xml:",chardata"`
							FieldName string `xml:"FieldName,attr"`
						} `xml:"FieldDefault"`
					} `xml:"FieldDefaults"`
					PropertyBagEntries struct {
						Text             string `xml:",chardata"`
						PropertyBagEntry struct {
							Text      string `xml:",chardata"`
							Key       string `xml:"Key,attr"`
							Value     string `xml:"Value,attr"`
							Overwrite string `xml:"Overwrite,attr"`
						} `xml:"PropertyBagEntry"`
					} `xml:"PropertyBagEntries"`
				} `xml:"ListInstance"`
			} `xml:"Lists"`
			Features struct {
				Text        string `xml:",chardata"`
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
						Section []struct {
							Text               string `xml:",chardata"`
							Order              string `xml:"Order,attr"`
							Type               string `xml:"Type,attr"`
							BackgroundEmphasis string `xml:"BackgroundEmphasis,attr"`
							Controls           struct {
								Text          string `xml:",chardata"`
								CanvasControl []struct {
									Text                    string `xml:",chardata"`
									WebPartType             string `xml:"WebPartType,attr"`
									JsonControlData         string `xml:"JsonControlData,attr"`
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
				Text      string `xml:",chardata"`
				Layout    string `xml:"Layout,attr"`
				MenuStyle string `xml:"MenuStyle,attr"`
			} `xml:"Header"`
			Footer struct {
				Text                string `xml:",chardata"`
				Enabled             string `xml:"Enabled,attr"`
				RemoveExistingNodes string `xml:"RemoveExistingNodes,attr"`
				Layout              string `xml:"Layout,attr"`
				BackgroundEmphasis  string `xml:"BackgroundEmphasis,attr"`
				FooterLinks         struct {
					Text       string `xml:",chardata"`
					FooterLink struct {
						Text        string `xml:",chardata"`
						DisplayName string `xml:"DisplayName,attr"`
						URL         string `xml:"Url,attr"`
					} `xml:"FooterLink"`
				} `xml:"FooterLinks"`
			} `xml:"Footer"`
		} `xml:"ProvisioningTemplate"`
	} `xml:"Templates"`
}
