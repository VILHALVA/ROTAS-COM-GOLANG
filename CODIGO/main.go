package main

import (
    "net/http"
    "path"
)

func main() {
    // Diretório onde os arquivos estáticos estão localizados
    staticDir := "public"

    // Configurar os handlers para as rotas
    http.HandleFunc("/", makeHandler(staticDir, "home.html"))
    http.HandleFunc("/curiosidades", makeHandler(staticDir, "curiosidades.html"))
    http.HandleFunc("/sobre", makeHandler(staticDir, "sobre.html"))

    // Configurar o servidor de arquivos estáticos
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(staticDir))))

    // Iniciar o servidor na porta 8080
    http.ListenAndServe(":8080", nil)
}

// Função auxiliar para criar handlers
func makeHandler(staticDir, page string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Redirecionar para a página HTML correspondente no diretório de arquivos estáticos
        http.ServeFile(w, r, path.Join(staticDir, page))
    }
}
