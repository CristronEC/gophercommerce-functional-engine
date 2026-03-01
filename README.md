# 📦 GopherCommerce – Sistema de E-Commerce en Go

## 📖 Descripción General

GopherCommerce es un sistema backend de comercio electrónico desarrollado en el lenguaje **Go**.  
El proyecto implementa servicios web REST que permiten la gestión de productos, creación de carritos de compra y simulación de pagos, siguiendo una arquitectura modular organizada por dominios.  

El sistema funciona completamente en memoria y expone endpoints HTTP que pueden ser consumidos mediante herramientas como Postman u otros clientes REST.

---

## 🚀 Principales Funcionalidades

### 1️⃣ Gestión de Productos (Catalog)
- Listar productos disponibles.
- Obtener información de cada producto (ID, nombre, precio y stock).
- Simulación de inventario en memoria.

### 2️⃣ Gestión de Carrito (Cart)
- Crear un carrito de compras.
- Agregar productos al carrito.
- Visualizar contenido del carrito.
- Marcar el carrito como pagado (checkout).

### 3️⃣ Simulación de Pago (Payment)
- Procesamiento simulado de pagos.
- Generación de un estado de pago (“approved”).
- Integración con el proceso de checkout.

### 4️⃣ Servicios Web REST
El sistema expone los siguientes endpoints HTTP:

- `GET /products`
- `POST /cart`
- `POST /cart/{id}/items`
- `POST /cart/{id}/checkout`

Los servicios reciben y devuelven datos en formato JSON.

---

## 🎯 Objetivo del Programa

El objetivo principal del programa es desarrollar un sistema de e-commerce backend utilizando el lenguaje Go, implementando correctamente la generación de servicios web REST, una arquitectura modular y la estructura básica de un flujo comercial (catálogo → carrito → pago).  

El proyecto busca demostrar:
- Desarrollo de servicios web.
- Organización del sistema por dominios.
- Manejo de estructuras de datos.
- Simulación de reglas de negocio.
- Separación de responsabilidades.

---

## 👥 Datos del Grupo

**Integrantes:**  
- Nombre 1  
- Nombre 2  
- Nombre 3  

**Carrera:** Ingeniería de Software  
**Materia:** (Nombre de la materia)  
**Docente:** (Nombre del docente)

---

## 📅 Fecha de Entrega

**01 de marzo de 2026**  
*(Reemplazar por la fecha real de entrega si es necesario)*

---
