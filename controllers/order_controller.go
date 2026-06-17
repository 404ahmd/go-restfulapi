package controllers

import (
	"go-restfulapi/config"
	"go-restfulapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var input models.CreateOrderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Validasi gagal: " + err.Error(),
		})
		return
	}

	// Cek customer ada
	var customer models.Customer
	if err := config.DB.First(&customer, input.CustomerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Customer tidak ditemukan",
		})
		return
	}

	// Mulai transaksi
	transaction := config.DB.Begin()
	if transaction.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal memulai transaksi",
		})
		return
	}

	// Buat order
	order := models.Order{
		UserID:     userID,
		CustomerID: input.CustomerID,
		Status:     "pending",
		Notes:      input.Notes,
		OrderedAt:  time.Now(),
	}

	if err := transaction.Create(&order).Error; err != nil {
		transaction.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal membuat order: " + err.Error(),
		})
		return
	}

	// Proses setiap item
	var totalPrice float64
	for _, itemReq := range input.Items {

		// Cek & lock produk
		var product models.Product
		if err := transaction.Set("gorm:query_option", "FOR UPDATE").
			First(&product, itemReq.ProductID).Error; err != nil {
			transaction.Rollback()
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Produk tidak ditemukan",
			})
			return
		}

		// ✅ Fix 1: Cek stok — bandingkan tipe yang sama (int vs int)
		if product.Stock < itemReq.Quantity {
			transaction.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Stok produk '" + product.Name + "' tidak mencukupi",
				"detail": gin.H{
					"requested": itemReq.Quantity,
					"available": product.Stock,
				},
			})
			return
		}

		// ✅ Fix 2: Rumus subtotal pakai * bukan +
		subTotal := product.Price * float64(itemReq.Quantity)
		totalPrice += subTotal

		// ✅ Fix 3: Syntax struct pakai {} bukan ()
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: itemReq.ProductID, // ← dari itemReq, bukan order
			Quantity:  itemReq.Quantity,
			UnitPrice: product.Price,
			Subtotal:  subTotal,
		}

		// Simpan order item
		if err := transaction.Create(&orderItem).Error; err != nil {
			transaction.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Gagal menyimpan item order: " + err.Error(),
			})
			return
		}

		// Kurangi stok produk
		if err := transaction.Model(&product).
			Update("stock", product.Stock-itemReq.Quantity).Error; err != nil {
			transaction.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Gagal mengupdate stok: " + err.Error(),
			})
			return
		}
	}

	// Update total price
	if err := transaction.Model(&order).
		Update("total_price", totalPrice).Error; err != nil {
		transaction.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal mengupdate total harga: " + err.Error(),
		})
		return
	}

	// Commit transaksi
	if err := transaction.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal commit transaksi: " + err.Error(),
		})
		return
	}

	// Ambil data lengkap dengan relasi
	var result models.Order
	config.DB.
		Preload("User").
		Preload("Customer").
		Preload("Items.Product").
		First(&result, order.ID)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Order berhasil dibuat",
		"data":    result,
	})
}