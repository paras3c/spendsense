# Spendsense: AI-Powered Financial Tutor

## Introduction
Spendsense is a financial analytics platform designed to solve the problem of financial illiteracy among young professionals. Traditional budgeting apps provide data but lack context. Spendsense bridges this gap by using Generative AI to analyze spending habits, identify waste (such as unused subscriptions), and generate personalized financial personas.

This project was built for the hackathon to demonstrate how AI can transform raw financial data into actionable, high-level economic intelligence.

## Features
- **Editorial E-Ink Dashboard**: A high-contrast, black-and-white interface designed for maximum accessibility and focus. Eliminates visual noise to prioritize data.
- **AI Intelligence Layer**: Automatically analyzes transaction history to flag "Wealth Leaks" (e.g., subscription waste) and "High Food Spending".
- **Savage Mode**: A unique, personality-driven AI analysis that provides "brutally honest" feedback on spending habits, gamifying the budgeting experience.
- **Expenditure Mapping**: Visualizes financial outflows using high-contrast doughnut charts for immediate category recognition.
- **Privacy-First Architecture**: Financial data is processed locally or in memory for the session duration, ensuring sensitive data is not permanently stored.

## Technical Stack

### Frontend
- **Framework**: SvelteKit (TypeScript) for performant, server-side rendered UI.
- **Styling**: Tailwind CSS with a custom "Editorial" design system (Playfair Display & Lora typography).
- **Visualization**: Chart.js for responsive, accessible data visualization.

### Backend
- **Language**: Go (Golang) 1.25+ for high-concurrency API processing.
- **Server**: Standard library `net/http` for minimal footprint and maximum control.
- **Data Parsing**: Custom CSV parser optimized for standard bank statement formats.

### Artificial Intelligence
- **Engine**: Large Language Model (LLM) integration for semantic analysis of transaction descriptions and natural language advice generation.
- **Persona Engine**: Dynamic prompting system to generate user archetypes (e.g., "The Impulse Shopper") based on spending patterns.

## Installation & Usage

### Prerequisites
- Node.js (v18+)
- Go (v1.21+)

### Backend Setup
1. **AI Setup (Ollama)**
   This project uses `Ollama` to run local LLMs.
   - Install Ollama from [ollama.com](https://ollama.com).
   - Pull the required model (we use `tinyllama` for speed, or `llama3` for better results):
     ```bash
     ollama pull tinyllama
     ```
   - Ensure Ollama is running (`ollama serve`).

2. Navigate to the backend directory:
   ```bash
   cd ai-finance-tutor/backend
   ```
3. Start the server:
   ```bash
   go run main.go
   ```
   The backend will start on `http://localhost:8000`.

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd ai-finance-tutor/frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```
   The application will be accessible at `http://localhost:5173`.

## How to Use
1. Open the application in your browser.
2. Click **Use Sample Data** to load a demonstration dataset, or upload a CSV file containing transaction history.
3. Explore the **AI Intelligence** section for automated insights.
4. Toggle **Savage Mode** in the top navigation to switch the AI personality from "Polite Financial Advisor" to "Brutal Money Critic".
5. Click **Generate Persona** to receive a comprehensive analysis of your spending character.
