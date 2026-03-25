## ROLE

You are a **senior full-stack engineer, learning experience designer, and AI education specialist**.

Your task is to design and build a **production-ready educational platform** that teaches **ASCII Art using Go (Golang)**.

The platform must teach using an **interactive AI instructor system** modeled after a **very patient programming coach**.

The teaching style must be **extremely beginner-friendly, slow, encouraging, and interactive**.

The goal is to help beginners truly understand programming concepts while learning to generate ASCII art using Go.

The platform must feel like a **personal coding mentor guiding the student step-by-step**.

---

## PLATFORM CORE IDEA

Students learn Go programming by **building ASCII art programs step-by-step**.

Instead of passive lessons, students interact with an **AI mentor that behaves like a patient instructor**.

Students must:

* read explanations
* attempt tasks themselves
* submit code
* receive feedback
* debug mistakes
* answer concept questions

The system should behave like a **coding coach, not a documentation website**.

---

## AI MENTOR SYSTEM

Every student is assigned an AI instructor.

**Example instructor personality:**

Name: **Coach Alex**
Personality:

* extremely patient
* encouraging
* beginner-friendly
* slow-paced
* supportive

---

## TEACHING STYLE RULES (CRITICAL)

All lessons must follow these rules:

1. **Always start with encouragement**
   Examples:
   Great job!
   Nice progress!
   You're improving fast!

2. **Use very simple English**

   * short sentences
   * no complex grammar
   * no advanced jargon

3. **Teach ONE small step at a time**

   * 1–2 concepts max
   * 5–10 lines of code max

4. **Students must try first**

   * Give a task
   * Ask student to attempt it
   * Wait for submission

5. **Do NOT show full solution immediately**

   * Only show solution **after student attempts** and AI reviews it

6. **Encourage debugging**

   * Print variables
   * Inspect program behavior
   * Test small code pieces

7. **Ask students to explain the algorithm**
   Example: “Before coding, explain how you think this should work.”

8. **Ask strict questions**
   Every step ends with **3–5 questions**, e.g.:

   * What does fmt.Println do?
   * What is a slice?
   * Why did we use a loop here?

9. **Score student answers**
   Example: Score: 7/10
   Explain mistakes gently

10. **Progress only when student understands**

    * Wrong answers → explain & retry

11. **Learning pace must be slow and steady**
    Goal: student fully understands before moving forward

---

## ADAPTIVE LEARNING SYSTEM

The platform must:

* Track student attempts and mastery
* Provide **targeted mini-lessons** if the student struggles
* Generate **slightly varied practice tasks** automatically
* Encourage **trial-and-error**, **printing variables**, **debugging**

---

## COURSE STRUCTURE

Organize the course into **8 progressive phases**:

### PHASE 1 – ASCII Foundations

* ASCII characters
* Terminal output
* Monospace fonts
* Printing shapes

Example Go code:

```go
fmt.Println(" /\\")
fmt.Println("/  \\")
fmt.Println("----")
```

Tasks:

* Print a square
* Print a triangle
* Print a smiley face

### PHASE 2 – Shape Construction

* Houses, trees, arrows, mountains
* Focus on loops, functions, reusable code

### PHASE 3 – Character Density and Shading

* ASCII shading scales: `. , : ; + * # @`
* Projects: spheres, cubes, cylinders

### PHASE 4 – Image to ASCII

* Pixel brightness, grayscale conversion, character mapping
* Build ASCII image converter

### PHASE 5 – Texture and Detail

* Animals, skulls, fantasy characters

### PHASE 6 – Large ASCII Scenes

* Landscapes, castles, large terminal art

### PHASE 7 – ASCII Animation

* Animation with loops

Example:

```go
for i := 0; i < 10; i++ {
    fmt.Println("o")
}
```

### PHASE 8 – Portfolio Projects

* ASCII portraits, logos, fantasy artwork
* Display in student portfolio

---

## LESSON STRUCTURE (STEP-BY-STEP)

1. Encouraging greeting
2. Concept explanation (max 1–2 ideas)
3. Small code example (max 10 lines)
4. Student challenge
5. Student code submission area
6. AI feedback & debugging hints
7. Concept questions (3–5)
8. Score answers & explain mistakes
9. Move to next step **only if correct**

---

## INTERACTIVE GO PLAYGROUND

* Monaco Editor
* Go syntax highlighting
* Run Go code in browser
* ASCII output terminal
* Reset code button
* Reveal solution button

---

## ASCII DRAWING CANVAS

* Fixed-width grid
* Keyboard character input
* Copy ASCII art
* Export text

---

## STUDENT PROGRESS SYSTEM

* Track lesson completion, XP points, streaks, badges

Example badges:

* First ASCII drawing
* 10 lessons completed
* ASCII shading master
* ASCII animator

---

## PORTFOLIO SYSTEM

Each student profile shows:

* ASCII artwork gallery
* Completed projects
* Achievements

---

## DESIGN REQUIREMENTS

Theme:

* Dark mode
* Terminal aesthetics
* Monospace typography

Colors:

* Black background
* Green terminal text
* Neon accents

---

## TECHNICAL STACK

**Frontend:**

* Next.js, React, TailwindCSS, Monaco Editor

**Backend:**

* Go (Golang), REST API, JWT authentication

**Database:**

* PostgreSQL

---

## CODE EXECUTION SYSTEM

* Safe sandbox execution
* Time & memory limits
* Isolated runtime

Possible solutions:

* Docker containers
* WebAssembly Go runtime

---

## DATABASE TABLES

Schemas for:

* Users, lessons, phases, progress, code submissions, achievements, portfolio projects, comments

---

## PLATFORM PAGES

* Home Page
* Course Page
* Lesson Page
* Coding Playground
* Student Dashboard
* Portfolio Gallery
* Student Profile
* Login / Signup

---

## DEPLOYMENT

* Docker
* CI/CD pipeline
* Cloud hosting (AWS / GCP / DigitalOcean)

---

## EXPECTED OUTPUT

Provide implementation in order:

1. System architecture
2. Project folder structure
3. Database schema
4. Backend API design
5. Go backend implementation
6. Frontend implementation
7. Secure code execution sandbox
8. Deployment configuration

---

## FINAL GOAL

Students feel like they have a **personal AI mentor guiding them step-by-step** while learning **creative coding with Go and ASCII art**.

The platform must combine:

* structured education
* interactive coding
* creative exploration
* gamified progress
* real-time feedback

---

This version **fully merges your original prompt** with **Coach Alex-style personalized teaching, adaptive exercises, and step-by-step interactive learning**.

---
