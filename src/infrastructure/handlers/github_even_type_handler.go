package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	domain "pull-request-checker/src/domain/value_objects"

	"github.com/gin-gonic/gin"
)

func handleGithubPingEvent(ctx *gin.Context) {
	log.Println("Ping recibido desde GitHub")
	ctx.JSON(http.StatusOK, gin.H{"status": "Pong!"})
}


func handleGithubPullRequestEvent(ctx *gin.Context, payload []byte) {
	var eventPayload domain.PullRequestEventPayload
	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("Error al deserializar el payload del pull request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el payload del pull request"})
		return
	}

	// Log del evento recibido
	log.Printf(
		"Evento Pull Request recibido: Acción=%s, PR Título='%s', Rama Base='%s', Repositorio='%s'",
		eventPayload.Action, eventPayload.PullRequest.Title, eventPayload.PullRequest.Base.Ref, eventPayload.Repository.FullName)


	// Filtrar solo pull requests con acción "closed"
	if eventPayload.Action == "closed" {
		fmt.Println("=== Pull Request Cerrado ===")
		fmt.Printf(" ->  Destino (Base Branch): %s\n", eventPayload.PullRequest.Base.Ref)
		fmt.Printf(" <-  Origen (Head Branch): %s\n", eventPayload.PullRequest.Head.Ref)
		fmt.Printf(" Usuario: %s\n", eventPayload.PullRequest.User.Login)
		fmt.Printf(" Repositorio: %s\n", eventPayload.Repository.FullName)
		fmt.Println("===========================")
	}

	// Verificar si el PR se dirige a la rama "develop"
	mainBranch := "develop"
	if eventPayload.PullRequest.Base.Ref == mainBranch {
		log.Printf("¡Pull Request a la rama '%s' detectado en el repositorio '%s'!", mainBranch, eventPayload.Repository.FullName)
		fmt.Printf("Pull Request detectado en la rama %s!\n", mainBranch)
	} else {
		log.Printf(
			"Pull Request detectado, pero no dirigido a la rama '%s'. Rama base: '%s'",
			mainBranch, eventPayload.PullRequest.Base.Ref)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Evento Pull Request recibido y procesado"})

}