/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ProductOrderInfo struct {
	ProductId int64 `json:"product_id"`

	ProductName string `json:"product_name"`

	ProductPrice int64 `json:"product_price"`

	ProductQuantity int64 `json:"product_quantity"`

	EventDiscountId int64 `json:"event_discount_id,omitempty"`

	EventDiscountMaxQuantity int32 `json:"event_discount_max_quantity,omitempty"`
}