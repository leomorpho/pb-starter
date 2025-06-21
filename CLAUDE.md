# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a full-stack application with two main components:
- **pb/**: PocketBase backend server (Go)
- **sk/**: SvelteKit frontend application (TypeScript/Svelte)

## Development Commands

### Frontend (sk/ directory)
```bash
cd sk
npm run dev          # Start development server
npm run build        # Build for production
npm run preview      # Preview production build
npm run check        # Type checking with svelte-check
npm run lint         # Lint and format check (Prettier + ESLint)
npm run format       # Format code with Prettier
npm run test         # Run all tests (unit + e2e)
npm run test:unit    # Run Vitest unit tests
npm run test:e2e     # Run Playwright e2e tests
npm run storybook    # Start Storybook development server
```

### Backend (pb/ directory)
```bash
cd pb
go build            # Build the PocketBase binary
./pocketbase serve --dev --http 0.0.0.0:8090 --publicDir ../sk/build
```

For development with auto-reload, use `modd` (if installed):
```bash
cd pb
modd               # Watch Go files and auto-restart server
```

## Architecture Overview

### Backend (PocketBase)
- Simple Go application using PocketBase framework
- Serves static files from `../sk/build` directory
- Runs on port 8090 in development
- Uses `modd.conf` for development auto-reload

### Frontend (SvelteKit)
- Uses SvelteKit with TypeScript and Tailwind CSS
- Configured as static site adapter (`@sveltejs/adapter-static`)
- Internationalization with Paraglide (English and French)
- Testing setup with Vitest (unit) and Playwright (e2e)
- Storybook for component development
- Uses pnpm as package manager

### Key Frontend Features
- **Internationalization**: Paraglide middleware handles locale routing and message interpolation
- **Component Library**: Storybook setup for isolated component development
- **Testing**: Comprehensive testing with unit tests (Vitest) and e2e tests (Playwright)
- **Styling**: Tailwind CSS with typography plugin

### File Structure Notes
- Messages for i18n are in `sk/messages/{locale}.json`
- Generated Paraglide files are in `sk/src/lib/paraglide/`
- Storybook stories are in `sk/src/stories/`
- E2E tests are in `sk/e2e/`

## Development Workflow

1. Start backend: `cd pb && modd` (or manually build and run)
2. Start frontend: `cd sk && npm run dev`
3. Frontend dev server proxies to backend on port 8090
4. Build frontend for production before running backend serve command to serve static files

## Testing

Always run the full test suite before commits:
```bash
cd sk
npm run test        # Runs both unit and e2e tests
```

For development, run tests individually:
```bash
npm run test:unit   # Faster feedback during development
npm run test:e2e    # Integration testing
```