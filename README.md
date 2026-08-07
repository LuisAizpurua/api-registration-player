<div align="center">

```text
███╗   ███╗██╗   ██╗████████╗███████╗██████╗ ███████╗██╗   ██╗
████╗ ████║██║   ██║╚══██╔══╝██╔════╝██╔══██╗██╔════╝██║   ██║
██╔████╔██║██║   ██║   ██║   █████╗  ██║  ██║█████╗  ██║   ██║
██║╚██╔╝██║██║   ██║   ██║   ██╔══╝  ██║  ██║██╔══╝  ╚██╗ ██╔╝
██║ ╚═╝ ██║╚██████╔╝   ██║   ███████╗██████╔╝███████╗ ╚████╔╝
╚═╝     ╚═╝ ╚═════╝    ╚═╝   ╚══════╝╚═════╝ ╚══════╝  ╚═══╝
```

</div>

<br>

<p align="center">
  <a href="https://github.com/LuisAizpurua/test">
    <img src="https://amzn-images-public.s3.us-east-1.amazonaws.com/images-edit-github/banner-github-golangapi.png"
         alt="Birria de Fútbol"
         width="100%">
  </a>
</p>

<h3 align="center">Inscripción online para birria de fútbol</h3>

<p align="center">
  <a href="https://go.dev/doc/"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white"></a>
  <a href="https://supabase.com/docs"><img src="https://img.shields.io/badge/Supabase-REST-3ECF8E?logo=supabase&logoColor=white"></a>
  <a href="https://docs.docker.com/"><img src="https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white"></a>
  <a href="https://kubernetes.io/docs/"><img src="https://img.shields.io/badge/Kubernetes-326CE5?logo=kubernetes&logoColor=white"></a>
  <a href="https://docs.github.com/actions"><img src="https://img.shields.io/badge/GitHub_Actions-2088FF?logo=githubactions&logoColor=white"></a>
  <img src="https://img.shields.io/badge/license-MIT-yellow">
</p>


---

### 👨🏻‍💻 Instalaciones

| Herramienta | Versión | Propósito |
|-------------|---------|-----------|
| [Go](https://go.dev/dl/) | 1.21+ | Compilar y ejecutar la aplicación |
| [Docker Desktop](https://www.docker.com/products/docker-desktop/) | Última | Contenerizar la aplicación |
| [Minikube](https://minikube.sigs.k8s.io/docs/start/) | Última | Cluster Kubernetes local |
| [Kubectl](https://kubernetes.io/docs/tasks/tools/) | Última | CLI para gestionar Kubernetes desde terminal|
| [ArgoCD CLI](https://argo-cd.readthedocs.io/en/stable/cli_installation/) | Última | Interaccion desde terminal con la UI argocd |

> **Instalar ArgoCD en Minikube:**
> ```bash
> kubectl create namespace argocd
> kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
> # Exponer el servidor via NodePort
> kubectl patch svc argocd-server -n argocd -p '{"spec":{"type":"NodePort"}}'
> # Obtener contraseña inicial del admin
> kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
> ```

---


### 🚀 Stack tecnológico 

| Capa | Tecnología |
|------|-----------|
| **Lenguaje** | Go 1.21+ |
| **Base de datos** | supabase (PostgreSQL vía REST) |
| **Templates** | `html/template` |
| **Docker** |  contenedor app |
| **Minikube** |  clúster de kubernetes de un solo nodo local |
| **Kubernetes** |  orquestador de contenedores |
| **Argocd** |  entrega continua `CD` (GitOps) |
| **GitHub Actions** |  integracion continua `CI` |
---

### ⚙️ Variables de entorno

| Variable | Descripción |
|----------|-------------|
| `APIKEY` | Anon key de Supabase (Configuración → API) |
| `ID` | ID del proyecto Supabase |
| `SVC_PORT` | Puerto del servicio en Kubernetes (opcional) |
| `NODE_IP` | IP del nodo Kubernetes (opcional) |
| `EMAIL` | Correo usado para autenticación en Supabase Auth |
| `PASSWORD` | Contraseña del usuario en Supabase Auth |


---

### 🏃 Inicio rápido

```bash
# 1. Clonar el repositorio
git clone https://github.com/tuusuario/turepo.git

# 2. Configurar variables de entorno
cp .env.example .env   # y completar con tus datos de Supabase

# 3a. Ejecutar directamente (sin compilar)
go run .

# 3b. O compilar y ejecutar
go build -o main . && ./main
```

Abre `http://localhost:8080` en tu navegador.

---

### 🧭 Rutas

| Ruta | Descripción |
|------|-------------|
| `/` | Página principal — muestra formulario, maneja cookie y redirige si ya estás registrado |
| `/submit` | Envía los datos del formulario a Supabase |
| `/notification` | Muestra mensajes de inscripcion (solo acceso interno) |
| `/health` | Health check |

---

### 🔐 Seguridad

- **Middleware de sesión** — middleware global que verifica la cookie `session_id`, consulta la existencia del usuario en Supabase y redirige según su estado (evita doble registro o acceso no autorizado).
- **Token cacheado con renovación** — `fetchToken` almacena el `access_token` y lo reutiliza mientras no expire. Cuando expira, lo renueva automáticamente mediante `refresh_token` sin reautenticar.
- **Cookie `session_id`** — configurada con las flags `HttpOnly`, `Secure`, `SameSite=Strict` y expiración de 1 año.

---

<div align="center">
  <sub>© 2026 — Birria de Fútbol</sub>
</div>