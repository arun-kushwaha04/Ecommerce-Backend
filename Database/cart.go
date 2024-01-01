package Database

import "errors"

var (
	ErrCantFindError = errors.New("can't find a product")
	ErrCantDecodeProduct = errors.New("can't find a product")
	ErrUserIdNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("can't update a user")
	ErrCantRemoveItemFromCart = errors.New("can't remove an item from cart")
	ErrCantGetItem = errors.New("can't get an item from cart")
	ErrCantBuyCartItem = errors.New("can't buy an item from cart")
)

// add product to cart
func AddProduct(){

}

// remove product from cart
func RemoveCartItem(){

}

// buy items from cart
func BuyItemsFromCart(){

}

// instant buy
func InstantBuy(){

}