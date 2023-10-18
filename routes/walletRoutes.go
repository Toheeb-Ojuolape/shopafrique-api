package routes

import (
	walletControllers "github.com/Toheeb-Ojuolape/shopafrique-api/controllers/wallet"
	middleware "github.com/Toheeb-Ojuolape/shopafrique-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func WalletRoutes(wallet fiber.Router) {
	wallet.Post("/fund-wallet", middleware.VerifyToken, walletControllers.FundWallet)
	wallet.Get("/transactions", middleware.VerifyToken, walletControllers.FetchTransactions)
}
