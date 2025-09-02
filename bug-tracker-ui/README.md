# Bug Tracker UI

This is the UI for the Bug Tracker created by members of Rewriting the Code. It's made in React with Vite, Typescript, TailwindCSS, and DaisyUI.

## Running the UI

You need the following

- Node JS (comes with NPM, the Node Package Manager)

## Run the UI Locally

Clone the project

```bash
  git clone https://github.com/WNBARookie/BugTracker.git
```

Go to the project directory

```bash
  cd BugTracker
  cd bug-tracker-ui
```

Install dependencies

```bash
  npm install
```

Start the application

```bash
  npm run dev
```

The project will spin up and run at http://localhost:5000/. Navigate to this URL in your browser

By default it will take you to the home page

## Environment Variables

Before running the project, create a `.env` file in the root of the `bug-tracker-ui` directory.

You can use the provided `.env.example` file as a reference:

```bash
cp .env.example .env
