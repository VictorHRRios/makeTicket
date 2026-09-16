# makeTicket

CLI pequeño escrito en Go para crear y gestionar carpetas de tickets/PMs. Sirve para generar automáticamente la estructura de archivos de un ticket (carpeta semanal, plantilla del PM, metadatos) y para abrir rápidamente el ticket más reciente desde cualquier carpeta.

Puedes guardar el proyecto donde prefieras: dentro de un OneDrive sincronizado, en el disco local, en una unidad de red, etc. La aplicación no depende de una ubicación fija.

---

## Instalación

Desde la raíz del proyecto:

```bash
go build -o makeTicket .
```

Esto genera el binario `makeTicket` (ya está ignorado en `.gitignore`).

Opcionalmente, puedes agregarlo al `PATH` o ejecutarlo directamente:

```bash
# Instalar para poder usarlo donde sea (se instala como makeTicket)
go install .
# Opcionalmente puedes hacer un alias para el comando como el alas mt
```

```bash
./makeTicket help
```

---

## Configuración (`config.json`)

La aplicación lee su configuración desde un archivo JSON global que contiene los proyectos, las versiones y el ticket actual.

### Ubicación del archivo

Busca `config.json` en este orden:

1. **Directorio de configuración del usuario:**
   - Linux / macOS: `~/.config/makeTicket/config.json`
   - Windows: `%APPDATA%\makeTicket\config.json`
2. **Fallback:** `config.json` en la carpeta desde donde ejecutes el comando.

> Si quieres usar el mismo catalogo de proyectos en cualquier parte, crea el archivo en la ubicación global del usuario. Si prefieres una configuración por carpeta o por máquina, basta con dejar un `config.json` junto al binario.

### Estructura

```json
{
    "projects": {
        "nombre-del-proyecto": {
            "name": "Nombre legible del proyecto",
            "svnProjectURL": "https://svn...",
            "programs": [
                {
                    "name": "Programa",
                    "typeOfFile": "Tipo",
                    "version": 1.0
                }
            ]
        }
    },
    "versions": {
        "nombre-de-version": {
            "evidenceURL": "https://...",
            "svnURL": "https://..."
        }
    },
    "current": "1234567",
    "cloudURL": "https://drive-or-cloud..."
}
```

| Campo | Descripción |
|-------|-------------|
| `projects` | Catálogo de proyectos disponibles al crear un ticket. |
| `versions` | Catálogo de versiones/líneas de trabajo disponibles. |
| `current` | Último ticket creado. Se actualiza automáticamente cada vez que ejecutas `init`. Es lo que permite abrir un ticket estando en cualquier carpeta. |
| `cloudURL` | URL base que se inyecta en la plantilla del PM para enlaces a vídeos de prueba. |

### Primer arranque

Si aún no existe el archivo global, crea la carpeta y el archivo manualmente. Por ejemplo, en Linux/macOS:

```bash
mkdir -p ~/.config/makeTicket
cat > ~/.config/makeTicket/config.json << 'EOF'
{
    "projects": {},
    "versions": {},
    "current": "",
    "cloudURL": ""
}
EOF
```

Luego rellena `projects` y `versions` con tus datos.

---

## Comandos básicos

```bash
makeTicket <comando> [argumento]
```

| Comando | Descripción |
|---------|-------------|
| `init <número-ticket>` | Crea la carpeta del ticket con la plantilla del PM y los metadatos. |
| `config` | Muestra el ticket, proyecto y versión actualmente cargados. |
| `edit` | Abre el archivo `<ticket>.txt` del ticket actual en el editor (`$EDITOR` o `nvim` por defecto). |
| `xml` | Abre el archivo `<ticket>.txt` del ticket actual en el editor. |
| `status` | Reservado; actualmente no devuelve información. |
| `help` | Lista los comandos disponibles. |

### Crear un ticket

```bash
./makeTicket init 1234567
```

El proceso te pedirá:

1. Seleccionar el proyecto a modificar.
2. Seleccionar la versión o línea de trabajo.

Después, crea la siguiente estructura (la semana se calcula automáticamente en formato `AAAA_S`):

```
./2026_37/
└── 1234567/
    ├── 1234567.txt   # Plantilla del PM rellena
    └── .config       # Metadatos: ticket,proyecto,versión
```

Además, actualiza automáticamente el campo `current` en tu `config.json` global con `1234567`.

### Ver la configuración cargada

```bash
./makeTicket config
```

### Editar el ticket actual

```bash
./makeTicket edit
```

Se abre el archivo `<ticket>.txt` del ticket que esté cargado como actual.

---

## Contexto: estar dentro del directorio o en un directorio general

`makeTicket` puede trabajar de dos formas:

### 1. Desde el interior de un ticket

Si ejecutas el comando dentro de una carpeta que contenga un archivo `.config`, la aplicación lee ese `.config` local y carga:

- Número de ticket
- Proyecto
- Versión

Esto permite ejecutar `config`, `edit` o `xml` directamente sin depender del ticket `current` del `config.json` global.

```bash
cd 2026_37/1234567
../../makeTicket edit
```

### 2. Desde cualquier otro directorio

Si no hay un `.config` en la carpeta actual, la aplicación usa el campo `current` del `config.json` global para localizar el ticket en:

```
./<año>_<semana>/<current>/.config
```

Por eso, después de crear un ticket con `init`, puedes moverte a otra carpeta y seguir abriéndolo con `edit` o `xml`.

### Ejemplo de flujo

```bash
# Creas un ticket desde la raíz de tu espacio de trabajo
./makeTicket init 1234567

# Más tarde, desde cualquier otra carpeta
./makeTicket edit          # Abre 1234567.txt gracias a "current"

# O entras directamente al ticket y editas desde ahí
cd 2026_37/1234567
../../makeTicket edit
```

---

## Variables de entorno

| Variable | Descripción |
|----------|-------------|
| `EDITOR` | Editor que se usará en los comandos `edit` y `xml`. Si no está definida, se usa `nvim`. |

---

## Notas

- El binario compilado se llama `makeTicket` y está ignorado en el repositorio.
- La carpeta del ticket se agrupa por año y número de semana ISO (`AAAA_S`).
- Los mensajes de la aplicación están en español.
- `status` está reservado para funcionalidad futura.
