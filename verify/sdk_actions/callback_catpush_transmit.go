package sdk_actions

import (
  "github.com/gudtech/scamp-go/scamp"
  "encoding/json"
)

type CatalogPushV1Input struct {
    Action string `json:"action"`
	Data   struct {
		Channel struct {
			ID     int `json:"id"`
			Params struct {
				AppKey             string `json:"appKey"`
				BreakdownInventory int    `json:"breakdown_inventory"`
				Tenant             string `json:"tenant"`
			} `json:"params"`
		} `json:"channel"`
		ClientID int `json:"client_id"`
		Feed     struct {
			Data struct {
				Product []struct {
					FeedData struct {
						Catalogs []struct {
							Catalog struct {
								Catalog_ID string `json:"Catalog ID"`
								Content    struct {
									Product_Images []struct {
										Product_Image struct {
											Alt_Text  string `json:"Alt Text"`
											Image_URL string `json:"Image URL"`
											Label     string `json:"Label"`
										} `json:"Product Image"`
									} `json:"Product Images"`
									Product_Name string `json:"Product Name"`
								} `json:"Content"`
								Price struct {
									Price string `json:"Price"`
								} `json:"Price"`
								Product_Categories []struct {
									Category_Code_s_ string `json:"Category Code(s)"`
								} `json:"Product Categories"`
								Status struct {
									Is_Active string `json:"Is Active"`
								} `json:"Status"`
							} `json:"Catalog"`
						} `json:"Catalogs"`
						Global struct {
							Extras []struct {
								Extra struct {
									Attribute_Name   string        `json:"Attribute Name"`
									Attribute_Values []interface{} `json:"Attribute Values"`
									Data_Type        string        `json:"Data Type"`
									Input_Type       string        `json:"Input Type"`
									Value_Type       string        `json:"Value Type"`
								} `json:"Extra"`
							} `json:"Extras"`
							General struct {
								Content struct {
									Full_Description string `json:"Full Description"`
									Product_Images   []struct {
										Product_Image struct {
											Alt_Text  string `json:"Alt Text"`
											Image_URL string `json:"Image URL"`
											Label     string `json:"Label"`
										} `json:"Product Image"`
									} `json:"Product Images"`
									Product_Name      string `json:"Product Name"`
									Short_Description string `json:"Short Description"`
								} `json:"Content"`
								Cost struct {
									Value string `json:"Value"`
								} `json:"Cost"`
								Pricing struct {
									Price string `json:"Price"`
								} `json:"Pricing"`
								Product_Code string `json:"Product Code"`
								Product_Type struct {
									Name string `json:"Name"`
								} `json:"Product Type"`
							} `json:"General"`
							Options []struct {
								FeedData struct {
									Variant_Options []struct {
										Option_Attribute struct {
											Attribute_Name  string `json:"Attribute Name"`
											Attribute_Value string `json:"Attribute Value"`
											Data_Type       string `json:"Data Type"`
										} `json:"Option Attribute"`
									} `json:"Variant Options"`
									Variation_Product_Code string `json:"Variation Product Code"`
								} `json:"feed_data"`
								ObjectData struct {
									ID         string `json:"id"`
									SourceType string `json:"source_type"`
								} `json:"object_data"`
							} `json:"Options"`
							Properties []struct {
								Property struct {
									Attribute_Name   string `json:"Attribute Name"`
									Attribute_Values []struct {
										Attribute_Value_s_ string `json:"Attribute Value(s)"`
									} `json:"Attribute Values"`
									Data_Type       string `json:"Data Type"`
									Input_Type      string `json:"Input Type"`
									Is_Multi_Select string `json:"Is Multi Select"`
									Value_Type      string `json:"Value Type"`
								} `json:"Property"`
							} `json:"Properties"`
							Shipping struct {
								Weight struct {
									Unit  string `json:"Unit"`
									Value string `json:"Value"`
								} `json:"Weight"`
							} `json:"Shipping"`
						} `json:"Global"`
						Publishing struct {
							Publish_Set_Code string `json:"Publish Set Code"`
							State            string `json:"State"`
						} `json:"Publishing"`
					} `json:"feed_data"`
					ObjectData struct {
						ID         string `json:"id"`
						SourceType string `json:"source_type"`
					} `json:"object_data"`
				} `json:"Product"`
			} `json:"data"`
		} `json:"feed"`
	} `json:"data"`
	Headers struct {
		ClientID int    `json:"client_id"`
		Ticket   string `json:"ticket"`
	} `json:"headers"`
	Version int `json:"version"`
}

type CatalogPushV1Output struct {
    Action       string `json:"action"`
	CatalogItems []CatalogItem `json:"catalog_items"`
	ChannelInfo struct {
		ID int `json:"id"`
	} `json:"channel_info"`
	ClientID             int    `json:"client_id"`
	IntegrationAuthToken string `json:"integration_auth_token"`
	Version              int    `json:"version"`
}

type CatalogItem struct {
    ItemInfo struct {
        ProductID      string `json:"product_id"`
        SkuNumber      string `json:"sku_number"`
        SkugroupNumber string `json:"skugroup_number"`
    } `json:"item_info"`
    ItemType       string   `json:"item_type"`
    TemplateOutput struct{} `json:"template_output"`
}

func CatalogPushV1(msg *scamp.Message, client *scamp.Client) {
    scamp.Info.Printf("incoming: %s", string(msg.Bytes()))
    var input CatalogPushV1Input

    err := json.Unmarshal(msg.Bytes(), &input)
    if err != nil {
        scamp.Info.Printf("Input Data Error: %s ", input)
    }

    //TODO: need to munge actual input data to output format for sdk
    var output CatalogPushV1Output

    respMsg := scamp.NewResponseMessage()
    respMsg.WriteJson(output)
    respMsg.SetRequestId(msg.RequestId)

    _,err = client.Send(respMsg)
    if err != nil {
      return
    }
}
