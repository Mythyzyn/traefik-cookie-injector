# traefik-cookieinjector

🧁 A simple Traefik plugin middleware that injects `Secure`, `HttpOnly`, and `SameSite=Strict` into all `Set-Cookie` headers.

## ✅ Features

- Automatically patches all cookies in the response
- Lightweight and easy to use

## 📦 Installation

Enable the plugin in Traefik static config:

```yaml
experimental:
  plugins:
    cookieinjector:
      moduleName: "github.com/<ton-utilisateur>/traefik-cookieinjector"
      version: "v0.0.1"
