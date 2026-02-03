Deployed on Railway.app finance-project.railway.internal

A minimal, production-oriented Expense Tracker that allows users to record, view, filter, and analyze personal expenses.
Built with Go (Gin) for the backend and a simple static frontend served from the same backend.

Features
Core Features

Add an expense (amount, category, description, date)

View list of expenses

Filter expenses by category

Sort expenses by date (newest first)

Display total of currently visible expenses

Reliability & Real-World Behavior

Idempotent API to handle retries and page refreshes

Safe handling of multiple submit clicks

Refresh-safe frontend

Handles slow or failed API responses

Nice-to-Have Enhancements

Mandatory category validation

Searchable category filter (prefix search)

Clear filter & toggle sort

Clean, minimal UI focused on correctness

 Project Structure
backend/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
│
├── internal/
│   ├── db/
│   │   └── db.go            # Database initialization
│   └── expense/
│       ├── model.go         # Expense data model
│       ├── repository.go    # DB access layer
│       └── handler.go       # HTTP handlers
│
├── static/
│   ├── index.html           # Frontend UI
│   └── app.js               # Frontend logic
│
├── go.mod
├── go.sum
└── Dockerfile

Why this structure?

Follows standard Go project conventions

Clear separation of concerns:

cmd/ → application entry point

internal/ → business logic (not externally importable)

static/ → frontend assets

Easy to scale and maintain

🌐 Why the Frontend Is Served from the Backend
Decision

The frontend (HTML + JS) is served directly from the Go backend using Gin’s static file support.

Why this approach?

Single deployment → one service, one URL

No CORS issues

Simpler configuration and fewer moving parts

Ideal for a small full-stack assignment

Benefits

Frontend and backend share the same origin

No need for separate hosting (e.g., Netlify/Vercel)

Production-ready for small tools and internal apps

How it works

Static assets are served under /static

Root path / serves index.html

API endpoints (/expenses) coexist cleanly

 Database Choice: SQLite
Why SQLite?

Lightweight, file-based relational database

Zero setup and operational overhead

ACID compliant

Perfect for:

Personal finance tools

Single-user or low-traffic apps

Assignments and prototypes that still require correctness

Trade-off

Not ideal for high concurrency or distributed systems

Easily replaceable with PostgreSQL or MySQL later

Production Thinking

The data access layer is isolated, making migration to another SQL database straightforward.

 Money Handling

Monetary values are stored as integers (paise) instead of floats

Avoids floating-point precision errors

Common best practice in financial systems

Example:

₹200.50 → stored as 20050

 Idempotency & Reliability
Problem Addressed

Users may:

Click submit multiple times

Refresh the page after submitting

Retry requests due to slow networks

Solution

Client sends an Idempotency-Key header

Backend uses this key as the unique expense ID

Duplicate requests with the same key do not create duplicates

Result

Safe retries

No duplicate expenses

Production-grade behavior

 Frontend Design Decisions

Plain HTML + JavaScript (no frameworks)

Focused on clarity and correctness, not styling

Client-side:

Input validation

Category search & filtering

Total calculation

Sort toggling

Why no React/Vue?

Overkill for the scope

Increases complexity without adding value here

The goal was to demonstrate system thinking, not framework usage

 Error Handling & UX

Loading indicators during API calls

Error messages on failed requests

Disabled submit button during submission

Clear form fields after successful submit

 Running Locally
Prerequisites

Go 1.22+

Run
cd backend/cmd/server
go run main.go


Open in browser:

http://localhost:8080

🐳 Deployment

Dockerized Go application

<img width="1768" height="726" alt="image" src="https://github.com/user-attachments/assets/10e306e9-2287-4a02-97de-fe45df06e96e" />


Successfully deployable on platforms like Railway

