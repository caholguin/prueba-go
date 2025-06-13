# 📦 User API - CRUD en Go + PostgreSQL + Docker

---

## 🚀 Características principales

- CRUD completo de usuarios
- Arquitectura limpia: Controllers, Services, Repository
- Persistencia con PostgreSQL
- Despliegue sencillo con Docker y Docker Compose

---

## 🏗️ Estructura del proyecto
```
├── cmd/ --> Punto de entrada principal (main.go)
├── config/ --> Configuración de la base de datos
├── controllers/ --> Controladores HTTP (entrada/salida)
├── models/ --> Modelos de entidades
├── repository/ --> Acceso a datos (persistencia)
├── service/ --> Lógica de negocio
├── routes/ --> Definición de rutas y middlewares
├── db/ --> Scripts de inicialización de base de datos
├── Dockerfile --> Imagen de Docker para la aplicación
├── docker-compose.yml --> Orquestación de contenedores
├── go.mod / go.sum --> Dependencias de Go
└── README.md --> Documentación
```
---

## 🖥️ Requisitos previos

- Go 1.21 o superior (solo si lo ejecutas sin Docker)
- Docker
- Docker Compose

---

## 🚀 Cómo levantar el proyecto

#### 1️⃣ Clonar el repositorio

```bash
git clone https://github.com/caholguin/prueba-go
cd prueba-go
```

#### 2️⃣ Levantar los servicios
```
docker-compose up --build
```

 👉 Nota: Si durante el primer build falla descargando la imagen golang:1.21-alpine, ejecutar manualmente:
 ```
 docker pull golang:1.21-alpine
 ```

#### 3️⃣ Detener los servicios
```
docker-compose down
```
---
## 🌐 Endpoints disponibles

#### 1️⃣ Listar usuarios
- Descripción: Obtiene la lista de todos los usuarios.
- Método: GET
- URL: http://localhost:8080/api/users

#### 2️⃣ Crear usuario
- Descripción: Descripción: Crea un nuevo usuario.
- Método: POST
- URL: http://localhost:8080/api/users

Body (JSON):
```
{
  "name": "Juan Perez",
  "email": "juan.perez@example.com"
}
```

#### 3️⃣ Consultar usuario por ID
- Descripción: Consulta los datos de un usuario específico por su ID.
- Método: GET
- URL: http://localhost:8080/api/users/1

#### 4️⃣ Actualizar usuario
- Descripción: Actualiza la información de un usuario existente.
- Método: PUT
- URL: http://localhost:8080/api/users/1

Body (JSON):
```
{
  "name": "Juan Perez Actualizado",
  "email": "juan.perez@nuevoemail.com"
}
```

#### 5️⃣ Eliminar usuario
- Descripción: Elimina un usuario por su ID.
- Método: DELETE
- URL: http://localhost:8080/api/users/1

---

### Que aprendi

- Configuración de GORM con PostgreSQL.

---

### Mejoras futuras

- Tests unitarios e integración.

- Middleware de autenticación.

---

## ✅ Patrones aplicados 
##### 1️⃣ Repository Pattern

Para que el código que consulta la base de datos esté separado de la lógica de negocio.


##### 2️⃣ Service Layer Pattern

Para que el controlador solo reciba las peticiones HTTP y llame al servicio.
Las reglas de negocio están en el servicio, lo que hace el código más limpio, mantenible y fácil de testear.

---

### ✅ Arquitectura aplicada


## 🎯 ¿Por qué usé arquitectura en capas?

Aplicar arquitectura en capas permite separar responsabilidades, hacer el código más limpio, mantenible y escalable. Cada capa (`controllers/`, `services/`, `repository/`, `models/`) tiene una función clara, facilitando futuras modificaciones, pruebas unitarias y el crecimiento ordenado del proyecto.

