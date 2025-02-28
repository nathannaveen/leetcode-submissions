func accountBalanceAfterPurchase(purchaseAmount int) int {
    x := (purchaseAmount + 5) / 10 * 10
    return 100 - x
}