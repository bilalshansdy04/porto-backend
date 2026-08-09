# Porto Backend API Documentation

This documentation provides details on how to interact with the backend API for the Porto project. The backend is built with Go, Gin framework, and MySQL. Base URL for all endpoints is `http://localhost:8080`.

## 1. Dashboard & Profile API

Endpoints to manage the user profile summary and retrieve dashboard statistics.

### A. Create Profile
Used to initialize the profile data if it does not exist.
- **Endpoint:** `POST /api/profile`
- **Headers:** `Content-Type: application/json`
- **Body:**
```json
{
  "summary": "Senior Developer specializing in scalable architectures and modern UI engineering.",
  "years_of_experience": 5
}
```

### B. Update Profile
- **Endpoint:** `PUT /api/profile`
- **Headers:** `Content-Type: application/json`
- **Body:** Same as the Create Profile payload.

### C. Get Dashboard Stats
- **Endpoint:** `GET /api/dashboard/stats`
- **Description:** Returns total projects, years of experience, profile summary, and 3 most recent projects.

---

## 2. Projects API

Endpoints for managing portfolio projects. Project creation is split into two parts: basic initialization with an image upload, and subsequent updating of text details (arrays/lists).

### A. Create New Project (Initialize)
- **Endpoint:** `POST /api/projects`
- **Headers:** `Content-Type: multipart/form-data`
- **Body Form-Data:**
  - `name` (text): Project name (e.g., "E-Commerce Dashboard")
  - `image` (file): Upload image file (PNG/JPG/SVG)

### B. Update Project Details
- **Endpoint:** `PUT /api/projects/:id`
- **Headers:** `Content-Type: application/json`
- **Body:**
```json
{
  "name": "E-Commerce Dashboard V2",
  "description": "Membangun dashboard yang scalable dan interaktif untuk manajemen e-commerce.",
  "status": "Live",
  "is_complete": true,
  "tech_stack": ["React", "Go", "MySQL"],
  "project_flow": ["Research phase", "Design Prototyping", "Implementation"],
  "jobdesc": ["Lead Frontend", "Backend Architect"]
}
```

### C. Retrieve Projects
- **Get All Projects:** `GET /api/projects`
- **Get Single Project:** `GET /api/projects/:id`

### D. Delete Project
- **Endpoint:** `DELETE /api/projects/:id`

---

## 3. Professional Journey (Experiences) API

Endpoints to manage work experiences shown in the timeline.

### A. Create Experience
- **Endpoint:** `POST /api/experiences`
- **Headers:** `Content-Type: application/json`
- **Body:**
```json
{
  "company_name": "Tech Solutions Inc.",
  "role": "Senior Developer",
  "start_date": "2021-01",
  "end_date": "",
  "is_current": true,
  "responsibilities": [
    "Spearheaded the migration to a modern React stack, improving load times by 40%.",
    "Membangun arsitektur microservices untuk fitur e-wallet."
  ]
}
```

### B. Update Experience
- **Endpoint:** `PUT /api/experiences/:id`
- **Headers:** `Content-Type: application/json`
- **Body:** Same as the Create Experience payload.

### C. Retrieve Experiences
- **Endpoint:** `GET /api/experiences`

### D. Delete Experience
- **Endpoint:** `DELETE /api/experiences/:id`

---

## 4. Skills API

Endpoints to manage skills categorized into different sections (e.g., Bahasa Pemrograman, Framework & Lingkungan).

### A. Create Skill
- **Endpoint:** `POST /api/skills`
- **Headers:** `Content-Type: application/json`
- **Body:**
```json
{
  "name": "Go",
  "category": "Bahasa Pemrograman"
}
```

### B. Update Skill
- **Endpoint:** `PUT /api/skills/:id`
- **Headers:** `Content-Type: application/json`
- **Body:** Same as the Create Skill payload.

### C. Retrieve Skills
- **Endpoint:** `GET /api/skills`

### D. Delete Skill
- **Endpoint:** `DELETE /api/skills/:id`
