package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdNotValid     = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductItem() {}

func RemoveCartItem() {}

func BuyItemFromCart() {}

func InstantBuyer() {}
