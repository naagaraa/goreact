# React + Vite

### run dev

```
bun dev
```

### build prod

```
bun build
```

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

# Backend go fiber

### run dev

```
bun go-dev
```

### build prod

```
bun go-build
```

# Backend Nestjs

### run dev

```
bun start-nest:dev
```

### build prod

```
bun build-nest
```

### run prod

```
bun start-nest:prod
```

## development and deploy ?

for development you need 2 to 3 terminal open, 1 for go apps, 1 for go nest (choose) want
you want or need. 1 for react apps.

for deploy i am really tested, this 1 repo for 3 codebase 2 backend nestjs and gofiber. for build
production with go fiber can use "go build -o server" or "bun go-build" file will create in root server, and then can run with "./server" if using nginx you can used this file

and for nestjs production build will be store at dist/server, so running with "bun nest-start:prod" before that run "bun build-nest"
