# 💼 Desafio: ConfigManager Singleton

## 📌 Objetivo

Criar uma estrutura `ConfigManager` que será responsável por **carregar e fornecer configurações da aplicação** (como se estivesse lendo de um `.env`, `.json` ou similar).  
Essa estrutura deve seguir o padrão **Singleton**, garantindo que:

- Apenas **uma instância** da configuração seja criada.
- Todas as partes da aplicação acessem a **mesma instância**.

---

## 🔹 Requisitos Funcionais
    ✅ A implementação deve ser thread-safe usando sync.Once.
    ✅ A estrutura deve permitir acesso global à instância.
    ✅ A leitura das configurações deve simular um "custo" de carregamento (ex: um fmt.Println("Lendo config...")).

```go
map[string]string{
    "DB_HOST": "localhost",
    "DB_PORT": "5432",
    "APP_ENV": "development",
}
