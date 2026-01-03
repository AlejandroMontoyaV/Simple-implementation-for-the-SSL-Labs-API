package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Estructura de la respuesta JSON de la API de SSL Labs
type sslResponse struct {
	Status    string `json:"status"`
	Endpoints []struct {
		Grade string `json:"grade"`
	} `json:"endpoints"`
}

func analyzeDomain(domain string) (string, error) {
	// Validar que el dominio no esté vacío
	if domain == "" {
		return "", errors.New("domain cannot be empty")
	}

	// StartNew=on solo se usa la primera vez para iniciar un nuevo análisis
	startNew := true

	// Se crea un cliente http con 10 segundos como maximo tiempo de espera
	client := &http.Client{Timeout: 10 * time.Second}

	//
	for {
		// Construimos la URL de la solicitud
		url := fmt.Sprintf("%sanalyze?host=%s&publish=off", entryPoint, domain)
		if startNew {
			url += "&startNew=on"
			startNew = false
		}

		resp, err := client.Get(url)
		// Manejamos errores en la solicitud HTTP
		if err != nil {
			return "", err
		}

		// Decodificamos la respuesta JSON
		var result sslResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		resp.Body.Close()

		if err != nil {
			return "", err
		}

		// Verificamos el estado del análisis
		switch result.Status {
		case "READY":
			if len(result.Endpoints) > 0 {
				return result.Endpoints[0].Grade, nil
			}
			return "", errors.New("no endpoints found")

		case "ERROR":
			return "", errors.New("analysis failed")

		// Con esto cubrimos IN_PROGRESS y cualquier otro estado
		default:
			fmt.Println("Analysis in progress, waiting... \nStatus:", result.Status)
			// Esperamos 10 segundos antes de volver a consultar el estado
			time.Sleep(20 * time.Second)

		}
	}
}
