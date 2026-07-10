<div align="center">

```
███╗   ███╗██╗   ██╗████████╗███████╗██████╗ ███████╗██╗   ██╗
████╗ ████║██║   ██║╚══██╔══╝██╔════╝██╔══██╗██╔════╝██║   ██║
██╔████╔██║██║   ██║   ██║   █████╗  ██║  ██║█████╗  ██║   ██║
██║╚██╔╝██║██║   ██║   ██║   ██╔══╝  ██║  ██║██╔══╝  ╚██╗ ██╔╝
██║ ╚═╝ ██║╚██████╔╝   ██║   ███████╗██████╔╝███████╗ ╚████╔╝ 
╚═╝     ╚═╝ ╚═════╝    ╚═╝   ╚══════╝╚═════╝ ╚══════╝  ╚═══╝  

</div>

<br/>

<!-- ===== HERO IMAGE — replace the URL below with your banner ===== -->
<p align="center">
  <a href="https://github.com/tuusuario/turepo">
    <img src="https://amzn-images-public.s3.us-east-1.amazonaws.com/images-edit-github/banner-github-golangapi.png" 
         alt="Birria de Fútbol" 
         width="100%">
  </a>
</p>
<!-- ===== END HERO ===== -->

<p align="center">
  <strong>Inscripción online para torneos de fútbol</strong>
  <br/>
  <sub>Registro seguro con Supabase</sub>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white" alt="Go version">
  <img src="https://img.shields.io/badge/Supabase-REST-3ECF8E?logo=supabase&logoColor=white" alt="Supabase">
  <img src="https://img.shields.io/badge/license-MIT-yellow" alt="License">
</p>

---

## ✨ Características

- 📋 Formulario de inscripción con validación
- 🔐 Autenticación via Supabase (password grant + refresh token)
- 🍪 Sesiones persistentes con cookie `session_id`
- 🚫 Protección de rutas internas
- 🗄️ Almacenamiento en Supabase REST API

---

## 🚀 Stack

| Capa | Tecnología |
|------|-----------|
| **Lenguaje** | Go 1.21+ |
| **Base de datos** | Supabase (PostgreSQL via REST) |
| **Templates** | `html/template` |
| **Autenticación** | Supabase Auth |

---

## ⚙️ Variables de entorno

```env
APIKEY=tu_anon_key
ID=tu_project_id
SVC_PORT=8090
NODE_IP=0.0.0.0
email=usuario@ejemplo.com
password=tu_contraseña
```

---

## 🏃 Inicio rápido

```bash
git clone https://github.com/tuusuario/turepo.git
cd test
go build -o main . && ./main
```

Abre `http://localhost:8080` en tu navegador.

---

## 🧭 Rutas

| Ruta | Descripción |
|------|-------------|
| `/` | Página principal — crea cookie y redirige si ya estás registrado |
| `/submit` | Envía los datos del formulario a Supabase |
| `/notification` | Muestra mensajes dinámicos (solo acceso interno) |
| `/health` | Health check → `{"status":"UP"}` |

---

## 📁 Estructura

```
.
├── main.go                 # Punto de entrada
├── public/
│   ├── index.html          # Formulario de inscripción
│   └── notification.html   # Página de confirmación
├── Dockerfile
└── README.md
```

---

## 🔐 Seguridad

- Las rutas `/notification` requieren el query param `?internal=true`
- El token de Supabase se cachea con renovación automática via `refresh_token`
- Cookie `session_id`: HttpOnly, Secure, SameSite=Strict, 1 año de expiración

---

<div align="center">
  <sub>Hecho con ❤️ por <strong>MuteDev</strong></sub>
  <br/>
  <sub>© 2026 — Birria de Fútbol</sub>
</div>
