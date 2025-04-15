# JotLang CLI Templates

The JotLang CLI provides a powerful templating system to quickly scaffold common application structures and components.

## Basic Commands

### Create New Project

```bash
jot new <project-name> [--template <template-name>]
```

Available project templates:
- `api` - RESTful API project
- `web` - Web application
- `cli` - Command-line application
- `lib` - Library project

### Create Component

```bash
jot generate <component-type> <name> [options]
```

## Available Templates

### API Project

```bash
jot new my-api --template api
```

Creates a new API project with the following structure:
```
my-api/
├── src/
│   ├── controllers/
│   ├── models/
│   ├── services/
│   └── main.jt
├── tests/
├── config.jt
└── README.md
```

### Web Application

```bash
jot new my-web --template web
```

Creates a new web application with:
```
my-web/
├── src/
│   ├── pages/
│   ├── components/
│   ├── services/
│   └── main.jt
├── public/
├── config.jt
└── README.md
```

### Library Project

```bash
jot new my-lib --template lib
```

Creates a new library project:
```
my-lib/
├── src/
│   ├── lib.jt
│   └── types.jt
├── examples/
├── tests/
└── README.md
```

## Component Generation

### API Components

```bash
# Generate controller
jot generate controller User

# Generate model
jot generate model User

# Generate service
jot generate service UserService

# Generate complete CRUD
jot generate crud User
```

### Web Components

```bash
# Generate page
jot generate page Home

# Generate component
jot generate component Button

# Generate service
jot generate service AuthService
```

## Custom Templates

You can create custom templates in `~/.jotlang/templates/`:

```
~/.jotlang/templates/
├── project/
│   └── custom-api/
│       ├── template.json
│       └── files/
└── component/
    └── custom-component/
        ├── template.json
        └── files/
```

### Template Configuration

Example `template.json`:
```json
{
    "name": "custom-api",
    "description": "Custom API template",
    "variables": [
        {
            "name": "projectName",
            "type": "string",
            "description": "Name of the project"
        },
        {
            "name": "port",
            "type": "number",
            "default": 8080,
            "description": "Port number"
        }
    ],
    "files": [
        {
            "source": "main.jt",
            "destination": "src/main.jt"
        }
    ]
}
```

## Best Practices

1. Use meaningful names for components
2. Follow the standard project structure
3. Include proper documentation
4. Add tests for generated components
5. Keep templates modular and reusable
6. Use consistent naming conventions
7. Include error handling in templates
8. Add proper logging
9. Follow security best practices
10. Include configuration files 