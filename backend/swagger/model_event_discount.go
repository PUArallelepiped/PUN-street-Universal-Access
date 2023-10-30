/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EventDiscount struct {

	DiscountId int64 `json:"discount_id"`

	DiscountName string `json:"discount_name"`

	DiscountDescription string `json:"discount_description"`

	DiscountMaxQuantity int64 `json:"discount_max_quantity"`

	ProductId int64 `json:"product_id"`
}