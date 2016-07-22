package sdk_actions

type RMA struct {
    DiscountAmt int `json:"discount_amt"`
    Items       []ReturnItem `json:"items"` //ReturnItem struct defined in callback_items_returned.go
    ProductAmt        int    `json:"product_amt"`
    RefundAction      string `json:"refund_action"`
    RefundAmt         int    `json:"refund_amt"`
    RetailopsReturnID int    `json:"retailops_return_id"`
    RetailopsRmaID    string `json:"retailops_rma_id"`
    ShippingAmt       int    `json:"shipping_amt"`
    SubtotalAmt       int    `json:"subtotal_amt"`
    TaxAmt            int    `json:"tax_amt"`
}
