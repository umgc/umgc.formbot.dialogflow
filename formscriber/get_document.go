package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getDocument(docid string) (get_doc string) {

	var t Tokenresponse
	var url = Oauth()
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &t)

	if err != nil {
		panic(err)
	}

	token := t.Access_token

	//RUle for tempaltes in google docs, the values must always be inside a table!
	url = "https://docs.googleapis.com/v1/documents/" + docid + "?fields=body(content/table/tableRows/tableCells/content/paragraph/elements/textRun/content)"
	method = "GET"

	client = &http.Client{}
	req, err = http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	get_doc = string(body)
	return get_doc

}

type Getdoc struct {
	Title string `json:"title"`
	Body  struct {
		Content []struct {
			EndIndex     int `json:"endIndex"`
			SectionBreak struct {
				SectionStyle struct {
					ColumnSeparatorStyle string `json:"columnSeparatorStyle"`
					ContentDirection     string `json:"contentDirection"`
					SectionType          string `json:"sectionType"`
				} `json:"sectionStyle"`
			} `json:"sectionBreak,omitempty"`
			StartIndex int `json:"startIndex,omitempty"`
			Paragraph  struct {
				Elements []struct {
					StartIndex int `json:"startIndex"`
					EndIndex   int `json:"endIndex"`
					TextRun    struct {
						Content   string `json:"content"`
						TextStyle struct {
							Underline bool `json:"underline"`
							FontSize  struct {
								Magnitude int    `json:"magnitude"`
								Unit      string `json:"unit"`
							} `json:"fontSize"`
						} `json:"textStyle"`
					} `json:"textRun,omitempty"`
					InlineObjectElement struct {
						InlineObjectID string `json:"inlineObjectId"`
						TextStyle      struct {
							FontSize struct {
								Magnitude int    `json:"magnitude"`
								Unit      string `json:"unit"`
							} `json:"fontSize"`
						} `json:"textStyle"`
					} `json:"inlineObjectElement,omitempty"`
				} `json:"elements"`
				ParagraphStyle struct {
					NamedStyleType string `json:"namedStyleType"`
					Direction      string `json:"direction"`
				} `json:"paragraphStyle"`
			} `json:"paragraph,omitempty"`
			Table struct {
				Rows      int `json:"rows"`
				Columns   int `json:"columns"`
				TableRows []struct {
					StartIndex int `json:"startIndex"`
					EndIndex   int `json:"endIndex"`
					TableCells []struct {
						StartIndex int `json:"startIndex"`
						EndIndex   int `json:"endIndex"`
						Content    []struct {
							StartIndex int `json:"startIndex"`
							EndIndex   int `json:"endIndex"`
							Paragraph  struct {
								Elements []struct {
									StartIndex int `json:"startIndex"`
									EndIndex   int `json:"endIndex"`
									TextRun    struct {
										Content   string `json:"content"`
										TextStyle struct {
										} `json:"textStyle"`
									} `json:"textRun"`
								} `json:"elements"`
								ParagraphStyle struct {
									NamedStyleType string `json:"namedStyleType"`
									Direction      string `json:"direction"`
								} `json:"paragraphStyle"`
							} `json:"paragraph"`
						} `json:"content"`
						TableCellStyle struct {
							RowSpan         int `json:"rowSpan"`
							ColumnSpan      int `json:"columnSpan"`
							BackgroundColor struct {
							} `json:"backgroundColor"`
							BorderBottom struct {
								Color struct {
									Color struct {
										RgbColor struct {
										} `json:"rgbColor"`
									} `json:"color"`
								} `json:"color"`
								Width struct {
									Magnitude int    `json:"magnitude"`
									Unit      string `json:"unit"`
								} `json:"width"`
								DashStyle string `json:"dashStyle"`
							} `json:"borderBottom"`
							PaddingLeft struct {
								Magnitude int    `json:"magnitude"`
								Unit      string `json:"unit"`
							} `json:"paddingLeft"`
							PaddingRight struct {
								Magnitude int    `json:"magnitude"`
								Unit      string `json:"unit"`
							} `json:"paddingRight"`
							PaddingTop struct {
								Magnitude int    `json:"magnitude"`
								Unit      string `json:"unit"`
							} `json:"paddingTop"`
							PaddingBottom struct {
								Magnitude int    `json:"magnitude"`
								Unit      string `json:"unit"`
							} `json:"paddingBottom"`
							ContentAlignment string `json:"contentAlignment"`
						} `json:"tableCellStyle"`
					} `json:"tableCells"`
					TableRowStyle struct {
						MinRowHeight struct {
							Unit string `json:"unit"`
						} `json:"minRowHeight"`
					} `json:"tableRowStyle"`
				} `json:"tableRows"`
				TableStyle struct {
					TableColumnProperties []struct {
						WidthType string `json:"widthType"`
						Width     struct {
							Magnitude float64 `json:"magnitude"`
							Unit      string  `json:"unit"`
						} `json:"width,omitempty"`
					} `json:"tableColumnProperties"`
				} `json:"tableStyle"`
			} `json:"table,omitempty"`
		} `json:"content"`
	} `json:"body"`
	DocumentStyle struct {
		Background struct {
			Color struct {
			} `json:"color"`
		} `json:"background"`
		PageNumberStart int `json:"pageNumberStart"`
		MarginTop       struct {
			Magnitude int    `json:"magnitude"`
			Unit      string `json:"unit"`
		} `json:"marginTop"`
		MarginBottom struct {
			Magnitude int    `json:"magnitude"`
			Unit      string `json:"unit"`
		} `json:"marginBottom"`
		MarginRight struct {
			Magnitude int    `json:"magnitude"`
			Unit      string `json:"unit"`
		} `json:"marginRight"`
		MarginLeft struct {
			Magnitude int    `json:"magnitude"`
			Unit      string `json:"unit"`
		} `json:"marginLeft"`
		PageSize struct {
			Height struct {
				Magnitude float64 `json:"magnitude"`
				Unit      string  `json:"unit"`
			} `json:"height"`
			Width struct {
				Magnitude float64 `json:"magnitude"`
				Unit      string  `json:"unit"`
			} `json:"width"`
		} `json:"pageSize"`
		MarginHeader struct {
			Magnitude int    `json:"magnitude"`
			Unit      string `json:"unit"`
		} `json:"marginHeader"`
		MarginFooter struct {
			Magnitude int    `json:"magnitude"`
			Unit      string `json:"unit"`
		} `json:"marginFooter"`
		UseCustomHeaderFooterMargins bool `json:"useCustomHeaderFooterMargins"`
	} `json:"documentStyle"`
	NamedStyles struct {
		Styles []struct {
			NamedStyleType string `json:"namedStyleType"`
			TextStyle2     struct {
				Bold            bool `json:"bold"`
				Italic          bool `json:"italic"`
				Underline       bool `json:"underline"`
				Strikethrough   bool `json:"strikethrough"`
				SmallCaps       bool `json:"smallCaps"`
				BackgroundColor struct {
				} `json:"backgroundColor"`
				ForegroundColor struct {
					Color struct {
						RgbColor struct {
						} `json:"rgbColor"`
					} `json:"color"`
				} `json:"foregroundColor"`
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
				WeightedFontFamily struct {
					FontFamily string `json:"fontFamily"`
					Weight     int    `json:"weight"`
				} `json:"weightedFontFamily"`
				BaselineOffset string `json:"baselineOffset"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle2 struct {
				NamedStyleType string `json:"namedStyleType"`
				Alignment      string `json:"alignment"`
				LineSpacing    int    `json:"lineSpacing"`
				Direction      string `json:"direction"`
				SpacingMode    string `json:"spacingMode"`
				SpaceAbove     struct {
					Unit string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Unit string `json:"unit"`
				} `json:"spaceBelow"`
				BorderBetween struct {
					Color struct {
					} `json:"color"`
					Width struct {
						Unit string `json:"unit"`
					} `json:"width"`
					Padding struct {
						Unit string `json:"unit"`
					} `json:"padding"`
					DashStyle string `json:"dashStyle"`
				} `json:"borderBetween"`
				BorderTop struct {
					Color struct {
					} `json:"color"`
					Width struct {
						Unit string `json:"unit"`
					} `json:"width"`
					Padding struct {
						Unit string `json:"unit"`
					} `json:"padding"`
					DashStyle string `json:"dashStyle"`
				} `json:"borderTop"`
				BorderBottom struct {
					Color struct {
					} `json:"color"`
					Width struct {
						Unit string `json:"unit"`
					} `json:"width"`
					Padding struct {
						Unit string `json:"unit"`
					} `json:"padding"`
					DashStyle string `json:"dashStyle"`
				} `json:"borderBottom"`
				BorderLeft struct {
					Color struct {
					} `json:"color"`
					Width struct {
						Unit string `json:"unit"`
					} `json:"width"`
					Padding struct {
						Unit string `json:"unit"`
					} `json:"padding"`
					DashStyle string `json:"dashStyle"`
				} `json:"borderLeft"`
				BorderRight struct {
					Color struct {
					} `json:"color"`
					Width struct {
						Unit string `json:"unit"`
					} `json:"width"`
					Padding struct {
						Unit string `json:"unit"`
					} `json:"padding"`
					DashStyle string `json:"dashStyle"`
				} `json:"borderRight"`
				IndentFirstLine struct {
					Unit string `json:"unit"`
				} `json:"indentFirstLine"`
				IndentStart struct {
					Unit string `json:"unit"`
				} `json:"indentStart"`
				IndentEnd struct {
					Unit string `json:"unit"`
				} `json:"indentEnd"`
				KeepLinesTogether   bool `json:"keepLinesTogether"`
				KeepWithNext        bool `json:"keepWithNext"`
				AvoidWidowAndOrphan bool `json:"avoidWidowAndOrphan"`
				Shading             struct {
					BackgroundColor struct {
					} `json:"backgroundColor"`
				} `json:"shading"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle3 struct {
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle3 struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle4 struct {
				Bold     bool `json:"bold"`
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle4 struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle5 struct {
				Bold            bool `json:"bold"`
				ForegroundColor struct {
					Color struct {
						RgbColor struct {
							Red   float64 `json:"red"`
							Green float64 `json:"green"`
							Blue  float64 `json:"blue"`
						} `json:"rgbColor"`
					} `json:"color"`
				} `json:"foregroundColor"`
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle5 struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle6 struct {
				ForegroundColor struct {
					Color struct {
						RgbColor struct {
							Red   float64 `json:"red"`
							Green float64 `json:"green"`
							Blue  float64 `json:"blue"`
						} `json:"rgbColor"`
					} `json:"color"`
				} `json:"foregroundColor"`
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle6 struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle7 struct {
				ForegroundColor struct {
					Color struct {
						RgbColor struct {
							Red   float64 `json:"red"`
							Green float64 `json:"green"`
							Blue  float64 `json:"blue"`
						} `json:"rgbColor"`
					} `json:"color"`
				} `json:"foregroundColor"`
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle7 struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle8 struct {
				Italic          bool `json:"italic"`
				ForegroundColor struct {
					Color struct {
						RgbColor struct {
							Red   float64 `json:"red"`
							Green float64 `json:"green"`
							Blue  float64 `json:"blue"`
						} `json:"rgbColor"`
					} `json:"color"`
				} `json:"foregroundColor"`
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle8 struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle9 struct {
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle9 struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Unit string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
			TextStyle10 struct {
				Italic          bool `json:"italic"`
				ForegroundColor struct {
					Color struct {
						RgbColor struct {
							Red   float64 `json:"red"`
							Green float64 `json:"green"`
							Blue  float64 `json:"blue"`
						} `json:"rgbColor"`
					} `json:"color"`
				} `json:"foregroundColor"`
				FontSize struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"fontSize"`
				WeightedFontFamily struct {
					FontFamily string `json:"fontFamily"`
					Weight     int    `json:"weight"`
				} `json:"weightedFontFamily"`
			} `json:"textStyle,omitempty"`
			ParagraphStyle struct {
				NamedStyleType string `json:"namedStyleType"`
				Direction      string `json:"direction"`
				SpaceAbove     struct {
					Unit string `json:"unit"`
				} `json:"spaceAbove"`
				SpaceBelow struct {
					Magnitude int    `json:"magnitude"`
					Unit      string `json:"unit"`
				} `json:"spaceBelow"`
				KeepLinesTogether bool `json:"keepLinesTogether"`
				KeepWithNext      bool `json:"keepWithNext"`
			} `json:"paragraphStyle,omitempty"`
		} `json:"styles"`
	} `json:"namedStyles"`
	RevisionID          string `json:"revisionId"`
	SuggestionsViewMode string `json:"suggestionsViewMode"`
	InlineObjects       struct {
		KixV4K3Xjjwmrt3 struct {
			ObjectID               string `json:"objectId"`
			InlineObjectProperties struct {
				EmbeddedObject struct {
					ImageProperties struct {
						ContentURI     string `json:"contentUri"`
						CropProperties struct {
						} `json:"cropProperties"`
					} `json:"imageProperties"`
					EmbeddedObjectBorder struct {
						Color struct {
							Color struct {
								RgbColor struct {
								} `json:"rgbColor"`
							} `json:"color"`
						} `json:"color"`
						Width struct {
							Unit string `json:"unit"`
						} `json:"width"`
						DashStyle     string `json:"dashStyle"`
						PropertyState string `json:"propertyState"`
					} `json:"embeddedObjectBorder"`
					Size struct {
						Height struct {
							Magnitude float64 `json:"magnitude"`
							Unit      string  `json:"unit"`
						} `json:"height"`
						Width struct {
							Magnitude float64 `json:"magnitude"`
							Unit      string  `json:"unit"`
						} `json:"width"`
					} `json:"size"`
					MarginTop struct {
						Magnitude int    `json:"magnitude"`
						Unit      string `json:"unit"`
					} `json:"marginTop"`
					MarginBottom struct {
						Magnitude int    `json:"magnitude"`
						Unit      string `json:"unit"`
					} `json:"marginBottom"`
					MarginRight struct {
						Magnitude int    `json:"magnitude"`
						Unit      string `json:"unit"`
					} `json:"marginRight"`
					MarginLeft struct {
						Magnitude int    `json:"magnitude"`
						Unit      string `json:"unit"`
					} `json:"marginLeft"`
				} `json:"embeddedObject"`
			} `json:"inlineObjectProperties"`
		} `json:"kix.v4k3xjjwmrt3"`
	} `json:"inlineObjects"`
	DocumentID string `json:"documentId"`
}
