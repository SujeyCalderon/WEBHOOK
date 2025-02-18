package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GithubWebhookHanlder(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")
	signature := ctx.GetHeader("X-Hub-Signature-256")

	log.Printf("Webhook recibido: Evento=%s, DeliveryID=%s, Firma=%s", eventType, deliveryID, signature)

	payload, err := ctx.GetRawData()

	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	switch eventType {
	case "ping":
		handleGithubPingEvent(ctx)
	case "pull_request":
		handleGithubPullRequestEvent(ctx, payload)
	}
}
